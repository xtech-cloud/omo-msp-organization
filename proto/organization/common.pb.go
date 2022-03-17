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
	return fileDescriptor_5b563812ef86dc2c, []int{1}
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
	return fileDescriptor_5b563812ef86dc2c, []int{2}
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
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
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
	return fileDescriptor_5b563812ef86dc2c, []int{3}
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

func (m *RequestFilter) GetParent() string {
	if m != nil {
		return m.Parent
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

type DisplayInfo struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Prepares             []string `protobuf:"bytes,2,rep,name=prepares,proto3" json:"prepares,omitempty"`
	Showings             []string `protobuf:"bytes,3,rep,name=showings,proto3" json:"showings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayInfo) Reset()         { *m = DisplayInfo{} }
func (m *DisplayInfo) String() string { return proto.CompactTextString(m) }
func (*DisplayInfo) ProtoMessage()    {}
func (*DisplayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{4}
}

func (m *DisplayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayInfo.Unmarshal(m, b)
}
func (m *DisplayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayInfo.Marshal(b, m, deterministic)
}
func (m *DisplayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayInfo.Merge(m, src)
}
func (m *DisplayInfo) XXX_Size() int {
	return xxx_messageInfo_DisplayInfo.Size(m)
}
func (m *DisplayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayInfo proto.InternalMessageInfo

func (m *DisplayInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *DisplayInfo) GetPrepares() []string {
	if m != nil {
		return m.Prepares
	}
	return nil
}

func (m *DisplayInfo) GetShowings() []string {
	if m != nil {
		return m.Showings
	}
	return nil
}

type ProductInfo struct {
	Type                 uint32         `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Uid                  string         `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Remark               string         `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Displays             []*DisplayInfo `protobuf:"bytes,4,rep,name=displays,proto3" json:"displays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProductInfo) Reset()         { *m = ProductInfo{} }
func (m *ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()    {}
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b563812ef86dc2c, []int{5}
}

func (m *ProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInfo.Unmarshal(m, b)
}
func (m *ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInfo.Marshal(b, m, deterministic)
}
func (m *ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfo.Merge(m, src)
}
func (m *ProductInfo) XXX_Size() int {
	return xxx_messageInfo_ProductInfo.Size(m)
}
func (m *ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfo proto.InternalMessageInfo

func (m *ProductInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ProductInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ProductInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ProductInfo) GetDisplays() []*DisplayInfo {
	if m != nil {
		return m.Displays
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
	return fileDescriptor_5b563812ef86dc2c, []int{6}
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
	return fileDescriptor_5b563812ef86dc2c, []int{7}
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
	return fileDescriptor_5b563812ef86dc2c, []int{8}
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
	return fileDescriptor_5b563812ef86dc2c, []int{9}
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
	return fileDescriptor_5b563812ef86dc2c, []int{10}
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
	return fileDescriptor_5b563812ef86dc2c, []int{11}
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
	return fileDescriptor_5b563812ef86dc2c, []int{12}
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
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.scene.RequestInfo")
	proto.RegisterType((*RequestPage)(nil), "omo.msp.scene.RequestPage")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.scene.RequestFilter")
	proto.RegisterType((*DisplayInfo)(nil), "omo.msp.scene.DisplayInfo")
	proto.RegisterType((*ProductInfo)(nil), "omo.msp.scene.ProductInfo")
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
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x4f, 0xdc, 0x30,
	0x10, 0x6d, 0x36, 0x1f, 0xc0, 0x84, 0xa5, 0x2b, 0x0b, 0xa1, 0x88, 0x4b, 0x51, 0x4e, 0xa8, 0x87,
	0x45, 0xa2, 0x52, 0x7b, 0x2e, 0x02, 0xa4, 0x4a, 0xb4, 0x85, 0x70, 0x6b, 0x4f, 0x26, 0x19, 0x52,
	0x8b, 0x24, 0x4e, 0x6d, 0x07, 0x08, 0xc7, 0xf6, 0xd0, 0x63, 0xff, 0x4e, 0x7f, 0x5e, 0x65, 0xc7,
	0x9b, 0xcd, 0xc2, 0x2e, 0xe2, 0x36, 0x6f, 0x9c, 0x79, 0x7e, 0xe3, 0x99, 0xa7, 0xc0, 0x9b, 0x5a,
	0x70, 0xc5, 0x0f, 0xb8, 0xc8, 0x69, 0xc5, 0x1e, 0xa8, 0x62, 0xbc, 0x3a, 0x48, 0x79, 0x59, 0xf2,
	0x6a, 0x6a, 0x4e, 0xc8, 0x98, 0x97, 0x7c, 0x5a, 0xca, 0x7a, 0x2a, 0x53, 0xac, 0x30, 0xfe, 0xe3,
	0x40, 0xf8, 0x31, 0xcb, 0x04, 0x4a, 0xf9, 0xa9, 0xba, 0xe6, 0x24, 0x82, 0xb5, 0x94, 0x37, 0x95,
	0x12, 0x6d, 0xe4, 0xec, 0x39, 0xfb, 0x1b, 0xc9, 0x0c, 0x92, 0x5d, 0x58, 0xaf, 0x05, 0xbf, 0x65,
	0x55, 0x8a, 0xd1, 0xc8, 0x1c, 0xf5, 0x98, 0x10, 0xf0, 0x52, 0xa6, 0xda, 0xc8, 0x35, 0x79, 0x13,
	0xeb, 0xdc, 0x03, 0xaf, 0x30, 0xf2, 0xba, 0x9c, 0x8e, 0x35, 0x47, 0xc1, 0x53, 0xa3, 0x2a, 0xf2,
	0x3b, 0x8e, 0x19, 0x8e, 0x2f, 0x21, 0x4c, 0xf0, 0x67, 0x83, 0x52, 0x19, 0x21, 0x3b, 0x10, 0xd4,
	0x54, 0x60, 0xa5, 0xac, 0x0e, 0x8b, 0xc8, 0x04, 0xdc, 0x86, 0x65, 0x56, 0x81, 0x0e, 0x35, 0x29,
	0xaf, 0x51, 0x50, 0xc5, 0x85, 0x15, 0xd0, 0xe3, 0xf8, 0xa2, 0x27, 0x3d, 0xa7, 0x39, 0xae, 0x24,
	0xdd, 0x81, 0xa0, 0x6a, 0xca, 0x2b, 0x14, 0x86, 0x77, 0x9c, 0x58, 0xa4, 0x7b, 0xa8, 0x69, 0x8e,
	0x86, 0x76, 0x9c, 0x98, 0x38, 0xfe, 0xeb, 0xc0, 0xd8, 0x72, 0x9e, 0xb2, 0x42, 0xa1, 0x78, 0x4e,
	0xea, 0x0d, 0xb6, 0x33, 0xa9, 0x37, 0xd8, 0x92, 0x6d, 0xf0, 0x6f, 0x69, 0xd1, 0xa0, 0xd5, 0xd9,
	0x81, 0xfe, 0x16, 0x6f, 0x7e, 0xcb, 0x40, 0x91, 0xff, 0x58, 0x51, 0xc1, 0xa4, 0x8a, 0x82, 0x3d,
	0x57, 0xbf, 0xaa, 0x8e, 0xe3, 0xef, 0x10, 0x1e, 0x33, 0x59, 0x17, 0xb4, 0x35, 0x2f, 0xb7, 0x0d,
	0x7e, 0x2e, 0x78, 0x53, 0x5b, 0x35, 0x1d, 0xe8, 0xc6, 0x87, 0x5a, 0x99, 0x8c, 0x46, 0xa6, 0xb8,
	0xc7, 0xfa, 0x4c, 0xfe, 0xe0, 0x77, 0xac, 0xca, 0x65, 0xe4, 0x76, 0x67, 0x33, 0x1c, 0xff, 0x76,
	0x20, 0x3c, 0x17, 0x3c, 0x6b, 0xd2, 0x6e, 0x2e, 0x04, 0x3c, 0xd5, 0xd6, 0x68, 0xc8, 0xc7, 0x89,
	0x89, 0x97, 0xcc, 0x64, 0x07, 0x02, 0x81, 0x25, 0x15, 0x37, 0xb6, 0x53, 0x8b, 0xc8, 0x7b, 0x58,
	0xcf, 0x3a, 0xa9, 0x32, 0xf2, 0xf6, 0xdc, 0xfd, 0xf0, 0x70, 0x77, 0xba, 0xb0, 0x90, 0xd3, 0x41,
	0x27, 0x49, 0xff, 0x6d, 0xfc, 0x41, 0xcf, 0xb1, 0x2e, 0xda, 0x4b, 0x45, 0x55, 0x23, 0xcd, 0xbe,
	0xf1, 0xac, 0x17, 0xa1, 0x63, 0xdd, 0x36, 0x0a, 0xc1, 0x85, 0x95, 0xd1, 0x81, 0xf8, 0x02, 0x36,
	0x4c, 0xa1, 0xd1, 0x6e, 0x75, 0x3a, 0x73, 0x9d, 0x87, 0x10, 0x48, 0x43, 0x69, 0xaa, 0x9e, 0xaa,
	0x19, 0x5c, 0x9a, 0xd8, 0x2f, 0xe3, 0x7f, 0x0e, 0x6c, 0xd9, 0x05, 0xb0, 0xce, 0x59, 0x42, 0x3c,
	0xf0, 0xd1, 0x68, 0xb5, 0x8f, 0xdc, 0x15, 0x3e, 0xf2, 0x96, 0xf8, 0xc8, 0x5f, 0xe1, 0xa3, 0x60,
	0xd1, 0x47, 0x0b, 0x76, 0x58, 0x7b, 0x64, 0x87, 0xaf, 0xbd, 0x1d, 0x4e, 0x0b, 0x9a, 0x2f, 0x91,
	0x4d, 0xc0, 0xbb, 0x2e, 0x68, 0x6e, 0x35, 0x9b, 0xf8, 0x59, 0x7f, 0xcd, 0x09, 0xcf, 0x98, 0x54,
	0x4b, 0x08, 0x87, 0xc5, 0xa3, 0xc5, 0xe2, 0x7e, 0x97, 0xdd, 0xc1, 0x2e, 0xa3, 0x9d, 0xd7, 0x0a,
	0xba, 0x59, 0xc9, 0x68, 0x5e, 0x32, 0x98, 0xa1, 0xfb, 0xe2, 0x19, 0xfe, 0x32, 0x33, 0xb4, 0x79,
	0x26, 0x15, 0x4b, 0x07, 0x34, 0xce, 0x4b, 0x69, 0x96, 0x3b, 0x9c, 0xdf, 0x55, 0x38, 0x7b, 0xa9,
	0x0e, 0xe8, 0xac, 0x19, 0xbf, 0xb5, 0x78, 0x07, 0xde, 0xa6, 0xb0, 0x99, 0xa0, 0x6c, 0x0a, 0x65,
	0xb7, 0x3a, 0x84, 0xb5, 0xcb, 0x26, 0x4d, 0x51, 0xca, 0xc9, 0x2b, 0xb2, 0x09, 0xeb, 0x9f, 0xe9,
	0xfd, 0x19, 0x2b, 0x99, 0x9a, 0x38, 0x1a, 0x25, 0x58, 0x23, 0x55, 0x98, 0x4d, 0x46, 0x64, 0x0b,
	0xe0, 0x0b, 0x57, 0x27, 0xf7, 0x4c, 0x6a, 0xec, 0x92, 0xd7, 0x10, 0x1e, 0x1f, 0x9d, 0xdc, 0xa7,
	0x58, 0xeb, 0x0d, 0x98, 0x78, 0x64, 0x03, 0xfc, 0x93, 0xb2, 0x56, 0xed, 0xc4, 0x3f, 0xda, 0xfe,
	0x46, 0x9e, 0xfe, 0x12, 0xae, 0x02, 0x93, 0x7b, 0xf7, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xb3,
	0x2b, 0x88, 0x2f, 0x06, 0x00, 0x00,
}
