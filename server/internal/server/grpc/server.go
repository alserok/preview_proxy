package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/service"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"google.golang.org/grpc"
)

func NewServer(srvc service.Service, cache cache.Cache) *server {
	return &server{
		cache: cache,
		srvc:  srvc,
	}
}

type server struct {
	srvc  service.Service
	cache cache.Cache

	proto.UnimplementedPreviewProxyServer
}

func (s *server) DownloadThumbnails(ctx context.Context, req *proto.DownloadThumbnailReq) (*proto.DownloadThumbnailRes, error) {

	return &proto.DownloadThumbnailRes{}, nil
}

func (s *server) MustServe(port string) {
	defer func() {
		if err := s.cache.Close(); err != nil {
			// TODO: log
		}
	}()

	serv := grpc.NewServer()

	proto.RegisterPreviewProxyServer(serv, s)
}
