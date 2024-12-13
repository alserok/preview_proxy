package redis

import (
	"context"
	"encoding/json"
	"github.com/alserok/preview_proxy/server/internal/utils"
	"github.com/go-redis/redis/v8"
	"time"
)

func NewCache(cl *redis.Client) *cache {
	return &cache{
		cl: cl,
	}
}

type cache struct {
	cl *redis.Client
}

func (c *cache) Set(ctx context.Context, key string, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = c.cl.Set(ctx, key, b, time.Hour*24).Err(); err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func (c *cache) Get(ctx context.Context, key string, target any) error {
	res, err := c.cl.Get(ctx, key).Result()
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = json.Unmarshal([]byte(res), target); err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func (c *cache) Close() error {
	if err := c.cl.Close(); err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}
