package service

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/api"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/service/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	s *service

	mocks struct {
		youtubeAPIClient *api.MockYoutubeAPI
		logger           *logger.MockLogger
	}
}

func (s *ServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mocks.youtubeAPIClient = api.NewMockYoutubeAPI(s.ctrl)
	s.mocks.logger = logger.NewMockLogger(s.ctrl)
	s.mocks.logger.EXPECT().
		Debug(gomock.Any(), gomock.Any()).
		AnyTimes()

	s.s = &service{
		youtubeAPIClient: s.mocks.youtubeAPIClient,
	}
}

func (s *ServiceSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *ServiceSuite) TestSyncGetThumbnails() {
	req := models.DownloadThumbnailsReq{
		VideoURLs: []string{"https://www.youtube.com/watch?v=lDLT7s0TAYs", "https://www.youtube.com/watch?v=3061eHuACEU"},
		Async:     false,
	}

	for _, videoURL := range req.VideoURLs {
		videoID, err := getVideoIDFromURL(videoURL)
		s.Require().NoError(err)
		s.Require().NotEmpty(videoID)

		s.mocks.youtubeAPIClient.EXPECT().
			GetThumbnail(gomock.Any(), gomock.Eq(videoID)).
			Times(1)
	}

	res, err := s.s.GetThumbnails(context.WithValue(context.Background(), logger.CtxLogger, s.mocks.logger), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(uint32(0), res.Failed)
	s.Require().Equal(len(req.VideoURLs), len(res.Videos))
	s.Require().Equal(uint32(len(req.VideoURLs)), res.Total)
}

func (s *ServiceSuite) TestAsyncGetThumbnails() {
	req := models.DownloadThumbnailsReq{
		VideoURLs: []string{"https://www.youtube.com/watch?v=lDLT7s0TAYs", "https://www.youtube.com/watch?v=3061eHuACEU"},
		Async:     true,
	}

	for _, videoURL := range req.VideoURLs {
		videoID, err := getVideoIDFromURL(videoURL)
		s.Require().NoError(err)
		s.Require().NotEmpty(videoID)

		s.mocks.youtubeAPIClient.EXPECT().
			GetThumbnail(gomock.Any(), gomock.Eq(videoID)).
			Times(1)
	}

	res, err := s.s.GetThumbnails(context.WithValue(context.Background(), logger.CtxLogger, s.mocks.logger), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(uint32(0), res.Failed)
	s.Require().Equal(len(req.VideoURLs), len(res.Videos))
	s.Require().Equal(uint32(len(req.VideoURLs)), res.Total)
}

func (s *ServiceSuite) TestGetVideoID() {
	videoIDs := [][2]string{
		{"https://www.youtube.com/watch?v=lDLT7s0TAYs", "lDLT7s0TAYs"},
		{"https://www.youtube.com/watch?v=3061eHuACEU", "3061eHuACEU"},
	}

	for _, video := range videoIDs {
		id, err := getVideoIDFromURL(video[0])
		s.Require().NoError(err)
		s.Require().Equal(video[1], id)
	}
}
