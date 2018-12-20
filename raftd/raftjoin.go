package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

// RaftJoin function
func (raft *Raftd) RaftJoin(ctx context.Context, req *raftdpb.RaftJoinRequest) (*raftdpb.RaftJoinResponse, error) {
	raft.raftStrg.GetLoggerHandle().Println("Raftd: Join Request")
	err := raft.raftStrg.RaftJoin(ctx, req.NodeID, req.NodeAddress)

	return &raftdpb.RaftJoinResponse{}, err
}
