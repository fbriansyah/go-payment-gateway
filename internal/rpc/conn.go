package rpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func clientErrorUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return errors.ReceiveGRPCError(invoker(ctx, method, req, reply, cc, opts...))
	}
}

func Dial(ctx context.Context, endpoint string) (conn *grpc.ClientConn, err error) {
	return grpc.DialContext(ctx, endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			clientErrorUnaryInterceptor(),
		),
	)
}