package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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

	clientStream, err := client.GreetClientStream(context.Background())
	if err != nil {
		log.Fatalf("failed to get client: %v", err)
	}

	for i := 0; i < 10; i++ {
		clientStream.Send(&greeter_server.GreetRequest{
			Name: "hello",
			Id:   fmt.Sprintf("%d", i),
			Date: "2006-01-02",
		})
		time.Sleep(50 * time.Millisecond)
	}

	resp, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close client: %v", err)
	}
	log.Printf("reply received: %v", resp.Reply)

	biStream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("failed to bidirectional stream: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			resp, err := biStream.Recv()
			if err == io.EOF {
				wg.Done()
				break
			}
			if err != nil {
				log.Fatalf("failed to receive response: %v", err)
			}

			log.Printf("client bidirectional stream received: %v", resp.Reply)
		}
	}()
	for i := 0; i < 10; i++ {
		biStream.Send(&greeter_server.GreetRequest{
			Name: "hello",
			Id:   fmt.Sprintf("%d", i),
			Date: "2006-01-02",
		})
	}

	wg.Wait()
}
