package userd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/web/userd/userdpb"
)

// UnFollowUser
func (usr *Userd) UnFollowUser(ctx context.Context, req *userdpb.UnFollowUserRequest) (*userdpb.UnFollowUserResponse, error) {

	log.Println("Userd: unfollow user")
	err := usr.usrStrg.UnFollowUser(ctx, req.UID, req.FollowedUID)

	return &userdpb.UnFollowUserResponse{}, err
}
