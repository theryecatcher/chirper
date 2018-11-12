package contentd

import (
	"github.com/distsys-project/web/contentd/storage"
)

type Contentd struct {
	strg storage.Storage
}

func New(cfg *Config) (*Contentd, error) {
	return &Contentd{
		strg: &storage.DummyStorage{},
	}, nil
}
