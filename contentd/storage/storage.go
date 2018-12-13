// Package contentstorage defines an interface, which is used
// by the Content package, for the storage and retrieval of tweets
package contentstorage

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/contentd/contentdpb"
)

// Storage is an interface for persisting tweets on behalf of the Content service
type Storage interface {
	// Commands
	NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error

	// Queries
	GetTweetsByUser(ctx context.Context, UID []string) ([]*contentdpb.Tweet, error)

	// Logger
	GetLoggerHandle() *log.Logger
}
