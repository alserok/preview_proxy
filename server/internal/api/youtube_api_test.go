package api

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestYoutubeAPIClient(t *testing.T) {
	suite.Run(t, new(YoutubeAPIClientSuite))
}

type YoutubeAPIClientSuite struct {
	suite.Suite
}

func (s *YoutubeAPIClientSuite) SetupTest() {

}

func (s *YoutubeAPIClientSuite) TestGetThumbnail() {

}
