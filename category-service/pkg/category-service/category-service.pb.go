// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/product-service/product-service.proto

package category_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryMethodV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryMethodV1Request) Reset() {
	*x = CategoryMethodV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryMethodV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryMethodV1Request) ProtoMessage() {}

func (x *CategoryMethodV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryMethodV1Request.ProtoReflect.Descriptor instead.
func (*CategoryMethodV1Request) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryMethodV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CategoryMethodV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Template `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CategoryMethodV1Response) Reset() {
	*x = CategoryMethodV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryMethodV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryMethodV1Response) ProtoMessage() {}

func (x *CategoryMethodV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryMethodV1Response.ProtoReflect.Descriptor instead.
func (*CategoryMethodV1Response) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryMethodV1Response) GetValue() *Template {
	if x != nil {
		return x.Value
	}
	return nil
}

type GetCategoryByIdV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryByIdV1Request) Reset() {
	*x = GetCategoryByIdV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdV1Request) ProtoMessage() {}

func (x *GetCategoryByIdV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdV1Request.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdV1Request) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCategoryByIdV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCategoryByIdV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Template `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetCategoryByIdV1Response) Reset() {
	*x = GetCategoryByIdV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdV1Response) ProtoMessage() {}

func (x *GetCategoryByIdV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdV1Response.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdV1Response) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryByIdV1Response) GetValue() *Template {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_api_category_service_category_service_proto protoreflect.FileDescriptor

var file_api_category_service_category_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x6b,
	0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x17, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x18, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76,
	0x64, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x9a, 0x03, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbe, 0x01, 0x0a, 0x10, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56, 0x31, 0x12,
	0x47, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e,
	0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0xc5, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x56,
	0x31, 0x12, 0x48, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x6b, 0x6f,
	0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_category_service_category_service_proto_rawDescOnce sync.Once
	file_api_category_service_category_service_proto_rawDescData = file_api_category_service_category_service_proto_rawDesc
)

func file_api_category_service_category_service_proto_rawDescGZIP() []byte {
	file_api_category_service_category_service_proto_rawDescOnce.Do(func() {
		file_api_category_service_category_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_category_service_category_service_proto_rawDescData)
	})
	return file_api_category_service_category_service_proto_rawDescData
}

var file_api_category_service_category_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_category_service_category_service_proto_goTypes = []interface{}{
	(*Template)(nil),                  // 0: kosenkovd.category_service.category_service.v1.Template
	(*CategoryMethodV1Request)(nil),   // 1: kosenkovd.category_service.category_service.v1.CategoryMethodV1Request
	(*CategoryMethodV1Response)(nil),  // 2: kosenkovd.category_service.category_service.v1.CategoryMethodV1Response
	(*GetCategoryByIdV1Request)(nil),  // 3: kosenkovd.category_service.category_service.v1.GetCategoryByIdV1Request
	(*GetCategoryByIdV1Response)(nil), // 4: kosenkovd.category_service.category_service.v1.GetCategoryByIdV1Response
}
var file_api_category_service_category_service_proto_depIdxs = []int32{
	0, // 0: kosenkovd.category_service.category_service.v1.CategoryMethodV1Response.value:type_name -> kosenkovd.category_service.category_service.v1.Template
	0, // 1: kosenkovd.category_service.category_service.v1.GetCategoryByIdV1Response.value:type_name -> kosenkovd.category_service.category_service.v1.Template
	1, // 2: kosenkovd.category_service.category_service.v1.CategoryService.CategoryMethodV1:input_type -> kosenkovd.category_service.category_service.v1.CategoryMethodV1Request
	3, // 3: kosenkovd.category_service.category_service.v1.CategoryService.GetCategoryByIdV1:input_type -> kosenkovd.category_service.category_service.v1.GetCategoryByIdV1Request
	2, // 4: kosenkovd.category_service.category_service.v1.CategoryService.CategoryMethodV1:output_type -> kosenkovd.category_service.category_service.v1.CategoryMethodV1Response
	4, // 5: kosenkovd.category_service.category_service.v1.CategoryService.GetCategoryByIdV1:output_type -> kosenkovd.category_service.category_service.v1.GetCategoryByIdV1Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_category_service_category_service_proto_init() }
func file_api_category_service_category_service_proto_init() {
	if File_api_category_service_category_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_category_service_category_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_api_category_service_category_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryMethodV1Request); i {
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
		file_api_category_service_category_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryMethodV1Response); i {
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
		file_api_category_service_category_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdV1Request); i {
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
		file_api_category_service_category_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdV1Response); i {
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
			RawDescriptor: file_api_category_service_category_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_category_service_category_service_proto_goTypes,
		DependencyIndexes: file_api_category_service_category_service_proto_depIdxs,
		MessageInfos:      file_api_category_service_category_service_proto_msgTypes,
	}.Build()
	File_api_category_service_category_service_proto = out.File
	file_api_category_service_category_service_proto_rawDesc = nil
	file_api_category_service_category_service_proto_goTypes = nil
	file_api_category_service_category_service_proto_depIdxs = nil
}
