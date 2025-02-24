// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.2
// source: api/dashBorad/v1/dataView.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DataView_GetAttackInfoFromDay_FullMethodName    = "/api.dashBorad.v1.DataView/GetAttackInfoFromDay"
	DataView_GetAttackInfoByTime_FullMethodName     = "/api.dashBorad.v1.DataView/GetAttackInfoByTime"
	DataView_GetAttackInfoFromServer_FullMethodName = "/api.dashBorad.v1.DataView/GetAttackInfoFromServer"
	DataView_GetAttackIpFromAddr_FullMethodName     = "/api.dashBorad.v1.DataView/GetAttackIpFromAddr"
	DataView_GetAttackDetail_FullMethodName         = "/api.dashBorad.v1.DataView/GetAttackDetail"
)

// DataViewClient is the client API for DataView service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataViewClient interface {
	GetAttackInfoFromDay(ctx context.Context, in *GetAttackInfoFromDayRequest, opts ...grpc.CallOption) (*GetAttackInfoFromDayReply, error)
	GetAttackInfoByTime(ctx context.Context, in *GetAttackInfoByTimeRequest, opts ...grpc.CallOption) (*GetAttackInfoByTimeReply, error)
	GetAttackInfoFromServer(ctx context.Context, in *GetAttackInfoFromServerRequest, opts ...grpc.CallOption) (*GetAttackInfoFromServerReply, error)
	GetAttackIpFromAddr(ctx context.Context, in *GetAttackIpFromAddrRequest, opts ...grpc.CallOption) (*GetAttackIpFromAddrReply, error)
	GetAttackDetail(ctx context.Context, in *GetAttackDetailRequest, opts ...grpc.CallOption) (*GetAttackDetailReply, error)
}

type dataViewClient struct {
	cc grpc.ClientConnInterface
}

func NewDataViewClient(cc grpc.ClientConnInterface) DataViewClient {
	return &dataViewClient{cc}
}

func (c *dataViewClient) GetAttackInfoFromDay(ctx context.Context, in *GetAttackInfoFromDayRequest, opts ...grpc.CallOption) (*GetAttackInfoFromDayReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttackInfoFromDayReply)
	err := c.cc.Invoke(ctx, DataView_GetAttackInfoFromDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataViewClient) GetAttackInfoByTime(ctx context.Context, in *GetAttackInfoByTimeRequest, opts ...grpc.CallOption) (*GetAttackInfoByTimeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttackInfoByTimeReply)
	err := c.cc.Invoke(ctx, DataView_GetAttackInfoByTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataViewClient) GetAttackInfoFromServer(ctx context.Context, in *GetAttackInfoFromServerRequest, opts ...grpc.CallOption) (*GetAttackInfoFromServerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttackInfoFromServerReply)
	err := c.cc.Invoke(ctx, DataView_GetAttackInfoFromServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataViewClient) GetAttackIpFromAddr(ctx context.Context, in *GetAttackIpFromAddrRequest, opts ...grpc.CallOption) (*GetAttackIpFromAddrReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttackIpFromAddrReply)
	err := c.cc.Invoke(ctx, DataView_GetAttackIpFromAddr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataViewClient) GetAttackDetail(ctx context.Context, in *GetAttackDetailRequest, opts ...grpc.CallOption) (*GetAttackDetailReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttackDetailReply)
	err := c.cc.Invoke(ctx, DataView_GetAttackDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataViewServer is the server API for DataView service.
// All implementations must embed UnimplementedDataViewServer
// for forward compatibility.
type DataViewServer interface {
	GetAttackInfoFromDay(context.Context, *GetAttackInfoFromDayRequest) (*GetAttackInfoFromDayReply, error)
	GetAttackInfoByTime(context.Context, *GetAttackInfoByTimeRequest) (*GetAttackInfoByTimeReply, error)
	GetAttackInfoFromServer(context.Context, *GetAttackInfoFromServerRequest) (*GetAttackInfoFromServerReply, error)
	GetAttackIpFromAddr(context.Context, *GetAttackIpFromAddrRequest) (*GetAttackIpFromAddrReply, error)
	GetAttackDetail(context.Context, *GetAttackDetailRequest) (*GetAttackDetailReply, error)
	mustEmbedUnimplementedDataViewServer()
}

// UnimplementedDataViewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataViewServer struct{}

func (UnimplementedDataViewServer) GetAttackInfoFromDay(context.Context, *GetAttackInfoFromDayRequest) (*GetAttackInfoFromDayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackInfoFromDay not implemented")
}
func (UnimplementedDataViewServer) GetAttackInfoByTime(context.Context, *GetAttackInfoByTimeRequest) (*GetAttackInfoByTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackInfoByTime not implemented")
}
func (UnimplementedDataViewServer) GetAttackInfoFromServer(context.Context, *GetAttackInfoFromServerRequest) (*GetAttackInfoFromServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackInfoFromServer not implemented")
}
func (UnimplementedDataViewServer) GetAttackIpFromAddr(context.Context, *GetAttackIpFromAddrRequest) (*GetAttackIpFromAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackIpFromAddr not implemented")
}
func (UnimplementedDataViewServer) GetAttackDetail(context.Context, *GetAttackDetailRequest) (*GetAttackDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackDetail not implemented")
}
func (UnimplementedDataViewServer) mustEmbedUnimplementedDataViewServer() {}
func (UnimplementedDataViewServer) testEmbeddedByValue()                  {}

// UnsafeDataViewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataViewServer will
// result in compilation errors.
type UnsafeDataViewServer interface {
	mustEmbedUnimplementedDataViewServer()
}

func RegisterDataViewServer(s grpc.ServiceRegistrar, srv DataViewServer) {
	// If the following call pancis, it indicates UnimplementedDataViewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataView_ServiceDesc, srv)
}

func _DataView_GetAttackInfoFromDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttackInfoFromDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewServer).GetAttackInfoFromDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataView_GetAttackInfoFromDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewServer).GetAttackInfoFromDay(ctx, req.(*GetAttackInfoFromDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataView_GetAttackInfoByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttackInfoByTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewServer).GetAttackInfoByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataView_GetAttackInfoByTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewServer).GetAttackInfoByTime(ctx, req.(*GetAttackInfoByTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataView_GetAttackInfoFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttackInfoFromServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewServer).GetAttackInfoFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataView_GetAttackInfoFromServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewServer).GetAttackInfoFromServer(ctx, req.(*GetAttackInfoFromServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataView_GetAttackIpFromAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttackIpFromAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewServer).GetAttackIpFromAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataView_GetAttackIpFromAddr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewServer).GetAttackIpFromAddr(ctx, req.(*GetAttackIpFromAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataView_GetAttackDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttackDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewServer).GetAttackDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataView_GetAttackDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewServer).GetAttackDetail(ctx, req.(*GetAttackDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataView_ServiceDesc is the grpc.ServiceDesc for DataView service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataView_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.dashBorad.v1.DataView",
	HandlerType: (*DataViewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAttackInfoFromDay",
			Handler:    _DataView_GetAttackInfoFromDay_Handler,
		},
		{
			MethodName: "GetAttackInfoByTime",
			Handler:    _DataView_GetAttackInfoByTime_Handler,
		},
		{
			MethodName: "GetAttackInfoFromServer",
			Handler:    _DataView_GetAttackInfoFromServer_Handler,
		},
		{
			MethodName: "GetAttackIpFromAddr",
			Handler:    _DataView_GetAttackIpFromAddr_Handler,
		},
		{
			MethodName: "GetAttackDetail",
			Handler:    _DataView_GetAttackDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/dashBorad/v1/dataView.proto",
}
