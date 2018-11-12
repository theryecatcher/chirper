package storage

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/adamsanghera/example-proj/web/contentd/contentdpb"
)

func TestDummyStorage_NewTweetParallel(t *testing.T) {
	ds := NewDummyStorage()

	wg := sync.WaitGroup{}

	for idx := 0; idx < 1000; idx++ {
		wg.Add(1)

		go func() {
			err := ds.NewTweet(context.Background(), &contentdpb.NewTweet{
				PosterUID: "a",
				Content:   "hi",
				Timestamp: time.Now().Unix(),
			})
			if err != nil {
				t.Fatalf("Unexpected error {%+v}", err)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	if len(ds.tweets) != 1000 {
		t.Fatal("Expected 1000 entries")
	}
}

func TestDummyStorage_NewTweetCancelledContext(t *testing.T) {
	ds := NewDummyStorage()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ds.NewTweet(ctx, &contentdpb.NewTweet{
		PosterUID: "a",
		Content:   "hi",
		Timestamp: time.Now().Unix(),
	})
	if err != context.Canceled {
		t.Fatalf("Unexpected error {%+v} instead of {%v}", err, context.Canceled)
	}

	ds.mut.Lock()
	defer ds.mut.Unlock()

	if len(ds.tweets) == 1 {
		t.Fatal("Expected tweet to not be written, because of cancelled context")
	}
}
