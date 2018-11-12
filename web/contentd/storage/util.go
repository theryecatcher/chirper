package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/distsys-project/web/contentd/contentdpb"
)

// DummyStorage is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
type DummyStorage struct {
	tweets map[string]*contentdpb.Tweet

	mut sync.Mutex
}

func NewDummyStorage() *DummyStorage {
	return &DummyStorage{
		tweets: make(map[string]*contentdpb.Tweet),
		mut:    sync.Mutex{},
	}
}

// NewTweet stores a tweet in-memory
func (ds *DummyStorage) NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error {
	done := make(chan string)

	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()

		TID := uuid.New().String()
		for {
			if _, exists := ds.tweets[TID]; exists {
				TID = uuid.New().String()
			} else {
				break
			}
		}

		ds.tweets[TID] = &contentdpb.Tweet{
			TID:       TID,
			Timestamp: tweet.Timestamp,
			Content:   tweet.Content,
			PosterUID: tweet.PosterUID,
		}
		done <- TID
	}()

	// Respect the context
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		// roll back TID if it was created
		TID := <-done

		ds.mut.Lock()
		defer ds.mut.Unlock()

		delete(ds.tweets, TID)
		return ctx.Err()
	}
}

// GetTweet returns a tweet given its TID
func (ds *DummyStorage) GetTweet(ctx context.Context, TID string) (*contentdpb.Tweet, error) {
	result := make(chan *contentdpb.Tweet)
	oops := make(chan error)

	// Go fetch the tweet
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		tweet, exists := ds.tweets[TID]
		if !exists {
			oops <- errors.New("Tweet could not be found")
			return
		}

		result <- tweet
	}()

	// Respect the context
	select {
	case res := <-result:
		return res, nil
	case err := <-oops:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
