package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theryecatcher/chirper/web/contentd"
	"github.com/theryecatcher/chirper/web/contentd/contentdpb"

	"google.golang.org/grpc"
)

func main() {
	cntCfg := &contentd.Config{}

	contentDb, err := contentd.New(cntCfg)
	if err != nil {
		panic(err)
	}

	lisCntD, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "5445"))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}

	grpcCntServer := grpc.NewServer()
	contentdpb.RegisterContentdServer(grpcCntServer, contentDb)
	grpcCntServer.Serve(lisCntD)
}
