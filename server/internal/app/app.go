package app

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/api"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/cache/redis"
	"github.com/alserok/preview_proxy/server/internal/config"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/server"
	"github.com/alserok/preview_proxy/server/internal/service"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	log := logger.NewLogger(logger.Slog, cfg.Env)
	defer func() {
		_ = log.Close()
	}()

	clients := service.Clients{
		YoutubeAPIClient: api.NewYoutubeAPIClient(),
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

	srvr.MustServe(ctx, cfg.Port)
}
