package app

import (
	"grpc-serv/internal/grpc_gen"
	"grpc-serv/pkg/rpc"
	"log"
)

func Run() {
	grpcServer, err := rpc.New()
	if err != nil {
		log.Fatalf("error in creating grp server: %s", err.Error())
	}
	helloServer := grpc_gen.NewHelloServerGRPC()
	err = grpcServer.Start(helloServer)

	if err != nil {
		log.Fatalf("error in start grpc_gen: %s", err.Error())
	}
}
