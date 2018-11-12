package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/distsys-project/web/userd/userdpb"
	"github.com/google/uuid"
)

// DummyStorage is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
type DummyStorage struct {
	users map[string]*userdpb.User
	mut   sync.Mutex
}

func NewDummyStorage() *DummyStorage {
	return &DummyStorage{
		users: make(map[string]*userdpb.User),
		mut:   sync.Mutex{},
	}
}

// NewUser stores a user in-memory
func (ds *DummyStorage) NewUser(ctx context.Context, user *userdpb.NewUser) error {
	done := make(chan string)

	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()

		UID := uuid.New().String()
		for {
			if _, exists := ds.users[UID]; exists {
				UID = uuid.New().String()
			} else {
				break
			}
		}

		ds.users[UID] = &userdpb.User{
			UID:      UID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		done <- UID
	}()

	// Respect the context
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		// roll back UID if it was created
		UID := <-done

		ds.mut.Lock()
		defer ds.mut.Unlock()

		delete(ds.users, UID)
		return ctx.Err()
	}
}

// GetUser returns a tweet given its UID
func (ds *DummyStorage) GetUser(ctx context.Context, UID string) (*userdpb.User, error) {
	result := make(chan *userdpb.User)
	oops := make(chan error)

	// Go fetch the tweet
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		user, exists := ds.users[UID]
		if !exists {
			oops <- errors.New("Tweet could not be found")
			return
		}

		result <- user
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
