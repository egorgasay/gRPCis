// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/proto/balancer.proto

package balancer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BalancerSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Server int32  `protobuf:"varint,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *BalancerSetRequest) Reset() {
	*x = BalancerSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerSetRequest) ProtoMessage() {}

func (x *BalancerSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerSetRequest.ProtoReflect.Descriptor instead.
func (*BalancerSetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *BalancerSetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BalancerSetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *BalancerSetRequest) GetServer() int32 {
	if x != nil {
		return x.Server
	}
	return 0
}

type BalancerGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Server int32  `protobuf:"varint,2,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *BalancerGetRequest) Reset() {
	*x = BalancerGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerGetRequest) ProtoMessage() {}

func (x *BalancerGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerGetRequest.ProtoReflect.Descriptor instead.
func (*BalancerGetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{1}
}

func (x *BalancerGetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BalancerGetRequest) GetServer() int32 {
	if x != nil {
		return x.Server
	}
	return 0
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Total     uint64 `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Available uint64 `protobuf:"varint,3,opt,name=Available,proto3" json:"Available,omitempty"`
	Server    int32  `protobuf:"varint,4,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectRequest) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ConnectRequest) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *ConnectRequest) GetServer() int32 {
	if x != nil {
		return x.Server
	}
	return 0
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerNumber int32 `protobuf:"varint,1,opt,name=serverNumber,proto3" json:"serverNumber,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectRequest) GetServerNumber() int32 {
	if x != nil {
		return x.ServerNumber
	}
	return 0
}

type ServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServersRequest) Reset() {
	*x = ServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersRequest) ProtoMessage() {}

func (x *ServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersRequest.ProtoReflect.Descriptor instead.
func (*ServersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{4}
}

type BalancerGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BalancerGetResponse) Reset() {
	*x = BalancerGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerGetResponse) ProtoMessage() {}

func (x *BalancerGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerGetResponse.ProtoReflect.Descriptor instead.
func (*BalancerGetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{5}
}

func (x *BalancerGetResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type BalancerSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SavedTo int32  `protobuf:"varint,2,opt,name=savedTo,proto3" json:"savedTo,omitempty"`
}

func (x *BalancerSetResponse) Reset() {
	*x = BalancerSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerSetResponse) ProtoMessage() {}

func (x *BalancerSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerSetResponse.ProtoReflect.Descriptor instead.
func (*BalancerSetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{6}
}

func (x *BalancerSetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BalancerSetResponse) GetSavedTo() int32 {
	if x != nil {
		return x.SavedTo
	}
	return 0
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ServerNumber int32  `protobuf:"varint,2,opt,name=serverNumber,proto3" json:"serverNumber,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{7}
}

func (x *ConnectResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ConnectResponse) GetServerNumber() int32 {
	if x != nil {
		return x.ServerNumber
	}
	return 0
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{8}
}

type ServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServersInfo string `protobuf:"bytes,1,opt,name=serversInfo,proto3" json:"serversInfo,omitempty"`
}

func (x *ServersResponse) Reset() {
	*x = ServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_balancer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersResponse) ProtoMessage() {}

func (x *ServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_balancer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersResponse.ProtoReflect.Descriptor instead.
func (*ServersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_balancer_proto_rawDescGZIP(), []int{9}
}

func (x *ServersResponse) GetServersInfo() string {
	if x != nil {
		return x.ServersInfo
	}
	return ""
}

var File_api_proto_balancer_proto protoreflect.FileDescriptor

var file_api_proto_balancer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22,
	0x54, 0x0a, 0x12, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x12, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x37, 0x0a,
	0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x61, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x22, 0x4d,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb3, 0x02, 0x0a, 0x08, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_balancer_proto_rawDescOnce sync.Once
	file_api_proto_balancer_proto_rawDescData = file_api_proto_balancer_proto_rawDesc
)

func file_api_proto_balancer_proto_rawDescGZIP() []byte {
	file_api_proto_balancer_proto_rawDescOnce.Do(func() {
		file_api_proto_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_balancer_proto_rawDescData)
	})
	return file_api_proto_balancer_proto_rawDescData
}

var file_api_proto_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_balancer_proto_goTypes = []interface{}{
	(*BalancerSetRequest)(nil),  // 0: api.BalancerSetRequest
	(*BalancerGetRequest)(nil),  // 1: api.BalancerGetRequest
	(*ConnectRequest)(nil),      // 2: api.ConnectRequest
	(*DisconnectRequest)(nil),   // 3: api.DisconnectRequest
	(*ServersRequest)(nil),      // 4: api.ServersRequest
	(*BalancerGetResponse)(nil), // 5: api.BalancerGetResponse
	(*BalancerSetResponse)(nil), // 6: api.BalancerSetResponse
	(*ConnectResponse)(nil),     // 7: api.ConnectResponse
	(*DisconnectResponse)(nil),  // 8: api.DisconnectResponse
	(*ServersResponse)(nil),     // 9: api.ServersResponse
}
var file_api_proto_balancer_proto_depIdxs = []int32{
	0, // 0: api.Balancer.Set:input_type -> api.BalancerSetRequest
	1, // 1: api.Balancer.Get:input_type -> api.BalancerGetRequest
	2, // 2: api.Balancer.Connect:input_type -> api.ConnectRequest
	3, // 3: api.Balancer.Disconnect:input_type -> api.DisconnectRequest
	4, // 4: api.Balancer.Servers:input_type -> api.ServersRequest
	6, // 5: api.Balancer.Set:output_type -> api.BalancerSetResponse
	5, // 6: api.Balancer.Get:output_type -> api.BalancerGetResponse
	7, // 7: api.Balancer.Connect:output_type -> api.ConnectResponse
	8, // 8: api.Balancer.Disconnect:output_type -> api.DisconnectResponse
	9, // 9: api.Balancer.Servers:output_type -> api.ServersResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_balancer_proto_init() }
func file_api_proto_balancer_proto_init() {
	if File_api_proto_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_balancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerSetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerGetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerSetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_balancer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_balancer_proto_goTypes,
		DependencyIndexes: file_api_proto_balancer_proto_depIdxs,
		MessageInfos:      file_api_proto_balancer_proto_msgTypes,
	}.Build()
	File_api_proto_balancer_proto = out.File
	file_api_proto_balancer_proto_rawDesc = nil
	file_api_proto_balancer_proto_goTypes = nil
	file_api_proto_balancer_proto_depIdxs = nil
}
