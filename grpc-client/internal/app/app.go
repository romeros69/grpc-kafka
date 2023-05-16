package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-client/internal/grpc_gen"
	"grpc-client/internal/grpc_gen/gen"
	"log"
	"time"
)

const GRPC_URL = "localhost:9011"

func Run() {
	conn, err := grpc.Dial(GRPC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Cannot dial with gRPC connection: %v", err)
	}
	defer conn.Close()

	grpcClient := grpc_gen.NewGRPCClient(conn)

	ctx := context.Background()

	response, err := grpcClient.Do(ctx, &gen.Request{Message: "куни чел"})

	if err != nil {
		log.Fatalf("Fuckup grpc: %v", err)
	}

	fmt.Printf("%s | Получено сообщение от сервера: %s\n", time.Now().String(), response.Message)
}
