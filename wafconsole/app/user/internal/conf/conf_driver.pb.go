// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: app/user/internal/conf/conf_driver.proto

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

type ConfDriver struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Consul        *ConfConsul            `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`
	Etcd          *ConfEtcd              `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfDriver) Reset() {
	*x = ConfDriver{}
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfDriver) ProtoMessage() {}

func (x *ConfDriver) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfDriver.ProtoReflect.Descriptor instead.
func (*ConfDriver) Descriptor() ([]byte, []int) {
	return file_app_user_internal_conf_conf_driver_proto_rawDescGZIP(), []int{0}
}

func (x *ConfDriver) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ConfDriver) GetConsul() *ConfConsul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *ConfDriver) GetEtcd() *ConfEtcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

// consul config
type ConfConsul struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme        string                 `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	PathPrefix    string                 `protobuf:"bytes,3,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Token         string                 `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Path          string                 `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfConsul) Reset() {
	*x = ConfConsul{}
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfConsul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfConsul) ProtoMessage() {}

func (x *ConfConsul) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfConsul.ProtoReflect.Descriptor instead.
func (*ConfConsul) Descriptor() ([]byte, []int) {
	return file_app_user_internal_conf_conf_driver_proto_rawDescGZIP(), []int{1}
}

func (x *ConfConsul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConfConsul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ConfConsul) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *ConfConsul) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConfConsul) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ConfEtcd struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme        string                 `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	PathPrefix    string                 `protobuf:"bytes,3,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Token         string                 `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Path          string                 `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfEtcd) Reset() {
	*x = ConfEtcd{}
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfEtcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfEtcd) ProtoMessage() {}

func (x *ConfEtcd) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_internal_conf_conf_driver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfEtcd.ProtoReflect.Descriptor instead.
func (*ConfEtcd) Descriptor() ([]byte, []int) {
	return file_app_user_internal_conf_conf_driver_proto_rawDescGZIP(), []int{2}
}

func (x *ConfEtcd) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConfEtcd) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ConfEtcd) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *ConfEtcd) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConfEtcd) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_app_user_internal_conf_conf_driver_proto protoreflect.FileDescriptor

var file_app_user_internal_conf_conf_driver_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x7a, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x87,
	0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x74, 0x63, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x28, 0x5a, 0x26, 0x77, 0x61, 0x66, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_app_user_internal_conf_conf_driver_proto_rawDescOnce sync.Once
	file_app_user_internal_conf_conf_driver_proto_rawDescData []byte
)

func file_app_user_internal_conf_conf_driver_proto_rawDescGZIP() []byte {
	file_app_user_internal_conf_conf_driver_proto_rawDescOnce.Do(func() {
		file_app_user_internal_conf_conf_driver_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_user_internal_conf_conf_driver_proto_rawDesc), len(file_app_user_internal_conf_conf_driver_proto_rawDesc)))
	})
	return file_app_user_internal_conf_conf_driver_proto_rawDescData
}

var file_app_user_internal_conf_conf_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_user_internal_conf_conf_driver_proto_goTypes = []any{
	(*ConfDriver)(nil), // 0: kratos.api.ConfDriver
	(*ConfConsul)(nil), // 1: kratos.api.ConfConsul
	(*ConfEtcd)(nil),   // 2: kratos.api.ConfEtcd
}
var file_app_user_internal_conf_conf_driver_proto_depIdxs = []int32{
	1, // 0: kratos.api.ConfDriver.consul:type_name -> kratos.api.ConfConsul
	2, // 1: kratos.api.ConfDriver.etcd:type_name -> kratos.api.ConfEtcd
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_user_internal_conf_conf_driver_proto_init() }
func file_app_user_internal_conf_conf_driver_proto_init() {
	if File_app_user_internal_conf_conf_driver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_user_internal_conf_conf_driver_proto_rawDesc), len(file_app_user_internal_conf_conf_driver_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_user_internal_conf_conf_driver_proto_goTypes,
		DependencyIndexes: file_app_user_internal_conf_conf_driver_proto_depIdxs,
		MessageInfos:      file_app_user_internal_conf_conf_driver_proto_msgTypes,
	}.Build()
	File_app_user_internal_conf_conf_driver_proto = out.File
	file_app_user_internal_conf_conf_driver_proto_goTypes = nil
	file_app_user_internal_conf_conf_driver_proto_depIdxs = nil
}
