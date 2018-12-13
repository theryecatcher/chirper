package contentd

import (
	"context"

	"github.com/theryecatcher/chirper/contentd/contentdpb"
)

func (cnt *Contentd) GetTweetsByUser(ctx context.Context, req *contentdpb.GetTweetsByUserRequest) (*contentdpb.GetTweetsByUserResponse, error) {
	cnt.strg.GetLoggerHandle().Println("Contentd: get tweets by user")
	t, err := cnt.strg.GetTweetsByUser(ctx, req.UID)

	return &contentdpb.GetTweetsByUserResponse{
		Tweets: t,
	}, err
}
