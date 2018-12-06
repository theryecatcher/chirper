package userd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func (usr *Userd) FollowUser(ctx context.Context, req *userdpb.FollowUserRequest) (*userdpb.FollowUserResponse, error) {

	log.Println("Userd: follow user")
	err := usr.usrStrg.FollowUser(ctx, req.UID, req.FollowingUID)

	return &userdpb.FollowUserResponse{}, err
}
