// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.8.0
// source: protos/dilithiumkeys.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DilithiumKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	DilithiumGroup []*DilithiumGroup `protobuf:"bytes,2,rep,name=dilithium_group,json=dilithiumGroup,proto3" json:"dilithium_group,omitempty"`
}

func (x *DilithiumKeys) Reset() {
	*x = DilithiumKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dilithiumkeys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DilithiumKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DilithiumKeys) ProtoMessage() {}

func (x *DilithiumKeys) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dilithiumkeys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DilithiumKeys.ProtoReflect.Descriptor instead.
func (*DilithiumKeys) Descriptor() ([]byte, []int) {
	return file_protos_dilithiumkeys_proto_rawDescGZIP(), []int{0}
}

func (x *DilithiumKeys) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DilithiumKeys) GetDilithiumGroup() []*DilithiumGroup {
	if x != nil {
		return x.DilithiumGroup
	}
	return nil
}

type DilithiumGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DilithiumInfo []*DilithiumInfo `protobuf:"bytes,1,rep,name=dilithium_info,json=dilithiumInfo,proto3" json:"dilithium_info,omitempty"`
}

func (x *DilithiumGroup) Reset() {
	*x = DilithiumGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dilithiumkeys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DilithiumGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DilithiumGroup) ProtoMessage() {}

func (x *DilithiumGroup) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dilithiumkeys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DilithiumGroup.ProtoReflect.Descriptor instead.
func (*DilithiumGroup) Descriptor() ([]byte, []int) {
	return file_protos_dilithiumkeys_proto_rawDescGZIP(), []int{1}
}

func (x *DilithiumGroup) GetDilithiumInfo() []*DilithiumInfo {
	if x != nil {
		return x.DilithiumInfo
	}
	return nil
}

type DilithiumInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PK string `protobuf:"bytes,1,opt,name=PK,proto3" json:"PK,omitempty"`
	SK string `protobuf:"bytes,2,opt,name=SK,proto3" json:"SK,omitempty"`
}

func (x *DilithiumInfo) Reset() {
	*x = DilithiumInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dilithiumkeys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DilithiumInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DilithiumInfo) ProtoMessage() {}

func (x *DilithiumInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dilithiumkeys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DilithiumInfo.ProtoReflect.Descriptor instead.
func (*DilithiumInfo) Descriptor() ([]byte, []int) {
	return file_protos_dilithiumkeys_proto_rawDescGZIP(), []int{2}
}

func (x *DilithiumInfo) GetPK() string {
	if x != nil {
		return x.PK
	}
	return ""
}

func (x *DilithiumInfo) GetSK() string {
	if x != nil {
		return x.SK
	}
	return ""
}

var File_protos_dilithiumkeys_proto protoreflect.FileDescriptor

var file_protos_dilithiumkeys_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69,
	0x75, 0x6d, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x6a, 0x0a, 0x0d, 0x44, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75,
	0x6d, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3f, 0x0a, 0x0f, 0x64, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x44, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0e, 0x64, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x4e, 0x0a, 0x0e, 0x44, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x3c, 0x0a, 0x0e, 0x64, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x64, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x2f, 0x0a, 0x0d, 0x44, 0x69, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x75, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x50, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x50,
	0x4b, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x53,
	0x4b, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x68, 0x65, 0x51, 0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_dilithiumkeys_proto_rawDescOnce sync.Once
	file_protos_dilithiumkeys_proto_rawDescData = file_protos_dilithiumkeys_proto_rawDesc
)

func file_protos_dilithiumkeys_proto_rawDescGZIP() []byte {
	file_protos_dilithiumkeys_proto_rawDescOnce.Do(func() {
		file_protos_dilithiumkeys_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_dilithiumkeys_proto_rawDescData)
	})
	return file_protos_dilithiumkeys_proto_rawDescData
}

var file_protos_dilithiumkeys_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_dilithiumkeys_proto_goTypes = []interface{}{
	(*DilithiumKeys)(nil),  // 0: protos.DilithiumKeys
	(*DilithiumGroup)(nil), // 1: protos.DilithiumGroup
	(*DilithiumInfo)(nil),  // 2: protos.DilithiumInfo
}
var file_protos_dilithiumkeys_proto_depIdxs = []int32{
	1, // 0: protos.DilithiumKeys.dilithium_group:type_name -> protos.DilithiumGroup
	2, // 1: protos.DilithiumGroup.dilithium_info:type_name -> protos.DilithiumInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_dilithiumkeys_proto_init() }
func file_protos_dilithiumkeys_proto_init() {
	if File_protos_dilithiumkeys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_dilithiumkeys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DilithiumKeys); i {
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
		file_protos_dilithiumkeys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DilithiumGroup); i {
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
		file_protos_dilithiumkeys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DilithiumInfo); i {
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
			RawDescriptor: file_protos_dilithiumkeys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_dilithiumkeys_proto_goTypes,
		DependencyIndexes: file_protos_dilithiumkeys_proto_depIdxs,
		MessageInfos:      file_protos_dilithiumkeys_proto_msgTypes,
	}.Build()
	File_protos_dilithiumkeys_proto = out.File
	file_protos_dilithiumkeys_proto_rawDesc = nil
	file_protos_dilithiumkeys_proto_goTypes = nil
	file_protos_dilithiumkeys_proto_depIdxs = nil
}
