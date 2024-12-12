package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"google.golang.org/grpc"
)

func WithLogger(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(logger.WrapLogger(ctx, log), req)
	}
}
