// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: internal/proto/system/system.proto

package system_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Informer_GetSystem_FullMethodName = "/Informer/GetSystem"
)

// InformerClient is the client API for Informer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformerClient interface {
	GetSystem(ctx context.Context, in *SystemRequest, opts ...grpc.CallOption) (*SystemResponse, error)
}

type informerClient struct {
	cc grpc.ClientConnInterface
}

func NewInformerClient(cc grpc.ClientConnInterface) InformerClient {
	return &informerClient{cc}
}

func (c *informerClient) GetSystem(ctx context.Context, in *SystemRequest, opts ...grpc.CallOption) (*SystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemResponse)
	err := c.cc.Invoke(ctx, Informer_GetSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformerServer is the server API for Informer service.
// All implementations must embed UnimplementedInformerServer
// for forward compatibility
type InformerServer interface {
	GetSystem(context.Context, *SystemRequest) (*SystemResponse, error)
	mustEmbedUnimplementedInformerServer()
}

// UnimplementedInformerServer must be embedded to have forward compatible implementations.
type UnimplementedInformerServer struct {
}

func (UnimplementedInformerServer) GetSystem(context.Context, *SystemRequest) (*SystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (UnimplementedInformerServer) mustEmbedUnimplementedInformerServer() {}

// UnsafeInformerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformerServer will
// result in compilation errors.
type UnsafeInformerServer interface {
	mustEmbedUnimplementedInformerServer()
}

func RegisterInformerServer(s grpc.ServiceRegistrar, srv InformerServer) {
	s.RegisterService(&Informer_ServiceDesc, srv)
}

func _Informer_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Informer_GetSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerServer).GetSystem(ctx, req.(*SystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Informer_ServiceDesc is the grpc.ServiceDesc for Informer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Informer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Informer",
	HandlerType: (*InformerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystem",
			Handler:    _Informer_GetSystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/system/system.proto",
}
