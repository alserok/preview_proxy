package redis

import "context"

func NewCache() *cache {
	return &cache{}
}

type cache struct {
}

func (c cache) Set(ctx context.Context, key string, value []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c cache) Get(ctx context.Context, key string, target any) error {
	//TODO implement me
	panic("implement me")
}

func (c cache) Close() error {
	//TODO implement me
	panic("implement me")
}
