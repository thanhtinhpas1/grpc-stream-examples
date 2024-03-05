// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: server/service.proto

package greeter_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	GreetServerStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreeterService_GreetServerStreamClient, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) GreetServerStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreeterService_GreetServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[0], "/greeter_server.GreeterService/GreetServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceGreetServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterService_GreetServerStreamClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greeterServiceGreetServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServiceGreetServerStreamClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	GreetServerStream(*GreetRequest, GreeterService_GreetServerStreamServer) error
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) GreetServerStream(*GreetRequest, GreeterService_GreetServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetServerStream not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_GreetServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServiceServer).GreetServerStream(m, &greeterServiceGreetServerStreamServer{stream})
}

type GreeterService_GreetServerStreamServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greeterServiceGreetServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServiceGreetServerStreamServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeter_server.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetServerStream",
			Handler:       _GreeterService_GreetServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/service.proto",
}