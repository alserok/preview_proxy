package grpc

import (
	"context"
	"google.golang.org/grpc"
)

func WithErrorHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			panic(err)
			// handle grpc error
		}

		return res, nil
	}
}
