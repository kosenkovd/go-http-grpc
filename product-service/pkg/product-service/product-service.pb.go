// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/product-service/product-service.proto

package product_service

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId uint64 `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_service_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_service_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_api_product_service_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type CreateProductV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId uint64 `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateProductV1Request) Reset() {
	*x = CreateProductV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_service_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductV1Request) ProtoMessage() {}

func (x *CreateProductV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_service_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductV1Request.ProtoReflect.Descriptor instead.
func (*CreateProductV1Request) Descriptor() ([]byte, []int) {
	return file_api_product_service_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductV1Request) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type CreateProductV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Product `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateProductV1Response) Reset() {
	*x = CreateProductV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_service_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductV1Response) ProtoMessage() {}

func (x *CreateProductV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_service_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductV1Response.ProtoReflect.Descriptor instead.
func (*CreateProductV1Response) Descriptor() ([]byte, []int) {
	return file_api_product_service_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductV1Response) GetResult() *Product {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_product_service_product_service_proto protoreflect.FileDescriptor

var file_api_product_service_product_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x6b, 0x6f, 0x73,
	0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x61, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x02, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd1, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xbe, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x56, 0x31, 0x12, 0x44, 0x2e, 0x6b, 0x6f, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x6b, 0x6f, 0x73,
	0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_product_service_product_service_proto_rawDescOnce sync.Once
	file_api_product_service_product_service_proto_rawDescData = file_api_product_service_product_service_proto_rawDesc
)

func file_api_product_service_product_service_proto_rawDescGZIP() []byte {
	file_api_product_service_product_service_proto_rawDescOnce.Do(func() {
		file_api_product_service_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_product_service_product_service_proto_rawDescData)
	})
	return file_api_product_service_product_service_proto_rawDescData
}

var file_api_product_service_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_product_service_product_service_proto_goTypes = []interface{}{
	(*Product)(nil),                 // 0: kosenkovd.product_service.product_service.v1.Product
	(*CreateProductV1Request)(nil),  // 1: kosenkovd.product_service.product_service.v1.CreateProductV1Request
	(*CreateProductV1Response)(nil), // 2: kosenkovd.product_service.product_service.v1.CreateProductV1Response
}
var file_api_product_service_product_service_proto_depIdxs = []int32{
	0, // 0: kosenkovd.product_service.product_service.v1.CreateProductV1Response.result:type_name -> kosenkovd.product_service.product_service.v1.Product
	1, // 1: kosenkovd.product_service.product_service.v1.ProductService.CreateProductV1:input_type -> kosenkovd.product_service.product_service.v1.CreateProductV1Request
	2, // 2: kosenkovd.product_service.product_service.v1.ProductService.CreateProductV1:output_type -> kosenkovd.product_service.product_service.v1.CreateProductV1Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_product_service_product_service_proto_init() }
func file_api_product_service_product_service_proto_init() {
	if File_api_product_service_product_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_product_service_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_api_product_service_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductV1Request); i {
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
		file_api_product_service_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductV1Response); i {
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
			RawDescriptor: file_api_product_service_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_product_service_product_service_proto_goTypes,
		DependencyIndexes: file_api_product_service_product_service_proto_depIdxs,
		MessageInfos:      file_api_product_service_product_service_proto_msgTypes,
	}.Build()
	File_api_product_service_product_service_proto = out.File
	file_api_product_service_product_service_proto_rawDesc = nil
	file_api_product_service_product_service_proto_goTypes = nil
	file_api_product_service_product_service_proto_depIdxs = nil
}
