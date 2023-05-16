package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-serv/internal/grpc_gen/gen"
	"log"
	"net"
)

const (
	GRPC_PROTOCOL = "tcp"
	GRPC_URL      = ":9011"
)

type ServerRpc struct {
	l net.Listener

	server *grpc.Server
}

func New() (*ServerRpc, error) {
	listen, err := net.Listen(GRPC_PROTOCOL, GRPC_URL)
	if err != nil {
		return nil, fmt.Errorf("error in init grpc_gen - lisening: %w", err)
	}

	return &ServerRpc{
		l:      listen,
		server: grpc.NewServer(),
	}, nil
}

func (r *ServerRpc) Start(sayHelloServ gen.SayHelloServer) error {
	//go func() {				<--- FIXME НЕ ЗАБУДЬ ВЕРНУТЬ (ЗАПУСК В MAIN ГОРУТИНЕ!)
	gen.RegisterSayHelloServer(r.server, sayHelloServ)
	if err := r.server.Serve(r.l); err != nil {
		log.Fatalf("Cannot start grpc_gen server: %v", err)
	}

	//}()

	return nil
}
