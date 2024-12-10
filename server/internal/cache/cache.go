package cache

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/cache/redis"
)

type Cache interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string, target any) error

	Close() error
}

const (
	Redis = iota
)

func New(t uint) Cache {
	switch t {
	case Redis:
		return redis.NewCache()
	default:
		panic("invalid cache type")
	}
}
