// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: necode/identity/group/group_api.proto

package group

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GroupAPI_GetGroup_FullMethodName        = "/necode.identity.group.GroupAPI/GetGroup"
	GroupAPI_GetGroupByClaim_FullMethodName = "/necode.identity.group.GroupAPI/GetGroupByClaim"
	GroupAPI_GetMembers_FullMethodName      = "/necode.identity.group.GroupAPI/GetMembers"
	GroupAPI_HasMember_FullMethodName       = "/necode.identity.group.GroupAPI/HasMember"
	GroupAPI_FindGroups_FullMethodName      = "/necode.identity.group.GroupAPI/FindGroups"
)

// GroupAPIClient is the client API for GroupAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupAPIClient interface {
	// Gets the information about a group by the group id.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	// Gets the information about a group based on a specified claim.
	GetGroupByClaim(ctx context.Context, in *GetGroupByClaimRequest, opts ...grpc.CallOption) (*GetGroupByClaimResponse, error)
	// Gets the members of a group.
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error)
	// Tells if the group has certain member.
	HasMember(ctx context.Context, in *HasMemberRequest, opts ...grpc.CallOption) (*HasMemberResponse, error)
	// Finds groups whose names match the specified filter.
	FindGroups(ctx context.Context, in *FindGroupsRequest, opts ...grpc.CallOption) (*FindGroupsResponse, error)
}

type groupAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupAPIClient(cc grpc.ClientConnInterface) GroupAPIClient {
	return &groupAPIClient{cc}
}

func (c *groupAPIClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, GroupAPI_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) GetGroupByClaim(ctx context.Context, in *GetGroupByClaimRequest, opts ...grpc.CallOption) (*GetGroupByClaimResponse, error) {
	out := new(GetGroupByClaimResponse)
	err := c.cc.Invoke(ctx, GroupAPI_GetGroupByClaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error) {
	out := new(GetMembersResponse)
	err := c.cc.Invoke(ctx, GroupAPI_GetMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) HasMember(ctx context.Context, in *HasMemberRequest, opts ...grpc.CallOption) (*HasMemberResponse, error) {
	out := new(HasMemberResponse)
	err := c.cc.Invoke(ctx, GroupAPI_HasMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) FindGroups(ctx context.Context, in *FindGroupsRequest, opts ...grpc.CallOption) (*FindGroupsResponse, error) {
	out := new(FindGroupsResponse)
	err := c.cc.Invoke(ctx, GroupAPI_FindGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupAPIServer is the server API for GroupAPI service.
// All implementations must embed UnimplementedGroupAPIServer
// for forward compatibility
type GroupAPIServer interface {
	// Gets the information about a group by the group id.
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	// Gets the information about a group based on a specified claim.
	GetGroupByClaim(context.Context, *GetGroupByClaimRequest) (*GetGroupByClaimResponse, error)
	// Gets the members of a group.
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error)
	// Tells if the group has certain member.
	HasMember(context.Context, *HasMemberRequest) (*HasMemberResponse, error)
	// Finds groups whose names match the specified filter.
	FindGroups(context.Context, *FindGroupsRequest) (*FindGroupsResponse, error)
	mustEmbedUnimplementedGroupAPIServer()
}

// UnimplementedGroupAPIServer must be embedded to have forward compatible implementations.
type UnimplementedGroupAPIServer struct {
}

func (UnimplementedGroupAPIServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupAPIServer) GetGroupByClaim(context.Context, *GetGroupByClaimRequest) (*GetGroupByClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupByClaim not implemented")
}
func (UnimplementedGroupAPIServer) GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedGroupAPIServer) HasMember(context.Context, *HasMemberRequest) (*HasMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasMember not implemented")
}
func (UnimplementedGroupAPIServer) FindGroups(context.Context, *FindGroupsRequest) (*FindGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGroups not implemented")
}
func (UnimplementedGroupAPIServer) mustEmbedUnimplementedGroupAPIServer() {}

// UnsafeGroupAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupAPIServer will
// result in compilation errors.
type UnsafeGroupAPIServer interface {
	mustEmbedUnimplementedGroupAPIServer()
}

func RegisterGroupAPIServer(s grpc.ServiceRegistrar, srv GroupAPIServer) {
	s.RegisterService(&GroupAPI_ServiceDesc, srv)
}

func _GroupAPI_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAPI_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_GetGroupByClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupByClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).GetGroupByClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAPI_GetGroupByClaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).GetGroupByClaim(ctx, req.(*GetGroupByClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAPI_GetMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_HasMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).HasMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAPI_HasMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).HasMember(ctx, req.(*HasMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_FindGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).FindGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupAPI_FindGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).FindGroups(ctx, req.(*FindGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupAPI_ServiceDesc is the grpc.ServiceDesc for GroupAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "necode.identity.group.GroupAPI",
	HandlerType: (*GroupAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _GroupAPI_GetGroup_Handler,
		},
		{
			MethodName: "GetGroupByClaim",
			Handler:    _GroupAPI_GetGroupByClaim_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _GroupAPI_GetMembers_Handler,
		},
		{
			MethodName: "HasMember",
			Handler:    _GroupAPI_HasMember_Handler,
		},
		{
			MethodName: "FindGroups",
			Handler:    _GroupAPI_FindGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "necode/identity/group/group_api.proto",
}
