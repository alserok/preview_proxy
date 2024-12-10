package cache

import "context"

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
	default:
		panic("invalid cache type")
	}
}
