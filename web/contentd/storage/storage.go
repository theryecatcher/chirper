// Package storage defines an interface, which is used
// by the Content package, for the storage and retrieval of tweets
package storage

import (
	"context"

	"github.com/distsys-project/web/contentd/contentdpb"
)

// Storage is an interface for persisting tweets on behalf of the Content service
type Storage interface {
	// Commands
	NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error

	// Queries
	GetTweet(ctx context.Context, TID string) (*contentdpb.Tweet, error)
}
