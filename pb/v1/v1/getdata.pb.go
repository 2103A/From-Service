// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: getdata.proto

package v1

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

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string `protobuf:"bytes,1,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_getdata_proto_rawDescGZIP(), []int{0}
}

func (x *GetDataRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

type GetDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string `protobuf:"bytes,1,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
}

func (x *GetDataReply) Reset() {
	*x = GetDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataReply) ProtoMessage() {}

func (x *GetDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_getdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataReply.ProtoReflect.Descriptor instead.
func (*GetDataReply) Descriptor() ([]byte, []int) {
	return file_getdata_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataReply) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

var File_getdata_proto protoreflect.FileDescriptor

var file_getdata_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x48, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x65,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x4f, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_getdata_proto_rawDescOnce sync.Once
	file_getdata_proto_rawDescData = file_getdata_proto_rawDesc
)

func file_getdata_proto_rawDescGZIP() []byte {
	file_getdata_proto_rawDescOnce.Do(func() {
		file_getdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_getdata_proto_rawDescData)
	})
	return file_getdata_proto_rawDescData
}

var file_getdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_getdata_proto_goTypes = []interface{}{
	(*GetDataRequest)(nil), // 0: getdata.v1.GetDataRequest
	(*GetDataReply)(nil),   // 1: getdata.v1.GetDataReply
}
var file_getdata_proto_depIdxs = []int32{
	0, // 0: getdata.v1.Get.GetData:input_type -> getdata.v1.GetDataRequest
	1, // 1: getdata.v1.Get.GetData:output_type -> getdata.v1.GetDataReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_getdata_proto_init() }
func file_getdata_proto_init() {
	if File_getdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_getdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataRequest); i {
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
		file_getdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataReply); i {
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
			RawDescriptor: file_getdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_getdata_proto_goTypes,
		DependencyIndexes: file_getdata_proto_depIdxs,
		MessageInfos:      file_getdata_proto_msgTypes,
	}.Build()
	File_getdata_proto = out.File
	file_getdata_proto_rawDesc = nil
	file_getdata_proto_goTypes = nil
	file_getdata_proto_depIdxs = nil
}
