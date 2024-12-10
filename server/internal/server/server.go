package server

import (
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/server/grpc"
	"github.com/alserok/preview_proxy/server/internal/service"
)

type Server interface {
	MustServe(port string)
}

const (
	GRPC = iota
)

func New(t uint, service service.Service, cache cache.Cache) Server {
	switch t {
	case GRPC:
		return grpc.NewServer(service, cache)
	default:
		panic("invalid server type")
	}
}
