// Package storage defines an interface, which is used
// by the Content package, for the storage and retrieval of tweets
package storage

import (
	"context"

	"github.com/distsys-project/web/userd/userdpb"
)

// Storage is an interface for persisting tweets on behalf of the Content service
type Storage interface {
	// Commands
	NewUser(ctx context.Context, user *userdpb.NewUser) error

	// Queries
	GetUser(ctx context.Context, TID string) (*userdpb.User, error)
}
