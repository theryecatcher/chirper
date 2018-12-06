package userd

import (
	"log"

	"github.com/theryecatcher/chirper/userd"
	"github.com/theryecatcher/chirper/userd/userdpb"
	"google.golang.org/grpc"
)

// RegisterGRPCServer returns the ContentD GRPC API to the caller
func RegisterGRPCServer(grpcCntServer *grpc.Server) error {
	usrCfg := &userd.Config{}

	userDb, err := userd.New(usrCfg)
	if err != nil {
		log.Println(err)
	}

	userdpb.RegisterUserdServer(grpcCntServer, userDb)

	return err
}
