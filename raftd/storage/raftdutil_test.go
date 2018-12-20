package raftstorage

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

// Test_StoreOpen tests that the Store can be opened.
func Test_RaftStoreCreate(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "raft_test")
	defer os.RemoveAll(tmpDir)

	s := NewRaftStore(tmpDir, "127.0.0.1:0")

	if s == nil {
		t.Fatalf("Failed to create Raft Store")
	}

	if err := s.RaftCreate(false, "node0"); err != nil {
		t.Fatalf("Failed to open Raft Store: %s", err)
	}
}

// Test_StoreOpenSingleNode tests that a command can be applied to the log
func Test_StoreCreateSingleNode(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "raft_test")
	defer os.RemoveAll(tmpDir)

	s := NewRaftStore(tmpDir, "127.0.0.1:0")

	if s == nil {
		t.Fatalf("Failed to create Raft Store")
	}

	if err := s.RaftCreate(true, "node0"); err != nil {
		t.Fatalf("Failed to open Raft Store: %s", err)
	}

	// Simple way to ensure there is a leader.
	time.Sleep(3 * time.Second)

	if err := s.RaftSetKeyValue(context.Background(), "abc", "xyz"); err != nil {
		t.Fatalf("Failed to set key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err := s.RaftGet(context.Background(), "abc")
	if err != nil {
		t.Fatalf("Failed to get key: %s", err.Error())
	}
	if value != "xyz" {
		t.Fatalf("key has wrong value: %s", value)
	}

	if err := s.RaftDelete(context.Background(), "abc"); err != nil {
		t.Fatalf("Failed to delete key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err = s.RaftGet(context.Background(), "abc")
	if err == nil {
		t.Fatalf("Got key: %s", value)
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

// Test_StoreOpenSingleNode tests that a command can be applied to the log
func Test_StoreCreateSingleNodeMultipleValues(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "raft_test")
	defer os.RemoveAll(tmpDir)

	s := NewRaftStore(tmpDir, "127.0.0.1:0")

	if s == nil {
		t.Fatalf("Failed to create Raft Store")
	}

	if err := s.RaftCreate(true, "node0"); err != nil {
		t.Fatalf("Failed to open Raft Store: %s", err)
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
		if err := s.RaftSetKeyValue(context.Background(), "usr:"+val, "xyz"); err != nil {
			t.Fatalf("Failed to set key: %s", err.Error())
		}
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, err := s.RaftGetAllUsrKeys(context.Background())
	if err != nil {
		t.Fatalf("Failed to get key: %s", err.Error())
	}
	if len(value) != len(uids) {
		t.Fatalf("Expected 100 entries")
	}
}
