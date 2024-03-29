// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/organization/group.proto

package organization

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

// Client API for GroupService service

type GroupService interface {
	AddOne(ctx context.Context, in *ReqGroupAdd, opts ...client.CallOption) (*ReplyGroupOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupOne, error)
	GetByUser(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupList, error)
	GetByContact(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupList, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyGroupList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqGroupUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateMaster(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateAssistant(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateContact(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateLocation(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateAddress(ctx context.Context, in *RequestAddress, opts ...client.CallOption) (*ReplyGroupOne, error)
	AppendMember(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
	SubtractMember(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
	UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error)
}

type groupService struct {
	c    client.Client
	name string
}

func NewGroupService(name string, c client.Client) GroupService {
	return &groupService{
		c:    c,
		name: name,
	}
}

func (c *groupService) AddOne(ctx context.Context, in *ReqGroupAdd, opts ...client.CallOption) (*ReplyGroupOne, error) {
	req := c.c.NewRequest(c.name, "GroupService.AddOne", in)
	out := new(ReplyGroupOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupOne, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetOne", in)
	out := new(ReplyGroupOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetByUser(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupList, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetByUser", in)
	out := new(ReplyGroupList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetByContact(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGroupList, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetByContact", in)
	out := new(ReplyGroupList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyGroupList, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetList", in)
	out := new(ReplyGroupList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateBase(ctx context.Context, in *ReqGroupUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateMaster(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateMaster", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateAssistant(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateAssistant", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateContact(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateContact", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateLocation(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateLocation", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateAddress(ctx context.Context, in *RequestAddress, opts ...client.CallOption) (*ReplyGroupOne, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateAddress", in)
	out := new(ReplyGroupOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) AppendMember(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "GroupService.AppendMember", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) SubtractMember(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "GroupService.SubtractMember", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GroupService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceHandler interface {
	AddOne(context.Context, *ReqGroupAdd, *ReplyGroupOne) error
	GetOne(context.Context, *RequestInfo, *ReplyGroupOne) error
	GetByUser(context.Context, *RequestInfo, *ReplyGroupList) error
	GetByContact(context.Context, *RequestInfo, *ReplyGroupList) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *RequestPage, *ReplyGroupList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqGroupUpdate, *ReplyInfo) error
	UpdateMaster(context.Context, *RequestFlag, *ReplyInfo) error
	UpdateAssistant(context.Context, *RequestFlag, *ReplyInfo) error
	UpdateContact(context.Context, *RequestFlag, *ReplyInfo) error
	UpdateLocation(context.Context, *RequestFlag, *ReplyInfo) error
	UpdateAddress(context.Context, *RequestAddress, *ReplyGroupOne) error
	AppendMember(context.Context, *RequestInfo, *ReplyList) error
	SubtractMember(context.Context, *RequestInfo, *ReplyList) error
	UpdateByFilter(context.Context, *ReqUpdateFilter, *ReplyInfo) error
}

func RegisterGroupServiceHandler(s server.Server, hdlr GroupServiceHandler, opts ...server.HandlerOption) error {
	type groupService interface {
		AddOne(ctx context.Context, in *ReqGroupAdd, out *ReplyGroupOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyGroupOne) error
		GetByUser(ctx context.Context, in *RequestInfo, out *ReplyGroupList) error
		GetByContact(ctx context.Context, in *RequestInfo, out *ReplyGroupList) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *RequestPage, out *ReplyGroupList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqGroupUpdate, out *ReplyInfo) error
		UpdateMaster(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		UpdateAssistant(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		UpdateContact(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		UpdateLocation(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		UpdateAddress(ctx context.Context, in *RequestAddress, out *ReplyGroupOne) error
		AppendMember(ctx context.Context, in *RequestInfo, out *ReplyList) error
		SubtractMember(ctx context.Context, in *RequestInfo, out *ReplyList) error
		UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error
	}
	type GroupService struct {
		groupService
	}
	h := &groupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GroupService{h}, opts...))
}

type groupServiceHandler struct {
	GroupServiceHandler
}

func (h *groupServiceHandler) AddOne(ctx context.Context, in *ReqGroupAdd, out *ReplyGroupOne) error {
	return h.GroupServiceHandler.AddOne(ctx, in, out)
}

func (h *groupServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyGroupOne) error {
	return h.GroupServiceHandler.GetOne(ctx, in, out)
}

func (h *groupServiceHandler) GetByUser(ctx context.Context, in *RequestInfo, out *ReplyGroupList) error {
	return h.GroupServiceHandler.GetByUser(ctx, in, out)
}

func (h *groupServiceHandler) GetByContact(ctx context.Context, in *RequestInfo, out *ReplyGroupList) error {
	return h.GroupServiceHandler.GetByContact(ctx, in, out)
}

func (h *groupServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.GroupServiceHandler.RemoveOne(ctx, in, out)
}

func (h *groupServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplyGroupList) error {
	return h.GroupServiceHandler.GetList(ctx, in, out)
}

func (h *groupServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.GroupServiceHandler.GetStatistic(ctx, in, out)
}

func (h *groupServiceHandler) UpdateBase(ctx context.Context, in *ReqGroupUpdate, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateBase(ctx, in, out)
}

func (h *groupServiceHandler) UpdateMaster(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateMaster(ctx, in, out)
}

func (h *groupServiceHandler) UpdateAssistant(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateAssistant(ctx, in, out)
}

func (h *groupServiceHandler) UpdateContact(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateContact(ctx, in, out)
}

func (h *groupServiceHandler) UpdateLocation(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateLocation(ctx, in, out)
}

func (h *groupServiceHandler) UpdateAddress(ctx context.Context, in *RequestAddress, out *ReplyGroupOne) error {
	return h.GroupServiceHandler.UpdateAddress(ctx, in, out)
}

func (h *groupServiceHandler) AppendMember(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.GroupServiceHandler.AppendMember(ctx, in, out)
}

func (h *groupServiceHandler) SubtractMember(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.GroupServiceHandler.SubtractMember(ctx, in, out)
}

func (h *groupServiceHandler) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error {
	return h.GroupServiceHandler.UpdateByFilter(ctx, in, out)
}
