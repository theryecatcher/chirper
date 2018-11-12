package contentd

import (
	"context"
	"log"
	"time"

	"github.com/distsys-project/web/contentd/contentdpb"
)

func (cnt *Contentd) NewTweet(ctx context.Context, req *contentdpb.NewTweetRequest) (*contentdpb.NewTweetResponse, error) {
	log.Println("Contentd: new tweet")

	t := &contentdpb.NewTweet{
		PosterUID: req.PosterUID,
		Content:   req.Content,
		Timestamp: time.Now().Unix(),
	}

	return &contentdpb.NewTweetResponse{}, cnt.strg.NewTweet(ctx, t)
}
