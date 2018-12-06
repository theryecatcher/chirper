package contentstorage

import (
	"context"
	"errors"
	"log"
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/theryecatcher/chirper/contentd/contentdpb"
)

// DummyStorage is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
// UID : TID , Tweet
type DummyStorage struct {
	tweets map[string]map[string]*contentdpb.Tweet
	mut    sync.Mutex
}

func NewDummyStorage() *DummyStorage {
	return &DummyStorage{
		tweets: make(map[string]map[string]*contentdpb.Tweet),
		mut:    sync.Mutex{},
	}
}

// NewTweet stores a tweet in-memory
func (ds *DummyStorage) NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error {
	done := make(chan string)

	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()

		if ds.tweets[tweet.PosterUID] == nil {
			ds.tweets[tweet.PosterUID] = make(map[string]*contentdpb.Tweet)
		}

		TID := uuid.New().String()
		for {
			if _, exists := ds.tweets[tweet.PosterUID][TID]; exists {
				TID = uuid.New().String()
			} else {
				break
			}
		}

		ds.tweets[tweet.PosterUID][TID] = &contentdpb.Tweet{
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

		delete(ds.tweets[tweet.PosterUID], TID)
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
		tweet, exists := ds.tweets[TID][TID]
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

// Custom Sorting Logic
type timeSortedTweets []*contentdpb.Tweet

func (t timeSortedTweets) Len() int {
	return len(t)
}

func (t timeSortedTweets) Less(i, j int) bool {
	return t[i].Timestamp > t[j].Timestamp
}

func (t timeSortedTweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// GetTweetsByUser returns tweets given the PosterUID's
func (ds *DummyStorage) GetTweetsByUser(ctx context.Context, UID []string) ([]*contentdpb.Tweet, error) {
	result := make(chan map[string]*contentdpb.Tweet)
	tweets := make(timeSortedTweets, 0)
	done := make(chan bool)
	tweetsNotFound := make(chan error)
	oops := make(chan error)

	log.Println(UID)

	for idx := range UID {
		go func(userID *string) {
			tweetMap, exists := ds.tweets[*userID]
			if !exists {
				tweetsNotFound <- errors.New("User " + *userID + "'s Content could not be found")
				return
			}
			result <- tweetMap
		}(&UID[idx])
	}

	go func() {
		count := 0
		for {
			select {
			case res := <-result:
				for _, tweet := range res {
					tweets = append(tweets, tweet)
				}
				count++
			case err := <-tweetsNotFound:
				log.Println(err)
				count++
			}
			if count == len(UID) {
				break
			}
		}
		done <- true
		return
	}()

	// Respect the context
	select {
	case <-done:
		sort.Sort(tweets)
		return tweets, nil
	case err := <-oops:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
