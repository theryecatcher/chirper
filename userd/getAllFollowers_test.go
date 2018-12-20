package userd

import (
	"context"
	"testing"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func TestUserD_GetAllFollowers(t *testing.T) {
	ds, err := New(&Config{})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}

	_, err = ds.NewUser(context.Background(), &userdpb.NewUserRequest{
		Name:     "Cool",
		Email:    "hi@work.com",
		Password: "Cool",
	})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}

	val, err := ds.GetAllFollowers(context.Background(), &userdpb.FollowerDetailsRequest{
		UID: "Cool",
	})

	if val != nil {
		t.Fatalf("Didn't expect any followers")
	}
}
