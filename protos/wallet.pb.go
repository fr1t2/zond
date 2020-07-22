// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.8.0
// source: protos/wallet.proto

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

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  string      `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XmssInfo []*XMSSInfo `protobuf:"bytes,2,rep,name=xmss_info,json=xmssInfo,proto3" json:"xmss_info,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_protos_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_protos_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Wallet) GetXmssInfo() []*XMSSInfo {
	if x != nil {
		return x.XmssInfo
	}
	return nil
}

type XMSSInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HexSeed  string `protobuf:"bytes,2,opt,name=hex_seed,json=hexSeed,proto3" json:"hex_seed,omitempty"`
	Mnemonic string `protobuf:"bytes,3,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
}

func (x *XMSSInfo) Reset() {
	*x = XMSSInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMSSInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMSSInfo) ProtoMessage() {}

func (x *XMSSInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMSSInfo.ProtoReflect.Descriptor instead.
func (*XMSSInfo) Descriptor() ([]byte, []int) {
	return file_protos_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *XMSSInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *XMSSInfo) GetHexSeed() string {
	if x != nil {
		return x.HexSeed
	}
	return ""
}

func (x *XMSSInfo) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

var File_protos_wallet_proto protoreflect.FileDescriptor

var file_protos_wallet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x51, 0x0a,
	0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x09, 0x78, 0x6d, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x58, 0x4d,
	0x53, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x78, 0x6d, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5b, 0x0a, 0x08, 0x58, 0x4d, 0x53, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x65, 0x78, 0x5f, 0x73, 0x65,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x78, 0x53, 0x65, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x42, 0x1f, 0x5a,
	0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x51,
	0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_wallet_proto_rawDescOnce sync.Once
	file_protos_wallet_proto_rawDescData = file_protos_wallet_proto_rawDesc
)

func file_protos_wallet_proto_rawDescGZIP() []byte {
	file_protos_wallet_proto_rawDescOnce.Do(func() {
		file_protos_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_wallet_proto_rawDescData)
	})
	return file_protos_wallet_proto_rawDescData
}

var file_protos_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_wallet_proto_goTypes = []interface{}{
	(*Wallet)(nil),   // 0: protos.Wallet
	(*XMSSInfo)(nil), // 1: protos.XMSSInfo
}
var file_protos_wallet_proto_depIdxs = []int32{
	1, // 0: protos.Wallet.xmss_info:type_name -> protos.XMSSInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_wallet_proto_init() }
func file_protos_wallet_proto_init() {
	if File_protos_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
		file_protos_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMSSInfo); i {
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
			RawDescriptor: file_protos_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_wallet_proto_goTypes,
		DependencyIndexes: file_protos_wallet_proto_depIdxs,
		MessageInfos:      file_protos_wallet_proto_msgTypes,
	}.Build()
	File_protos_wallet_proto = out.File
	file_protos_wallet_proto_rawDesc = nil
	file_protos_wallet_proto_goTypes = nil
	file_protos_wallet_proto_depIdxs = nil
}
