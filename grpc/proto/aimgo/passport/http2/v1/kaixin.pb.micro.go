// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: kaixin.proto

package aimgo_passport_http2_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for KaixinService service

type KaixinService interface {
	Kaixin(ctx context.Context, in *KaixinRequest, opts ...client.CallOption) (*KaixinResponse, error)
}

type kaixinService struct {
	c    client.Client
	name string
}

func NewKaixinService(name string, c client.Client) KaixinService {
	return &kaixinService{
		c:    c,
		name: name,
	}
}

func (c *kaixinService) Kaixin(ctx context.Context, in *KaixinRequest, opts ...client.CallOption) (*KaixinResponse, error) {
	req := c.c.NewRequest(c.name, "KaixinService.Kaixin", in)
	out := new(KaixinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KaixinService service

type KaixinServiceHandler interface {
	Kaixin(context.Context, *KaixinRequest, *KaixinResponse) error
}

func RegisterKaixinServiceHandler(s server.Server, hdlr KaixinServiceHandler, opts ...server.HandlerOption) error {
	type kaixinService interface {
		Kaixin(ctx context.Context, in *KaixinRequest, out *KaixinResponse) error
	}
	type KaixinService struct {
		kaixinService
	}
	h := &kaixinServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&KaixinService{h}, opts...))
}

type kaixinServiceHandler struct {
	KaixinServiceHandler
}

func (h *kaixinServiceHandler) Kaixin(ctx context.Context, in *KaixinRequest, out *KaixinResponse) error {
	return h.KaixinServiceHandler.Kaixin(ctx, in, out)
}

// Client API for WeiwuService service

type WeiwuService interface {
	Haha(ctx context.Context, in *HahaRequest, opts ...client.CallOption) (*HahaResponse, error)
}

type weiwuService struct {
	c    client.Client
	name string
}

func NewWeiwuService(name string, c client.Client) WeiwuService {
	return &weiwuService{
		c:    c,
		name: name,
	}
}

func (c *weiwuService) Haha(ctx context.Context, in *HahaRequest, opts ...client.CallOption) (*HahaResponse, error) {
	req := c.c.NewRequest(c.name, "WeiwuService.Haha", in)
	out := new(HahaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WeiwuService service

type WeiwuServiceHandler interface {
	Haha(context.Context, *HahaRequest, *HahaResponse) error
}

func RegisterWeiwuServiceHandler(s server.Server, hdlr WeiwuServiceHandler, opts ...server.HandlerOption) error {
	type weiwuService interface {
		Haha(ctx context.Context, in *HahaRequest, out *HahaResponse) error
	}
	type WeiwuService struct {
		weiwuService
	}
	h := &weiwuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WeiwuService{h}, opts...))
}

type weiwuServiceHandler struct {
	WeiwuServiceHandler
}

func (h *weiwuServiceHandler) Haha(ctx context.Context, in *HahaRequest, out *HahaResponse) error {
	return h.WeiwuServiceHandler.Haha(ctx, in, out)
}
