package grpc

import (
	"context"
	"errors"
	"github.com/alserok/preview_proxy/server/internal/cache"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/service"
	"github.com/alserok/preview_proxy/server/internal/service/models"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestGRPCServerSuite(t *testing.T) {
	suite.Run(t, new(GRPCServerSuite))
}

type GRPCServerSuite struct {
	suite.Suite

	ctrl  *gomock.Controller
	mocks struct {
		service *service.MockService
		cache   *cache.MockCache
		logger  *logger.MockLogger
	}

	port string
	s    *server
}

func (s *GRPCServerSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mocks.service = service.NewMockService(s.ctrl)
	s.mocks.cache = cache.NewMockCache(s.ctrl)
	s.mocks.logger = logger.NewMockLogger(s.ctrl)

	s.port = "3001"
	s.s = NewServer(s.mocks.service, s.mocks.cache, s.mocks.logger)
}

func (s *GRPCServerSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *GRPCServerSuite) TestMustServe() {
	s.mocks.logger.EXPECT().
		Info("server is running", gomock.Eq(logger.WithArg("port", s.port))).
		Times(1)

	s.mocks.cache.EXPECT().
		Close().
		Return(nil).
		Times(1)

	defer func() {
		s.Require().Nil(recover())
	}()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	s.s.MustServe(ctx, s.port)
}

func (s *GRPCServerSuite) TestSyncGetThumbnails() {
	req := proto.GetThumbnailReq{
		VideoUrls: []string{"a", "b"},
		Async:     false,
	}

	s.mocks.service.EXPECT().
		GetThumbnails(gomock.Any(), gomock.Eq(models.DownloadThumbnailsReq{
			VideoURLs: req.VideoUrls,
			Async:     req.Async,
		})).
		Return(models.DownloadThumbnailsRes{
			Failed: uint32(0),
			Total:  uint32(len(req.VideoUrls)),
			Videos: []models.Video{
				{
					VideoURL:     req.VideoUrls[0],
					ThumbnailURL: req.VideoUrls[0],
				},
				{
					VideoURL:     req.VideoUrls[1],
					ThumbnailURL: req.VideoUrls[1],
				},
			},
		}, nil).
		Times(1)

	for _, url := range req.VideoUrls {
		s.mocks.cache.EXPECT().
			Get(gomock.Any(), gomock.Eq(url), gomock.Any()).
			Return(errors.New("not found")).
			Times(1)

		s.mocks.cache.EXPECT().
			Set(gomock.Any(), gomock.Eq(url), gomock.Eq(models.Video{
				VideoURL:     url,
				ThumbnailURL: url,
			})).
			Return(nil).
			Times(1)
	}

	res, err := s.s.GetThumbnails(context.Background(), &req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(len(req.VideoUrls), len(res.Videos))
	s.Require().Equal(uint32(0), res.Failed)
	s.Require().Equal(uint32(len(req.VideoUrls)), res.Total)
}
