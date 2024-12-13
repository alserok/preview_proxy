package grpc

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/utils"
	"google.golang.org/grpc"
)

func WithErrorHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			return nil, utils.ErrToGRPCError(ctx, err)
		}

		return res, nil
	}
}
