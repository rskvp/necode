// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/gateway/resources.proto

package gateway

import (
	provider "github.com/rskvp/necode/proto/gen/necode/storage/provider"
	types "github.com/rskvp/necode/proto/gen/necode/types"
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

// A file upload protocol object stores information about
// uploading resources using a specific protocol.
type FileUploadProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The protocol to be followed.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// REQUIRED.
	// The endpoint where to upload the data.
	// The value MUST be a Uniform Resource Identifier (URI)
	// as specified in RFC 3986.
	UploadEndpoint string `protobuf:"bytes,3,opt,name=upload_endpoint,json=uploadEndpoint,proto3" json:"upload_endpoint,omitempty"`
	// REQUIRED.
	// List of available checksums
	// the client can use when sending
	// the file.
	AvailableChecksums []*provider.ResourceChecksumPriority `protobuf:"bytes,4,rep,name=available_checksums,json=availableChecksums,proto3" json:"available_checksums,omitempty"`
	// OPTIONAL.
	// A token that MUST be validated by the data gateway for the upload.
	// Only makes sense for uploads passing through the data gateway.
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *FileUploadProtocol) Reset() {
	*x = FileUploadProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_gateway_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadProtocol) ProtoMessage() {}

func (x *FileUploadProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_necode_gateway_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadProtocol.ProtoReflect.Descriptor instead.
func (*FileUploadProtocol) Descriptor() ([]byte, []int) {
	return file_necode_gateway_resources_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadProtocol) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *FileUploadProtocol) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *FileUploadProtocol) GetUploadEndpoint() string {
	if x != nil {
		return x.UploadEndpoint
	}
	return ""
}

func (x *FileUploadProtocol) GetAvailableChecksums() []*provider.ResourceChecksumPriority {
	if x != nil {
		return x.AvailableChecksums
	}
	return nil
}

func (x *FileUploadProtocol) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// A file download protocol object stores information about
// downloading resources using a specific protocol.
type FileDownloadProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The protocol to be followed.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// REQUIRED.
	// The endpoint where to download the data.
	// The value MUST be a Uniform Resource Identifier (URI)
	// as specified in RFC 3986.
	DownloadEndpoint string `protobuf:"bytes,3,opt,name=download_endpoint,json=downloadEndpoint,proto3" json:"download_endpoint,omitempty"`
	// OPTIONAL.
	// A token that MUST be validated by the data gateway for the download.
	// Only makes sense for downloads passing through the data gateway.
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *FileDownloadProtocol) Reset() {
	*x = FileDownloadProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_gateway_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownloadProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadProtocol) ProtoMessage() {}

func (x *FileDownloadProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_necode_gateway_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadProtocol.ProtoReflect.Descriptor instead.
func (*FileDownloadProtocol) Descriptor() ([]byte, []int) {
	return file_necode_gateway_resources_proto_rawDescGZIP(), []int{1}
}

func (x *FileDownloadProtocol) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *FileDownloadProtocol) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *FileDownloadProtocol) GetDownloadEndpoint() string {
	if x != nil {
		return x.DownloadEndpoint
	}
	return ""
}

func (x *FileDownloadProtocol) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_necode_gateway_resources_proto protoreflect.FileDescriptor

var file_necode_gateway_resources_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x1a, 0x27, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x62, 0x0a,
	0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x12, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0xaf, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xa2, 0x02, 0x03, 0x4e, 0x47, 0x58, 0xaa, 0x02,
	0x0e, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xca,
	0x02, 0x0e, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0xe2, 0x02, 0x1a, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f,
	0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_gateway_resources_proto_rawDescOnce sync.Once
	file_necode_gateway_resources_proto_rawDescData = file_necode_gateway_resources_proto_rawDesc
)

func file_necode_gateway_resources_proto_rawDescGZIP() []byte {
	file_necode_gateway_resources_proto_rawDescOnce.Do(func() {
		file_necode_gateway_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_gateway_resources_proto_rawDescData)
	})
	return file_necode_gateway_resources_proto_rawDescData
}

var file_necode_gateway_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_necode_gateway_resources_proto_goTypes = []interface{}{
	(*FileUploadProtocol)(nil),                // 0: necode.gateway.FileUploadProtocol
	(*FileDownloadProtocol)(nil),              // 1: necode.gateway.FileDownloadProtocol
	(*types.Opaque)(nil),                      // 2: necode.types.Opaque
	(*provider.ResourceChecksumPriority)(nil), // 3: necode.storage.provider.ResourceChecksumPriority
}
var file_necode_gateway_resources_proto_depIdxs = []int32{
	2, // 0: necode.gateway.FileUploadProtocol.opaque:type_name -> necode.types.Opaque
	3, // 1: necode.gateway.FileUploadProtocol.available_checksums:type_name -> necode.storage.provider.ResourceChecksumPriority
	2, // 2: necode.gateway.FileDownloadProtocol.opaque:type_name -> necode.types.Opaque
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_necode_gateway_resources_proto_init() }
func file_necode_gateway_resources_proto_init() {
	if File_necode_gateway_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_necode_gateway_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadProtocol); i {
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
		file_necode_gateway_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownloadProtocol); i {
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
			RawDescriptor: file_necode_gateway_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_necode_gateway_resources_proto_goTypes,
		DependencyIndexes: file_necode_gateway_resources_proto_depIdxs,
		MessageInfos:      file_necode_gateway_resources_proto_msgTypes,
	}.Build()
	File_necode_gateway_resources_proto = out.File
	file_necode_gateway_resources_proto_rawDesc = nil
	file_necode_gateway_resources_proto_goTypes = nil
	file_necode_gateway_resources_proto_depIdxs = nil
}