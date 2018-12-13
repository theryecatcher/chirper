package contentstorage

import (
	"context"
	"errors"
	"log"
	"os"
	"sort"

	"google.golang.org/grpc/codes"

	"github.com/google/uuid"
	"github.com/theryecatcher/chirper/contentd/contentdpb"
	"github.com/theryecatcher/chirper/raftd/raftdpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ContentStore is an in-memory implementation
// of the Storage interface, which is used for unit-testing
// functions that depend on a "real" implementation of Storage.
// UID : TID , Tweet
type ContentStore struct {
	leader    raftdpb.RaftdClient
	follower1 raftdpb.RaftdClient
	follower2 raftdpb.RaftdClient

	logger *log.Logger
}

// NewContentStore NewContentStore
func NewContentStore() *ContentStore {

	loclLoggger := log.New(os.Stderr, "[contentdWrapper] ", log.LstdFlags)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	leaderConn, err := grpc.Dial("localhost:45000", opts...)
	if err != nil {
		loclLoggger.Fatalf("failure while dialing: %v", err)
	}
	// defer cntdConn.Close()
	// Need to figure out adding this I keep getting error
	// rpc error: code = Canceled desc = grpc: the client connection is closing

	follower1Conn, err := grpc.Dial("localhost:45001", opts...)
	if err != nil {
		loclLoggger.Fatalf("failure while dialing: %v", err)
	}
	follower2Conn, err := grpc.Dial("localhost:45002", opts...)
	if err != nil {
		loclLoggger.Fatalf("failure while dialing: %v", err)
	}

	return &ContentStore{
		leader:    raftdpb.NewRaftdClient(leaderConn),
		follower1: raftdpb.NewRaftdClient(follower1Conn),
		follower2: raftdpb.NewRaftdClient(follower2Conn),

		logger: loclLoggger,
	}
}

// GetLoggerHandle return the logger handle
func (ds *ContentStore) GetLoggerHandle() *log.Logger {
	return ds.logger
}

func (ds *ContentStore) raftget(ctx context.Context, key string) (string, error) {

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

func (ds *ContentStore) raftdel(ctx context.Context, key string) error {

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

func (ds *ContentStore) raftset(ctx context.Context, k string, v string) error {

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

// NewTweet stores a tweet
func (ds *ContentStore) NewTweet(ctx context.Context, tweet *contentdpb.NewTweet) error {
	done := make(chan string)
	oops := make(chan error)

	go func() {
		var tweets TweetMap

		data, err := ds.raftget(ctx, "cnt:"+tweet.PosterUID)

		if err != nil {
			if err.Error() == "rpc error: code = Unknown desc = Value not found" {
				ds.logger.Println(err)
			} else {
				oops <- err
				return
			}
		}

		if data != "" {
			tweets = FromGOB64(data)
		} else {
			tweets = make(TweetMap)
		}

		TID := uuid.New().String()
		for {
			if _, exists := tweets[TID]; exists {
				TID = uuid.New().String()
			} else {
				break
			}
		}

		tweets[TID] = &contentdpb.Tweet{
			TID:       TID,
			Timestamp: tweet.Timestamp,
			Content:   tweet.Content,
			PosterUID: tweet.PosterUID,
		}

		err = ds.raftset(ctx, "cnt:"+tweet.PosterUID, ToGOB64(tweets))
		if err != nil {
			oops <- err
			return
		}

		done <- TID
	}()

	// Respect the context
	select {
	case <-done:
		return nil
	case err := <-oops:
		return err
	case <-ctx.Done():
		// roll back TID if it was created
		// TID := <-done
		// ds.raftdel(ctx, tweet.PosterUID)
		return ctx.Err()
	}
}

// GetTweet returns a tweet given its TID
// func (ds *ContentStore) GetTweet(ctx context.Context, TID string) (*contentdpb.Tweet, error) {
// 	result := make(chan *contentdpb.Tweet)
// 	oops := make(chan error)

// 	// Go fetch the tweet
// 	go func() {

// 		data, err := ds.raftget(ctx, tweet.PosterUID)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if data == nil {
// 			tweets = make(map[string]*contentdpb.Tweet)
// 		} else {
// 			var bD bytes.Buffer
// 			bD.Write(*data)
// 			d := gob.NewDecoder(&bD)
// 			if err = d.Decode(&tweets); err != nil {
// 				panic(err)
// 			}

// 			fmt.Println("Decoded Tweets ", tweets)
// 		}

// 		tweet, e := ds.tweets[TID][TID]
// 		if !exists {
// 			oops <- errors.New("Tweet could not be found")
// 			return
// 		}
// 		result <- tweet
// 	}()

// 	// Respect the context
// 	select {
// 	case res := <-result:
// 		return res, nil
// 	case err := <-oops:
// 		return nil, err
// 	case <-ctx.Done():
// 		return nil, ctx.Err()
// 	}
// }

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
func (ds *ContentStore) GetTweetsByUser(ctx context.Context, UID []string) ([]*contentdpb.Tweet, error) {
	result := make(chan TweetMap)
	tweets := make(timeSortedTweets, 0)
	done := make(chan bool)
	tweetsNotFound := make(chan error)
	oops := make(chan error)

	for idx := range UID {
		go func(userID *string) {

			ds.logger.Println(userID)
			var gettweets TweetMap

			data, err := ds.raftget(ctx, "cnt:"+*userID)

			if err != nil {
				if err.Error() == "rpc error: code = Unknown desc = Value not found" {
					ds.logger.Println(err)
				} else {
					panic(err)
				}
			}

			if data == "" {
				tweetsNotFound <- errors.New("User " + *userID + "'s Content could not be found")
				return
			}

			gettweets = FromGOB64(data)
			result <- gettweets
			return

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
				ds.logger.Println(err)
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
