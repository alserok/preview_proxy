package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/logger"
	mw "github.com/alserok/preview_proxy/server/internal/middleware/grpc"
	"github.com/alserok/preview_proxy/server/internal/service"
	"github.com/alserok/preview_proxy/server/internal/service/models"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"google.golang.org/grpc"
	"net"
)

func NewServer(srvc service.Service, cache cache.Cache, log logger.Logger) *server {
	return &server{
		cache: cache,
		srvc:  srvc,
		log:   log,
	}
}

type server struct {
	srvc  service.Service
	cache cache.Cache
	log   logger.Logger

	proto.UnimplementedPreviewProxyServer
}

func (s *server) GetThumbnails(ctx context.Context, req *proto.GetThumbnailReq) (*proto.GetThumbnailRes, error) {
	var (
		videos    []models.Video
		videoURLs []string
	)
	for _, url := range req.VideoUrls {
		var video models.Video
		if err := s.cache.Get(ctx, url, &video); err == nil {
			videos = append(videos, video)
			continue
		}

		videoURLs = append(videoURLs, url)
	}

	data, err := s.srvc.GetThumbnails(ctx, models.DownloadThumbnailsReq{
		VideoURLs: videoURLs,
		Async:     req.Async,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download thumbnails: %w", err)
	}

	var res proto.GetThumbnailRes
	res.Total = data.Total
	res.Failed = data.Failed
	for _, video := range append(data.Videos, videos...) {
		res.Videos = append(res.Videos, &proto.Video{
			VideoUrl:  video.VideoURL,
			Thumbnail: video.Thumbnail,
		})

		if err = s.cache.Set(ctx, video.VideoURL, video); err != nil {
			s.log.Warn("failed to set cache value", logger.WithArg("error", err.Error()))
		}
	}

	return &res, nil
}

func (s *server) MustServe(ctx context.Context, port string) {
	defer func() {
		if err := s.cache.Close(); err != nil {
			s.log.Error("failed to close cache", logger.WithArg("error", err.Error()))
		}
	}()

	serv := grpc.NewServer(mw.WithMiddlewareChain(
		mw.WithRecovery(s.log),
		mw.WithLogger(s.log),
		mw.WithErrorHandler(),
	))
	defer serv.GracefulStop()

	proto.RegisterPreviewProxyServer(serv, s)

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	defer func() {
		_ = l.Close()
	}()

	s.log.Info("server is running", logger.WithArg("port", port))
	go func() {
		if err = serv.Serve(l); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			panic("failed to serve: " + err.Error())
		}
	}()

	<-ctx.Done()
}
