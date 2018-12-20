package userd

import (
	"context"
	"testing"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func TestUserD_ValidateFollowUser(t *testing.T) {
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

	val, err := ds.ValidateUser(context.Background(), &userdpb.ValidateUserRequest{
		Email:    "hi@work.com",
		Password: "Cool",
	})

	if val.User == nil {
		t.Fatalf("Expected the Cool user to be persisted")
	}
}
