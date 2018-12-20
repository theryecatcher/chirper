package userd

import (
	"context"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

// GetAllFollowers Function to retuurn all users with Follower details
func (usr *Userd) GetAllFollowers(ctx context.Context, req *userdpb.FollowerDetailsRequest) (*userdpb.FollowerDetailsResponse, error) {

	usr.usrStrg.GetLoggerHandle().Println("Userd: get all followers")
	f, err := usr.usrStrg.GetAllFollowers(ctx, req.UID)

	return &userdpb.FollowerDetailsResponse{
		Followers: f,
	}, err
}
