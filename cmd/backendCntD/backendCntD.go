package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theryecatcher/chirper/contentd/api/grpc"
	"google.golang.org/grpc"
)

func main() {

	lisCntD, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "5445"))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}

	contentServer := grpc.NewServer()
	contentd.RegisterGRPCServer(contentServer)
	contentServer.Serve(lisCntD)
}
