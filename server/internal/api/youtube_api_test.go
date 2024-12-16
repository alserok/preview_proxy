package api

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestYoutubeAPIClient(t *testing.T) {
	suite.Run(t, new(YoutubeAPIClientSuite))
}

type YoutubeAPIClientSuite struct {
	suite.Suite

	ctrl  *gomock.Controller
	mocks struct {
		logger *logger.MockLogger
	}

	client *youtubeAPIClient
}

func (s *YoutubeAPIClientSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mocks.logger = logger.NewMockLogger(s.ctrl)
	s.mocks.logger.EXPECT().
		Debug(gomock.Any(), gomock.Any()).
		AnyTimes()

	s.mocks.logger.EXPECT().
		Debug(gomock.Any(), gomock.Any()).
		AnyTimes()

	s.client = NewYoutubeAPIClient()
}

func (s *YoutubeAPIClientSuite) TestGetThumbnail() {
	previewURL, err := s.client.GetThumbnail(context.WithValue(context.Background(), logger.CtxLogger, s.mocks.logger), "3061eHuACEU")
	s.Require().NoError(err)
	s.Require().NotEmpty(previewURL)
}
