// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/proto/storage.proto

package storage

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

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Set(ctx context.Context, opts ...grpc.CallOption) (Storage_SetClient, error)
	Get(ctx context.Context, opts ...grpc.CallOption) (Storage_GetClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Set(ctx context.Context, opts ...grpc.CallOption) (Storage_SetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], "/api.Storage/Set", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageSetClient{stream}
	return x, nil
}

type Storage_SetClient interface {
	Send(*SetRequest) error
	Recv() (*SetResponse, error)
	grpc.ClientStream
}

type storageSetClient struct {
	grpc.ClientStream
}

func (x *storageSetClient) Send(m *SetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageSetClient) Recv() (*SetResponse, error) {
	m := new(SetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Get(ctx context.Context, opts ...grpc.CallOption) (Storage_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[1], "/api.Storage/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetClient{stream}
	return x, nil
}

type Storage_GetClient interface {
	Send(*GetRequest) error
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type storageGetClient struct {
	grpc.ClientStream
}

func (x *storageGetClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageGetClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Set(Storage_SetServer) error
	Get(Storage_GetServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Set(Storage_SetServer) error {
	return status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStorageServer) Get(Storage_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Set_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Set(&storageSetServer{stream})
}

type Storage_SetServer interface {
	Send(*SetResponse) error
	Recv() (*SetRequest, error)
	grpc.ServerStream
}

type storageSetServer struct {
	grpc.ServerStream
}

func (x *storageSetServer) Send(m *SetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageSetServer) Recv() (*SetRequest, error) {
	m := new(SetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Get(&storageGetServer{stream})
}

type Storage_GetServer interface {
	Send(*GetResponse) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type storageGetServer struct {
	grpc.ServerStream
}

func (x *storageGetServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageGetServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Set",
			Handler:       _Storage_Set_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _Storage_Get_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/storage.proto",
}
