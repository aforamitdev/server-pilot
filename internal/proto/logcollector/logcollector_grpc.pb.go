// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: internal/proto/logcollector/logcollector.proto

package logcollector_proto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// LogCollectorClient is the client API for LogCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogCollectorClient interface {
}

type logCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewLogCollectorClient(cc grpc.ClientConnInterface) LogCollectorClient {
	return &logCollectorClient{cc}
}

// LogCollectorServer is the server API for LogCollector service.
// All implementations must embed UnimplementedLogCollectorServer
// for forward compatibility
type LogCollectorServer interface {
	mustEmbedUnimplementedLogCollectorServer()
}

// UnimplementedLogCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedLogCollectorServer struct {
}

func (UnimplementedLogCollectorServer) mustEmbedUnimplementedLogCollectorServer() {}

// UnsafeLogCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogCollectorServer will
// result in compilation errors.
type UnsafeLogCollectorServer interface {
	mustEmbedUnimplementedLogCollectorServer()
}

func RegisterLogCollectorServer(s grpc.ServiceRegistrar, srv LogCollectorServer) {
	s.RegisterService(&LogCollector_ServiceDesc, srv)
}

// LogCollector_ServiceDesc is the grpc.ServiceDesc for LogCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LogCollector",
	HandlerType: (*LogCollectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "internal/proto/logcollector/logcollector.proto",
}