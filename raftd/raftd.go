package raftd

import (
	"context"

	"github.com/theryecatcher/chirper/raftd/raftdpb"
	"github.com/theryecatcher/chirper/raftd/storage"
	"google.golang.org/grpc"
)

// Raftd Struct
type Raftd struct {
	raftStrg raftstorage.Storage
}

// New function
func New(cfg *Config) (*Raftd, error) {

	r := &Raftd{
		raftStrg: raftstorage.NewRaftStore(cfg.LocalDir, cfg.NodeAddr),
	}

	if err := r.raftStrg.RaftCreate(cfg.JoinAddr == "", cfg.NodeID); err != nil {
		r.raftStrg.GetLoggerHandle().Fatalf("Failed to create store: %s", err.Error())
	}

	if cfg.JoinAddr != "" {

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		raftConn, err := grpc.Dial(cfg.JoinAddr, opts...)
		if err != nil {
			r.raftStrg.GetLoggerHandle().Fatalf("failure while dialing: %v", err)
		}

		c := raftdpb.NewRaftdClient(raftConn)

		if _, err := c.RaftJoin(context.Background(), &raftdpb.RaftJoinRequest{
			NodeID:      cfg.NodeID,
			NodeAddress: cfg.NodeAddr,
		}); err != nil {
			r.raftStrg.GetLoggerHandle().Fatalf("failed to join node at %s: %s", cfg.JoinAddr, err.Error())
		}

		if err = raftConn.Close(); err != nil {
			r.raftStrg.GetLoggerHandle().Fatalf("Failed to Close GRPC Client: %s", err.Error())
		}
	}

	return r, nil
}
