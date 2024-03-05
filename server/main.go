package main

import (
	"flag"
	"fmt"
	"io"
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
	for i := 0; i < 100; i++ {
		stream.Send(&pb.GreetResponse{
			Reply: fmt.Sprintf("%d", i),
		})

		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func (s *serverImpl) GreetClientStream(stream pb.GreeterService_GreetClientStreamServer) error {
	var reqs []*pb.GreetRequest
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{Reply: fmt.Sprintf("received total %d requests", len(reqs))})
			break
		}

		if err != nil {
			log.Fatalf("failed to received requests: %v", err)
		}

		log.Printf("received requests %v\n", req)
		reqs = append(reqs, req)
	}

	return nil
}

func (s *serverImpl) BidirectionalStream(stream pb.GreeterService_BidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive requests: %v", err)
		}

		log.Printf("server bidirection received: %v", req)
		time.Sleep(50 * time.Millisecond)
		stream.Send(&pb.GreetResponse{Reply: fmt.Sprintf("stream received request has id %s", req.Id)})
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
