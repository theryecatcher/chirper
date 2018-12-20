package raftd

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/theryecatcher/chirper/raftd/raftdpb"

	"github.com/google/uuid"
)

// Test_RaftDSingleNode tests that a command can be applied to the log
func Test_RaftDSingleNode(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "raft_test")
	defer os.RemoveAll(tmpDir)

	raftCfg := &Config{
		NodeAddr: "127.0.0.1:0",
		LocalDir: tmpDir,
		NodeID:   "node0",
	}

	s, err := New(raftCfg)

	if err != nil {
		t.Fatalf("Failed to create Raft Store")
	}

	// Simple way to ensure there is a leader.
	time.Sleep(3 * time.Second)

	if _, err := s.RaftSetKeyValue(context.Background(), &raftdpb.RaftSetKeyValueRequest{
		Key:   "abc",
		Value: "xyz",
	}); err != nil {
		t.Fatalf("Failed to set key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err := s.RaftGet(context.Background(), &raftdpb.RaftGetRequest{
		Key: "abc",
	})
	if err != nil {
		t.Fatalf("Failed to get key: %s", err.Error())
	}
	if value.Value != "xyz" {
		t.Fatalf("key has wrong value: %s", value)
	}

	if _, err := s.RaftDelete(context.Background(), &raftdpb.RaftDeleteRequest{
		Key: "abc",
	}); err != nil {
		t.Fatalf("Failed to delete key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err = s.RaftGet(context.Background(), &raftdpb.RaftGetRequest{
		Key: "abc",
	})
	if err == nil {
		t.Fatalf("Got key: %s", value.Value)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Test_RaftDSCreateSingleNodeMultipleValues tests that a multiple commands can be applied to the log
func Test_RaftDSCreateSingleNodeMultipleValues(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "raft_test")
	defer os.RemoveAll(tmpDir)

	raftCfg := &Config{
		NodeAddr: "127.0.0.1:0",
		LocalDir: tmpDir,
		NodeID:   "node0",
	}

	s, err := New(raftCfg)

	if err != nil {
		t.Fatalf("Failed to create Raft Store")
	}

	// Simple way to ensure there is a leader.
	time.Sleep(3 * time.Second)

	// Generate 100 UUID's
	var uids []string
	for i := 1; i <= 100; i++ {
		uid := uuid.New().String()
		if contains(uids, uid) {
			uid = uuid.New().String()
		}
		uids = append(uids, uid)
	}

	for _, val := range uids {
		if _, err := s.RaftSetKeyValue(context.Background(), &raftdpb.RaftSetKeyValueRequest{
			Key:   "usr:" + val,
			Value: "xyz",
		}); err != nil {
			t.Fatalf("Failed to set key: %s", err.Error())
		}
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err := s.RaftGetAllUsrKeys(context.Background(), &raftdpb.RaftGetAllUsrKeysRequest{})
	if err != nil {
		t.Fatalf("Failed to get key: %s", err.Error())
	}
	if len(value.Value) != len(uids) {
		t.Fatalf("Expected 100 entries")
	}
}
