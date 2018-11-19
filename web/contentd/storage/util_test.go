package contentstorage

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/theryecatcher/chirper/web/contentd/contentdpb"
)

func TestDummyStorage_NewTweetParallel(t *testing.T) {
	ds := NewDummyStorage()

	wg := sync.WaitGroup{}

	wg.Add(1000)
	for _, usr := range []string{"abc"} {
		for idx := 0; idx < 1000; idx++ {
			go func() {
				err := ds.NewTweet(context.Background(), &contentdpb.NewTweet{
					PosterUID: usr,
					Content:   "hi",
					Timestamp: time.Now().Unix(),
				})
				if err != nil {
					t.Fatalf("Unexpected error {%+v}", err)
				}

				wg.Done()
			}()
		}
	}

	wg.Wait()

	if len(ds.tweets["abc"]) != 1000 {
		fmt.Println(len(ds.tweets["abc"]))
		t.Fatal("Expected 1000 entries for abc user")
	}
	// if len(ds.tweets["hij"]) != 1000 {
	// 	t.Fatal("Expected 1000 entries for hij user")
	// }
	if len(ds.tweets) != 1 {
		t.Fatal("Expected entries for two users")
	}
}

func TestDummyStorage_NewTweetCancelledContext(t *testing.T) {
	ds := NewDummyStorage()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ds.NewTweet(ctx, &contentdpb.NewTweet{
		PosterUID: "abc",
		Content:   "hi",
		Timestamp: time.Now().Unix(),
	})
	if err != context.Canceled {
		t.Fatalf("Unexpected error {%+v} instead of {%v}", err, context.Canceled)
	}

	ds.mut.Lock()
	defer ds.mut.Unlock()

	fmt.Println(len(ds.tweets["a"]))
	if len(ds.tweets["a"]) == 1 {
		t.Fatal("Expected tweet to not be written, because of cancelled context")
	}
}

func TestDummyStorage_GetTweetsbyUser(t *testing.T) {
	ds := NewDummyStorage()

	wg := sync.WaitGroup{}

	for idx := 0; idx < 1000; idx++ {
		wg.Add(1)
		go func() {
			err := ds.NewTweet(context.Background(), &contentdpb.NewTweet{
				PosterUID: "abc",
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

	tweets, err := ds.GetTweetsByUser(context.Background(), []string{"abc"})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}

	if len(tweets) != 1000 {
		t.Fatal("Expected 1000 entries")
	}

	tweets, err = ds.GetTweetsByUser(context.Background(), []string{"hij"})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}

	if len(tweets) != 0 {
		t.Fatal("Expected no entries")
	}
}

// A test for cehking if the returned array is sorted or not
