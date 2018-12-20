package userd

import (
	"context"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func (usr *Userd) NewUser(ctx context.Context, req *userdpb.NewUserRequest) (*userdpb.NewUserResponse, error) {

	usr.usrStrg.GetLoggerHandle().Println("Userd: New User")

	u := &userdpb.NewUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	return &userdpb.NewUserResponse{}, usr.usrStrg.NewUser(ctx, u)
}
