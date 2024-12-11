package grpc

import "google.golang.org/grpc"

func WithMiddlewareChain(inters ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(inters...)
}
