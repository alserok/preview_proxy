package app

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/api"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/cache/redis"
	"github.com/alserok/preview_proxy/server/internal/config"
	"github.com/alserok/preview_proxy/server/internal/server"
	"github.com/alserok/preview_proxy/server/internal/service"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	log := 0

	clients := service.Clients{
		YoutubeAPIClient: api.NewYoutubeAPIClient(cfg.API.YoutubeAddr),
	}

	srvc := service.New(clients)
	srvr := server.New(
		server.GRPC,
		srvc,
		cache.New(cache.Redis, redis.MustConnect(cfg.Cache.Addr)),
		log,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go srvr.MustServe(cfg.Port)

	<-ctx.Done()
}
