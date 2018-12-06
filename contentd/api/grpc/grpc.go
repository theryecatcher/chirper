package contentd

import (
	"log"

	"github.com/theryecatcher/chirper/contentd"
	"github.com/theryecatcher/chirper/contentd/contentdpb"
	"google.golang.org/grpc"
)

// RegisterGRPCServer returns the ContentD GRPC API to the caller
func RegisterGRPCServer(grpcCntServer *grpc.Server) error {
	cntCfg := &contentd.Config{}

	contentDb, err := contentd.New(cntCfg)
	if err != nil {
		log.Println(err)
	}

	contentdpb.RegisterContentdServer(grpcCntServer, contentDb)

	return err
}
