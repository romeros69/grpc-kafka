package grpc_gen

import (
	"context"
	gen2 "grpc-serv/internal/grpc_gen/gen"
)

type HelloServerGRPC struct {
	gen2.UnimplementedSayHelloServer
}

func NewHelloServerGRPC() *HelloServerGRPC {
	return &HelloServerGRPC{}
}

func (hs *HelloServerGRPC) Do(ctx context.Context, req *gen2.Request) (*gen2.Response, error) {
	return &gen2.Response{
		Message: "Hello " + req.Message,
	}, nil
}
