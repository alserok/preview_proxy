package api

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestYoutubeAPIClient(t *testing.T) {
	suite.Run(t, new(YoutubeAPIClientSuite))
}

type YoutubeAPIClientSuite struct {
	suite.Suite

	client *youtubeAPIClient
}

func (s *YoutubeAPIClientSuite) SetupTest() {
	s.client = NewYoutubeAPIClient()
}

func (s *YoutubeAPIClientSuite) TestGetThumbnail() {
	previewURL, err := s.client.GetThumbnail(context.Background(), "ps--Onn3p_s")
	s.Require().NoError(err)
	s.Require().NotEmpty(previewURL)
}
