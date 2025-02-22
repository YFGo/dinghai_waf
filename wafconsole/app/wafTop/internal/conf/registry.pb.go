// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: app/wafTop/internal/conf/registry.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Registry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Consul        *RegistryConsul        `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`
	Etcd          *RegistryEtcd          `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registry) Reset() {
	*x = Registry{}
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_app_wafTop_internal_conf_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Registry) GetConsul() *RegistryConsul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetEtcd() *RegistryEtcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

// consul Registryig
type RegistryConsul struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryConsul) Reset() {
	*x = RegistryConsul{}
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryConsul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConsul) ProtoMessage() {}

func (x *RegistryConsul) ProtoReflect() protoreflect.Message {
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConsul.ProtoReflect.Descriptor instead.
func (*RegistryConsul) Descriptor() ([]byte, []int) {
	return file_app_wafTop_internal_conf_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryConsul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegistryEtcd struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryEtcd) Reset() {
	*x = RegistryEtcd{}
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryEtcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryEtcd) ProtoMessage() {}

func (x *RegistryEtcd) ProtoReflect() protoreflect.Message {
	mi := &file_app_wafTop_internal_conf_registry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryEtcd.ProtoReflect.Descriptor instead.
func (*RegistryEtcd) Descriptor() ([]byte, []int) {
	return file_app_wafTop_internal_conf_registry_proto_rawDescGZIP(), []int{2}
}

func (x *RegistryEtcd) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_app_wafTop_internal_conf_registry_proto protoreflect.FileDescriptor

var file_app_wafTop_internal_conf_registry_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x61, 0x70, 0x70, 0x2f, 0x77, 0x61, 0x66, 0x54, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x2c, 0x0a, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x74,
	0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x45, 0x74, 0x63, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x2a,
	0x5a, 0x28, 0x77, 0x61, 0x66, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x77, 0x61, 0x66, 0x54, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_app_wafTop_internal_conf_registry_proto_rawDescOnce sync.Once
	file_app_wafTop_internal_conf_registry_proto_rawDescData []byte
)

func file_app_wafTop_internal_conf_registry_proto_rawDescGZIP() []byte {
	file_app_wafTop_internal_conf_registry_proto_rawDescOnce.Do(func() {
		file_app_wafTop_internal_conf_registry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_wafTop_internal_conf_registry_proto_rawDesc), len(file_app_wafTop_internal_conf_registry_proto_rawDesc)))
	})
	return file_app_wafTop_internal_conf_registry_proto_rawDescData
}

var file_app_wafTop_internal_conf_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_wafTop_internal_conf_registry_proto_goTypes = []any{
	(*Registry)(nil),       // 0: kratos.api.Registry
	(*RegistryConsul)(nil), // 1: kratos.api.RegistryConsul
	(*RegistryEtcd)(nil),   // 2: kratos.api.RegistryEtcd
}
var file_app_wafTop_internal_conf_registry_proto_depIdxs = []int32{
	1, // 0: kratos.api.Registry.consul:type_name -> kratos.api.RegistryConsul
	2, // 1: kratos.api.Registry.etcd:type_name -> kratos.api.RegistryEtcd
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_wafTop_internal_conf_registry_proto_init() }
func file_app_wafTop_internal_conf_registry_proto_init() {
	if File_app_wafTop_internal_conf_registry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_wafTop_internal_conf_registry_proto_rawDesc), len(file_app_wafTop_internal_conf_registry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_wafTop_internal_conf_registry_proto_goTypes,
		DependencyIndexes: file_app_wafTop_internal_conf_registry_proto_depIdxs,
		MessageInfos:      file_app_wafTop_internal_conf_registry_proto_msgTypes,
	}.Build()
	File_app_wafTop_internal_conf_registry_proto = out.File
	file_app_wafTop_internal_conf_registry_proto_goTypes = nil
	file_app_wafTop_internal_conf_registry_proto_depIdxs = nil
}
