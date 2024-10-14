// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v4.25.2
// source: api/wafTop/v1/wafApp.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWafAppCreateWafApp = "/api.wafTop.v1.WafApp/CreateWafApp"
const OperationWafAppDeleteWafApp = "/api.wafTop.v1.WafApp/DeleteWafApp"
const OperationWafAppGetWafApp = "/api.wafTop.v1.WafApp/GetWafApp"
const OperationWafAppListWafApp = "/api.wafTop.v1.WafApp/ListWafApp"
const OperationWafAppUpdateWafApp = "/api.wafTop.v1.WafApp/UpdateWafApp"

type WafAppHTTPServer interface {
	CreateWafApp(context.Context, *ChangeServerRequest) (*CreateWafAppReply, error)
	DeleteWafApp(context.Context, *DeleteRequest) (*DeleteReply, error)
	GetWafApp(context.Context, *GetWafAppRequest) (*GetWafAppReply, error)
	ListWafApp(context.Context, *ListWafAppRequest) (*ListWafAppReply, error)
	UpdateWafApp(context.Context, *ChangeServerRequest) (*UpdateWafAppReply, error)
}

func RegisterWafAppHTTPServer(s *http.Server, srv WafAppHTTPServer) {
	r := s.Route("/")
	r.POST("/app/wafTop/v1/wafApp", _WafApp_CreateWafApp0_HTTP_Handler(srv))
	r.PATCH("/app/wafTop/v1/wafApp", _WafApp_UpdateWafApp0_HTTP_Handler(srv))
	r.DELETE("/app/wafTop/v1/wafApp", _WafApp_DeleteWafApp0_HTTP_Handler(srv))
	r.GET("/app/wafTop/v1/wafApp/{id}", _WafApp_GetWafApp0_HTTP_Handler(srv))
	r.GET("/app/wafTop/v1/wafApps", _WafApp_ListWafApp0_HTTP_Handler(srv))
}

func _WafApp_CreateWafApp0_HTTP_Handler(srv WafAppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangeServerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWafAppCreateWafApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateWafApp(ctx, req.(*ChangeServerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateWafAppReply)
		return ctx.Result(200, reply)
	}
}

func _WafApp_UpdateWafApp0_HTTP_Handler(srv WafAppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangeServerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWafAppUpdateWafApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWafApp(ctx, req.(*ChangeServerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateWafAppReply)
		return ctx.Result(200, reply)
	}
}

func _WafApp_DeleteWafApp0_HTTP_Handler(srv WafAppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWafAppDeleteWafApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteWafApp(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteReply)
		return ctx.Result(200, reply)
	}
}

func _WafApp_GetWafApp0_HTTP_Handler(srv WafAppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWafAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWafAppGetWafApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWafApp(ctx, req.(*GetWafAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWafAppReply)
		return ctx.Result(200, reply)
	}
}

func _WafApp_ListWafApp0_HTTP_Handler(srv WafAppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListWafAppRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWafAppListWafApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListWafApp(ctx, req.(*ListWafAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListWafAppReply)
		return ctx.Result(200, reply)
	}
}

type WafAppHTTPClient interface {
	CreateWafApp(ctx context.Context, req *ChangeServerRequest, opts ...http.CallOption) (rsp *CreateWafAppReply, err error)
	DeleteWafApp(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	GetWafApp(ctx context.Context, req *GetWafAppRequest, opts ...http.CallOption) (rsp *GetWafAppReply, err error)
	ListWafApp(ctx context.Context, req *ListWafAppRequest, opts ...http.CallOption) (rsp *ListWafAppReply, err error)
	UpdateWafApp(ctx context.Context, req *ChangeServerRequest, opts ...http.CallOption) (rsp *UpdateWafAppReply, err error)
}

type WafAppHTTPClientImpl struct {
	cc *http.Client
}

func NewWafAppHTTPClient(client *http.Client) WafAppHTTPClient {
	return &WafAppHTTPClientImpl{client}
}

func (c *WafAppHTTPClientImpl) CreateWafApp(ctx context.Context, in *ChangeServerRequest, opts ...http.CallOption) (*CreateWafAppReply, error) {
	var out CreateWafAppReply
	pattern := "/app/wafTop/v1/wafApp"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWafAppCreateWafApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WafAppHTTPClientImpl) DeleteWafApp(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*DeleteReply, error) {
	var out DeleteReply
	pattern := "/app/wafTop/v1/wafApp"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWafAppDeleteWafApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WafAppHTTPClientImpl) GetWafApp(ctx context.Context, in *GetWafAppRequest, opts ...http.CallOption) (*GetWafAppReply, error) {
	var out GetWafAppReply
	pattern := "/app/wafTop/v1/wafApp/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWafAppGetWafApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WafAppHTTPClientImpl) ListWafApp(ctx context.Context, in *ListWafAppRequest, opts ...http.CallOption) (*ListWafAppReply, error) {
	var out ListWafAppReply
	pattern := "/app/wafTop/v1/wafApps"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWafAppListWafApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WafAppHTTPClientImpl) UpdateWafApp(ctx context.Context, in *ChangeServerRequest, opts ...http.CallOption) (*UpdateWafAppReply, error) {
	var out UpdateWafAppReply
	pattern := "/app/wafTop/v1/wafApp"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWafAppUpdateWafApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
