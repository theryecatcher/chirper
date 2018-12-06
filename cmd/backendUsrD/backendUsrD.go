package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theryecatcher/chirper/userd/api/grpc"

	"google.golang.org/grpc"
)

func main() {

	lisUsrD, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "5446"))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}

	userdServer := grpc.NewServer()
	userd.RegisterGRPCServer(userdServer)
	userdServer.Serve(lisUsrD)
}
