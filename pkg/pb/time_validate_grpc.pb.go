// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: pkg/proto/time_validate.proto

package pb

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

// TimeValidateServiceClient is the client API for TimeValidateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeValidateServiceClient interface {
	Validate(ctx context.Context, opts ...grpc.CallOption) (TimeValidateService_ValidateClient, error)
}

type timeValidateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeValidateServiceClient(cc grpc.ClientConnInterface) TimeValidateServiceClient {
	return &timeValidateServiceClient{cc}
}

func (c *timeValidateServiceClient) Validate(ctx context.Context, opts ...grpc.CallOption) (TimeValidateService_ValidateClient, error) {
	stream, err := c.cc.NewStream(ctx, &TimeValidateService_ServiceDesc.Streams[0], "/TimeValidateService/Validate", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeValidateServiceValidateClient{stream}
	return x, nil
}

type TimeValidateService_ValidateClient interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ClientStream
}

type timeValidateServiceValidateClient struct {
	grpc.ClientStream
}

func (x *timeValidateServiceValidateClient) Send(m *Response) error {
	return x.ClientStream.SendMsg(m)
}

func (x *timeValidateServiceValidateClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeValidateServiceServer is the server API for TimeValidateService service.
// All implementations should embed UnimplementedTimeValidateServiceServer
// for forward compatibility
type TimeValidateServiceServer interface {
	Validate(TimeValidateService_ValidateServer) error
}

// UnimplementedTimeValidateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTimeValidateServiceServer struct {
}

func (UnimplementedTimeValidateServiceServer) Validate(TimeValidateService_ValidateServer) error {
	return status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

// UnsafeTimeValidateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeValidateServiceServer will
// result in compilation errors.
type UnsafeTimeValidateServiceServer interface {
	mustEmbedUnimplementedTimeValidateServiceServer()
}

func RegisterTimeValidateServiceServer(s grpc.ServiceRegistrar, srv TimeValidateServiceServer) {
	s.RegisterService(&TimeValidateService_ServiceDesc, srv)
}

func _TimeValidateService_Validate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TimeValidateServiceServer).Validate(&timeValidateServiceValidateServer{stream})
}

type TimeValidateService_ValidateServer interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ServerStream
}

type timeValidateServiceValidateServer struct {
	grpc.ServerStream
}

func (x *timeValidateServiceValidateServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func (x *timeValidateServiceValidateServer) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeValidateService_ServiceDesc is the grpc.ServiceDesc for TimeValidateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeValidateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimeValidateService",
	HandlerType: (*TimeValidateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Validate",
			Handler:       _TimeValidateService_Validate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/time_validate.proto",
}
