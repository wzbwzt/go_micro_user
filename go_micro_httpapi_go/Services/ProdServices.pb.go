// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.0
// source: ProdServices.proto

package Services

import (
	proto "github.com/golang/protobuf/proto"
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

type ProdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag:form:"size"
	Size int32 `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty" form:"size"`
}

func (x *ProdReq) Reset() {
	*x = ProdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdServices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdReq) ProtoMessage() {}

func (x *ProdReq) ProtoReflect() protoreflect.Message {
	mi := &file_ProdServices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdReq.ProtoReflect.Descriptor instead.
func (*ProdReq) Descriptor() ([]byte, []int) {
	return file_ProdServices_proto_rawDescGZIP(), []int{0}
}

func (x *ProdReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ProdRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProdModel `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ProdRsp) Reset() {
	*x = ProdRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdServices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRsp) ProtoMessage() {}

func (x *ProdRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ProdServices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRsp.ProtoReflect.Descriptor instead.
func (*ProdRsp) Descriptor() ([]byte, []int) {
	return file_ProdServices_proto_rawDescGZIP(), []int{1}
}

func (x *ProdRsp) GetData() []*ProdModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ProdServices_proto protoreflect.FileDescriptor

var file_ProdServices_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0b,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x32, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x52, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x42,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52,
	0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProdServices_proto_rawDescOnce sync.Once
	file_ProdServices_proto_rawDescData = file_ProdServices_proto_rawDesc
)

func file_ProdServices_proto_rawDescGZIP() []byte {
	file_ProdServices_proto_rawDescOnce.Do(func() {
		file_ProdServices_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProdServices_proto_rawDescData)
	})
	return file_ProdServices_proto_rawDescData
}

var file_ProdServices_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProdServices_proto_goTypes = []interface{}{
	(*ProdReq)(nil),   // 0: Services.ProdReq
	(*ProdRsp)(nil),   // 1: Services.ProdRsp
	(*ProdModel)(nil), // 2: Services.ProdModel
}
var file_ProdServices_proto_depIdxs = []int32{
	2, // 0: Services.ProdRsp.Data:type_name -> Services.ProdModel
	0, // 1: Services.ProdService.GetProdList:input_type -> Services.ProdReq
	1, // 2: Services.ProdService.GetProdList:output_type -> Services.ProdRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProdServices_proto_init() }
func file_ProdServices_proto_init() {
	if File_ProdServices_proto != nil {
		return
	}
	file_Model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ProdServices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdReq); i {
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
		file_ProdServices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRsp); i {
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
			RawDescriptor: file_ProdServices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProdServices_proto_goTypes,
		DependencyIndexes: file_ProdServices_proto_depIdxs,
		MessageInfos:      file_ProdServices_proto_msgTypes,
	}.Build()
	File_ProdServices_proto = out.File
	file_ProdServices_proto_rawDesc = nil
	file_ProdServices_proto_goTypes = nil
	file_ProdServices_proto_depIdxs = nil
}
