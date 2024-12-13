package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WithRecovery(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
		defer func() {
			if e := recover(); e != nil {
				log.Error("panic recovery", logger.WithArg("error", e))
				err = status.Errorf(codes.Internal, "internal error")
			}
		}()

		res, err = handler(ctx, req)

		return
	}
}
