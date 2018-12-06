package contentd

import (
	"github.com/theryecatcher/chirper/contentd/storage"
)

type Contentd struct {
	strg contentstorage.Storage
}

func New(cfg *Config) (*Contentd, error) {
	return &Contentd{
		strg: contentstorage.NewDummyStorage(),
	}, nil
}
