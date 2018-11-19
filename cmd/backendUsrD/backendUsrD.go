package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theryecatcher/chirper/web/userd"
	"github.com/theryecatcher/chirper/web/userd/userdpb"

	"google.golang.org/grpc"
)

func main() {
	usrCfg := &userd.Config{}

	usrDb, err := userd.New(usrCfg)
	if err != nil {
		panic(err)
	}

	lisUsrD, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "5446"))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}

	grpcUsrServer := grpc.NewServer()
	userdpb.RegisterUserdServer(grpcUsrServer, usrDb)
	grpcUsrServer.Serve(lisUsrD)

}
