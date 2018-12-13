package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

// RaftSetKeyValue function
func (raft *Raftd) RaftSetKeyValue(ctx context.Context, req *raftdpb.RaftSetKeyValueRequest) (*raftdpb.RaftSetKeyValueResponse, error) {
	raft.raftStrg.GetLoggerHandle().Println("Raftd: Set Request")
	err := raft.raftStrg.RaftSetKeyValue(ctx, req.Key, req.Value)

	return &raftdpb.RaftSetKeyValueResponse{}, err
}
