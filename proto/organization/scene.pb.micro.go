// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/organization/scene.proto

package organization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SceneService service

func NewSceneServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SceneService service

type SceneService interface {
	AddOne(ctx context.Context, in *ReqSceneAdd, opts ...client.CallOption) (*ReplySceneOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySceneOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	IsMasterUsed(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyMasterUsed, error)
	GetList(ctx context.Context, in *ReqScenePage, opts ...client.CallOption) (*ReplySceneList, error)
	UpdateBase(ctx context.Context, in *ReqSceneUpdate, opts ...client.CallOption) (*ReplySceneOne, error)
	UpdateStatus(ctx context.Context, in *ReqSceneStatus, opts ...client.CallOption) (*ReplySceneOne, error)
	AppendMember(ctx context.Context, in *ReqSceneMember, opts ...client.CallOption) (*ReplySceneMembers, error)
	SubtractMember(ctx context.Context, in *ReqSceneMember, opts ...client.CallOption) (*ReplySceneMembers, error)
}

type sceneService struct {
	c    client.Client
	name string
}

func NewSceneService(name string, c client.Client) SceneService {
	return &sceneService{
		c:    c,
		name: name,
	}
}

func (c *sceneService) AddOne(ctx context.Context, in *ReqSceneAdd, opts ...client.CallOption) (*ReplySceneOne, error) {
	req := c.c.NewRequest(c.name, "SceneService.AddOne", in)
	out := new(ReplySceneOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySceneOne, error) {
	req := c.c.NewRequest(c.name, "SceneService.GetOne", in)
	out := new(ReplySceneOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SceneService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) IsMasterUsed(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyMasterUsed, error) {
	req := c.c.NewRequest(c.name, "SceneService.IsMasterUsed", in)
	out := new(ReplyMasterUsed)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) GetList(ctx context.Context, in *ReqScenePage, opts ...client.CallOption) (*ReplySceneList, error) {
	req := c.c.NewRequest(c.name, "SceneService.GetList", in)
	out := new(ReplySceneList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) UpdateBase(ctx context.Context, in *ReqSceneUpdate, opts ...client.CallOption) (*ReplySceneOne, error) {
	req := c.c.NewRequest(c.name, "SceneService.UpdateBase", in)
	out := new(ReplySceneOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) UpdateStatus(ctx context.Context, in *ReqSceneStatus, opts ...client.CallOption) (*ReplySceneOne, error) {
	req := c.c.NewRequest(c.name, "SceneService.UpdateStatus", in)
	out := new(ReplySceneOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) AppendMember(ctx context.Context, in *ReqSceneMember, opts ...client.CallOption) (*ReplySceneMembers, error) {
	req := c.c.NewRequest(c.name, "SceneService.AppendMember", in)
	out := new(ReplySceneMembers)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) SubtractMember(ctx context.Context, in *ReqSceneMember, opts ...client.CallOption) (*ReplySceneMembers, error) {
	req := c.c.NewRequest(c.name, "SceneService.SubtractMember", in)
	out := new(ReplySceneMembers)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SceneService service

type SceneServiceHandler interface {
	AddOne(context.Context, *ReqSceneAdd, *ReplySceneOne) error
	GetOne(context.Context, *RequestInfo, *ReplySceneOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	IsMasterUsed(context.Context, *RequestInfo, *ReplyMasterUsed) error
	GetList(context.Context, *ReqScenePage, *ReplySceneList) error
	UpdateBase(context.Context, *ReqSceneUpdate, *ReplySceneOne) error
	UpdateStatus(context.Context, *ReqSceneStatus, *ReplySceneOne) error
	AppendMember(context.Context, *ReqSceneMember, *ReplySceneMembers) error
	SubtractMember(context.Context, *ReqSceneMember, *ReplySceneMembers) error
}

func RegisterSceneServiceHandler(s server.Server, hdlr SceneServiceHandler, opts ...server.HandlerOption) error {
	type sceneService interface {
		AddOne(ctx context.Context, in *ReqSceneAdd, out *ReplySceneOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplySceneOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		IsMasterUsed(ctx context.Context, in *RequestInfo, out *ReplyMasterUsed) error
		GetList(ctx context.Context, in *ReqScenePage, out *ReplySceneList) error
		UpdateBase(ctx context.Context, in *ReqSceneUpdate, out *ReplySceneOne) error
		UpdateStatus(ctx context.Context, in *ReqSceneStatus, out *ReplySceneOne) error
		AppendMember(ctx context.Context, in *ReqSceneMember, out *ReplySceneMembers) error
		SubtractMember(ctx context.Context, in *ReqSceneMember, out *ReplySceneMembers) error
	}
	type SceneService struct {
		sceneService
	}
	h := &sceneServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SceneService{h}, opts...))
}

type sceneServiceHandler struct {
	SceneServiceHandler
}

func (h *sceneServiceHandler) AddOne(ctx context.Context, in *ReqSceneAdd, out *ReplySceneOne) error {
	return h.SceneServiceHandler.AddOne(ctx, in, out)
}

func (h *sceneServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplySceneOne) error {
	return h.SceneServiceHandler.GetOne(ctx, in, out)
}

func (h *sceneServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.SceneServiceHandler.RemoveOne(ctx, in, out)
}

func (h *sceneServiceHandler) IsMasterUsed(ctx context.Context, in *RequestInfo, out *ReplyMasterUsed) error {
	return h.SceneServiceHandler.IsMasterUsed(ctx, in, out)
}

func (h *sceneServiceHandler) GetList(ctx context.Context, in *ReqScenePage, out *ReplySceneList) error {
	return h.SceneServiceHandler.GetList(ctx, in, out)
}

func (h *sceneServiceHandler) UpdateBase(ctx context.Context, in *ReqSceneUpdate, out *ReplySceneOne) error {
	return h.SceneServiceHandler.UpdateBase(ctx, in, out)
}

func (h *sceneServiceHandler) UpdateStatus(ctx context.Context, in *ReqSceneStatus, out *ReplySceneOne) error {
	return h.SceneServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *sceneServiceHandler) AppendMember(ctx context.Context, in *ReqSceneMember, out *ReplySceneMembers) error {
	return h.SceneServiceHandler.AppendMember(ctx, in, out)
}

func (h *sceneServiceHandler) SubtractMember(ctx context.Context, in *ReqSceneMember, out *ReplySceneMembers) error {
	return h.SceneServiceHandler.SubtractMember(ctx, in, out)
}
