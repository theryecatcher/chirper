package userd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/web/userd/userdpb"
)

// GetAllFollowers Function to retuurn all users with Follower details
func (usr *Userd) GetAllFollowers(ctx context.Context, req *userdpb.FollowerDetailsRequest) (*userdpb.FollowerDetailsResponse, error) {

	log.Println("Userd: get all followers")
	f, err := usr.usrStrg.GetAllFollowers(ctx, req.UID)

	return &userdpb.FollowerDetailsResponse{
		Followers: f,
	}, err
}
