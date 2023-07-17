// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/preferences/resources.proto

package preferences

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

// Represents a key object consisting of a namespace and a string key.
type PreferenceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The namespace to which the key belongs.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// REQUIRED.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PreferenceKey) Reset() {
	*x = PreferenceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_preferences_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreferenceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreferenceKey) ProtoMessage() {}

func (x *PreferenceKey) ProtoReflect() protoreflect.Message {
	mi := &file_necode_preferences_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreferenceKey.ProtoReflect.Descriptor instead.
func (*PreferenceKey) Descriptor() ([]byte, []int) {
	return file_necode_preferences_resources_proto_rawDescGZIP(), []int{0}
}

func (x *PreferenceKey) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PreferenceKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_necode_preferences_resources_proto protoreflect.FileDescriptor

var file_necode_preferences_resources_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x4e,
	0x50, 0x58, 0xaa, 0x02, 0x12, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0xca, 0x02, 0x12, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x5c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1e, 0x4e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_preferences_resources_proto_rawDescOnce sync.Once
	file_necode_preferences_resources_proto_rawDescData = file_necode_preferences_resources_proto_rawDesc
)

func file_necode_preferences_resources_proto_rawDescGZIP() []byte {
	file_necode_preferences_resources_proto_rawDescOnce.Do(func() {
		file_necode_preferences_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_preferences_resources_proto_rawDescData)
	})
	return file_necode_preferences_resources_proto_rawDescData
}

var file_necode_preferences_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_necode_preferences_resources_proto_goTypes = []interface{}{
	(*PreferenceKey)(nil), // 0: necode.preferences.PreferenceKey
}
var file_necode_preferences_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_necode_preferences_resources_proto_init() }
func file_necode_preferences_resources_proto_init() {
	if File_necode_preferences_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_necode_preferences_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreferenceKey); i {
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
			RawDescriptor: file_necode_preferences_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_necode_preferences_resources_proto_goTypes,
		DependencyIndexes: file_necode_preferences_resources_proto_depIdxs,
		MessageInfos:      file_necode_preferences_resources_proto_msgTypes,
	}.Build()
	File_necode_preferences_resources_proto = out.File
	file_necode_preferences_resources_proto_rawDesc = nil
	file_necode_preferences_resources_proto_goTypes = nil
	file_necode_preferences_resources_proto_depIdxs = nil
}