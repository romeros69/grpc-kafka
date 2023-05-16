package grpc_gen

import (
	"context"
	"google.golang.org/grpc"
	"grpc-client/internal/grpc_gen/gen"
)

type HelloClientGRPC struct {
	client gen.SayHelloClient
}

func NewGRPCClient(rpcConn *grpc.ClientConn) *HelloClientGRPC {
	return &HelloClientGRPC{
		client: gen.NewSayHelloClient(rpcConn),
	}
}

func (hc *HelloClientGRPC) Do(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	return hc.client.Do(ctx, req)
}
