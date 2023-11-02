// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/common.proto

package organization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ResultStatus int32

const (
	ResultStatus_Success     ResultStatus = 0
	ResultStatus_MaxLimit    ResultStatus = 1
	ResultStatus_Repeated    ResultStatus = 2
	ResultStatus_NotExisted  ResultStatus = 3
	ResultStatus_DBException ResultStatus = 4
	ResultStatus_Empty       ResultStatus = 5
)

var ResultStatus_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultStatus_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultStatus) String() string {
	return proto.EnumName(ResultStatus_name, int32(x))
}

func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{0}
}

type AddressInfo struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,2,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Zone                 string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Location             string   `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressInfo) Reset()         { *m = AddressInfo{} }
func (m *AddressInfo) String() string { return proto.CompactTextString(m) }
func (*AddressInfo) ProtoMessage()    {}
func (*AddressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{0}
}

func (m *AddressInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressInfo.Unmarshal(m, b)
}
func (m *AddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressInfo.Marshal(b, m, deterministic)
}
func (m *AddressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressInfo.Merge(m, src)
}
func (m *AddressInfo) XXX_Size() int {
	return xxx_messageInfo_AddressInfo.Size(m)
}
func (m *AddressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddressInfo proto.InternalMessageInfo

func (m *AddressInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AddressInfo) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *AddressInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AddressInfo) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *AddressInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type PairInfo struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInfo) Reset()         { *m = PairInfo{} }
func (m *PairInfo) String() string { return proto.CompactTextString(m) }
func (*PairInfo) ProtoMessage()    {}
func (*PairInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{1}
}

func (m *PairInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInfo.Unmarshal(m, b)
}
func (m *PairInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInfo.Marshal(b, m, deterministic)
}
func (m *PairInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInfo.Merge(m, src)
}
func (m *PairInfo) XXX_Size() int {
	return xxx_messageInfo_PairInfo.Size(m)
}
func (m *PairInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PairInfo proto.InternalMessageInfo

func (m *PairInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PairInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AreaInfo struct {
	Id                   uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64      `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64      `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string      `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string      `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Remark               string      `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string      `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Parent               string      `protobuf:"bytes,10,opt,name=parent,proto3" json:"parent,omitempty"`
	Template             string      `protobuf:"bytes,11,opt,name=template,proto3" json:"template,omitempty"`
	Device               string      `protobuf:"bytes,12,opt,name=device,proto3" json:"device,omitempty"`
	Question             string      `protobuf:"bytes,13,opt,name=question,proto3" json:"question,omitempty"`
	Type                 uint32      `protobuf:"varint,14,opt,name=type,proto3" json:"type,omitempty"`
	Catalog              string      `protobuf:"bytes,15,opt,name=catalog,proto3" json:"catalog,omitempty"`
	Sn                   string      `protobuf:"bytes,16,opt,name=sn,proto3" json:"sn,omitempty"`
	Aspect               string      `protobuf:"bytes,17,opt,name=aspect,proto3" json:"aspect,omitempty"`
	Limit                uint32      `protobuf:"varint,18,opt,name=limit,proto3" json:"limit,omitempty"`
	Displays             []string    `protobuf:"bytes,21,rep,name=displays,proto3" json:"displays,omitempty"`
	Modules              []*PairInfo `protobuf:"bytes,31,rep,name=modules,proto3" json:"modules,omitempty"`
	Sources              []*PairInfo `protobuf:"bytes,32,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AreaInfo) Reset()         { *m = AreaInfo{} }
func (m *AreaInfo) String() string { return proto.CompactTextString(m) }
func (*AreaInfo) ProtoMessage()    {}
func (*AreaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{2}
}

func (m *AreaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaInfo.Unmarshal(m, b)
}
func (m *AreaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaInfo.Marshal(b, m, deterministic)
}
func (m *AreaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaInfo.Merge(m, src)
}
func (m *AreaInfo) XXX_Size() int {
	return xxx_messageInfo_AreaInfo.Size(m)
}
func (m *AreaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AreaInfo proto.InternalMessageInfo

func (m *AreaInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AreaInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AreaInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AreaInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AreaInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *AreaInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AreaInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *AreaInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *AreaInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *AreaInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *AreaInfo) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *AreaInfo) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *AreaInfo) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *AreaInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AreaInfo) GetCatalog() string {
	if m != nil {
		return m.Catalog
	}
	return ""
}

func (m *AreaInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *AreaInfo) GetAspect() string {
	if m != nil {
		return m.Aspect
	}
	return ""
}

func (m *AreaInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AreaInfo) GetDisplays() []string {
	if m != nil {
		return m.Displays
	}
	return nil
}

func (m *AreaInfo) GetModules() []*PairInfo {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *AreaInfo) GetSources() []*PairInfo {
	if m != nil {
		return m.Sources
	}
	return nil
}

type RequestInfo struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{3}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestPage struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Page                 uint32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPage) Reset()         { *m = RequestPage{} }
func (m *RequestPage) String() string { return proto.CompactTextString(m) }
func (*RequestPage) ProtoMessage()    {}
func (*RequestPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{4}
}

func (m *RequestPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPage.Unmarshal(m, b)
}
func (m *RequestPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPage.Marshal(b, m, deterministic)
}
func (m *RequestPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPage.Merge(m, src)
}
func (m *RequestPage) XXX_Size() int {
	return xxx_messageInfo_RequestPage.Size(m)
}
func (m *RequestPage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPage proto.InternalMessageInfo

func (m *RequestPage) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RequestPage) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type RequestFilter struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []string `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{5}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *RequestFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestFilter) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqUpdateFilter struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateFilter) Reset()         { *m = ReqUpdateFilter{} }
func (m *ReqUpdateFilter) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateFilter) ProtoMessage()    {}
func (*ReqUpdateFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{6}
}

func (m *ReqUpdateFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateFilter.Unmarshal(m, b)
}
func (m *ReqUpdateFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateFilter.Marshal(b, m, deterministic)
}
func (m *ReqUpdateFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateFilter.Merge(m, src)
}
func (m *ReqUpdateFilter) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateFilter.Size(m)
}
func (m *ReqUpdateFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateFilter proto.InternalMessageInfo

func (m *ReqUpdateFilter) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqUpdateFilter) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqUpdateFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqUpdateFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqUpdateFilter) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqUpdateFilter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{7}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{8}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type RequestAddress struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Zone                 string   `protobuf:"bytes,5,opt,name=zone,proto3" json:"zone,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAddress) Reset()         { *m = RequestAddress{} }
func (m *RequestAddress) String() string { return proto.CompactTextString(m) }
func (*RequestAddress) ProtoMessage()    {}
func (*RequestAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{9}
}

func (m *RequestAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAddress.Unmarshal(m, b)
}
func (m *RequestAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAddress.Marshal(b, m, deterministic)
}
func (m *RequestAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAddress.Merge(m, src)
}
func (m *RequestAddress) XXX_Size() int {
	return xxx_messageInfo_RequestAddress.Size(m)
}
func (m *RequestAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAddress.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAddress proto.InternalMessageInfo

func (m *RequestAddress) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestAddress) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *RequestAddress) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *RequestAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RequestAddress) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *RequestAddress) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *RequestAddress) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFlag) Reset()         { *m = RequestFlag{} }
func (m *RequestFlag) String() string { return proto.CompactTextString(m) }
func (*RequestFlag) ProtoMessage()    {}
func (*RequestFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{10}
}

func (m *RequestFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFlag.Unmarshal(m, b)
}
func (m *RequestFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFlag.Marshal(b, m, deterministic)
}
func (m *RequestFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFlag.Merge(m, src)
}
func (m *RequestFlag) XXX_Size() int {
	return xxx_messageInfo_RequestFlag.Size(m)
}
func (m *RequestFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFlag.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFlag proto.InternalMessageInfo

func (m *RequestFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestFlag) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *RequestFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{11}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []string     `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{12}
}

func (m *ReplyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyList.Unmarshal(m, b)
}
func (m *ReplyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyList.Marshal(b, m, deterministic)
}
func (m *ReplyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyList.Merge(m, src)
}
func (m *ReplyList) XXX_Size() int {
	return xxx_messageInfo_ReplyList.Size(m)
}
func (m *ReplyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyList proto.InternalMessageInfo

func (m *ReplyList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{13}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("omo.msp.scene.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterType((*AddressInfo)(nil), "omo.msp.scene.AddressInfo")
	proto.RegisterType((*PairInfo)(nil), "omo.msp.scene.PairInfo")
	proto.RegisterType((*AreaInfo)(nil), "omo.msp.scene.AreaInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.scene.RequestInfo")
	proto.RegisterType((*RequestPage)(nil), "omo.msp.scene.RequestPage")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.scene.RequestFilter")
	proto.RegisterType((*ReqUpdateFilter)(nil), "omo.msp.scene.ReqUpdateFilter")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.scene.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.scene.ReplyInfo")
	proto.RegisterType((*RequestAddress)(nil), "omo.msp.scene.RequestAddress")
	proto.RegisterType((*RequestFlag)(nil), "omo.msp.scene.RequestFlag")
	proto.RegisterType((*RequestList)(nil), "omo.msp.scene.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.scene.ReplyList")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.scene.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/organization/common.proto", fileDescriptor_5b563812ef86dc2c)
}

var fileDescriptor_5b563812ef86dc2c = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x3b, 0x8f, 0x1b, 0x37,
	0x10, 0xce, 0xbe, 0xf4, 0xa0, 0x4e, 0x3a, 0x65, 0x71, 0x71, 0x08, 0x37, 0x16, 0xb6, 0x3a, 0xa4,
	0x90, 0x11, 0xa5, 0x48, 0x6d, 0x23, 0x67, 0x20, 0x80, 0x93, 0xd8, 0x7b, 0x48, 0x93, 0x8e, 0xde,
	0x1d, 0x0b, 0x84, 0x77, 0x97, 0x6b, 0x92, 0x7b, 0x39, 0x5d, 0x99, 0x26, 0x55, 0xca, 0xfc, 0x8d,
	0xd4, 0xf9, 0x79, 0x06, 0x87, 0xa4, 0xb4, 0xd2, 0x49, 0x77, 0xd7, 0xcd, 0xb7, 0xe4, 0x7c, 0xf3,
	0xcd, 0x83, 0x23, 0x91, 0x17, 0xad, 0x14, 0x5a, 0xbc, 0x14, 0x72, 0xcd, 0x1a, 0x7e, 0xc7, 0x34,
	0x17, 0xcd, 0xcb, 0x42, 0xd4, 0xb5, 0x68, 0x96, 0x78, 0x92, 0x4e, 0x45, 0x2d, 0x96, 0xb5, 0x6a,
	0x97, 0xaa, 0x80, 0x06, 0xb2, 0xbf, 0x03, 0x32, 0x79, 0x55, 0x96, 0x12, 0x94, 0xfa, 0xb9, 0xf9,
	0x28, 0x52, 0x4a, 0x86, 0x85, 0xe8, 0x1a, 0x2d, 0x37, 0x34, 0x58, 0x04, 0x97, 0xe3, 0xdc, 0xc3,
	0xf4, 0x39, 0x19, 0xb5, 0x52, 0xdc, 0xf0, 0xa6, 0x00, 0x1a, 0xe2, 0xd1, 0x16, 0xa7, 0x29, 0x89,
	0x0b, 0xae, 0x37, 0x34, 0xc2, 0xef, 0x68, 0x9b, 0x6f, 0x77, 0xa2, 0x01, 0x1a, 0xdb, 0x6f, 0xc6,
	0x36, 0x1c, 0x95, 0x28, 0x50, 0x15, 0x4d, 0x2c, 0x87, 0xc7, 0xd9, 0x8a, 0x8c, 0xde, 0x31, 0x2e,
	0x51, 0xc5, 0x9c, 0x44, 0x9f, 0xc0, 0x2b, 0x30, 0x66, 0x7a, 0x41, 0x92, 0x1b, 0x56, 0x75, 0x3e,
	0xb4, 0x05, 0xd9, 0x7f, 0x31, 0x19, 0xbd, 0x92, 0xc0, 0xd0, 0x69, 0x46, 0x42, 0x5e, 0xa2, 0x4f,
	0x9c, 0x87, 0xbc, 0x34, 0x24, 0x1d, 0x2f, 0x9d, 0x83, 0x31, 0x8d, 0xa4, 0x86, 0xd5, 0xe0, 0x65,
	0x1a, 0x1b, 0x13, 0x96, 0xc0, 0x34, 0x94, 0xa8, 0x34, 0xce, 0x3d, 0x34, 0x27, 0x5d, 0x5b, 0xe2,
	0x49, 0x62, 0x4f, 0x1c, 0xdc, 0xfa, 0x08, 0x49, 0x07, 0xae, 0x48, 0x16, 0x9a, 0x04, 0x45, 0x0b,
	0x12, 0x8f, 0x86, 0x36, 0x41, 0x8f, 0xd3, 0x67, 0x64, 0x20, 0xa1, 0x66, 0xf2, 0x13, 0x1d, 0xe1,
	0x89, 0x43, 0x26, 0x35, 0xf1, 0x67, 0x03, 0x92, 0x8e, 0x6d, 0x6a, 0x08, 0xcc, 0xed, 0x96, 0x49,
	0x68, 0x34, 0x25, 0xf6, 0xb6, 0x45, 0x26, 0x82, 0x86, 0xba, 0xad, 0x98, 0x06, 0x3a, 0xb1, 0x11,
	0x3c, 0x36, 0x3e, 0x25, 0xdc, 0xf0, 0x02, 0xe8, 0x99, 0xf5, 0xb1, 0xc8, 0xf8, 0x7c, 0xee, 0x40,
	0x61, 0xd9, 0xa7, 0xd6, 0xc7, 0x63, 0x53, 0x13, 0xbd, 0x69, 0x81, 0xce, 0x16, 0xc1, 0xe5, 0x34,
	0x47, 0x1b, 0xf3, 0x63, 0x9a, 0x55, 0x62, 0x4d, 0xcf, 0x5d, 0x7e, 0x16, 0x9a, 0x1a, 0xab, 0x86,
	0xce, 0xf1, 0x63, 0xa8, 0x1a, 0x13, 0x91, 0xa9, 0x16, 0x0a, 0x4d, 0xbf, 0xb6, 0x11, 0x2d, 0x32,
	0x39, 0x55, 0xbc, 0xe6, 0x9a, 0xa6, 0x48, 0x6b, 0x81, 0xd1, 0x51, 0x72, 0xd5, 0x56, 0x6c, 0xa3,
	0xe8, 0x37, 0x8b, 0xc8, 0xe8, 0xf0, 0x38, 0xfd, 0x9e, 0x0c, 0x6b, 0x51, 0x76, 0x15, 0x28, 0xfa,
	0x62, 0x11, 0x5d, 0x4e, 0x56, 0xdf, 0x2e, 0xf7, 0x26, 0x75, 0xe9, 0x87, 0x23, 0xf7, 0xf7, 0x8c,
	0x8b, 0x12, 0x9d, 0x2c, 0x40, 0xd1, 0xc5, 0x23, 0x2e, 0xee, 0x5e, 0x76, 0x4d, 0x26, 0x39, 0x60,
	0xee, 0x38, 0x32, 0xbb, 0x22, 0x07, 0x7b, 0x45, 0xbe, 0x3f, 0x3a, 0xfd, 0xc6, 0x46, 0xfb, 0x8d,
	0xcd, 0xde, 0x6f, 0x49, 0xdf, 0xb1, 0x35, 0x9c, 0x24, 0x7d, 0x46, 0x06, 0x4d, 0x57, 0x7f, 0x00,
	0x89, 0xbc, 0xd3, 0xdc, 0x21, 0xd3, 0x81, 0x96, 0xad, 0xed, 0x54, 0x4e, 0x73, 0xb4, 0xb3, 0x7f,
	0x02, 0x32, 0x75, 0x9c, 0x6f, 0x78, 0xa5, 0x41, 0x9a, 0x8a, 0x62, 0x52, 0x8e, 0xd4, 0x02, 0xff,
	0x50, 0xc2, 0x23, 0x0f, 0x25, 0xea, 0x3d, 0x94, 0x6d, 0x8c, 0x78, 0x17, 0xa3, 0xa7, 0x27, 0x39,
	0xd4, 0x53, 0x71, 0xa5, 0xe9, 0x00, 0x3b, 0x84, 0x76, 0xf6, 0x6f, 0x40, 0xce, 0x73, 0xf8, 0xfc,
	0x3b, 0x3e, 0x80, 0xc7, 0x14, 0x1d, 0x94, 0xce, 0x69, 0x8c, 0x8e, 0x68, 0x8c, 0xfb, 0x1a, 0xfb,
	0x25, 0x4e, 0xee, 0xbf, 0x1d, 0xbc, 0xa4, 0x9c, 0x2a, 0x87, 0xb2, 0x1f, 0x4d, 0xe9, 0xdb, 0x6a,
	0x73, 0xad, 0x99, 0xee, 0x14, 0xee, 0x21, 0x51, 0x5a, 0x45, 0xd3, 0x1c, 0x6d, 0x13, 0x0c, 0xa4,
	0x14, 0xd2, 0x6f, 0x0e, 0x04, 0xd9, 0x7b, 0x32, 0x46, 0x47, 0xbf, 0x6e, 0x3a, 0xb7, 0x3a, 0x9c,
	0xe6, 0x15, 0x19, 0x28, 0xa4, 0x44, 0xaf, 0xc9, 0xea, 0xf9, 0xc1, 0x64, 0xf5, 0x82, 0xe6, 0xee,
	0x66, 0xf6, 0x7f, 0x40, 0x66, 0xae, 0x67, 0x6e, 0xa3, 0x1e, 0x21, 0xee, 0xed, 0xd7, 0xf0, 0xf4,
	0x7e, 0x8d, 0x4e, 0xec, 0xd7, 0xf8, 0xc8, 0x7e, 0x4d, 0x4e, 0xec, 0xd7, 0xc1, 0xfe, 0x7e, 0x7d,
	0x68, 0x35, 0x65, 0xbf, 0x6d, 0x27, 0xf8, 0x4d, 0xc5, 0xd6, 0x47, 0x64, 0xa7, 0x24, 0xfe, 0x58,
	0xb1, 0xb5, 0xd3, 0x8c, 0xf6, 0x83, 0x4f, 0x62, 0x47, 0xf8, 0x96, 0x2b, 0x7d, 0x84, 0xb0, 0xef,
	0x1c, 0x1e, 0x34, 0xdb, 0x0f, 0x60, 0xd4, 0x1b, 0x40, 0x70, 0xfd, 0x3a, 0x41, 0xe7, 0x5d, 0xc2,
	0x9d, 0x4b, 0xaf, 0x87, 0xd1, 0x93, 0x7b, 0xf8, 0x17, 0xf6, 0xd0, 0x7d, 0xe7, 0x4a, 0xf3, 0xa2,
	0x47, 0x13, 0x3c, 0x95, 0xe6, 0xf8, 0xb3, 0xb4, 0x4b, 0x3e, 0xea, 0x2f, 0xf9, 0x0b, 0x92, 0x60,
	0xfb, 0xdd, 0xbb, 0xb4, 0xe0, 0xbb, 0x82, 0x9c, 0xe5, 0xa0, 0xba, 0x4a, 0xbb, 0xa9, 0x9e, 0x90,
	0xe1, 0x75, 0x57, 0x14, 0xa0, 0xd4, 0xfc, 0xab, 0xf4, 0x8c, 0x8c, 0x7e, 0x61, 0xb7, 0x6f, 0xcd,
	0x3e, 0x9d, 0x07, 0x06, 0xe5, 0xd0, 0xe2, 0xef, 0xd5, 0x3c, 0x4c, 0x67, 0x84, 0xfc, 0x2a, 0xf4,
	0xd5, 0x2d, 0x57, 0x06, 0x47, 0xe9, 0x39, 0x99, 0xfc, 0xf4, 0xfa, 0xea, 0xb6, 0x80, 0xd6, 0x4c,
	0xc0, 0x3c, 0x4e, 0xc7, 0x24, 0xb9, 0xaa, 0x5b, 0xbd, 0x99, 0x27, 0xaf, 0x2f, 0xfe, 0x48, 0xef,
	0xff, 0x55, 0xf8, 0x30, 0xc0, 0x6f, 0x3f, 0x7c, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x34, 0xf7,
	0xc1, 0x47, 0x08, 0x00, 0x00,
}
