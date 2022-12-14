// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: logger/logger.proto

package logger

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

// LoggerServiceClient is the client API for LoggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerServiceClient interface {
	Streamer(ctx context.Context, opts ...grpc.CallOption) (LoggerService_StreamerClient, error)
}

type loggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerServiceClient(cc grpc.ClientConnInterface) LoggerServiceClient {
	return &loggerServiceClient{cc}
}

func (c *loggerServiceClient) Streamer(ctx context.Context, opts ...grpc.CallOption) (LoggerService_StreamerClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggerService_ServiceDesc.Streams[0], "/logger.LoggerService/Streamer", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggerServiceStreamerClient{stream}
	return x, nil
}

type LoggerService_StreamerClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type loggerServiceStreamerClient struct {
	grpc.ClientStream
}

func (x *loggerServiceStreamerClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loggerServiceStreamerClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggerServiceServer is the server API for LoggerService service.
// All implementations must embed UnimplementedLoggerServiceServer
// for forward compatibility
type LoggerServiceServer interface {
	Streamer(LoggerService_StreamerServer) error
	mustEmbedUnimplementedLoggerServiceServer()
}

// UnimplementedLoggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServiceServer struct {
}

func (UnimplementedLoggerServiceServer) Streamer(LoggerService_StreamerServer) error {
	return status.Errorf(codes.Unimplemented, "method Streamer not implemented")
}
func (UnimplementedLoggerServiceServer) mustEmbedUnimplementedLoggerServiceServer() {}

// UnsafeLoggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServiceServer will
// result in compilation errors.
type UnsafeLoggerServiceServer interface {
	mustEmbedUnimplementedLoggerServiceServer()
}

func RegisterLoggerServiceServer(s grpc.ServiceRegistrar, srv LoggerServiceServer) {
	s.RegisterService(&LoggerService_ServiceDesc, srv)
}

func _LoggerService_Streamer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoggerServiceServer).Streamer(&loggerServiceStreamerServer{stream})
}

type LoggerService_StreamerServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type loggerServiceStreamerServer struct {
	grpc.ServerStream
}

func (x *loggerServiceStreamerServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loggerServiceStreamerServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggerService_ServiceDesc is the grpc.ServiceDesc for LoggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.LoggerService",
	HandlerType: (*LoggerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Streamer",
			Handler:       _LoggerService_Streamer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "logger/logger.proto",
}
