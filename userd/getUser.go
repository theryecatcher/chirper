package userd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func (usr *Userd) GetUser(ctx context.Context, req *userdpb.GetUserRequest) (*userdpb.GetUserResponse, error) {

	log.Println("Userd: get user")
	u, err := usr.usrStrg.GetUser(ctx, req.UID)

	return &userdpb.GetUserResponse{
		User: u,
	}, err
}
