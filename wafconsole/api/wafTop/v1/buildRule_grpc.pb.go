// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.2
// source: api/wafTop/v1/buildRule.proto

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
	BuildRule_GetBuildRule_FullMethodName  = "/api.wafTop.v1.BuildRule/GetBuildRule"
	BuildRule_ListBuildRule_FullMethodName = "/api.wafTop.v1.BuildRule/ListBuildRule"
)

// BuildRuleClient is the client API for BuildRule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildRuleClient interface {
	GetBuildRule(ctx context.Context, in *GetBuildRuleRequest, opts ...grpc.CallOption) (*GetBuildRuleReply, error)
	ListBuildRule(ctx context.Context, in *ListBuildRuleRequest, opts ...grpc.CallOption) (*ListBuildRuleReply, error)
}

type buildRuleClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildRuleClient(cc grpc.ClientConnInterface) BuildRuleClient {
	return &buildRuleClient{cc}
}

func (c *buildRuleClient) GetBuildRule(ctx context.Context, in *GetBuildRuleRequest, opts ...grpc.CallOption) (*GetBuildRuleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBuildRuleReply)
	err := c.cc.Invoke(ctx, BuildRule_GetBuildRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildRuleClient) ListBuildRule(ctx context.Context, in *ListBuildRuleRequest, opts ...grpc.CallOption) (*ListBuildRuleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBuildRuleReply)
	err := c.cc.Invoke(ctx, BuildRule_ListBuildRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildRuleServer is the server API for BuildRule service.
// All implementations must embed UnimplementedBuildRuleServer
// for forward compatibility.
type BuildRuleServer interface {
	GetBuildRule(context.Context, *GetBuildRuleRequest) (*GetBuildRuleReply, error)
	ListBuildRule(context.Context, *ListBuildRuleRequest) (*ListBuildRuleReply, error)
	mustEmbedUnimplementedBuildRuleServer()
}

// UnimplementedBuildRuleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBuildRuleServer struct{}

func (UnimplementedBuildRuleServer) GetBuildRule(context.Context, *GetBuildRuleRequest) (*GetBuildRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildRule not implemented")
}
func (UnimplementedBuildRuleServer) ListBuildRule(context.Context, *ListBuildRuleRequest) (*ListBuildRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuildRule not implemented")
}
func (UnimplementedBuildRuleServer) mustEmbedUnimplementedBuildRuleServer() {}
func (UnimplementedBuildRuleServer) testEmbeddedByValue()                   {}

// UnsafeBuildRuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildRuleServer will
// result in compilation errors.
type UnsafeBuildRuleServer interface {
	mustEmbedUnimplementedBuildRuleServer()
}

func RegisterBuildRuleServer(s grpc.ServiceRegistrar, srv BuildRuleServer) {
	// If the following call pancis, it indicates UnimplementedBuildRuleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BuildRule_ServiceDesc, srv)
}

func _BuildRule_GetBuildRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRuleServer).GetBuildRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuildRule_GetBuildRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRuleServer).GetBuildRule(ctx, req.(*GetBuildRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildRule_ListBuildRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildRuleServer).ListBuildRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuildRule_ListBuildRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildRuleServer).ListBuildRule(ctx, req.(*ListBuildRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuildRule_ServiceDesc is the grpc.ServiceDesc for BuildRule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildRule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wafTop.v1.BuildRule",
	HandlerType: (*BuildRuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuildRule",
			Handler:    _BuildRule_GetBuildRule_Handler,
		},
		{
			MethodName: "ListBuildRule",
			Handler:    _BuildRule_ListBuildRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/wafTop/v1/buildRule.proto",
}
