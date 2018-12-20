package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

// RaftDelete function
func (raft *Raftd) RaftDelete(ctx context.Context, req *raftdpb.RaftDeleteRequest) (*raftdpb.RaftDeleteResponse, error) {
	raft.raftStrg.GetLoggerHandle().Println("Raftd: Delete Request")
	err := raft.raftStrg.RaftDelete(ctx, req.Key)

	return &raftdpb.RaftDeleteResponse{}, err
}
