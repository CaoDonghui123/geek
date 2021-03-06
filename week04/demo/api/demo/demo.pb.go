// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/demo/demo.proto

package demo

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

type CreateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDemoRequest) Reset() {
	*x = CreateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoRequest) ProtoMessage() {}

func (x *CreateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoRequest.ProtoReflect.Descriptor instead.
func (*CreateDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{0}
}

type CreateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDemoReply) Reset() {
	*x = CreateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoReply) ProtoMessage() {}

func (x *CreateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoReply.ProtoReflect.Descriptor instead.
func (*CreateDemoReply) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{1}
}

type UpdateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDemoRequest) Reset() {
	*x = UpdateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoRequest) ProtoMessage() {}

func (x *UpdateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoRequest.ProtoReflect.Descriptor instead.
func (*UpdateDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{2}
}

type UpdateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDemoReply) Reset() {
	*x = UpdateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoReply) ProtoMessage() {}

func (x *UpdateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoReply.ProtoReflect.Descriptor instead.
func (*UpdateDemoReply) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{3}
}

type DeleteDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDemoRequest) Reset() {
	*x = DeleteDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoRequest) ProtoMessage() {}

func (x *DeleteDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoRequest.ProtoReflect.Descriptor instead.
func (*DeleteDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{4}
}

type DeleteDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDemoReply) Reset() {
	*x = DeleteDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoReply) ProtoMessage() {}

func (x *DeleteDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoReply.ProtoReflect.Descriptor instead.
func (*DeleteDemoReply) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{5}
}

type GetDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDemoRequest) Reset() {
	*x = GetDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoRequest) ProtoMessage() {}

func (x *GetDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoRequest.ProtoReflect.Descriptor instead.
func (*GetDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{6}
}

type GetDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDemoReply) Reset() {
	*x = GetDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoReply) ProtoMessage() {}

func (x *GetDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoReply.ProtoReflect.Descriptor instead.
func (*GetDemoReply) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{7}
}

type ListDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDemoRequest) Reset() {
	*x = ListDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoRequest) ProtoMessage() {}

func (x *ListDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoRequest.ProtoReflect.Descriptor instead.
func (*ListDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{8}
}

type ListDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDemoReply) Reset() {
	*x = ListDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_demo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoReply) ProtoMessage() {}

func (x *ListDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_demo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoReply.ProtoReflect.Descriptor instead.
func (*ListDemoReply) Descriptor() ([]byte, []int) {
	return file_api_demo_demo_proto_rawDescGZIP(), []int{9}
}

var File_api_demo_demo_proto protoreflect.FileDescriptor

var file_api_demo_demo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x22,
	0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xd5, 0x02,
	0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f,
	0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x6d, 0x6f, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d,
	0x6f, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x20, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x50, 0x01, 0x5a, 0x12, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65,
	0x6d, 0x6f, 0x3b, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_demo_demo_proto_rawDescOnce sync.Once
	file_api_demo_demo_proto_rawDescData = file_api_demo_demo_proto_rawDesc
)

func file_api_demo_demo_proto_rawDescGZIP() []byte {
	file_api_demo_demo_proto_rawDescOnce.Do(func() {
		file_api_demo_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_demo_demo_proto_rawDescData)
	})
	return file_api_demo_demo_proto_rawDescData
}

var file_api_demo_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_demo_demo_proto_goTypes = []interface{}{
	(*CreateDemoRequest)(nil), // 0: api.demo.CreateDemoRequest
	(*CreateDemoReply)(nil),   // 1: api.demo.CreateDemoReply
	(*UpdateDemoRequest)(nil), // 2: api.demo.UpdateDemoRequest
	(*UpdateDemoReply)(nil),   // 3: api.demo.UpdateDemoReply
	(*DeleteDemoRequest)(nil), // 4: api.demo.DeleteDemoRequest
	(*DeleteDemoReply)(nil),   // 5: api.demo.DeleteDemoReply
	(*GetDemoRequest)(nil),    // 6: api.demo.GetDemoRequest
	(*GetDemoReply)(nil),      // 7: api.demo.GetDemoReply
	(*ListDemoRequest)(nil),   // 8: api.demo.ListDemoRequest
	(*ListDemoReply)(nil),     // 9: api.demo.ListDemoReply
}
var file_api_demo_demo_proto_depIdxs = []int32{
	0, // 0: api.demo.Demo.CreateDemo:input_type -> api.demo.CreateDemoRequest
	2, // 1: api.demo.Demo.UpdateDemo:input_type -> api.demo.UpdateDemoRequest
	4, // 2: api.demo.Demo.DeleteDemo:input_type -> api.demo.DeleteDemoRequest
	6, // 3: api.demo.Demo.GetDemo:input_type -> api.demo.GetDemoRequest
	8, // 4: api.demo.Demo.ListDemo:input_type -> api.demo.ListDemoRequest
	1, // 5: api.demo.Demo.CreateDemo:output_type -> api.demo.CreateDemoReply
	3, // 6: api.demo.Demo.UpdateDemo:output_type -> api.demo.UpdateDemoReply
	5, // 7: api.demo.Demo.DeleteDemo:output_type -> api.demo.DeleteDemoReply
	7, // 8: api.demo.Demo.GetDemo:output_type -> api.demo.GetDemoReply
	9, // 9: api.demo.Demo.ListDemo:output_type -> api.demo.ListDemoReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_demo_demo_proto_init() }
func file_api_demo_demo_proto_init() {
	if File_api_demo_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_demo_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDemoRequest); i {
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
		file_api_demo_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDemoReply); i {
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
		file_api_demo_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDemoRequest); i {
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
		file_api_demo_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDemoReply); i {
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
		file_api_demo_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDemoRequest); i {
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
		file_api_demo_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDemoReply); i {
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
		file_api_demo_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoRequest); i {
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
		file_api_demo_demo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoReply); i {
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
		file_api_demo_demo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemoRequest); i {
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
		file_api_demo_demo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemoReply); i {
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
			RawDescriptor: file_api_demo_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_demo_demo_proto_goTypes,
		DependencyIndexes: file_api_demo_demo_proto_depIdxs,
		MessageInfos:      file_api_demo_demo_proto_msgTypes,
	}.Build()
	File_api_demo_demo_proto = out.File
	file_api_demo_demo_proto_rawDesc = nil
	file_api_demo_demo_proto_goTypes = nil
	file_api_demo_demo_proto_depIdxs = nil
}
