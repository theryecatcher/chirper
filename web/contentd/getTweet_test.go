package contentd

import (
	"context"
	"testing"

	"github.com/adamsanghera/example-proj/web/contentd/storage"
)

func TestContentd_GetTweet(t *testing.T) {
	c := Contentd{
		strg: &storage.DummyStorage{},
	}

	c.GetTweet(context.Background())

}
