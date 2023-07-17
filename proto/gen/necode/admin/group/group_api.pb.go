// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/admin/group/group_api.proto

package group

import (
	group "github.com/rskvp/necode/proto/gen/necode/identity/group"
	user "github.com/rskvp/necode/proto/gen/necode/identity/user"
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

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The information of group to be created.
	Group *group.Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *CreateGroupRequest) GetGroup() *group.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupResponse struct {
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
	// The group information.
	Group *group.Group `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreateGroupResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *CreateGroupResponse) GetGroup() *group.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The group to be deleted, given their ID.
	GroupId *group.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGroupRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *DeleteGroupRequest) GetGroupId() *group.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

type DeleteGroupResponse struct {
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

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGroupResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeleteGroupResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type AddUserToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// ID of the user to add to the group
	UserId *user.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *group.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *AddUserToGroupRequest) Reset() {
	*x = AddUserToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupRequest) ProtoMessage() {}

func (x *AddUserToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserToGroupRequest) GetUserId() *user.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *AddUserToGroupRequest) GetGroupId() *group.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *AddUserToGroupRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type AddUserToGroupResponse struct {
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

func (x *AddUserToGroupResponse) Reset() {
	*x = AddUserToGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupResponse) ProtoMessage() {}

func (x *AddUserToGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{5}
}

func (x *AddUserToGroupResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AddUserToGroupResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type RemoveUserFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// ID of the user to add to the group
	UserId *user.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *group.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *RemoveUserFromGroupRequest) Reset() {
	*x = RemoveUserFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupRequest) ProtoMessage() {}

func (x *RemoveUserFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveUserFromGroupRequest) GetUserId() *user.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *RemoveUserFromGroupRequest) GetGroupId() *group.GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *RemoveUserFromGroupRequest) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type RemoveUserFromGroupResponse struct {
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

func (x *RemoveUserFromGroupResponse) Reset() {
	*x = RemoveUserFromGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_necode_admin_group_group_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupResponse) ProtoMessage() {}

func (x *RemoveUserFromGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_necode_admin_group_group_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupResponse) Descriptor() ([]byte, []int) {
	return file_necode_admin_group_group_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveUserFromGroupResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RemoveUserFromGroupResponse) GetOpaque() *types.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

var File_necode_admin_group_group_api_proto protoreflect.FileDescriptor

var file_necode_admin_group_group_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x25, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0xa3, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x7d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x22, 0x72, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x22, 0x77, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x32, 0xab, 0x03, 0x0a,
	0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x50, 0x49, 0x12, 0x5e, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x29, 0x2e, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x76, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xa2, 0x02, 0x03, 0x4e,
	0x41, 0x47, 0xaa, 0x02, 0x12, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xca, 0x02, 0x12, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xe2, 0x02, 0x1e, 0x4e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14,
	0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_admin_group_group_api_proto_rawDescOnce sync.Once
	file_necode_admin_group_group_api_proto_rawDescData = file_necode_admin_group_group_api_proto_rawDesc
)

func file_necode_admin_group_group_api_proto_rawDescGZIP() []byte {
	file_necode_admin_group_group_api_proto_rawDescOnce.Do(func() {
		file_necode_admin_group_group_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_admin_group_group_api_proto_rawDescData)
	})
	return file_necode_admin_group_group_api_proto_rawDescData
}

var file_necode_admin_group_group_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_necode_admin_group_group_api_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),          // 0: necode.admin.group.CreateGroupRequest
	(*CreateGroupResponse)(nil),         // 1: necode.admin.group.CreateGroupResponse
	(*DeleteGroupRequest)(nil),          // 2: necode.admin.group.DeleteGroupRequest
	(*DeleteGroupResponse)(nil),         // 3: necode.admin.group.DeleteGroupResponse
	(*AddUserToGroupRequest)(nil),       // 4: necode.admin.group.AddUserToGroupRequest
	(*AddUserToGroupResponse)(nil),      // 5: necode.admin.group.AddUserToGroupResponse
	(*RemoveUserFromGroupRequest)(nil),  // 6: necode.admin.group.RemoveUserFromGroupRequest
	(*RemoveUserFromGroupResponse)(nil), // 7: necode.admin.group.RemoveUserFromGroupResponse
	(*types.Opaque)(nil),                // 8: necode.types.Opaque
	(*group.Group)(nil),                 // 9: necode.identity.group.Group
	(*rpc.Status)(nil),                  // 10: necode.rpc.Status
	(*group.GroupId)(nil),               // 11: necode.identity.group.GroupId
	(*user.UserId)(nil),                 // 12: necode.identity.user.UserId
}
var file_necode_admin_group_group_api_proto_depIdxs = []int32{
	8,  // 0: necode.admin.group.CreateGroupRequest.opaque:type_name -> necode.types.Opaque
	9,  // 1: necode.admin.group.CreateGroupRequest.group:type_name -> necode.identity.group.Group
	10, // 2: necode.admin.group.CreateGroupResponse.status:type_name -> necode.rpc.Status
	8,  // 3: necode.admin.group.CreateGroupResponse.opaque:type_name -> necode.types.Opaque
	9,  // 4: necode.admin.group.CreateGroupResponse.group:type_name -> necode.identity.group.Group
	8,  // 5: necode.admin.group.DeleteGroupRequest.opaque:type_name -> necode.types.Opaque
	11, // 6: necode.admin.group.DeleteGroupRequest.group_id:type_name -> necode.identity.group.GroupId
	10, // 7: necode.admin.group.DeleteGroupResponse.status:type_name -> necode.rpc.Status
	8,  // 8: necode.admin.group.DeleteGroupResponse.opaque:type_name -> necode.types.Opaque
	12, // 9: necode.admin.group.AddUserToGroupRequest.user_id:type_name -> necode.identity.user.UserId
	11, // 10: necode.admin.group.AddUserToGroupRequest.group_id:type_name -> necode.identity.group.GroupId
	8,  // 11: necode.admin.group.AddUserToGroupRequest.opaque:type_name -> necode.types.Opaque
	10, // 12: necode.admin.group.AddUserToGroupResponse.status:type_name -> necode.rpc.Status
	8,  // 13: necode.admin.group.AddUserToGroupResponse.opaque:type_name -> necode.types.Opaque
	12, // 14: necode.admin.group.RemoveUserFromGroupRequest.user_id:type_name -> necode.identity.user.UserId
	11, // 15: necode.admin.group.RemoveUserFromGroupRequest.group_id:type_name -> necode.identity.group.GroupId
	8,  // 16: necode.admin.group.RemoveUserFromGroupRequest.opaque:type_name -> necode.types.Opaque
	10, // 17: necode.admin.group.RemoveUserFromGroupResponse.status:type_name -> necode.rpc.Status
	8,  // 18: necode.admin.group.RemoveUserFromGroupResponse.opaque:type_name -> necode.types.Opaque
	0,  // 19: necode.admin.group.GroupAPI.CreateGroup:input_type -> necode.admin.group.CreateGroupRequest
	2,  // 20: necode.admin.group.GroupAPI.DeleteGroup:input_type -> necode.admin.group.DeleteGroupRequest
	4,  // 21: necode.admin.group.GroupAPI.AddUserToGroup:input_type -> necode.admin.group.AddUserToGroupRequest
	6,  // 22: necode.admin.group.GroupAPI.RemoveUserFromGroup:input_type -> necode.admin.group.RemoveUserFromGroupRequest
	1,  // 23: necode.admin.group.GroupAPI.CreateGroup:output_type -> necode.admin.group.CreateGroupResponse
	3,  // 24: necode.admin.group.GroupAPI.DeleteGroup:output_type -> necode.admin.group.DeleteGroupResponse
	5,  // 25: necode.admin.group.GroupAPI.AddUserToGroup:output_type -> necode.admin.group.AddUserToGroupResponse
	7,  // 26: necode.admin.group.GroupAPI.RemoveUserFromGroup:output_type -> necode.admin.group.RemoveUserFromGroupResponse
	23, // [23:27] is the sub-list for method output_type
	19, // [19:23] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_necode_admin_group_group_api_proto_init() }
func file_necode_admin_group_group_api_proto_init() {
	if File_necode_admin_group_group_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_necode_admin_group_group_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupResponse); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupRequest); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupResponse); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupRequest); i {
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
		file_necode_admin_group_group_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupResponse); i {
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
			RawDescriptor: file_necode_admin_group_group_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_necode_admin_group_group_api_proto_goTypes,
		DependencyIndexes: file_necode_admin_group_group_api_proto_depIdxs,
		MessageInfos:      file_necode_admin_group_group_api_proto_msgTypes,
	}.Build()
	File_necode_admin_group_group_api_proto = out.File
	file_necode_admin_group_group_api_proto_rawDesc = nil
	file_necode_admin_group_group_api_proto_goTypes = nil
	file_necode_admin_group_group_api_proto_depIdxs = nil
}