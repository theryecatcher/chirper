package userd

import (
	"context"
	"log"

	"github.com/distsys-project/web/userd/userdpb"
)

func (cnt *userd) GetUser(ctx context.Context, req *userdpb.GetUserRequest) (*userdpb.GetUserResponse, error) {
	log.Println("Userd: get tweet")
	t, err := cnt.strg.GetUser(ctx, req.TID)

	return &userdpb.GetUserResponse{
		Tweet: t,
	}, err
}
