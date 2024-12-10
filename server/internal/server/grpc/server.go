package grpc

import (
	"context"
	"fmt"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/service"
	"github.com/alserok/preview_proxy/server/internal/service/models"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"google.golang.org/grpc"
	"net"
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

func (s *server) GetThumbnails(ctx context.Context, req *proto.GetThumbnailReq) (*proto.GetThumbnailRes, error) {
	data, err := s.srvc.GetThumbnails(ctx, models.DownloadThumbnailsReq{
		VideoURLs: req.VideoUrls,
		Async:     req.Async,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download thumbnails: %w", err)
	}

	var res proto.GetThumbnailRes

	res.Total = data.Total
	res.Failed = data.Failed
	for _, video := range data.Videos {
		res.Videos = append(res.Videos, &proto.Video{
			VideoUrl:     video.VideoURL,
			ThumbnailUrl: video.ThumbnailURL,
		})
	}

	return &res, nil
}

func (s *server) MustServe(port string) {
	defer func() {
		if err := s.cache.Close(); err != nil {
			// TODO: log
			return
		}
	}()

	serv := grpc.NewServer()

	proto.RegisterPreviewProxyServer(serv, s)

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	if err = serv.Serve(l); err != nil {
		panic("failed to serve: " + err.Error())
	}
}
