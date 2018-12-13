package contentd

import (
	"context"
	"time"

	"github.com/theryecatcher/chirper/contentd/contentdpb"
)

// NewTweet exposes the new tweet method
func (cnt *Contentd) NewTweet(ctx context.Context, req *contentdpb.NewTweetRequest) (*contentdpb.NewTweetResponse, error) {
	cnt.strg.GetLoggerHandle().Println("Contentd: new tweet")

	t := &contentdpb.NewTweet{
		PosterUID: req.PosterUID,
		Content:   req.Content,
		Timestamp: time.Now().Unix(),
	}

	return &contentdpb.NewTweetResponse{}, cnt.strg.NewTweet(ctx, t)
}
