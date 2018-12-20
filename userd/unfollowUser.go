package userd

import (
	"context"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

// UnFollowUser
func (usr *Userd) UnFollowUser(ctx context.Context, req *userdpb.UnFollowUserRequest) (*userdpb.UnFollowUserResponse, error) {

	usr.usrStrg.GetLoggerHandle().Println("Userd: unfollow user")
	err := usr.usrStrg.UnFollowUser(ctx, req.UID, req.FollowedUID)

	return &userdpb.UnFollowUserResponse{}, err
}
