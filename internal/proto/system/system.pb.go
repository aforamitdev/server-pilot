// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/proto/system/system.proto

package system_proto

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

type SystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SystemRequest) Reset() {
	*x = SystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_system_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRequest) ProtoMessage() {}

func (x *SystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_system_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRequest.ProtoReflect.Descriptor instead.
func (*SystemRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_system_system_proto_rawDescGZIP(), []int{0}
}

type SystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SystemResponse) Reset() {
	*x = SystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_system_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemResponse) ProtoMessage() {}

func (x *SystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_system_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemResponse.ProtoReflect.Descriptor instead.
func (*SystemResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_system_system_proto_rawDescGZIP(), []int{1}
}

func (x *SystemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_internal_proto_system_system_proto protoreflect.FileDescriptor

var file_internal_proto_system_system_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x38, 0x0a, 0x08, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x66, 0x6f, 0x72, 0x61, 0x6d, 0x69, 0x74, 0x64, 0x65, 0x76, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_system_system_proto_rawDescOnce sync.Once
	file_internal_proto_system_system_proto_rawDescData = file_internal_proto_system_system_proto_rawDesc
)

func file_internal_proto_system_system_proto_rawDescGZIP() []byte {
	file_internal_proto_system_system_proto_rawDescOnce.Do(func() {
		file_internal_proto_system_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_system_system_proto_rawDescData)
	})
	return file_internal_proto_system_system_proto_rawDescData
}

var file_internal_proto_system_system_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_system_system_proto_goTypes = []interface{}{
	(*SystemRequest)(nil),  // 0: SystemRequest
	(*SystemResponse)(nil), // 1: SystemResponse
}
var file_internal_proto_system_system_proto_depIdxs = []int32{
	0, // 0: Informer.GetSystem:input_type -> SystemRequest
	1, // 1: Informer.GetSystem:output_type -> SystemResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_system_system_proto_init() }
func file_internal_proto_system_system_proto_init() {
	if File_internal_proto_system_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_system_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRequest); i {
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
		file_internal_proto_system_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemResponse); i {
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
			RawDescriptor: file_internal_proto_system_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_system_system_proto_goTypes,
		DependencyIndexes: file_internal_proto_system_system_proto_depIdxs,
		MessageInfos:      file_internal_proto_system_system_proto_msgTypes,
	}.Build()
	File_internal_proto_system_system_proto = out.File
	file_internal_proto_system_system_proto_rawDesc = nil
	file_internal_proto_system_system_proto_goTypes = nil
	file_internal_proto_system_system_proto_depIdxs = nil
}