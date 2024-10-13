// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: wafTop/v1/message.proto

package v1

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

type CreateWafAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWafAppRequest) Reset() {
	*x = CreateWafAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWafAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWafAppRequest) ProtoMessage() {}

func (x *CreateWafAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWafAppRequest.ProtoReflect.Descriptor instead.
func (*CreateWafAppRequest) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{0}
}

type CreateWafAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWafAppReply) Reset() {
	*x = CreateWafAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWafAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWafAppReply) ProtoMessage() {}

func (x *CreateWafAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWafAppReply.ProtoReflect.Descriptor instead.
func (*CreateWafAppReply) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{1}
}

type UpdateWafAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateWafAppRequest) Reset() {
	*x = UpdateWafAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWafAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWafAppRequest) ProtoMessage() {}

func (x *UpdateWafAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWafAppRequest.ProtoReflect.Descriptor instead.
func (*UpdateWafAppRequest) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{2}
}

type UpdateWafAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateWafAppReply) Reset() {
	*x = UpdateWafAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWafAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWafAppReply) ProtoMessage() {}

func (x *UpdateWafAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWafAppReply.ProtoReflect.Descriptor instead.
func (*UpdateWafAppReply) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{3}
}

type DeleteWafAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWafAppRequest) Reset() {
	*x = DeleteWafAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWafAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWafAppRequest) ProtoMessage() {}

func (x *DeleteWafAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWafAppRequest.ProtoReflect.Descriptor instead.
func (*DeleteWafAppRequest) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{4}
}

type DeleteWafAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWafAppReply) Reset() {
	*x = DeleteWafAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWafAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWafAppReply) ProtoMessage() {}

func (x *DeleteWafAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWafAppReply.ProtoReflect.Descriptor instead.
func (*DeleteWafAppReply) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{5}
}

type GetWafAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWafAppRequest) Reset() {
	*x = GetWafAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWafAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWafAppRequest) ProtoMessage() {}

func (x *GetWafAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWafAppRequest.ProtoReflect.Descriptor instead.
func (*GetWafAppRequest) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{6}
}

type GetWafAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWafAppReply) Reset() {
	*x = GetWafAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWafAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWafAppReply) ProtoMessage() {}

func (x *GetWafAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWafAppReply.ProtoReflect.Descriptor instead.
func (*GetWafAppReply) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{7}
}

type ListWafAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListWafAppRequest) Reset() {
	*x = ListWafAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWafAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWafAppRequest) ProtoMessage() {}

func (x *ListWafAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWafAppRequest.ProtoReflect.Descriptor instead.
func (*ListWafAppRequest) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{8}
}

type ListWafAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListWafAppReply) Reset() {
	*x = ListWafAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wafTop_v1_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWafAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWafAppReply) ProtoMessage() {}

func (x *ListWafAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_wafTop_v1_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWafAppReply.ProtoReflect.Descriptor instead.
func (*ListWafAppReply) Descriptor() ([]byte, []int) {
	return file_wafTop_v1_message_proto_rawDescGZIP(), []int{9}
}

var File_wafTop_v1_message_proto protoreflect.FileDescriptor

var file_wafTop_v1_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x61, 0x66, 0x54, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x77,
	0x61, 0x66, 0x54, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x66, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x61, 0x66, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2e, 0x0a, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x61, 0x66, 0x54, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x77,
	0x61, 0x66, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61,
	0x66, 0x54, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wafTop_v1_message_proto_rawDescOnce sync.Once
	file_wafTop_v1_message_proto_rawDescData = file_wafTop_v1_message_proto_rawDesc
)

func file_wafTop_v1_message_proto_rawDescGZIP() []byte {
	file_wafTop_v1_message_proto_rawDescOnce.Do(func() {
		file_wafTop_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_wafTop_v1_message_proto_rawDescData)
	})
	return file_wafTop_v1_message_proto_rawDescData
}

var file_wafTop_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wafTop_v1_message_proto_goTypes = []any{
	(*CreateWafAppRequest)(nil), // 0: api.wafTop.v1.CreateWafAppRequest
	(*CreateWafAppReply)(nil),   // 1: api.wafTop.v1.CreateWafAppReply
	(*UpdateWafAppRequest)(nil), // 2: api.wafTop.v1.UpdateWafAppRequest
	(*UpdateWafAppReply)(nil),   // 3: api.wafTop.v1.UpdateWafAppReply
	(*DeleteWafAppRequest)(nil), // 4: api.wafTop.v1.DeleteWafAppRequest
	(*DeleteWafAppReply)(nil),   // 5: api.wafTop.v1.DeleteWafAppReply
	(*GetWafAppRequest)(nil),    // 6: api.wafTop.v1.GetWafAppRequest
	(*GetWafAppReply)(nil),      // 7: api.wafTop.v1.GetWafAppReply
	(*ListWafAppRequest)(nil),   // 8: api.wafTop.v1.ListWafAppRequest
	(*ListWafAppReply)(nil),     // 9: api.wafTop.v1.ListWafAppReply
}
var file_wafTop_v1_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wafTop_v1_message_proto_init() }
func file_wafTop_v1_message_proto_init() {
	if File_wafTop_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wafTop_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWafAppRequest); i {
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
		file_wafTop_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWafAppReply); i {
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
		file_wafTop_v1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateWafAppRequest); i {
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
		file_wafTop_v1_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateWafAppReply); i {
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
		file_wafTop_v1_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteWafAppRequest); i {
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
		file_wafTop_v1_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteWafAppReply); i {
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
		file_wafTop_v1_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetWafAppRequest); i {
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
		file_wafTop_v1_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetWafAppReply); i {
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
		file_wafTop_v1_message_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListWafAppRequest); i {
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
		file_wafTop_v1_message_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListWafAppReply); i {
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
			RawDescriptor: file_wafTop_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wafTop_v1_message_proto_goTypes,
		DependencyIndexes: file_wafTop_v1_message_proto_depIdxs,
		MessageInfos:      file_wafTop_v1_message_proto_msgTypes,
	}.Build()
	File_wafTop_v1_message_proto = out.File
	file_wafTop_v1_message_proto_rawDesc = nil
	file_wafTop_v1_message_proto_goTypes = nil
	file_wafTop_v1_message_proto_depIdxs = nil
}
