// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/ocm/provider/provider_api.proto

package provider

import (
	rpc "github.com/rskvp/necode/proto/gen/necode/rpc"
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

type IsProviderAllowedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The provider that we need to check against the list of verified mesh providers.
	Provider *ProviderInfo `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *IsProviderAllowedRequest) Reset() {
	*x = IsProviderAllowedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsProviderAllowedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsProviderAllowedRequest) ProtoMessage() {}

func (x *IsProviderAllowedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsProviderAllowedRequest.ProtoReflect.Descriptor instead.
func (*IsProviderAllowedRequest) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{0}
}

func (x *IsProviderAllowedRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *IsProviderAllowedRequest) GetProvider() *ProviderInfo {
	if x != nil {
		return x.Provider
	}
	return nil
}

type IsProviderAllowedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *IsProviderAllowedResponse) Reset() {
	*x = IsProviderAllowedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsProviderAllowedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsProviderAllowedResponse) ProtoMessage() {}

func (x *IsProviderAllowedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsProviderAllowedResponse.ProtoReflect.Descriptor instead.
func (*IsProviderAllowedResponse) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{1}
}

func (x *IsProviderAllowedResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IsProviderAllowedResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type GetInfoByDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The domain of the system provider.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *GetInfoByDomainRequest) Reset() {
	*x = GetInfoByDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoByDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoByDomainRequest) ProtoMessage() {}

func (x *GetInfoByDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoByDomainRequest.ProtoReflect.Descriptor instead.
func (*GetInfoByDomainRequest) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoByDomainRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetInfoByDomainRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type GetInfoByDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The info of the provider
	ProviderInfo *ProviderInfo `protobuf:"bytes,3,opt,name=provider_info,json=providerInfo,proto3" json:"provider_info,omitempty"`
}

func (x *GetInfoByDomainResponse) Reset() {
	*x = GetInfoByDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoByDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoByDomainResponse) ProtoMessage() {}

func (x *GetInfoByDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoByDomainResponse.ProtoReflect.Descriptor instead.
func (*GetInfoByDomainResponse) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoByDomainResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetInfoByDomainResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetInfoByDomainResponse) GetProviderInfo() *ProviderInfo {
	if x != nil {
		return x.ProviderInfo
	}
	return nil
}

type ListAllProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *ListAllProvidersRequest) Reset() {
	*x = ListAllProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllProvidersRequest) ProtoMessage() {}

func (x *ListAllProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListAllProvidersRequest) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllProvidersRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type ListAllProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The share.
	Providers []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *ListAllProvidersResponse) Reset() {
	*x = ListAllProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllProvidersResponse) ProtoMessage() {}

func (x *ListAllProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_ocm_provider_provider_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListAllProvidersResponse) Descriptor() ([]byte, []int) {
	return file_necode_ocm_provider_provider_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListAllProvidersResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListAllProvidersResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *ListAllProvidersResponse) GetProviders() []*ProviderInfo {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_necode_ocm_provider_provider_api_proto protoreflect.FileDescriptor

var file_necode_ocm_provider_provider_api_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x6f, 0x63, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x23, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x6f, 0x63, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x18, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22,
	0x75, 0x0a, 0x19, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xb5, 0x01,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x32, 0xe0, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x72, 0x0a, 0x11, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x2d, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd0, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x42, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2f, 0x6f, 0x63, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xa2,
	0x02, 0x03, 0x4e, 0x4f, 0x50, 0xaa, 0x02, 0x13, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x4f,
	0x63, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xca, 0x02, 0x13, 0x4e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x5c, 0x4f, 0x63, 0x6d, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0xe2, 0x02, 0x1f, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x4f, 0x63, 0x6d, 0x5c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x4f, 0x63,
	0x6d, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_necode_ocm_provider_provider_api_proto_rawDescOnce sync.Once
	file_necode_ocm_provider_provider_api_proto_rawDescData = file_necode_ocm_provider_provider_api_proto_rawDesc
)

func file_necode_ocm_provider_provider_api_proto_rawDescGZIP() []byte {
	file_necode_ocm_provider_provider_api_proto_rawDescOnce.Do(func() {
		file_necode_ocm_provider_provider_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_ocm_provider_provider_api_proto_rawDescData)
	})
	return file_necode_ocm_provider_provider_api_proto_rawDescData
}

var file_necode_ocm_provider_provider_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_necode_ocm_provider_provider_api_proto_goTypes = []interface{}{
	(*IsProviderAllowedRequest)(nil),  // 0: necode.ocm.provider.IsProviderAllowedRequest
	(*IsProviderAllowedResponse)(nil), // 1: necode.ocm.provider.IsProviderAllowedResponse
	(*GetInfoByDomainRequest)(nil),    // 2: necode.ocm.provider.GetInfoByDomainRequest
	(*GetInfoByDomainResponse)(nil),   // 3: necode.ocm.provider.GetInfoByDomainResponse
	(*ListAllProvidersRequest)(nil),   // 4: necode.ocm.provider.ListAllProvidersRequest
	(*ListAllProvidersResponse)(nil),  // 5: necode.ocm.provider.ListAllProvidersResponse
	(*types.Opaque)(nil),              // 6: necode.types.Opaque
	(*ProviderInfo)(nil),              // 7: necode.ocm.provider.ProviderInfo
	(*rpc.Status)(nil),                // 8: necode.rpc.Status
}
var file_necode_ocm_provider_provider_api_proto_depIdxs = []int32{
	6,  // 0: necode.ocm.provider.IsProviderAllowedRequest.opaque:type_name -> necode.types.Opaque
	7,  // 1: necode.ocm.provider.IsProviderAllowedRequest.provider:type_name -> necode.ocm.provider.ProviderInfo
	8,  // 2: necode.ocm.provider.IsProviderAllowedResponse.status:type_name -> necode.rpc.Status
	6,  // 3: necode.ocm.provider.IsProviderAllowedResponse.opaque:type_name -> necode.types.Opaque
	6,  // 4: necode.ocm.provider.GetInfoByDomainRequest.opaque:type_name -> necode.types.Opaque
	8,  // 5: necode.ocm.provider.GetInfoByDomainResponse.status:type_name -> necode.rpc.Status
	6,  // 6: necode.ocm.provider.GetInfoByDomainResponse.opaque:type_name -> necode.types.Opaque
	7,  // 7: necode.ocm.provider.GetInfoByDomainResponse.provider_info:type_name -> necode.ocm.provider.ProviderInfo
	6,  // 8: necode.ocm.provider.ListAllProvidersRequest.opaque:type_name -> necode.types.Opaque
	8,  // 9: necode.ocm.provider.ListAllProvidersResponse.status:type_name -> necode.rpc.Status
	6,  // 10: necode.ocm.provider.ListAllProvidersResponse.opaque:type_name -> necode.types.Opaque
	7,  // 11: necode.ocm.provider.ListAllProvidersResponse.providers:type_name -> necode.ocm.provider.ProviderInfo
	0,  // 12: necode.ocm.provider.ProviderAPI.IsProviderAllowed:input_type -> necode.ocm.provider.IsProviderAllowedRequest
	2,  // 13: necode.ocm.provider.ProviderAPI.GetInfoByDomain:input_type -> necode.ocm.provider.GetInfoByDomainRequest
	4,  // 14: necode.ocm.provider.ProviderAPI.ListAllProviders:input_type -> necode.ocm.provider.ListAllProvidersRequest
	1,  // 15: necode.ocm.provider.ProviderAPI.IsProviderAllowed:output_type -> necode.ocm.provider.IsProviderAllowedResponse
	3,  // 16: necode.ocm.provider.ProviderAPI.GetInfoByDomain:output_type -> necode.ocm.provider.GetInfoByDomainResponse
	5,  // 17: necode.ocm.provider.ProviderAPI.ListAllProviders:output_type -> necode.ocm.provider.ListAllProvidersResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_necode_ocm_provider_provider_api_proto_init() }
func file_necode_ocm_provider_provider_api_proto_init() {
	if File_necode_ocm_provider_provider_api_proto != nil {
		return
	}
	file_necode_ocm_provider_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_necode_ocm_provider_provider_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsProviderAllowedRequest); i {
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
		file_necode_ocm_provider_provider_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsProviderAllowedResponse); i {
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
		file_necode_ocm_provider_provider_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoByDomainRequest); i {
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
		file_necode_ocm_provider_provider_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoByDomainResponse); i {
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
		file_necode_ocm_provider_provider_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllProvidersRequest); i {
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
		file_necode_ocm_provider_provider_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllProvidersResponse); i {
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
			RawDescriptor: file_necode_ocm_provider_provider_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_necode_ocm_provider_provider_api_proto_goTypes,
		DependencyIndexes: file_necode_ocm_provider_provider_api_proto_depIdxs,
		MessageInfos:      file_necode_ocm_provider_provider_api_proto_msgTypes,
	}.Build()
	File_necode_ocm_provider_provider_api_proto = out.File
	file_necode_ocm_provider_provider_api_proto_rawDesc = nil
	file_necode_ocm_provider_provider_api_proto_goTypes = nil
	file_necode_ocm_provider_provider_api_proto_depIdxs = nil
}
