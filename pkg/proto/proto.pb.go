// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: proto.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileState int32

const (
	FileState_FILE_RECEIVED FileState = 0
	FileState_FILE_FAILED   FileState = 1
)

// Enum value maps for FileState.
var (
	FileState_name = map[int32]string{
		0: "FILE_RECEIVED",
		1: "FILE_FAILED",
	}
	FileState_value = map[string]int32{
		"FILE_RECEIVED": 0,
		"FILE_FAILED":   1,
	}
)

func (x FileState) Enum() *FileState {
	p := new(FileState)
	*p = x
	return p
}

func (x FileState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_enumTypes[0].Descriptor()
}

func (FileState) Type() protoreflect.EnumType {
	return &file_proto_proto_enumTypes[0]
}

func (x FileState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileState.Descriptor instead.
func (FileState) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

type ResourceState int32

const (
	ResourceState_RES_SUCCESS ResourceState = 0
	ResourceState_RES_PENDING ResourceState = 1
	ResourceState_RES_ERROR   ResourceState = 2
	ResourceState_RES_OTHER   ResourceState = 3
)

// Enum value maps for ResourceState.
var (
	ResourceState_name = map[int32]string{
		0: "RES_SUCCESS",
		1: "RES_PENDING",
		2: "RES_ERROR",
		3: "RES_OTHER",
	}
	ResourceState_value = map[string]int32{
		"RES_SUCCESS": 0,
		"RES_PENDING": 1,
		"RES_ERROR":   2,
		"RES_OTHER":   3,
	}
)

func (x ResourceState) Enum() *ResourceState {
	p := new(ResourceState)
	*p = x
	return p
}

func (x ResourceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_enumTypes[1].Descriptor()
}

func (ResourceState) Type() protoreflect.EnumType {
	return &file_proto_proto_enumTypes[1]
}

func (x ResourceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceState.Descriptor instead.
func (ResourceState) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Remain   int32  `protobuf:"varint,3,opt,name=remain,proto3" json:"remain,omitempty"`
	Compress string `protobuf:"bytes,4,opt,name=compress,proto3" json:"compress,omitempty"`
	Digest   string `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
	Data     string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetRemain() int32 {
	if x != nil {
		return x.Remain
	}
	return 0
}

func (x *File) GetCompress() string {
	if x != nil {
		return x.Compress
	}
	return ""
}

func (x *File) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *File) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type FileStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path  string    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	State FileState `protobuf:"varint,3,opt,name=state,proto3,enum=FileState" json:"state,omitempty"`
	Error string    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FileStatus) Reset() {
	*x = FileStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStatus) ProtoMessage() {}

func (x *FileStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStatus.ProtoReflect.Descriptor instead.
func (*FileStatus) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

func (x *FileStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileStatus) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileStatus) GetState() FileState {
	if x != nil {
		return x.State
	}
	return FileState_FILE_RECEIVED
}

func (x *FileStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResourceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   ResourceState `protobuf:"varint,1,opt,name=state,proto3,enum=ResourceState" json:"state,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResourceStatus) Reset() {
	*x = ResourceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatus) ProtoMessage() {}

func (x *ResourceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatus.ProtoReflect.Descriptor instead.
func (*ResourceStatus) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceStatus) GetState() ResourceState {
	if x != nil {
		return x.State
	}
	return ResourceState_RES_SUCCESS
}

func (x *ResourceStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeployStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Resources map[string]ResourceState `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=ResourceState"`
}

func (x *DeployStatus) Reset() {
	*x = DeployStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployStatus) ProtoMessage() {}

func (x *DeployStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployStatus.ProtoReflect.Descriptor instead.
func (*DeployStatus) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{3}
}

func (x *DeployStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeployStatus) GetResources() map[string]ResourceState {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{4}
}

func (x *Reply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x50, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x4c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x35, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2f, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45,
	0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x4f, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45,
	0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x45, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x45, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x52,
	0x45, 0x53, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x03, 0x32, 0x37, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x32, 0x32, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x0e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_proto_goTypes = []interface{}{
	(FileState)(0),         // 0: FileState
	(ResourceState)(0),     // 1: ResourceState
	(*File)(nil),           // 2: File
	(*FileStatus)(nil),     // 3: FileStatus
	(*ResourceStatus)(nil), // 4: ResourceStatus
	(*DeployStatus)(nil),   // 5: DeployStatus
	(*Reply)(nil),          // 6: Reply
	nil,                    // 7: DeployStatus.ResourcesEntry
}
var file_proto_proto_depIdxs = []int32{
	0, // 0: FileStatus.state:type_name -> FileState
	1, // 1: ResourceStatus.state:type_name -> ResourceState
	7, // 2: DeployStatus.resources:type_name -> DeployStatus.ResourcesEntry
	1, // 3: DeployStatus.ResourcesEntry.value:type_name -> ResourceState
	5, // 4: Server.UpdateDeployStatus:input_type -> DeployStatus
	2, // 5: Worker.SendDeployFile:input_type -> File
	6, // 6: Server.UpdateDeployStatus:output_type -> Reply
	3, // 7: Worker.SendDeployFile:output_type -> FileStatus
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStatus); i {
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
		file_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceStatus); i {
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
		file_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployStatus); i {
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
		file_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		EnumInfos:         file_proto_proto_enumTypes,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	UpdateDeployStatus(ctx context.Context, in *DeployStatus, opts ...grpc.CallOption) (*Reply, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) UpdateDeployStatus(ctx context.Context, in *DeployStatus, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Server/UpdateDeployStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	UpdateDeployStatus(context.Context, *DeployStatus) (*Reply, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) UpdateDeployStatus(context.Context, *DeployStatus) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployStatus not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_UpdateDeployStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).UpdateDeployStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/UpdateDeployStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).UpdateDeployStatus(ctx, req.(*DeployStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDeployStatus",
			Handler:    _Server_UpdateDeployStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	SendDeployFile(ctx context.Context, opts ...grpc.CallOption) (Worker_SendDeployFileClient, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) SendDeployFile(ctx context.Context, opts ...grpc.CallOption) (Worker_SendDeployFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[0], "/Worker/SendDeployFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerSendDeployFileClient{stream}
	return x, nil
}

type Worker_SendDeployFileClient interface {
	Send(*File) error
	CloseAndRecv() (*FileStatus, error)
	grpc.ClientStream
}

type workerSendDeployFileClient struct {
	grpc.ClientStream
}

func (x *workerSendDeployFileClient) Send(m *File) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerSendDeployFileClient) CloseAndRecv() (*FileStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	SendDeployFile(Worker_SendDeployFileServer) error
}

// UnimplementedWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) SendDeployFile(Worker_SendDeployFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SendDeployFile not implemented")
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_SendDeployFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).SendDeployFile(&workerSendDeployFileServer{stream})
}

type Worker_SendDeployFileServer interface {
	SendAndClose(*FileStatus) error
	Recv() (*File, error)
	grpc.ServerStream
}

type workerSendDeployFileServer struct {
	grpc.ServerStream
}

func (x *workerSendDeployFileServer) SendAndClose(m *FileStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerSendDeployFileServer) Recv() (*File, error) {
	m := new(File)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendDeployFile",
			Handler:       _Worker_SendDeployFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto.proto",
}
