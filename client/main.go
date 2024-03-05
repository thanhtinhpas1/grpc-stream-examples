package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greeter_server "io.examples.greeter/grpc/examples/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()
	client := greeter_server.NewGreeterServiceClient(conn)

	stream, err := client.GreetServerStream(context.Background(), &greeter_server.GreetRequest{
		Name: "abc",
		Id:   "1",
		Date: "01/01/2024",
	})

	if err != nil {
		log.Fatalf("cannot call to grpc: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}

		log.Printf("reply received: %v", resp.Reply)
	}
}
