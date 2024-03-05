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
