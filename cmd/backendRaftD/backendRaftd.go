package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/theryecatcher/chirper/raftd/api/grpc"
	"google.golang.org/grpc"
)

// Application parameters
var serverAddr string
var raftAddr string
var joinAddr string
var nodeID string

func init() {
	flag.StringVar(&serverAddr, "grpc", "45000", "Set the GRPC bind address")
	flag.StringVar(&raftAddr, "node", "46000", "Set Raft bind address")
	flag.StringVar(&joinAddr, "leader", "", "Set leader address, if any")
	flag.StringVar(&nodeID, "id", "", "Node ID")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <raft-local-storage-path> \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Please run <cmd> -h (help) for usage instructions\n")
		os.Exit(1)
	}

	if nodeID == "" {
		fmt.Fprintf(os.Stderr, "Please specify the Node ID/Name\n")
		os.Exit(1)
	}

	// Create Raft Directory for Snapshots.
	raftDir := flag.Arg(0)
	if raftDir == "" {
		fmt.Fprintf(os.Stderr, "Need to have a Raft storage directory\n")
		os.Exit(1)
	}
	os.MkdirAll(raftDir, 0700)

	lisraftD, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", serverAddr))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}
	raftNodeServer := grpc.NewServer()

	err = raftdgrpc.RegisterGRPCServer(&raftdgrpc.Config{
		RaftLocalDir: raftDir,
		RaftNodeAddr: raftAddr,
		RaftJoinAddr: joinAddr,
		RaftNodeID:   nodeID,
	}, raftNodeServer)
	if err != nil {
		log.Fatalf("Failed to start Raft Node: %s", err.Error())
	}
	raftNodeServer.Serve(lisraftD)
}
