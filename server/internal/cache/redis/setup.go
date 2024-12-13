package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func MustConnect(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	return client
}
