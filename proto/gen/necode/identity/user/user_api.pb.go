// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/identity/user/user_api.proto

package user

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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	UserId *UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// OPTIONAL.
	// Whether to skip fetching user groups along with the user object.
	SkipFetchingUserGroups bool `protobuf:"varint,3,opt,name=skip_fetching_user_groups,json=skipFetchingUserGroups,proto3" json:"skip_fetching_user_groups,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserRequest) GetUserId() *UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *GetUserRequest) GetSkipFetchingUserGroups() bool {
	if x != nil {
		return x.SkipFetchingUserGroups
	}
	return false
}

type GetUserResponse struct {
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
	// The user information.
	User *User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetUserResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The claim on the basis of which users will be filtered.
	Claim string `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
	// REQUIRED.
	// The value of the claim to find the specific user.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// OPTIONAL.
	// Whether to skip fetching user groups along with the user object.
	SkipFetchingUserGroups bool `protobuf:"varint,4,opt,name=skip_fetching_user_groups,json=skipFetchingUserGroups,proto3" json:"skip_fetching_user_groups,omitempty"`
}

func (x *GetUserByClaimRequest) Reset() {
	*x = GetUserByClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByClaimRequest) ProtoMessage() {}

func (x *GetUserByClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByClaimRequest.ProtoReflect.Descriptor instead.
func (*GetUserByClaimRequest) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByClaimRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserByClaimRequest) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

func (x *GetUserByClaimRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetUserByClaimRequest) GetSkipFetchingUserGroups() bool {
	if x != nil {
		return x.SkipFetchingUserGroups
	}
	return false
}

type GetUserByClaimResponse struct {
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
	// The user information.
	User *User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByClaimResponse) Reset() {
	*x = GetUserByClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByClaimResponse) ProtoMessage() {}

func (x *GetUserByClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByClaimResponse.ProtoReflect.Descriptor instead.
func (*GetUserByClaimResponse) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByClaimResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetUserByClaimResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserByClaimResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The id of the user.
	UserId *UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserGroupsRequest) Reset() {
	*x = GetUserGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGroupsRequest) ProtoMessage() {}

func (x *GetUserGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetUserGroupsRequest) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserGroupsRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserGroupsRequest) GetUserId() *UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

type GetUserGroupsResponse struct {
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
	// The groups for the user.
	Groups []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetUserGroupsResponse) Reset() {
	*x = GetUserGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGroupsResponse) ProtoMessage() {}

func (x *GetUserGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetUserGroupsResponse) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserGroupsResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetUserGroupsResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *GetUserGroupsResponse) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type FindUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED. TODO(labkode): create proper filters for most common searches.
	// The filter to apply.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// OPTIONAL.
	// Whether to skip fetching user groups along with the user object.
	SkipFetchingUserGroups bool `protobuf:"varint,3,opt,name=skip_fetching_user_groups,json=skipFetchingUserGroups,proto3" json:"skip_fetching_user_groups,omitempty"`
}

func (x *FindUsersRequest) Reset() {
	*x = FindUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUsersRequest) ProtoMessage() {}

func (x *FindUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUsersRequest.ProtoReflect.Descriptor instead.
func (*FindUsersRequest) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{6}
}

func (x *FindUsersRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *FindUsersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *FindUsersRequest) GetSkipFetchingUserGroups() bool {
	if x != nil {
		return x.SkipFetchingUserGroups
	}
	return false
}

type FindUsersResponse struct {
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
	// The users matching the specified filter.
	Users []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *FindUsersResponse) Reset() {
	*x = FindUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_identity_user_user_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUsersResponse) ProtoMessage() {}

func (x *FindUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_identity_user_user_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUsersResponse.ProtoReflect.Descriptor instead.
func (*FindUsersResponse) Descriptor() ([]byte, []int) {
	return file_necode_identity_user_user_api_proto_rawDescGZIP(), []int{7}
}

func (x *FindUsersResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *FindUsersResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *FindUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_necode_identity_user_user_api_proto protoreflect.FileDescriptor

var file_necode_identity_user_user_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x24, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x19,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x73, 0x6b, 0x69, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x6b,
	0x69, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a,
	0x19, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x73, 0x6b, 0x69, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x96, 0x03, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2a, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xd2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b,
	0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4e, 0x49, 0x55, 0xaa, 0x02,
	0x14, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x14, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x20, 0x4e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x16, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_identity_user_user_api_proto_rawDescOnce sync.Once
	file_necode_identity_user_user_api_proto_rawDescData = file_necode_identity_user_user_api_proto_rawDesc
)

func file_necode_identity_user_user_api_proto_rawDescGZIP() []byte {
	file_necode_identity_user_user_api_proto_rawDescOnce.Do(func() {
		file_necode_identity_user_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_identity_user_user_api_proto_rawDescData)
	})
	return file_necode_identity_user_user_api_proto_rawDescData
}

var file_necode_identity_user_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_necode_identity_user_user_api_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),         // 0: necode.identity.user.GetUserRequest
	(*GetUserResponse)(nil),        // 1: necode.identity.user.GetUserResponse
	(*GetUserByClaimRequest)(nil),  // 2: necode.identity.user.GetUserByClaimRequest
	(*GetUserByClaimResponse)(nil), // 3: necode.identity.user.GetUserByClaimResponse
	(*GetUserGroupsRequest)(nil),   // 4: necode.identity.user.GetUserGroupsRequest
	(*GetUserGroupsResponse)(nil),  // 5: necode.identity.user.GetUserGroupsResponse
	(*FindUsersRequest)(nil),       // 6: necode.identity.user.FindUsersRequest
	(*FindUsersResponse)(nil),      // 7: necode.identity.user.FindUsersResponse
	(*types.Opaque)(nil),           // 8: necode.types.Opaque
	(*UserId)(nil),                 // 9: necode.identity.user.UserId
	(*rpc.Status)(nil),             // 10: necode.rpc.Status
	(*User)(nil),                   // 11: necode.identity.user.User
}
var file_necode_identity_user_user_api_proto_depIdxs = []int32{
	8,  // 0: necode.identity.user.GetUserRequest.opaque:type_name -> necode.types.Opaque
	9,  // 1: necode.identity.user.GetUserRequest.user_id:type_name -> necode.identity.user.UserId
	10, // 2: necode.identity.user.GetUserResponse.status:type_name -> necode.rpc.Status
	8,  // 3: necode.identity.user.GetUserResponse.opaque:type_name -> necode.types.Opaque
	11, // 4: necode.identity.user.GetUserResponse.user:type_name -> necode.identity.user.User
	8,  // 5: necode.identity.user.GetUserByClaimRequest.opaque:type_name -> necode.types.Opaque
	10, // 6: necode.identity.user.GetUserByClaimResponse.status:type_name -> necode.rpc.Status
	8,  // 7: necode.identity.user.GetUserByClaimResponse.opaque:type_name -> necode.types.Opaque
	11, // 8: necode.identity.user.GetUserByClaimResponse.user:type_name -> necode.identity.user.User
	8,  // 9: necode.identity.user.GetUserGroupsRequest.opaque:type_name -> necode.types.Opaque
	9,  // 10: necode.identity.user.GetUserGroupsRequest.user_id:type_name -> necode.identity.user.UserId
	10, // 11: necode.identity.user.GetUserGroupsResponse.status:type_name -> necode.rpc.Status
	8,  // 12: necode.identity.user.GetUserGroupsResponse.opaque:type_name -> necode.types.Opaque
	8,  // 13: necode.identity.user.FindUsersRequest.opaque:type_name -> necode.types.Opaque
	10, // 14: necode.identity.user.FindUsersResponse.status:type_name -> necode.rpc.Status
	8,  // 15: necode.identity.user.FindUsersResponse.opaque:type_name -> necode.types.Opaque
	11, // 16: necode.identity.user.FindUsersResponse.users:type_name -> necode.identity.user.User
	0,  // 17: necode.identity.user.UserAPI.GetUser:input_type -> necode.identity.user.GetUserRequest
	2,  // 18: necode.identity.user.UserAPI.GetUserByClaim:input_type -> necode.identity.user.GetUserByClaimRequest
	4,  // 19: necode.identity.user.UserAPI.GetUserGroups:input_type -> necode.identity.user.GetUserGroupsRequest
	6,  // 20: necode.identity.user.UserAPI.FindUsers:input_type -> necode.identity.user.FindUsersRequest
	1,  // 21: necode.identity.user.UserAPI.GetUser:output_type -> necode.identity.user.GetUserResponse
	3,  // 22: necode.identity.user.UserAPI.GetUserByClaim:output_type -> necode.identity.user.GetUserByClaimResponse
	5,  // 23: necode.identity.user.UserAPI.GetUserGroups:output_type -> necode.identity.user.GetUserGroupsResponse
	7,  // 24: necode.identity.user.UserAPI.FindUsers:output_type -> necode.identity.user.FindUsersResponse
	21, // [21:25] is the sub-list for method output_type
	17, // [17:21] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_necode_identity_user_user_api_proto_init() }
func file_necode_identity_user_user_api_proto_init() {
	if File_necode_identity_user_user_api_proto != nil {
		return
	}
	file_necode_identity_user_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_necode_identity_user_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByClaimRequest); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByClaimResponse); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserGroupsRequest); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserGroupsResponse); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUsersRequest); i {
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
		file_necode_identity_user_user_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUsersResponse); i {
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
			RawDescriptor: file_necode_identity_user_user_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_necode_identity_user_user_api_proto_goTypes,
		DependencyIndexes: file_necode_identity_user_user_api_proto_depIdxs,
		MessageInfos:      file_necode_identity_user_user_api_proto_msgTypes,
	}.Build()
	File_necode_identity_user_user_api_proto = out.File
	file_necode_identity_user_user_api_proto_rawDesc = nil
	file_necode_identity_user_user_api_proto_goTypes = nil
	file_necode_identity_user_user_api_proto_depIdxs = nil
}
