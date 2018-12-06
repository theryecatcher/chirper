package userstorage

import (
	"context"
	"sync"
	"testing"

	"github.com/theryecatcher/chirper/web/userd/userdpb"
)

func TestDummyStorage_NewUserParallel(t *testing.T) {
	ds := NewDummyUserStorage()

	wg := sync.WaitGroup{}

	wg.Add(1000)
	for idx := 0; idx < 1000; idx++ {
		go func() {
			err := ds.NewUser(context.Background(), &userdpb.NewUser{
				Name:     "Cool",
				Email:    "hi@work.com",
				Password: "Cool",
			})
			if err != nil {
				t.Fatalf("Unexpected error {%+v}", err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	var count int
	for range ds.users {
		count = count + 1
		if count != 1000 {
			t.Fatal("Expected 1000 entries Mr. Cool")
		}
	}
}

func TestDummyStorage_NewUserCancelledContext(t *testing.T) {
	ds := NewDummyUserStorage()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ds.NewUser(ctx, &userdpb.NewUser{
		Name:     "Cool",
		Email:    "hi@work.com",
		Password: "Cool",
	})
	if err != context.Canceled {
		t.Fatalf("Unexpected error {%+v} instead of {%v}", err, context.Canceled)
	}

	ds.mut.Lock()
	defer ds.mut.Unlock()

	var count int
	for range ds.users {
		count = count + 1
	}
	if count == 1 {
		t.Fatal("Mr. Cool cannot be added because of cancelled context")
	}
}

func TestDummyStorage_GetUser(t *testing.T) {
	ds := NewDummyUserStorage()

	wg := sync.WaitGroup{}
	var count int

	wg.Add(1000)
	for k := range ds.users {
		count = count + 1
		go func(UID *string) {
			_, err := ds.GetUser(context.Background(), *UID)
			if err != nil {
				t.Fatalf("Unexpected error {%+v}", err)
			}
			wg.Done()
		}(&k)
	}

	wg.Wait()

	if count != 1000 {
		t.Fatal("Expected 1000 entries")
	}
}

func TestDummyStorage_FollowUserParallel(t *testing.T) {
	ds := NewDummyUserStorage()

	wg := sync.WaitGroup{}
	var uid string
	for k := range ds.users {
		uid = k
		break
	}

	wg.Add(1000)
	for k := range ds.users {
		go func(FollowingUID string) {
			err := ds.FollowUser(context.Background(), uid, FollowingUID)
			//Password: "Cool",

			if err != nil {
				t.Fatalf("Unexpected error {%+v}", err)
			}

			wg.Done()
		}(k)
	}

	wg.Wait()
	if len(ds.users[uid].FollowingUID) != 1000 {
		t.Fatal("Expected 1000 entries")
	}
}

func TestDummyStorage_UnFollowUserParallel(t *testing.T) {
	ds := NewDummyUserStorage()

	wg := sync.WaitGroup{}
	var uid string
	for k := range ds.users {
		uid = k
		break
	}

	wg.Add(1000)
	for k := range ds.users {
		go func(FollowedUID string) {
			err := ds.UnFollowUser(context.Background(), uid, FollowedUID)

			if err != nil {
				t.Fatalf("Unexpected error {%+v}", err)
			}

			wg.Done()
		}(k)
	}

	wg.Wait()
	if len(ds.users[uid].FollowingUID) != 0 {
		t.Fatal("Expected 0 entries")
	}
}
