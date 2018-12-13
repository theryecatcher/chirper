package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
)

// RaftGetAllUsrKeys function
func (raft *Raftd) RaftGetAllUsrKeys(ctx context.Context, req *raftdpb.RaftGetAllUsrKeysRequest) (*raftdpb.RaftGetAllUsrKeysResponse, error) {
	raft.raftStrg.GetLoggerHandle().Println("Raftd: Get all User Keys Request")
	v, err := raft.raftStrg.RaftGetAllUsrKeys(ctx)

	return &raftdpb.RaftGetAllUsrKeysResponse{
		Value: v,
	}, err
}
