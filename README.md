# GRPC Stream Example

This repository contains examples about grpc Server stream, Client stream and Bidirection Grpc.

### Server Stream Proto
```proto
syntax = "proto3";

option go_package = "./greeter_server";
option java_multiple_files = true;
option java_package = "io.grpc.greeter.server";
option java_outer_classname = "ServerProto";

package greeter_server;

service GreeterService {
    rpc GreetServerStream (GreetRequest) returns (stream GreetResponse);
}

message GreetRequest {
    string id = 1;
    string name = 2;
    string date = 3;
}

message GreetResponse {
    string reply = 1;
}
```

### Server Implement
```go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "io.examples.greeter/grpc/examples/proto"
)

var (
	flagPort int64
)

func init() {
	flag.Int64Var(&flagPort, "port", 50051, "port to listen on server")
}

type serverImpl struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *serverImpl) GreetServerStream(req *pb.GreetRequest, stream pb.GreeterService_GreetServerStreamServer) error {
	for i := 0; i < 10000; i++ {
		stream.Send(&pb.GreetResponse{
			Reply: fmt.Sprintf("%d", i),
		})

		if i%100 == 0 {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(300 * time.Millisecond)
		}
	}

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", flagPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &serverImpl{})
	log.Printf("sever listening at port %d", flagPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

```

### Client Implement
```go
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

```