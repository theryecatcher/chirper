package userd

import (
	"context"
	"testing"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func TestUserD_UnFollowUser(t *testing.T) {
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

	_, err = ds.NewUser(context.Background(), &userdpb.NewUserRequest{
		Name:     "Hello",
		Email:    "hi@work.com",
		Password: "Hello",
	})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}

	_, err = ds.FollowUser(context.Background(), &userdpb.FollowUserRequest{
		FollowingUID: "hello",
		UID:          "Cool",
	})

	_, err = ds.UnFollowUser(context.Background(), &userdpb.UnFollowUserRequest{
		FollowedUID: "Hello",
		UID:         "Cool",
	})
}
