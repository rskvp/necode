// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/enums/update.proto

package enums

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

// UpdateWorkflowExecutionLifecycleStage is specified by clients invoking
// workflow execution updates and used to indicate to the server how long the
// client wishes to wait for a return value from the RPC. If any value other
// than UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED is sent by the
// client then the RPC will complete before the update is finished and will
// return a handle to the running update so that it can later be polled for
// completion.
type UpdateWorkflowExecutionLifecycleStage int32

const (
	// An unspecified vale for this enum.
	UpdateWorkflowExecutionLifecycleStage_UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED UpdateWorkflowExecutionLifecycleStage = 0
	// The gRPC call will not return until the update request has been admitted
	// by the server - it may be the case that due to a considerations like load
	// or resource limits that an update is made to wait before the server will
	// indicate that it has been received and will be processed. This value
	// does not wait for any sort of acknowledgement from a worker.
	UpdateWorkflowExecutionLifecycleStage_UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED UpdateWorkflowExecutionLifecycleStage = 1
	// The gRPC call will not return until the update has passed validation on
	// a worker.
	UpdateWorkflowExecutionLifecycleStage_UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED UpdateWorkflowExecutionLifecycleStage = 2
	// The gRPC call will not return until the update has executed to completion
	// on a worker and has either been rejected or returned a value or an error.
	UpdateWorkflowExecutionLifecycleStage_UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED UpdateWorkflowExecutionLifecycleStage = 3
)

// Enum value maps for UpdateWorkflowExecutionLifecycleStage.
var (
	UpdateWorkflowExecutionLifecycleStage_name = map[int32]string{
		0: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED",
		1: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED",
		2: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED",
		3: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED",
	}
	UpdateWorkflowExecutionLifecycleStage_value = map[string]int32{
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED": 0,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED":    1,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED":    2,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED":   3,
	}
)

func (x UpdateWorkflowExecutionLifecycleStage) Enum() *UpdateWorkflowExecutionLifecycleStage {
	p := new(UpdateWorkflowExecutionLifecycleStage)
	*p = x
	return p
}

func (x UpdateWorkflowExecutionLifecycleStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateWorkflowExecutionLifecycleStage) Descriptor() protoreflect.EnumDescriptor {
	return file_necode_enums_update_proto_enumTypes[0].Descriptor()
}

func (UpdateWorkflowExecutionLifecycleStage) Type() protoreflect.EnumType {
	return &file_necode_enums_update_proto_enumTypes[0]
}

func (x UpdateWorkflowExecutionLifecycleStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateWorkflowExecutionLifecycleStage.Descriptor instead.
func (UpdateWorkflowExecutionLifecycleStage) EnumDescriptor() ([]byte, []int) {
	return file_necode_enums_update_proto_rawDescGZIP(), []int{0}
}

var File_necode_enums_update_proto protoreflect.FileDescriptor

var file_necode_enums_update_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x8b, 0x02, 0x0a, 0x25, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x35, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x36,
	0x0a, 0x32, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x46, 0x45,
	0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49,
	0x54, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x36, 0x0a, 0x32, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x37,
	0x0a, 0x33, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x46, 0x45,
	0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0xa0, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x4e, 0x45,
	0x58, 0xaa, 0x02, 0x0c, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x0c, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2,
	0x02, 0x18, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_necode_enums_update_proto_rawDescOnce sync.Once
	file_necode_enums_update_proto_rawDescData = file_necode_enums_update_proto_rawDesc
)

func file_necode_enums_update_proto_rawDescGZIP() []byte {
	file_necode_enums_update_proto_rawDescOnce.Do(func() {
		file_necode_enums_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_enums_update_proto_rawDescData)
	})
	return file_necode_enums_update_proto_rawDescData
}

var file_necode_enums_update_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_necode_enums_update_proto_goTypes = []interface{}{
	(UpdateWorkflowExecutionLifecycleStage)(0), // 0: necode.enums.UpdateWorkflowExecutionLifecycleStage
}
var file_necode_enums_update_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_necode_enums_update_proto_init() }
func file_necode_enums_update_proto_init() {
	if File_necode_enums_update_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_necode_enums_update_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_necode_enums_update_proto_goTypes,
		DependencyIndexes: file_necode_enums_update_proto_depIdxs,
		EnumInfos:         file_necode_enums_update_proto_enumTypes,
	}.Build()
	File_necode_enums_update_proto = out.File
	file_necode_enums_update_proto_rawDesc = nil
	file_necode_enums_update_proto_goTypes = nil
	file_necode_enums_update_proto_depIdxs = nil
}
