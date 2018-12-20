package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

// RaftGet function
func (raft *Raftd) RaftGet(ctx context.Context, req *raftdpb.RaftGetRequest) (*raftdpb.RaftGetResponse, error) {
	raft.raftStrg.GetLoggerHandle().Println("Raftd: Get Request")
	v, err := raft.raftStrg.RaftGet(ctx, req.Key)

	return &raftdpb.RaftGetResponse{
		Value: v,
	}, err
}
