package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/logger"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"testing"
)

func TestLoggerMWSuite(t *testing.T) {
	suite.Run(t, new(LoggerMWSuite))
}

type LoggerMWSuite struct {
	suite.Suite

	ctrl  *gomock.Controller
	mocks struct {
		logger *logger.MockLogger
	}
}

func (s *LoggerMWSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mocks.logger = logger.NewMockLogger(s.ctrl)
}

func (s *LoggerMWSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *LoggerMWSuite) TestToAndFromContext() {
	s.mocks.logger.EXPECT().
		Info(gomock.Eq("1")).
		Times(1)
	s.mocks.logger.EXPECT().
		Error(gomock.Eq("1")).
		Times(1)
	s.mocks.logger.EXPECT().
		Debug(gomock.Eq("1")).
		Times(1)
	s.mocks.logger.EXPECT().
		Warn(gomock.Eq("1")).
		Times(1)

	res, err := WithRecovery(s.mocks.logger)(
		logger.WrapLogger(context.Background(), s.mocks.logger),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			log := logger.FromContext(ctx)
			s.Require().NotNil(log)

			log.Info("1")
			log.Error("1")
			log.Debug("1")
			log.Warn("1")

			return nil, nil
		},
	)
	s.Require().Nil(res)
	s.Require().Nil(err)
}
