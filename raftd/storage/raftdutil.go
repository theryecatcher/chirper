package raftstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/raft"
)

const (
	noOfSnapshots = 3
	raftTimeout   = 10 * time.Second
)

type command struct {
	Op    string `json:"op,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// RaftStore is a simple key-value store, where all changes are made via Raft consensur.
type RaftStore struct {
	raftLocalDir  string
	raftNodeAddr  string
	raftMut       sync.Mutex
	keyValueStore map[string]string
	raft          *raft.Raft

	logger *log.Logger
}

// NewRaftStore returns a new Raft Storage.
func NewRaftStore(raftLocalDir string, raftNodeAddr string) *RaftStore {
	return &RaftStore{
		keyValueStore: make(map[string]string),
		raftLocalDir:  raftLocalDir,
		raftNodeAddr:  raftNodeAddr,
		logger:        log.New(os.Stderr, "[raftdb] ", log.LstdFlags),
	}
}

// GetLoggerHandle return the logger handle
func (r *RaftStore) GetLoggerHandle() *log.Logger {
	return r.logger
}

// RaftCreate creates the store. If master is set, and there are no existing peers,
// then this node becomes the first node, and therefore leader, of the cluster.
// localID should be the server identifier for this node.
func (r *RaftStore) RaftCreate(master bool, nodeID string) error {
	// Setup Raft configuration.
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(nodeID)

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", r.raftNodeAddr)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(r.raftNodeAddr, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(r.raftLocalDir, noOfSnapshots, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}

	// Create the log store and stable store.
	var logStore raft.LogStore
	var stableStore raft.StableStore

	logStore = raft.NewInmemStore()
	stableStore = raft.NewInmemStore()

	// Instantiate the Raft system.
	raftNode, err := raft.NewRaft(config, (*fsm)(r), logStore, stableStore, snapshots, transport)
	if err != nil {
		return fmt.Errorf("Error in creating Raft Node %s : %s", nodeID, err)
	}
	r.raft = raftNode

	if master {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		raftNode.BootstrapCluster(configuration)
	}
	return nil
}

// RaftGet returns the value for the given key.
func (r *RaftStore) RaftGet(ctx context.Context, key string) (string, error) {

	result := make(chan string)
	oops := make(chan error)

	r.logger.Println(key)

	go func() {
		r.raftMut.Lock()
		defer r.raftMut.Unlock()
		val, exists := r.keyValueStore[key]
		if !exists {
			oops <- errors.New("Value not found")
			return
		}
		result <- val
	}()

	// Respect the context
	select {
	case res := <-result:
		return res, nil
	case err := <-oops:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// RaftGetAllUsrKeys returns All User Keys.
func (r *RaftStore) RaftGetAllUsrKeys(ctx context.Context) ([]string, error) {

	result := make(chan []string)

	go func() {
		r.raftMut.Lock()
		defer r.raftMut.Unlock()

		var keys []string
		for k := range r.keyValueStore {
			if strings.HasPrefix(k, "usr") {
				keys = append(keys, k)
			}
		}
		result <- keys
	}()

	// Respect the context
	select {
	case res := <-result:
		return res, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RaftSetKeyValue sets the value for the given key.
func (r *RaftStore) RaftSetKeyValue(ctx context.Context, key string, value string) error {
	result := make(chan error)
	oops := make(chan error)

	go func() {
		if r.raft.State() != raft.Leader {
			oops <- errors.New("Not Leader")
			return
		}

		setCommand := &command{
			Op:    "set",
			Key:   key,
			Value: value,
		}
		byteStream, err := json.Marshal(setCommand)
		if err != nil {
			oops <- err
			return
		}

		applyFuture := r.raft.Apply(byteStream, raftTimeout)
		result <- applyFuture.Error()
	}()

	// Respect the context
	select {
	case res := <-result:
		return res
	case err := <-oops:
		r.logger.Println(err)
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RaftDelete deletes the given key.
func (r *RaftStore) RaftDelete(ctx context.Context, key string) error {
	if r.raft.State() != raft.Leader {
		return fmt.Errorf("Not Leader")
	}

	deleteCommand := &command{
		Op:  "delete",
		Key: key,
	}
	byteStream, err := json.Marshal(deleteCommand)
	if err != nil {
		return err
	}

	applyFuture := r.raft.Apply(byteStream, raftTimeout)
	return applyFuture.Error()
}

// RaftJoin joins a node, identified by nodeID and located at addr, to this store.
// The node must be ready to respond to Raft communications at that address.
func (r *RaftStore) RaftJoin(ctx context.Context, nodeID, addr string) error {
	r.logger.Printf("received join request for remote node %s at %s", nodeID, addr)

	config := r.raft.GetConfiguration()
	if err := config.Error(); err != nil {
		r.logger.Printf("failed to get raft configuration: %v", err)
		return err
	}

	for _, servers := range config.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if servers.ID == raft.ServerID(nodeID) || servers.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if servers.Address == raft.ServerAddress(addr) && servers.ID == raft.ServerID(nodeID) {
				r.logger.Printf("Node %s:%s already member of Raft Servers", nodeID, addr)
				return nil
			}

			removeServer := r.raft.RemoveServer(servers.ID, 0, 0)
			if err := removeServer.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := r.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	r.logger.Printf("Node %s:%s added successfully", nodeID, addr)
	return nil
}

type fsm RaftStore

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	var cmd command
	if err := json.Unmarshal(l.Data, &cmd); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}

	switch cmd.Op {
	case "set":
		return f.applySet(cmd.Key, cmd.Value)
	case "delete":
		return f.applyDelete(cmd.Key)
	default:
		panic(fmt.Sprintf("unrecognized command op: %s", cmd.Op))
	}
}

// Snapshot returns a snapshot of the key-value store.
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	f.raftMut.Lock()
	defer f.raftMut.Unlock()

	// Clone the map.
	o := make(map[string]string)
	for k, v := range f.keyValueStore {
		o[k] = v
	}
	return &fsmSnapshot{store: o}, nil
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	o := make(map[string]string)
	if err := json.NewDecoder(rc).Decode(&o); err != nil {
		return err
	}

	// Set the state from the snapshot, no lock required according to
	// Hashicorp docr.
	f.keyValueStore = o
	return nil
}

func (f *fsm) applySet(key, value string) interface{} {
	f.raftMut.Lock()
	defer f.raftMut.Unlock()
	f.keyValueStore[key] = value
	return nil
}

func (f *fsm) applyDelete(key string) interface{} {
	f.raftMut.Lock()
	defer f.raftMut.Unlock()
	delete(f.keyValueStore, key)
	return nil
}

type fsmSnapshot struct {
	store map[string]string
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		byteStream, err := json.Marshal(f.store)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(byteStream); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (f *fsmSnapshot) Release() {}
