package raftd

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

func Test_RaftDDeleteOnSingleNode(t *testing.T) {
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

	if _, err := s.RaftDelete(context.Background(), &raftdpb.RaftDeleteRequest{
		Key: "abc",
	}); err != nil {
		t.Fatalf("Failed to delete key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err := s.RaftGet(context.Background(), &raftdpb.RaftGetRequest{
		Key: "abc",
	})
	if err == nil {
		t.Fatalf("Got key: %s", value.Value)
	}
}
