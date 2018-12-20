package raftdgrpc

import (
	"log"

	"github.com/theryecatcher/chirper/raftd"
	"github.com/theryecatcher/chirper/raftd/raftdpb"
	"google.golang.org/grpc"
)

// RegisterGRPCServer returns the ContentD GRPC API to the caller
func RegisterGRPCServer(grpcCfg *Config, grpcraftServer *grpc.Server) error {

	raftNode, err := raftd.New(&raftd.Config{
		LocalDir: grpcCfg.RaftLocalDir,
		NodeAddr: grpcCfg.RaftNodeAddr,
		JoinAddr: grpcCfg.RaftJoinAddr,
		NodeID:   grpcCfg.RaftNodeID,
	})
	if err != nil {
		log.Println(err)
	}

	raftdpb.RegisterRaftdServer(grpcraftServer, raftNode)

	return err
}
