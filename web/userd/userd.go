package userd

import (
	"github.com/theryecatcher/chirper/web/userd/storage"
)

type Userd struct {
	usrStrg userstorage.Storage
}

func New(cfg *Config) (*Userd, error) {
	return &Userd{
		usrStrg: userstorage.NewDummyUserStorage(),
	}, nil
}
