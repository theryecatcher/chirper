package userstorage

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/theryecatcher/chirper/raftd/raftdpb"
	"github.com/theryecatcher/chirper/userd/userdpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserStoreWrapper is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
type UserStoreWrapper struct {
	leader    raftdpb.RaftdClient
	follower1 raftdpb.RaftdClient
	follower2 raftdpb.RaftdClient

	logger *log.Logger
}

// NewUserStoreWrapper Default Storage Interfacce for In Memory Implementation
func NewUserStoreWrapper() *UserStoreWrapper {

	loclLogger := log.New(os.Stderr, "[userdWrapper] ", log.LstdFlags)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	leaderConn, err := grpc.Dial("localhost:45000", opts...)
	if err != nil {
		loclLogger.Fatalf("failure while dialing: %v", err)
	}
	// defer cntdConn.Close()
	// Need to figure out adding this I keep getting error
	// rpc error: code = Canceled desc = grpc: the client connection is closing

	follower1Conn, err := grpc.Dial("localhost:45001", opts...)
	if err != nil {
		loclLogger.Fatalf("failure while dialing: %v", err)
	}
	follower2Conn, err := grpc.Dial("localhost:45002", opts...)
	if err != nil {
		loclLogger.Fatalf("failure while dialing: %v", err)
	}

	return &UserStoreWrapper{
		leader:    raftdpb.NewRaftdClient(leaderConn),
		follower1: raftdpb.NewRaftdClient(follower1Conn),
		follower2: raftdpb.NewRaftdClient(follower2Conn),

		logger: loclLogger,
	}
}

// GetLoggerHandle return the logger handle
func (ds *UserStoreWrapper) GetLoggerHandle() *log.Logger {
	return ds.logger
}

func (ds *UserStoreWrapper) raftget(ctx context.Context, key string) (string, error) {

	value, err := ds.leader.RaftGet(ctx, &raftdpb.RaftGetRequest{
		Key: key,
	})
	if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
		value, err = ds.follower1.RaftGet(ctx, &raftdpb.RaftGetRequest{
			Key: key,
		})
		if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
			value, err = ds.follower2.RaftGet(ctx, &raftdpb.RaftGetRequest{
				Key: key,
			})
		}
	}

	if err == nil {
		return value.Value, err
	}

	return "", err
}

func (ds *UserStoreWrapper) raftgetallusrkeys(ctx context.Context) ([]string, error) {

	value, err := ds.leader.RaftGetAllUsrKeys(ctx, &raftdpb.RaftGetAllUsrKeysRequest{})
	if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
		value, err = ds.follower1.RaftGetAllUsrKeys(ctx, &raftdpb.RaftGetAllUsrKeysRequest{})
		if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
			value, err = ds.follower2.RaftGetAllUsrKeys(ctx, &raftdpb.RaftGetAllUsrKeysRequest{})
		}
	}

	if err == nil {
		return value.Value, err
	}

	return nil, err
}

func (ds *UserStoreWrapper) raftdel(ctx context.Context, key string) error {

	var err error

	_, err = ds.leader.RaftDelete(ctx, &raftdpb.RaftDeleteRequest{
		Key: key,
	})
	if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
		_, err = ds.follower1.RaftDelete(ctx, &raftdpb.RaftDeleteRequest{
			Key: key,
		})
		if errStatus, _ := status.FromError(err); codes.Unavailable == errStatus.Code() {
			_, err = ds.follower2.RaftDelete(ctx, &raftdpb.RaftDeleteRequest{
				Key: key,
			})
		}
	}

	return err
}

func (ds *UserStoreWrapper) raftset(ctx context.Context, k string, v string) error {

	var errldr, errf1, errf2 error

	_, errldr = ds.leader.RaftSetKeyValue(ctx, &raftdpb.RaftSetKeyValueRequest{
		Key:   k,
		Value: v,
	})
	if errldr != nil {
		ds.logger.Println("1st Node Errored")
		if errStatus, _ := status.FromError(errldr); codes.Unavailable == errStatus.Code() || errldr.Error() == "rpc error: code = Unknown desc = Not Leader" {
			_, errf1 = ds.follower1.RaftSetKeyValue(ctx, &raftdpb.RaftSetKeyValueRequest{
				Key:   k,
				Value: v,
			})
		}
		if errf1 != nil {
			ds.logger.Println("2nd Node Errored")
			if errStatus, _ := status.FromError(errf1); codes.Unavailable == errStatus.Code() || errf1.Error() == "rpc error: code = Unknown desc = Not Leader" {
				_, errf2 = ds.follower2.RaftSetKeyValue(ctx, &raftdpb.RaftSetKeyValueRequest{
					Key:   k,
					Value: v,
				})
				if errf2 != nil {
					return errf2
				}
				ds.leader, ds.follower2 = ds.follower2, ds.leader
				return errf2
			}
		}
		ds.leader, ds.follower1 = ds.follower1, ds.leader
		return errf1
	}

	return errldr
}

// NewUser stores a user in-memory
func (ds *UserStoreWrapper) NewUser(ctx context.Context, user *userdpb.NewUser) error {
	done := make(chan string)
	oops := make(chan error)

	// if ds.users == nil {
	// 	fmt.Println("Nil Set")
	// }

	go func() {

		userkeys, err := ds.raftgetallusrkeys(ctx)
		if err != nil {
			oops <- err
			return
		}
		for _, k := range userkeys {
			usr, err := ds.GetUser(ctx, k)

			if err != nil {
				oops <- err
				return
			}
			if usr.Email == user.Email {
				oops <- errors.New("User already exists")
				return
			}
		}

		UID := "usr:" + uuid.New().String()
		alluserkeys, err := ds.raftgetallusrkeys(ctx)
		for {
			if alluserkeys != nil {
				if contains(alluserkeys, UID) {
					UID = uuid.New().String()
				} else {
					break
				}
			} else {
				break
			}
		}
		var init []string
		newusr := &userdpb.User{
			UID:          UID,
			Name:         user.Name,
			Email:        user.Email,
			Password:     user.Password,
			FollowingUID: init,
		}

		err = ds.raftset(ctx, UID, ToGOB64(newusr))
		if err != nil {
			oops <- err
		}
		done <- UID
	}()

	// Respect the context
	select {
	case <-done:
		return nil
	case err := <-oops:
		return err
	case <-ctx.Done():
		// roll back UID if it was created
		// UID := <-done
		// delete(ds.users, UID)
		return ctx.Err()
	}
}

// GetUser returns a USer Details given its UID
func (ds *UserStoreWrapper) GetUser(ctx context.Context, UID string) (*userdpb.User, error) {

	result := make(chan *userdpb.User)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		var data string
		var err error

		ds.logger.Println(UID)

		if strings.HasPrefix(UID, "usr") {
			data, err = ds.raftget(ctx, UID)
		} else {
			data, err = ds.raftget(ctx, "usr:"+UID)
		}

		if err != nil {
			if err.Error() == "rpc error: code = Unknown desc = Value not found" {
				ds.logger.Println(err)
			} else {
				ds.logger.Panicln(err)
			}
		}
		if data == "" {
			oops <- errors.New("User not found")
			return
		}
		// return user
		result <- FromGOB64(data)
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
func (ds *UserStoreWrapper) ValidateUser(ctx context.Context, chkUser *userdpb.CheckUser) (*userdpb.User, error) {

	// fmt.Println(ds.users)
	// fmt.Println(chkUser.Email)

	result := make(chan *userdpb.User)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		userkeys, err := ds.raftgetallusrkeys(ctx)
		if err != nil {
			oops <- err
			return
		}
		for _, k := range userkeys {
			usr, err := ds.GetUser(ctx, k)

			if err != nil {
				oops <- err
				return
			}
			if usr.Email == chkUser.Email {
				if usr.Password == chkUser.Password {
					result <- usr
					return
				}
				oops <- errors.New("Incorrect password")
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
func (ds *UserStoreWrapper) FollowUser(ctx context.Context, UID string, FollowingUID string) error {

	//fmt.Println(ds.users)

	result := make(chan bool)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		user, err := ds.GetUser(ctx, UID)
		if err != nil {
			oops <- err
			return
		}
		user.FollowingUID = append(user.FollowingUID, FollowingUID)
		// Update DB
		err = ds.raftset(ctx, UID, ToGOB64(user))
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
func (ds *UserStoreWrapper) GetAllFollowers(ctx context.Context, UID string) ([]*userdpb.FollowerDetails, error) {

	followers := make([]*userdpb.FollowerDetails, 0)

	result := make(chan bool)
	oops := make(chan error)

	// Go fetch the user
	go func() {
		currentuser, err := ds.GetUser(ctx, UID)
		if err != nil {
			oops <- err
			return
		}
		userkeys, err := ds.raftgetallusrkeys(ctx)
		for _, k := range userkeys {
			user, err := ds.GetUser(ctx, k)
			if err != nil {
				oops <- err
				return
			}
			if user.UID != currentuser.UID {
				followers = append(followers, &userdpb.FollowerDetails{
					Name:     user.Name,
					UID:      k,
					Followed: contains(currentuser.FollowingUID, user.UID),
				})
			}
		}
		ds.logger.Println(followers)
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
func (ds *UserStoreWrapper) UnFollowUser(ctx context.Context, UID string, FollowedUID string) error {

	result := make(chan bool)
	oops := make(chan error)

	ds.logger.Println(FollowedUID)
	ds.logger.Println(UID)

	// Go fetch the user
	go func() {
		user, err := ds.GetUser(ctx, UID)
		if err != nil {
			oops <- err
			return
		}
		ds.logger.Println(user)
		for k, v := range user.FollowingUID {
			if v == FollowedUID {
				user.FollowingUID = append(user.FollowingUID[:k], user.FollowingUID[k+1:]...)
				break
			}
		}
		// Update DB
		err = ds.raftset(ctx, UID, ToGOB64(user))
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
