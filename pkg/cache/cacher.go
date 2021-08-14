package cache

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/stoppage"
)

// Cacher defines everything that an adapter needs to provide to every cacher client.
type Cacher interface {
	Read(ctx context.Context, key string) ([]byte, error)
	Store(ctx context.Context, key string, value []byte) error
	GetKeys(ctx context.Context, pattern string) ([]string, error)
	Delete(ctx context.Context, key string) (int, error)
	DeleteMany(ctx context.Context, keys []string) (int, error)
	stoppage.Stopper
}
