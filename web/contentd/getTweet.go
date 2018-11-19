package contentd

import (
	"context"
	"log"

	"github.com/theryecatcher/chirper/web/contentd/contentdpb"
)

func (cnt *Contentd) GetTweet(ctx context.Context, req *contentdpb.GetTweetRequest) (*contentdpb.GetTweetResponse, error) {
	log.Println("Contentd: get tweet")
	t, err := cnt.strg.GetTweet(ctx, req.TID)

	return &contentdpb.GetTweetResponse{
		Tweet: t,
	}, err
}
