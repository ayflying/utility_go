// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: system/v1/cron.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义 CronType 计划任务 枚举
type CronType int32

const (
	CronType_UNKNOWN   CronType = 0  // 未知类型
	CronType_SECOND    CronType = 1  // 每秒触发一次
	CronType_MINUTE    CronType = 2  // 每分钟触发一次
	CronType_HOUR      CronType = 3  // 每小时触发一次
	CronType_DAILY     CronType = 4  // 每天触发一次
	CronType_WEEK      CronType = 5  // 每周触发一次
	CronType_MONTH     CronType = 6  // 每月触发一次
	CronType_YEAR      CronType = 7  // 每年触发一次
	CronType_MONDAY    CronType = 8  // 每周一触发一次
	CronType_TUESDAY   CronType = 9  // 每周二触发一次
	CronType_WEDNESDAY CronType = 10 // 每周三触发一次
	CronType_THURSDAY  CronType = 11 // 每周四触发一次
	CronType_FRIDAY    CronType = 12 // 每周五触发一次
	CronType_SATURDAY  CronType = 13 // 每周六触发一次
	CronType_SUNDAY    CronType = 14 // 每周日触发一次
)

// Enum value maps for CronType.
var (
	CronType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "SECOND",
		2:  "MINUTE",
		3:  "HOUR",
		4:  "DAILY",
		5:  "WEEK",
		6:  "MONTH",
		7:  "YEAR",
		8:  "MONDAY",
		9:  "TUESDAY",
		10: "WEDNESDAY",
		11: "THURSDAY",
		12: "FRIDAY",
		13: "SATURDAY",
		14: "SUNDAY",
	}
	CronType_value = map[string]int32{
		"UNKNOWN":   0,
		"SECOND":    1,
		"MINUTE":    2,
		"HOUR":      3,
		"DAILY":     4,
		"WEEK":      5,
		"MONTH":     6,
		"YEAR":      7,
		"MONDAY":    8,
		"TUESDAY":   9,
		"WEDNESDAY": 10,
		"THURSDAY":  11,
		"FRIDAY":    12,
		"SATURDAY":  13,
		"SUNDAY":    14,
	}
)

func (x CronType) Enum() *CronType {
	p := new(CronType)
	*p = x
	return p
}

func (x CronType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CronType) Descriptor() protoreflect.EnumDescriptor {
	return file_system_v1_cron_proto_enumTypes[0].Descriptor()
}

func (CronType) Type() protoreflect.EnumType {
	return &file_system_v1_cron_proto_enumTypes[0]
}

func (x CronType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CronType.Descriptor instead.
func (CronType) EnumDescriptor() ([]byte, []int) {
	return file_system_v1_cron_proto_rawDescGZIP(), []int{0}
}

var File_system_v1_cron_proto protoreflect.FileDescriptor

var file_system_v1_cron_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2a, 0xbf,
	0x01, 0x0a, 0x08, 0x43, 0x72, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41,
	0x49, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45,
	0x41, 0x52, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x08,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x09, 0x12, 0x0d, 0x0a,
	0x09, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52,
	0x49, 0x44, 0x41, 0x59, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44,
	0x41, 0x59, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x0e,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x79, 0x66, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1_cron_proto_rawDescOnce sync.Once
	file_system_v1_cron_proto_rawDescData = file_system_v1_cron_proto_rawDesc
)

func file_system_v1_cron_proto_rawDescGZIP() []byte {
	file_system_v1_cron_proto_rawDescOnce.Do(func() {
		file_system_v1_cron_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1_cron_proto_rawDescData)
	})
	return file_system_v1_cron_proto_rawDescData
}

var file_system_v1_cron_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_system_v1_cron_proto_goTypes = []interface{}{
	(CronType)(0), // 0: system.CronType
}
var file_system_v1_cron_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_v1_cron_proto_init() }
func file_system_v1_cron_proto_init() {
	if File_system_v1_cron_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_v1_cron_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_v1_cron_proto_goTypes,
		DependencyIndexes: file_system_v1_cron_proto_depIdxs,
		EnumInfos:         file_system_v1_cron_proto_enumTypes,
	}.Build()
	File_system_v1_cron_proto = out.File
	file_system_v1_cron_proto_rawDesc = nil
	file_system_v1_cron_proto_goTypes = nil
	file_system_v1_cron_proto_depIdxs = nil
}
