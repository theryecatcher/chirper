package contentd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/web/contentd/contentdpb"
)

func (cnt *Contentd) GetTweetsByUser(ctx context.Context, req *contentdpb.GetTweetsByUserRequest) (*contentdpb.GetTweetsByUserResponse, error) {
	log.Println("Contentd: get tweets by user")
	t, err := cnt.strg.GetTweetsByUser(ctx, req.UID)

	return &contentdpb.GetTweetsByUserResponse{
		Tweets: t,
	}, err
}
