// Package userstorage defines an interface, which is used
// by the Content package, for the storage and retrieval of userdata
package userstorage

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

// Storage is an interface for persisting users on behalf of the User service
type Storage interface {
	// Commands
	NewUser(ctx context.Context, user *userdpb.NewUser) error
	FollowUser(ctx context.Context, UID string, FollowingUID string) error
	UnFollowUser(ctx context.Context, UID string, FollowingUID string) error
	// Queries
	GetUser(ctx context.Context, UID string) (*userdpb.User, error)
	ValidateUser(ctx context.Context, user *userdpb.CheckUser) (*userdpb.User, error)
	GetAllFollowers(ctx context.Context, UID string) ([]*userdpb.FollowerDetails, error)
	// Logger
	GetLoggerHandle() *log.Logger
}
