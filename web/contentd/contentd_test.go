package contentd

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/theryecatcher/chirper/web/contentd/contentdpb"
)

func TestDummyStorage_NewTweetParallel(t *testing.T) {
	cntCfg := &Config{}

	contentDb, err := New(cntCfg)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1000)
	for _, usr := range []string{"abc"} {
		for idx := 0; idx < 1000; idx++ {
			go func() {
				_, err := contentDb.NewTweet(context.Background(), &contentdpb.NewTweetRequest{
					PosterUID: usr,
					Content:   "hi",
				})

				if err != nil {
					t.Fatalf("Unexpected error {%+v}", err)
				}

				wg.Done()
			}()
		}
	}

	wg.Wait()

	tweets, err := contentDb.GetTweetsByUser(context.Background(), &contentdpb.GetTweetsByUserRequest{
		UID: []string{"abc"},
	})

	if len(tweets.Tweets) != 1000 {
		fmt.Println(len(tweets.Tweets))
		t.Fatal("Expected 1000 entries for abc user")
	}
}

func TestDummyStorage_NewTweetCancelledContext(t *testing.T) {
	cntCfg := &Config{}

	contentDb, err := New(cntCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, qErr := contentDb.NewTweet(ctx, &contentdpb.NewTweetRequest{
		PosterUID: "abc",
		Content:   "hi",
	})

	if qErr != context.Canceled {
		t.Fatalf("Unexpected error {%+v} instead of {%v}", err, context.Canceled)
	}

	tweets, err := contentDb.GetTweetsByUser(context.Background(), &contentdpb.GetTweetsByUserRequest{
		UID: []string{"abc"},
	})

	if len(tweets.Tweets) == 1 {
		t.Fatal("Expected tweet to not be written, because of cancelled context")
	}
}
