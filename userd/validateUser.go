package userd

import (
	"context"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func (usr *Userd) ValidateUser(ctx context.Context, req *userdpb.ValidateUserRequest) (*userdpb.GetUserResponse, error) {

	usr.usrStrg.GetLoggerHandle().Println("Userd: Validate user")

	cu := &userdpb.CheckUser{
		Email:    req.Email,
		Password: req.Password,
	}

	u, err := usr.usrStrg.ValidateUser(ctx, cu)

	return &userdpb.GetUserResponse{
		User: u,
	}, err
}
