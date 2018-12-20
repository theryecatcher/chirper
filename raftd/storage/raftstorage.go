package raftstorage

import (
	"context"
	"log"
)

// Storage is an interface for persisting users on behalf of the User service
type Storage interface {
	// Methods
	RaftSetKeyValue(ctx context.Context, key string, value string) error
	RaftJoin(ctx context.Context, nodeID, addr string) error
	RaftGet(ctx context.Context, key string) (string, error)
	RaftDelete(ctx context.Context, key string) error
	RaftGetAllUsrKeys(ctx context.Context) ([]string, error)
	// Internal Methods
	RaftCreate(master bool, nodeID string) error
	// Logger
	GetLoggerHandle() *log.Logger
}
