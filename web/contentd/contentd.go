package contentd

import (
	"github.com/theryecatcher/chirper/web/contentd/storage"
)

type Contentd struct {
	strg contentstorage.Storage
}

func New(cfg *Config) (*Contentd, error) {
	return &Contentd{
		strg: contentstorage.NewDummyStorage(),
	}, nil
}
