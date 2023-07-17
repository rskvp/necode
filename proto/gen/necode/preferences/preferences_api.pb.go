// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/preferences/preferences_api.proto

package preferences

import (
	rpc "github.com/rskvp/necode/proto/gen/necode/rpc"
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

type SetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Key *PreferenceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// REQUIRED.
	// The value associated with the key.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *SetKeyRequest) Reset() {
	*x = SetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_preferences_preferences_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyRequest) ProtoMessage() {}

func (x *SetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_preferences_preferences_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyRequest.ProtoReflect.Descriptor instead.
func (*SetKeyRequest) Descriptor() ([]byte, []int) {
	return file_necode_preferences_preferences_api_proto_rawDescGZIP(), []int{0}
}

func (x *SetKeyRequest) GetKey() *PreferenceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SetKeyRequest) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type SetKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetKeyResponse) Reset() {
	*x = SetKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_preferences_preferences_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyResponse) ProtoMessage() {}

func (x *SetKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_preferences_preferences_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyResponse.ProtoReflect.Descriptor instead.
func (*SetKeyResponse) Descriptor() ([]byte, []int) {
	return file_necode_preferences_preferences_api_proto_rawDescGZIP(), []int{1}
}

func (x *SetKeyResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Key *PreferenceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKeyRequest) Reset() {
	*x = GetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_preferences_preferences_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyRequest) ProtoMessage() {}

func (x *GetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_preferences_preferences_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyRequest.ProtoReflect.Descriptor instead.
func (*GetKeyRequest) Descriptor() ([]byte, []int) {
	return file_necode_preferences_preferences_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetKeyRequest) GetKey() *PreferenceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// REQUIRED.
	// The value associated with the key.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *GetKeyResponse) Reset() {
	*x = GetKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_preferences_preferences_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyResponse) ProtoMessage() {}

func (x *GetKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_preferences_preferences_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyResponse.ProtoReflect.Descriptor instead.
func (*GetKeyResponse) Descriptor() ([]byte, []int) {
	return file_necode_preferences_preferences_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetKeyResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetKeyResponse) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_necode_preferences_preferences_api_proto protoreflect.FileDescriptor

var file_necode_preferences_preferences_api_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x22,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0d, 0x53,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x32, 0xb2, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x41, 0x50, 0x49, 0x12, 0x4f, 0x0a, 0x06, 0x53, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xcc, 0x01, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x13, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70,
	0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x4e, 0x50, 0x58, 0xaa, 0x02, 0x12, 0x4e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x12, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1e, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_necode_preferences_preferences_api_proto_rawDescOnce sync.Once
	file_necode_preferences_preferences_api_proto_rawDescData = file_necode_preferences_preferences_api_proto_rawDesc
)

func file_necode_preferences_preferences_api_proto_rawDescGZIP() []byte {
	file_necode_preferences_preferences_api_proto_rawDescOnce.Do(func() {
		file_necode_preferences_preferences_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_preferences_preferences_api_proto_rawDescData)
	})
	return file_necode_preferences_preferences_api_proto_rawDescData
}

var file_necode_preferences_preferences_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_necode_preferences_preferences_api_proto_goTypes = []interface{}{
	(*SetKeyRequest)(nil),  // 0: necode.preferences.SetKeyRequest
	(*SetKeyResponse)(nil), // 1: necode.preferences.SetKeyResponse
	(*GetKeyRequest)(nil),  // 2: necode.preferences.GetKeyRequest
	(*GetKeyResponse)(nil), // 3: necode.preferences.GetKeyResponse
	(*PreferenceKey)(nil),  // 4: necode.preferences.PreferenceKey
	(*rpc.Status)(nil),     // 5: necode.rpc.Status
}
var file_necode_preferences_preferences_api_proto_depIdxs = []int32{
	4, // 0: necode.preferences.SetKeyRequest.key:type_name -> necode.preferences.PreferenceKey
	5, // 1: necode.preferences.SetKeyResponse.status:type_name -> necode.rpc.Status
	4, // 2: necode.preferences.GetKeyRequest.key:type_name -> necode.preferences.PreferenceKey
	5, // 3: necode.preferences.GetKeyResponse.status:type_name -> necode.rpc.Status
	0, // 4: necode.preferences.PreferencesAPI.SetKey:input_type -> necode.preferences.SetKeyRequest
	2, // 5: necode.preferences.PreferencesAPI.GetKey:input_type -> necode.preferences.GetKeyRequest
	1, // 6: necode.preferences.PreferencesAPI.SetKey:output_type -> necode.preferences.SetKeyResponse
	3, // 7: necode.preferences.PreferencesAPI.GetKey:output_type -> necode.preferences.GetKeyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_necode_preferences_preferences_api_proto_init() }
func file_necode_preferences_preferences_api_proto_init() {
	if File_necode_preferences_preferences_api_proto != nil {
		return
	}
	file_necode_preferences_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_necode_preferences_preferences_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyRequest); i {
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
		file_necode_preferences_preferences_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyResponse); i {
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
		file_necode_preferences_preferences_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyRequest); i {
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
		file_necode_preferences_preferences_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyResponse); i {
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
			RawDescriptor: file_necode_preferences_preferences_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_necode_preferences_preferences_api_proto_goTypes,
		DependencyIndexes: file_necode_preferences_preferences_api_proto_depIdxs,
		MessageInfos:      file_necode_preferences_preferences_api_proto_msgTypes,
	}.Build()
	File_necode_preferences_preferences_api_proto = out.File
	file_necode_preferences_preferences_api_proto_rawDesc = nil
	file_necode_preferences_preferences_api_proto_goTypes = nil
	file_necode_preferences_preferences_api_proto_depIdxs = nil
}