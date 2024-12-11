package cache

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/cache/redis"
	r "github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string, target any) error

	Close() error
}

const (
	Redis = iota
)

func New(t uint, cl ...any) Cache {
	switch t {
	case Redis:
		return redis.NewCache(cl[0].(*r.Client))
	default:
		panic("invalid cache type")
	}
}
