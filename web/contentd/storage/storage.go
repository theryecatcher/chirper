// Package storage defines an interface, which is used
// by the Content package, for the storage and retrieval of tweets
package contentstorage

import (
	"context"

	"github.com/theryecatcher/chirper/web/contentd/contentdpb"
)

// Storage is an interface for persisting tweets on behalf of the Content service
type Storage interface {
	// Commands
	NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error

	// Queries
	GetTweet(ctx context.Context, TID string) (*contentdpb.Tweet, error)
	GetTweetsByUser(ctx context.Context, UID []string) ([]*contentdpb.Tweet, error)
}
