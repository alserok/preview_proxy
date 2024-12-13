package grpc

import (
	"context"
	"errors"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/utils"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
)

func TestErrorsMWSuite(t *testing.T) {
	suite.Run(t, new(ErrorsMWSuite))
}

type ErrorsMWSuite struct {
	suite.Suite

	ctrl  *gomock.Controller
	mocks struct {
		logger *logger.MockLogger
	}
}

func (s *ErrorsMWSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mocks.logger = logger.NewMockLogger(s.ctrl)
}

func (s *ErrorsMWSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *ErrorsMWSuite) TestInternalError() {
	msg := "internal error"
	err := utils.NewError(msg, utils.Internal)

	s.mocks.logger.EXPECT().
		Error(gomock.Eq(msg)).
		Times(1)

	res, err := WithErrorHandler()(
		logger.WrapLogger(context.Background(), s.mocks.logger),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			return nil, err
		},
	)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(msg, st.Message())
}

func (s *ErrorsMWSuite) TestBadRequestError() {
	msg := "bad request"
	err := utils.NewError(msg, utils.BadRequest)

	res, err := WithErrorHandler()(
		logger.WrapLogger(context.Background(), s.mocks.logger),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			return nil, err
		},
	)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(msg, st.Message())
}

func (s *ErrorsMWSuite) TestNotFoundError() {
	msg := "not found"
	err := utils.NewError(msg, utils.NotFound)

	res, err := WithErrorHandler()(
		logger.WrapLogger(context.Background(), s.mocks.logger),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			return nil, err
		},
	)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(msg, st.Message())
}

func (s *ErrorsMWSuite) TestInvalidTypeError() {
	msg := "internal error"
	err := errors.New("some error with invalid type")

	s.mocks.logger.EXPECT().
		Error(gomock.Eq(msg)).
		Times(1)

	res, err := WithErrorHandler()(
		logger.WrapLogger(context.Background(), s.mocks.logger),
		&proto.GetThumbnailReq{},
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req any) (interface{}, error) {
			return nil, err
		},
	)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(msg, st.Message())
}
