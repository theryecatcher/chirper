package userd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func (usr *Userd) NewUser(ctx context.Context, req *userdpb.NewUserRequest) (*userdpb.NewUserResponse, error) {
	log.Println("Userd: New User")

	u := &userdpb.NewUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	return &userdpb.NewUserResponse{}, usr.usrStrg.NewUser(ctx, u)
}
