package server

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/server/grpc"
	"github.com/alserok/preview_proxy/server/internal/service"
)

type Server interface {
	MustServe(ctx context.Context, port string)
}

const (
	GRPC = iota
)

func New(t uint, service service.Service, cache cache.Cache, log logger.Logger) Server {
	switch t {
	case GRPC:
		return grpc.NewServer(service, cache, log)
	default:
		panic("invalid server type")
	}
}
