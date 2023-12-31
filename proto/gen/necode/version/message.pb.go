// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/version/message.proto

package version

import (
	enums "github.com/rskvp/necode/proto/gen/necode/enums"
	_ "github.com/rskvp/necode/proto/gen/necode/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ReleaseInfo contains information about specific version of temporal.
type ReleaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	Notes       string                 `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ReleaseInfo) Reset() {
	*x = ReleaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_version_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseInfo) ProtoMessage() {}

func (x *ReleaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_necode_version_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseInfo.ProtoReflect.Descriptor instead.
func (*ReleaseInfo) Descriptor() ([]byte, []int) {
	return file_necode_version_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReleaseInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ReleaseInfo) GetReleaseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseTime
	}
	return nil
}

func (x *ReleaseInfo) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

// Alert contains notification and severity.
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Severity enums.Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=necode.enums.Severity" json:"severity,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_version_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_necode_version_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_necode_version_message_proto_rawDescGZIP(), []int{1}
}

func (x *Alert) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Alert) GetSeverity() enums.Severity {
	if x != nil {
		return x.Severity
	}
	return enums.Severity(0)
}

// VersionInfo contains details about current and recommended release versions as well as alerts and upgrade instructions.
type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current        *ReleaseInfo           `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Recommended    *ReleaseInfo           `protobuf:"bytes,2,opt,name=recommended,proto3" json:"recommended,omitempty"`
	Instructions   string                 `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	Alerts         []*Alert               `protobuf:"bytes,4,rep,name=alerts,proto3" json:"alerts,omitempty"`
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_version_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_necode_version_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_necode_version_message_proto_rawDescGZIP(), []int{2}
}

func (x *VersionInfo) GetCurrent() *ReleaseInfo {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *VersionInfo) GetRecommended() *ReleaseInfo {
	if x != nil {
		return x.Recommended
	}
	return nil
}

func (x *VersionInfo) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *VersionInfo) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *VersionInfo) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

var File_necode_version_message_proto protoreflect.FileDescriptor

var file_necode_version_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x05,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x10,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0xad, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76,
	0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x4e, 0x56, 0x58, 0xaa, 0x02, 0x0e, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0e, 0x4e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x5c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1a, 0x4e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x5c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a,
	0x3a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_version_message_proto_rawDescOnce sync.Once
	file_necode_version_message_proto_rawDescData = file_necode_version_message_proto_rawDesc
)

func file_necode_version_message_proto_rawDescGZIP() []byte {
	file_necode_version_message_proto_rawDescOnce.Do(func() {
		file_necode_version_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_version_message_proto_rawDescData)
	})
	return file_necode_version_message_proto_rawDescData
}

var file_necode_version_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_necode_version_message_proto_goTypes = []interface{}{
	(*ReleaseInfo)(nil),           // 0: necode.version.ReleaseInfo
	(*Alert)(nil),                 // 1: necode.version.Alert
	(*VersionInfo)(nil),           // 2: necode.version.VersionInfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(enums.Severity)(0),           // 4: necode.enums.Severity
}
var file_necode_version_message_proto_depIdxs = []int32{
	3, // 0: necode.version.ReleaseInfo.release_time:type_name -> google.protobuf.Timestamp
	4, // 1: necode.version.Alert.severity:type_name -> necode.enums.Severity
	0, // 2: necode.version.VersionInfo.current:type_name -> necode.version.ReleaseInfo
	0, // 3: necode.version.VersionInfo.recommended:type_name -> necode.version.ReleaseInfo
	1, // 4: necode.version.VersionInfo.alerts:type_name -> necode.version.Alert
	3, // 5: necode.version.VersionInfo.last_update_time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_necode_version_message_proto_init() }
func file_necode_version_message_proto_init() {
	if File_necode_version_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_necode_version_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseInfo); i {
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
		file_necode_version_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_necode_version_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
			RawDescriptor: file_necode_version_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_necode_version_message_proto_goTypes,
		DependencyIndexes: file_necode_version_message_proto_depIdxs,
		MessageInfos:      file_necode_version_message_proto_msgTypes,
	}.Build()
	File_necode_version_message_proto = out.File
	file_necode_version_message_proto_rawDesc = nil
	file_necode_version_message_proto_goTypes = nil
	file_necode_version_message_proto_depIdxs = nil
}
