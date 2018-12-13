package contentd

import (
	"github.com/theryecatcher/chirper/contentd/storage"
)

// Contentd Contentd
type Contentd struct {
	strg contentstorage.Storage
}

// New new
func New(cfg *Config) (*Contentd, error) {
	return &Contentd{
		strg: contentstorage.NewContentStore(),
	}, nil
}
