// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: necode/enums/schedule.proto

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

// ScheduleOverlapPolicy controls what happens when a workflow would be started
// by a schedule, and is already running.
type ScheduleOverlapPolicy int32

const (
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_UNSPECIFIED ScheduleOverlapPolicy = 0
	// SCHEDULE_OVERLAP_POLICY_SKIP (default) means don't start anything. When the
	// workflow completes, the next scheduled event after that time will be considered.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_SKIP ScheduleOverlapPolicy = 1
	// SCHEDULE_OVERLAP_POLICY_BUFFER_ONE means start the workflow again soon as the
	// current one completes, but only buffer one start in this way. If another start is
	// supposed to happen when the workflow is running, and one is already buffered, then
	// only the first one will be started after the running workflow finishes.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_BUFFER_ONE ScheduleOverlapPolicy = 2
	// SCHEDULE_OVERLAP_POLICY_BUFFER_ALL means buffer up any number of starts to all
	// happen sequentially, immediately after the running workflow completes.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_BUFFER_ALL ScheduleOverlapPolicy = 3
	// SCHEDULE_OVERLAP_POLICY_CANCEL_OTHER means that if there is another workflow
	// running, cancel it, and start the new one after the old one completes cancellation.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_CANCEL_OTHER ScheduleOverlapPolicy = 4
	// SCHEDULE_OVERLAP_POLICY_TERMINATE_OTHER means that if there is another workflow
	// running, terminate it and start the new one immediately.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_TERMINATE_OTHER ScheduleOverlapPolicy = 5
	// SCHEDULE_OVERLAP_POLICY_ALLOW_ALL means start any number of concurrent workflows.
	// Note that with this policy, last completion result and last failure will not be
	// available since workflows are not sequential.
	ScheduleOverlapPolicy_SCHEDULE_OVERLAP_POLICY_ALLOW_ALL ScheduleOverlapPolicy = 6
)

// Enum value maps for ScheduleOverlapPolicy.
var (
	ScheduleOverlapPolicy_name = map[int32]string{
		0: "SCHEDULE_OVERLAP_POLICY_UNSPECIFIED",
		1: "SCHEDULE_OVERLAP_POLICY_SKIP",
		2: "SCHEDULE_OVERLAP_POLICY_BUFFER_ONE",
		3: "SCHEDULE_OVERLAP_POLICY_BUFFER_ALL",
		4: "SCHEDULE_OVERLAP_POLICY_CANCEL_OTHER",
		5: "SCHEDULE_OVERLAP_POLICY_TERMINATE_OTHER",
		6: "SCHEDULE_OVERLAP_POLICY_ALLOW_ALL",
	}
	ScheduleOverlapPolicy_value = map[string]int32{
		"SCHEDULE_OVERLAP_POLICY_UNSPECIFIED":     0,
		"SCHEDULE_OVERLAP_POLICY_SKIP":            1,
		"SCHEDULE_OVERLAP_POLICY_BUFFER_ONE":      2,
		"SCHEDULE_OVERLAP_POLICY_BUFFER_ALL":      3,
		"SCHEDULE_OVERLAP_POLICY_CANCEL_OTHER":    4,
		"SCHEDULE_OVERLAP_POLICY_TERMINATE_OTHER": 5,
		"SCHEDULE_OVERLAP_POLICY_ALLOW_ALL":       6,
	}
)

func (x ScheduleOverlapPolicy) Enum() *ScheduleOverlapPolicy {
	p := new(ScheduleOverlapPolicy)
	*p = x
	return p
}

func (x ScheduleOverlapPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduleOverlapPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_necode_enums_schedule_proto_enumTypes[0].Descriptor()
}

func (ScheduleOverlapPolicy) Type() protoreflect.EnumType {
	return &file_necode_enums_schedule_proto_enumTypes[0]
}

func (x ScheduleOverlapPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduleOverlapPolicy.Descriptor instead.
func (ScheduleOverlapPolicy) EnumDescriptor() ([]byte, []int) {
	return file_necode_enums_schedule_proto_rawDescGZIP(), []int{0}
}

var File_necode_enums_schedule_proto protoreflect.FileDescriptor

var file_necode_enums_schedule_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0xb0, 0x02, 0x0a, 0x15,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c,
	0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x01,
	0x12, 0x26, 0x0a, 0x22, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x42, 0x55, 0x46, 0x46,
	0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x03,
	0x12, 0x28, 0x0a, 0x24, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x04, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x43, 0x48, 0x45, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49,
	0x43, 0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x06, 0x42, 0xa2,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x42, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x73, 0x6b, 0x76, 0x70, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x4e, 0x45, 0x58, 0xaa, 0x02, 0x0c, 0x4e, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x0c, 0x4e, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x18, 0x4e, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_necode_enums_schedule_proto_rawDescOnce sync.Once
	file_necode_enums_schedule_proto_rawDescData = file_necode_enums_schedule_proto_rawDesc
)

func file_necode_enums_schedule_proto_rawDescGZIP() []byte {
	file_necode_enums_schedule_proto_rawDescOnce.Do(func() {
		file_necode_enums_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_necode_enums_schedule_proto_rawDescData)
	})
	return file_necode_enums_schedule_proto_rawDescData
}

var file_necode_enums_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_necode_enums_schedule_proto_goTypes = []interface{}{
	(ScheduleOverlapPolicy)(0), // 0: necode.enums.ScheduleOverlapPolicy
}
var file_necode_enums_schedule_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_necode_enums_schedule_proto_init() }
func file_necode_enums_schedule_proto_init() {
	if File_necode_enums_schedule_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_necode_enums_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_necode_enums_schedule_proto_goTypes,
		DependencyIndexes: file_necode_enums_schedule_proto_depIdxs,
		EnumInfos:         file_necode_enums_schedule_proto_enumTypes,
	}.Build()
	File_necode_enums_schedule_proto = out.File
	file_necode_enums_schedule_proto_rawDesc = nil
	file_necode_enums_schedule_proto_goTypes = nil
	file_necode_enums_schedule_proto_depIdxs = nil
}
