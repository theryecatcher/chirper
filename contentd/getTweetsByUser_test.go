package contentd

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/theryecatcher/chirper/contentd/contentdpb"
	"github.com/theryecatcher/chirper/contentd/storage"
)

func TestContentd_GetTweetByUser(t *testing.T) {
	c := &Contentd{
		strg: contentstorage.NewContentStore(),
	}

	uid := uuid.New().String()
	tweet := "hello"

	c.NewTweet(context.Background(), &contentdpb.NewTweetRequest{
		PosterUID: uid,
		Content:   tweet,
	})

	var uids []string
	uids = append(uids, uid)

	res, err := c.GetTweetsByUser(context.Background(), &contentdpb.GetTweetsByUserRequest{
		UID: uids,
	})

	if err != nil {
		if res.Tweets[0].Content != tweet {
			t.Fatal("Expected for entries one tweet")
		}
	}
}
