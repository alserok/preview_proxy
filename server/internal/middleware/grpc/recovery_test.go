package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/logger"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
)

func TestRecoveryMWSuite(t *testing.T) {
	suite.Run(t, new(RecoveryMWSuite))
}

type RecoveryMWSuite struct {
	suite.Suite

	ctrl  *gomock.Controller
	mocks struct {
		logger *logger.MockLogger
	}
}

func (s *RecoveryMWSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mocks.logger = logger.NewMockLogger(s.ctrl)
}

func (s *RecoveryMWSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *RecoveryMWSuite) TestWithPanic() {
	msg := "internal error"

	s.mocks.logger.EXPECT().
		Error(gomock.Eq("panic recovery"), gomock.Eq(logger.WithArg("error", "some panic"))).
		Times(1)

	res, err := WithRecovery(s.mocks.logger)(
		context.Background(),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			defer func() {
				panic("some panic")
			}()
			return nil, nil
		},
	)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(msg, st.Message())
}
