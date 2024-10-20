// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v4.25.2
// source: api/wafTop/v1/strategy.proto

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

const OperationStrategyCreateStrategy = "/api.wafTop.v1.Strategy/CreateStrategy"
const OperationStrategyDeleteStrategy = "/api.wafTop.v1.Strategy/DeleteStrategy"
const OperationStrategyGetStrategy = "/api.wafTop.v1.Strategy/GetStrategy"
const OperationStrategyListStrategy = "/api.wafTop.v1.Strategy/ListStrategy"
const OperationStrategyUpdateStrategy = "/api.wafTop.v1.Strategy/UpdateStrategy"

type StrategyHTTPServer interface {
	CreateStrategy(context.Context, *CreateStrategyRequest) (*CreateStrategyReply, error)
	DeleteStrategy(context.Context, *DeleteStrategyRequest) (*DeleteStrategyReply, error)
	GetStrategy(context.Context, *GetStrategyRequest) (*GetStrategyReply, error)
	ListStrategy(context.Context, *ListStrategyRequest) (*ListStrategyReply, error)
	UpdateStrategy(context.Context, *UpdateStrategyRequest) (*UpdateStrategyReply, error)
}

func RegisterStrategyHTTPServer(s *http.Server, srv StrategyHTTPServer) {
	r := s.Route("/")
	r.POST("/app/wafTop/v1/strategy", _Strategy_CreateStrategy0_HTTP_Handler(srv))
	r.PATCH("/app/wafTop/v1/strategy", _Strategy_UpdateStrategy0_HTTP_Handler(srv))
	r.DELETE("/app/wafTop/v1/strategy", _Strategy_DeleteStrategy0_HTTP_Handler(srv))
	r.GET("/app/wafTop/v1/strategy/{id}", _Strategy_GetStrategy0_HTTP_Handler(srv))
	r.GET("/app/wafTop/v1/strategies", _Strategy_ListStrategy0_HTTP_Handler(srv))
}

func _Strategy_CreateStrategy0_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateStrategyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStrategyCreateStrategy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateStrategy(ctx, req.(*CreateStrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateStrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_UpdateStrategy0_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateStrategyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStrategyUpdateStrategy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateStrategy(ctx, req.(*UpdateStrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateStrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_DeleteStrategy0_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteStrategyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStrategyDeleteStrategy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteStrategy(ctx, req.(*DeleteStrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteStrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_GetStrategy0_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStrategyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStrategyGetStrategy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStrategy(ctx, req.(*GetStrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_ListStrategy0_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListStrategyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStrategyListStrategy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListStrategy(ctx, req.(*ListStrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListStrategyReply)
		return ctx.Result(200, reply)
	}
}

type StrategyHTTPClient interface {
	CreateStrategy(ctx context.Context, req *CreateStrategyRequest, opts ...http.CallOption) (rsp *CreateStrategyReply, err error)
	DeleteStrategy(ctx context.Context, req *DeleteStrategyRequest, opts ...http.CallOption) (rsp *DeleteStrategyReply, err error)
	GetStrategy(ctx context.Context, req *GetStrategyRequest, opts ...http.CallOption) (rsp *GetStrategyReply, err error)
	ListStrategy(ctx context.Context, req *ListStrategyRequest, opts ...http.CallOption) (rsp *ListStrategyReply, err error)
	UpdateStrategy(ctx context.Context, req *UpdateStrategyRequest, opts ...http.CallOption) (rsp *UpdateStrategyReply, err error)
}

type StrategyHTTPClientImpl struct {
	cc *http.Client
}

func NewStrategyHTTPClient(client *http.Client) StrategyHTTPClient {
	return &StrategyHTTPClientImpl{client}
}

func (c *StrategyHTTPClientImpl) CreateStrategy(ctx context.Context, in *CreateStrategyRequest, opts ...http.CallOption) (*CreateStrategyReply, error) {
	var out CreateStrategyReply
	pattern := "/app/wafTop/v1/strategy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStrategyCreateStrategy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *StrategyHTTPClientImpl) DeleteStrategy(ctx context.Context, in *DeleteStrategyRequest, opts ...http.CallOption) (*DeleteStrategyReply, error) {
	var out DeleteStrategyReply
	pattern := "/app/wafTop/v1/strategy"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStrategyDeleteStrategy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *StrategyHTTPClientImpl) GetStrategy(ctx context.Context, in *GetStrategyRequest, opts ...http.CallOption) (*GetStrategyReply, error) {
	var out GetStrategyReply
	pattern := "/app/wafTop/v1/strategy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStrategyGetStrategy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *StrategyHTTPClientImpl) ListStrategy(ctx context.Context, in *ListStrategyRequest, opts ...http.CallOption) (*ListStrategyReply, error) {
	var out ListStrategyReply
	pattern := "/app/wafTop/v1/strategies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStrategyListStrategy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *StrategyHTTPClientImpl) UpdateStrategy(ctx context.Context, in *UpdateStrategyRequest, opts ...http.CallOption) (*UpdateStrategyReply, error) {
	var out UpdateStrategyReply
	pattern := "/app/wafTop/v1/strategy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStrategyUpdateStrategy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}