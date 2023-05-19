package grpc_gen

import (
	"context"
	"fmt"
	gen2 "grpc-serv/internal/grpc_gen/gen"
	"time"
)

type HelloServerGRPC struct {
	gen2.SayHelloServer
}

func NewHelloServerGRPC() *HelloServerGRPC {
	return &HelloServerGRPC{}
}

func (hs *HelloServerGRPC) Do(ctx context.Context, req *gen2.Request) (*gen2.Response, error) {
	fmt.Printf("%s | Получено сообщение от клиента: %s\n", time.Now().String(), req.Message)
	return &gen2.Response{
		Message: "Привет " + req.Message,
	}, nil
}
