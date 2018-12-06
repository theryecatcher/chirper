package userstorage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/theryecatcher/chirper/web/userd/userdpb"
)

// DummyUserStorage is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
type DummyUserStorage struct {
	users map[string]*userdpb.User
	mut   sync.Mutex
}

// NewDummyUserStorage Default Storage Interfacce for In Memory Implementation
func NewDummyUserStorage() *DummyUserStorage {
	return &DummyUserStorage{
		users: make(map[string]*userdpb.User),
		mut:   sync.Mutex{},
	}
}

// NewUser stores a user in-memory
func (ds *DummyUserStorage) NewUser(ctx context.Context, user *userdpb.NewUser) error {
	done := make(chan string)

	if ds.users == nil {
		fmt.Println("Nil Set")
	}

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
		var init []string
		ds.users[UID] = &userdpb.User{
			UID:          UID,
			Name:         user.Name,
			Email:        user.Email,
			Password:     user.Password,
			FollowingUID: init,
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

// GetUser returns a USer Details given its UID
func (ds *DummyUserStorage) GetUser(ctx context.Context, UID string) (*userdpb.User, error) {

	fmt.Println(ds.users)

	result := make(chan *userdpb.User)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		// ds.mut.Lock()
		// defer ds.mut.Unlock()
		user, exists := ds.users[UID]
		if !exists {
			oops <- errors.New("User not found")
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

// ValidateUser returns a User is present or not given its Email
func (ds *DummyUserStorage) ValidateUser(ctx context.Context, chkUser *userdpb.CheckUser) (*userdpb.User, error) {

	fmt.Println(ds.users)
	fmt.Println(chkUser.Email)

	result := make(chan *userdpb.User)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		// user, exists := ds.users
		for k, v := range ds.users {
			if v.Email == chkUser.Email {
				if v.Password == chkUser.Password {
					result <- ds.users[k]
					return
				}
				result <- nil
				return
			}
		}
		oops <- errors.New("User not found")
		return
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

// FollowUser follows user
func (ds *DummyUserStorage) FollowUser(ctx context.Context, UID string, FollowingUID string) error {

	//fmt.Println(ds.users)

	result := make(chan bool)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		user, exists := ds.users[UID]
		if !exists {
			oops <- errors.New("User not found")
			return
		}
		user.FollowingUID = append(user.FollowingUID, FollowingUID)
		result <- true
	}()

	// Respect the context
	select {
	case <-result:
		return nil
	case err := <-oops:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func contains(s []string, item string) bool {
	for _, a := range s {
		if a == item {
			return true
		}
	}
	return false
}

// GetAllFollowers Method to populate all followers
func (ds *DummyUserStorage) GetAllFollowers(ctx context.Context, UID string) ([]*userdpb.FollowerDetails, error) {

	fmt.Println("in test get all followers")
	followers := make([]*userdpb.FollowerDetails, 0)

	result := make(chan bool)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		user, exists := ds.users[UID]
		if !exists {
			oops <- errors.New("User not found")
		} else {
			for k, v := range ds.users {
				if v.UID != user.UID {
					followers = append(followers, &userdpb.FollowerDetails{
						Name:     v.Name,
						UID:      k,
						Followed: contains(user.FollowingUID, v.UID),
					})
				}
			}
		}
		result <- true
	}()

	// Respect the context
	select {
	case <-result:
		return followers, nil
	case err := <-oops:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

//UnFollowUser Unfollows a user
func (ds *DummyUserStorage) UnFollowUser(ctx context.Context, UID string, FollowedUID string) error {

	result := make(chan bool)
	oops := make(chan error)

	log.Println(FollowedUID)
	log.Println(UID)

	// Go fetch the user
	go func() {
		ds.mut.Lock()
		defer ds.mut.Unlock()
		user, exists := ds.users[UID]
		if !exists {
			oops <- errors.New("User not found")
			return
		}
		for k, v := range user.FollowingUID {
			if v == FollowedUID {
				user.FollowingUID = append(user.FollowingUID[:k], user.FollowingUID[k+1:]...)
				break
			}
		}
		result <- true
	}()

	// Respect the context
	select {
	case <-result:
		return nil
	case err := <-oops:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
