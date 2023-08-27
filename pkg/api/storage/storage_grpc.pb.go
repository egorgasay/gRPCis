// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
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

const (
	Storage_Set_FullMethodName            = "/api.Storage/Set"
	Storage_Get_FullMethodName            = "/api.Storage/Get"
	Storage_SetToObject_FullMethodName    = "/api.Storage/SetToObject"
	Storage_AttachToObject_FullMethodName = "/api.Storage/AttachToObject"
	Storage_GetFromObject_FullMethodName  = "/api.Storage/GetFromObject"
	Storage_ObjectToJSON_FullMethodName   = "/api.Storage/ObjectToJSON"
	Storage_IsObject_FullMethodName       = "/api.Storage/IsObject"
	Storage_NewObject_FullMethodName      = "/api.Storage/NewObject"
	Storage_Size_FullMethodName           = "/api.Storage/Size"
	Storage_Delete_FullMethodName         = "/api.Storage/Delete"
	Storage_DeleteObject_FullMethodName   = "/api.Storage/DeleteObject"
	Storage_DeleteAttr_FullMethodName     = "/api.Storage/DeleteAttr"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	SetToObject(ctx context.Context, in *SetToObjectRequest, opts ...grpc.CallOption) (*SetResponse, error)
	AttachToObject(ctx context.Context, in *AttachToObjectRequest, opts ...grpc.CallOption) (*AttachToObjectResponse, error)
	GetFromObject(ctx context.Context, in *GetFromObjectRequest, opts ...grpc.CallOption) (*GetResponse, error)
	ObjectToJSON(ctx context.Context, in *ObjectToJSONRequest, opts ...grpc.CallOption) (*ObjectToJSONResponse, error)
	IsObject(ctx context.Context, in *IsObjectRequest, opts ...grpc.CallOption) (*IsObjectResponse, error)
	NewObject(ctx context.Context, in *NewObjectRequest, opts ...grpc.CallOption) (*NewObjectResponse, error)
	Size(ctx context.Context, in *ObjectSizeRequest, opts ...grpc.CallOption) (*ObjectSizeResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	DeleteAttr(ctx context.Context, in *DeleteAttrRequest, opts ...grpc.CallOption) (*DeleteAttrResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, Storage_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Storage_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SetToObject(ctx context.Context, in *SetToObjectRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, Storage_SetToObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) AttachToObject(ctx context.Context, in *AttachToObjectRequest, opts ...grpc.CallOption) (*AttachToObjectResponse, error) {
	out := new(AttachToObjectResponse)
	err := c.cc.Invoke(ctx, Storage_AttachToObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetFromObject(ctx context.Context, in *GetFromObjectRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Storage_GetFromObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) ObjectToJSON(ctx context.Context, in *ObjectToJSONRequest, opts ...grpc.CallOption) (*ObjectToJSONResponse, error) {
	out := new(ObjectToJSONResponse)
	err := c.cc.Invoke(ctx, Storage_ObjectToJSON_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) IsObject(ctx context.Context, in *IsObjectRequest, opts ...grpc.CallOption) (*IsObjectResponse, error) {
	out := new(IsObjectResponse)
	err := c.cc.Invoke(ctx, Storage_IsObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) NewObject(ctx context.Context, in *NewObjectRequest, opts ...grpc.CallOption) (*NewObjectResponse, error) {
	out := new(NewObjectResponse)
	err := c.cc.Invoke(ctx, Storage_NewObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Size(ctx context.Context, in *ObjectSizeRequest, opts ...grpc.CallOption) (*ObjectSizeResponse, error) {
	out := new(ObjectSizeResponse)
	err := c.cc.Invoke(ctx, Storage_Size_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Storage_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteAttr(ctx context.Context, in *DeleteAttrRequest, opts ...grpc.CallOption) (*DeleteAttrResponse, error) {
	out := new(DeleteAttrResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteAttr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	SetToObject(context.Context, *SetToObjectRequest) (*SetResponse, error)
	AttachToObject(context.Context, *AttachToObjectRequest) (*AttachToObjectResponse, error)
	GetFromObject(context.Context, *GetFromObjectRequest) (*GetResponse, error)
	ObjectToJSON(context.Context, *ObjectToJSONRequest) (*ObjectToJSONResponse, error)
	IsObject(context.Context, *IsObjectRequest) (*IsObjectResponse, error)
	NewObject(context.Context, *NewObjectRequest) (*NewObjectResponse, error)
	Size(context.Context, *ObjectSizeRequest) (*ObjectSizeResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	DeleteAttr(context.Context, *DeleteAttrRequest) (*DeleteAttrResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStorageServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) SetToObject(context.Context, *SetToObjectRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToObject not implemented")
}
func (UnimplementedStorageServer) AttachToObject(context.Context, *AttachToObjectRequest) (*AttachToObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachToObject not implemented")
}
func (UnimplementedStorageServer) GetFromObject(context.Context, *GetFromObjectRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromObject not implemented")
}
func (UnimplementedStorageServer) ObjectToJSON(context.Context, *ObjectToJSONRequest) (*ObjectToJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectToJSON not implemented")
}
func (UnimplementedStorageServer) IsObject(context.Context, *IsObjectRequest) (*IsObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsObject not implemented")
}
func (UnimplementedStorageServer) NewObject(context.Context, *NewObjectRequest) (*NewObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewObject not implemented")
}
func (UnimplementedStorageServer) Size(context.Context, *ObjectSizeRequest) (*ObjectSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
}
func (UnimplementedStorageServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStorageServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedStorageServer) DeleteAttr(context.Context, *DeleteAttrRequest) (*DeleteAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttr not implemented")
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

func _Storage_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SetToObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetToObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SetToObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SetToObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SetToObject(ctx, req.(*SetToObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_AttachToObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachToObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).AttachToObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_AttachToObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).AttachToObject(ctx, req.(*AttachToObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetFromObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFromObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetFromObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetFromObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetFromObject(ctx, req.(*GetFromObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_ObjectToJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectToJSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).ObjectToJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_ObjectToJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).ObjectToJSON(ctx, req.(*ObjectToJSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_IsObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).IsObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_IsObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).IsObject(ctx, req.(*IsObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_NewObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).NewObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_NewObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).NewObject(ctx, req.(*NewObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Size_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Size_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Size(ctx, req.(*ObjectSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteAttr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteAttr(ctx, req.(*DeleteAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Storage_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "SetToObject",
			Handler:    _Storage_SetToObject_Handler,
		},
		{
			MethodName: "AttachToObject",
			Handler:    _Storage_AttachToObject_Handler,
		},
		{
			MethodName: "GetFromObject",
			Handler:    _Storage_GetFromObject_Handler,
		},
		{
			MethodName: "ObjectToJSON",
			Handler:    _Storage_ObjectToJSON_Handler,
		},
		{
			MethodName: "IsObject",
			Handler:    _Storage_IsObject_Handler,
		},
		{
			MethodName: "NewObject",
			Handler:    _Storage_NewObject_Handler,
		},
		{
			MethodName: "Size",
			Handler:    _Storage_Size_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Storage_Delete_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Storage_DeleteObject_Handler,
		},
		{
			MethodName: "DeleteAttr",
			Handler:    _Storage_DeleteAttr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/storage.proto",
}
