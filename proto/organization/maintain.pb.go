// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/maintain.proto

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

type MaintainInfo struct {
	Uid                  string             `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32             `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64              `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string             `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string             `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string             `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Device               string             `protobuf:"bytes,8,opt,name=device,proto3" json:"device,omitempty"`
	Scene                string             `protobuf:"bytes,9,opt,name=scene,proto3" json:"scene,omitempty"`
	Date                 string             `protobuf:"bytes,10,opt,name=date,proto3" json:"date,omitempty"`
	Area                 string             `protobuf:"bytes,11,opt,name=area,proto3" json:"area,omitempty"`
	Submitter            string             `protobuf:"bytes,12,opt,name=submitter,proto3" json:"submitter,omitempty"`
	Contacts             string             `protobuf:"bytes,13,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Contents             []*MaintainContent `protobuf:"bytes,21,rep,name=contents,proto3" json:"contents,omitempty"`
	Maintainers          []string           `protobuf:"bytes,22,rep,name=maintainers,proto3" json:"maintainers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MaintainInfo) Reset()         { *m = MaintainInfo{} }
func (m *MaintainInfo) String() string { return proto.CompactTextString(m) }
func (*MaintainInfo) ProtoMessage()    {}
func (*MaintainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b76e010cb7bf06, []int{0}
}

func (m *MaintainInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintainInfo.Unmarshal(m, b)
}
func (m *MaintainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintainInfo.Marshal(b, m, deterministic)
}
func (m *MaintainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintainInfo.Merge(m, src)
}
func (m *MaintainInfo) XXX_Size() int {
	return xxx_messageInfo_MaintainInfo.Size(m)
}
func (m *MaintainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MaintainInfo proto.InternalMessageInfo

func (m *MaintainInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MaintainInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MaintainInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MaintainInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MaintainInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MaintainInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MaintainInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *MaintainInfo) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *MaintainInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *MaintainInfo) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *MaintainInfo) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *MaintainInfo) GetSubmitter() string {
	if m != nil {
		return m.Submitter
	}
	return ""
}

func (m *MaintainInfo) GetContacts() string {
	if m != nil {
		return m.Contacts
	}
	return ""
}

func (m *MaintainInfo) GetContents() []*MaintainContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *MaintainInfo) GetMaintainers() []string {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

type MaintainContent struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Assets               []string `protobuf:"bytes,10,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaintainContent) Reset()         { *m = MaintainContent{} }
func (m *MaintainContent) String() string { return proto.CompactTextString(m) }
func (*MaintainContent) ProtoMessage()    {}
func (*MaintainContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b76e010cb7bf06, []int{1}
}

func (m *MaintainContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintainContent.Unmarshal(m, b)
}
func (m *MaintainContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintainContent.Marshal(b, m, deterministic)
}
func (m *MaintainContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintainContent.Merge(m, src)
}
func (m *MaintainContent) XXX_Size() int {
	return xxx_messageInfo_MaintainContent.Size(m)
}
func (m *MaintainContent) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintainContent.DiscardUnknown(m)
}

var xxx_messageInfo_MaintainContent proto.InternalMessageInfo

func (m *MaintainContent) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MaintainContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MaintainContent) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqMaintainAdd struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string             `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Type                 uint32             `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Scene                string             `protobuf:"bytes,4,opt,name=scene,proto3" json:"scene,omitempty"`
	Operator             string             `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Date                 string             `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Area                 string             `protobuf:"bytes,8,opt,name=area,proto3" json:"area,omitempty"`
	Submitter            string             `protobuf:"bytes,9,opt,name=submitter,proto3" json:"submitter,omitempty"`
	Contacts             string             `protobuf:"bytes,10,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Contents             []*MaintainContent `protobuf:"bytes,20,rep,name=contents,proto3" json:"contents,omitempty"`
	Maintainers          []string           `protobuf:"bytes,21,rep,name=maintainers,proto3" json:"maintainers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReqMaintainAdd) Reset()         { *m = ReqMaintainAdd{} }
func (m *ReqMaintainAdd) String() string { return proto.CompactTextString(m) }
func (*ReqMaintainAdd) ProtoMessage()    {}
func (*ReqMaintainAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b76e010cb7bf06, []int{2}
}

func (m *ReqMaintainAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMaintainAdd.Unmarshal(m, b)
}
func (m *ReqMaintainAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMaintainAdd.Marshal(b, m, deterministic)
}
func (m *ReqMaintainAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMaintainAdd.Merge(m, src)
}
func (m *ReqMaintainAdd) XXX_Size() int {
	return xxx_messageInfo_ReqMaintainAdd.Size(m)
}
func (m *ReqMaintainAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMaintainAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMaintainAdd proto.InternalMessageInfo

func (m *ReqMaintainAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqMaintainAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqMaintainAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqMaintainAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqMaintainAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqMaintainAdd) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ReqMaintainAdd) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *ReqMaintainAdd) GetSubmitter() string {
	if m != nil {
		return m.Submitter
	}
	return ""
}

func (m *ReqMaintainAdd) GetContacts() string {
	if m != nil {
		return m.Contacts
	}
	return ""
}

func (m *ReqMaintainAdd) GetContents() []*MaintainContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *ReqMaintainAdd) GetMaintainers() []string {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

type ReplyMaintainInfo struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *MaintainInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyMaintainInfo) Reset()         { *m = ReplyMaintainInfo{} }
func (m *ReplyMaintainInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyMaintainInfo) ProtoMessage()    {}
func (*ReplyMaintainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b76e010cb7bf06, []int{3}
}

func (m *ReplyMaintainInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMaintainInfo.Unmarshal(m, b)
}
func (m *ReplyMaintainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMaintainInfo.Marshal(b, m, deterministic)
}
func (m *ReplyMaintainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMaintainInfo.Merge(m, src)
}
func (m *ReplyMaintainInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyMaintainInfo.Size(m)
}
func (m *ReplyMaintainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMaintainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMaintainInfo proto.InternalMessageInfo

func (m *ReplyMaintainInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMaintainInfo) GetInfo() *MaintainInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyMaintainList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32          `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32          `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32          `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*MaintainInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyMaintainList) Reset()         { *m = ReplyMaintainList{} }
func (m *ReplyMaintainList) String() string { return proto.CompactTextString(m) }
func (*ReplyMaintainList) ProtoMessage()    {}
func (*ReplyMaintainList) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b76e010cb7bf06, []int{4}
}

func (m *ReplyMaintainList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMaintainList.Unmarshal(m, b)
}
func (m *ReplyMaintainList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMaintainList.Marshal(b, m, deterministic)
}
func (m *ReplyMaintainList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMaintainList.Merge(m, src)
}
func (m *ReplyMaintainList) XXX_Size() int {
	return xxx_messageInfo_ReplyMaintainList.Size(m)
}
func (m *ReplyMaintainList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMaintainList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMaintainList proto.InternalMessageInfo

func (m *ReplyMaintainList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMaintainList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMaintainList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyMaintainList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyMaintainList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyMaintainList) GetList() []*MaintainInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MaintainInfo)(nil), "omo.msp.scene.MaintainInfo")
	proto.RegisterType((*MaintainContent)(nil), "omo.msp.scene.MaintainContent")
	proto.RegisterType((*ReqMaintainAdd)(nil), "omo.msp.scene.ReqMaintainAdd")
	proto.RegisterType((*ReplyMaintainInfo)(nil), "omo.msp.scene.ReplyMaintainInfo")
	proto.RegisterType((*ReplyMaintainList)(nil), "omo.msp.scene.ReplyMaintainList")
}

func init() {
	proto.RegisterFile("proto/organization/maintain.proto", fileDescriptor_22b76e010cb7bf06)
}

var fileDescriptor_22b76e010cb7bf06 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xc1, 0x4e, 0xdb, 0x4c,
	0x10, 0xc7, 0x71, 0x1c, 0x4c, 0x32, 0x21, 0x7c, 0x7c, 0xab, 0x80, 0x56, 0x29, 0x50, 0xd7, 0x27,
	0x9f, 0x82, 0x94, 0xde, 0x7a, 0x03, 0xa4, 0xd2, 0xaa, 0x20, 0xa4, 0x45, 0x55, 0xa5, 0xde, 0x16,
	0x7b, 0x40, 0xab, 0xc6, 0x5e, 0xe3, 0xdd, 0xa0, 0xd2, 0xc7, 0xe8, 0x23, 0xf5, 0x09, 0x7a, 0xef,
	0xcb, 0x54, 0xbb, 0xeb, 0x98, 0x90, 0x98, 0x50, 0xb5, 0xb7, 0x99, 0xff, 0xcc, 0xfe, 0x37, 0x33,
	0xbf, 0xb5, 0x02, 0xaf, 0x8a, 0x52, 0x6a, 0x79, 0x28, 0xcb, 0x1b, 0x9e, 0x8b, 0x6f, 0x5c, 0x0b,
	0x99, 0x1f, 0x66, 0x5c, 0xe4, 0x9a, 0x8b, 0x7c, 0x64, 0x6b, 0xa4, 0x2f, 0x33, 0x39, 0xca, 0x54,
	0x31, 0x52, 0x09, 0xe6, 0x38, 0x7c, 0xd9, 0x70, 0x22, 0x91, 0x59, 0x26, 0xab, 0xfe, 0xe8, 0xbb,
	0x0f, 0x9b, 0xe7, 0x95, 0xc5, 0xfb, 0xfc, 0x5a, 0x92, 0x6d, 0xf0, 0xa7, 0x22, 0xa5, 0x5e, 0xe8,
	0xc5, 0x5d, 0x66, 0x42, 0xb2, 0x05, 0x2d, 0x91, 0xd2, 0x56, 0xe8, 0xc5, 0x6d, 0xd6, 0x12, 0x29,
	0x21, 0xd0, 0xd6, 0xf7, 0x05, 0x52, 0x3f, 0xf4, 0xe2, 0x3e, 0xb3, 0x31, 0xa1, 0xb0, 0x91, 0x94,
	0xc8, 0x35, 0xa6, 0xb4, 0x1d, 0x7a, 0xb1, 0xcf, 0x66, 0x69, 0x5d, 0x91, 0x25, 0x5d, 0xb7, 0x9e,
	0xb3, 0xd4, 0xf8, 0xe4, 0x3c, 0x43, 0x1a, 0x58, 0xd9, 0xc6, 0x64, 0x17, 0x82, 0x12, 0x33, 0x5e,
	0x7e, 0xa1, 0x1b, 0x56, 0xad, 0x32, 0xa3, 0xa7, 0x78, 0x27, 0x12, 0xa4, 0x1d, 0xa7, 0xbb, 0x8c,
	0x0c, 0x60, 0xdd, 0x0e, 0x4a, 0xbb, 0x56, 0x76, 0x89, 0x71, 0x4e, 0xb9, 0x46, 0x0a, 0xce, 0xd9,
	0xc4, 0x46, 0xe3, 0x25, 0x72, 0xda, 0x73, 0x9a, 0x89, 0xc9, 0x1e, 0x74, 0xd5, 0xf4, 0x2a, 0x13,
	0x5a, 0x63, 0x49, 0x37, 0x6d, 0xe1, 0x41, 0x20, 0x43, 0xe8, 0x24, 0x32, 0xd7, 0x3c, 0xd1, 0x8a,
	0xf6, 0x6d, 0xb1, 0xce, 0xc9, 0x1b, 0x57, 0xc3, 0x5c, 0x2b, 0xba, 0x13, 0xfa, 0x71, 0x6f, 0x7c,
	0x30, 0x7a, 0xb4, 0xf9, 0xd1, 0x6c, 0xa9, 0x27, 0xae, 0x8d, 0xd5, 0xfd, 0x24, 0x84, 0xde, 0x0c,
	0x1a, 0x96, 0x8a, 0xee, 0x86, 0x7e, 0xdc, 0x65, 0xf3, 0x52, 0xf4, 0x09, 0xfe, 0x5b, 0x38, 0x5e,
	0x2f, 0xdd, 0x5b, 0x58, 0xba, 0x2b, 0x5b, 0x3a, 0x66, 0xb5, 0x55, 0xf7, 0x2e, 0x04, 0x5c, 0x29,
	0xd4, 0x8a, 0x82, 0x75, 0xaf, 0xb2, 0xe8, 0x47, 0x0b, 0xb6, 0x18, 0xde, 0xce, 0xcc, 0x8f, 0xd2,
	0xb4, 0xa6, 0xe0, 0x35, 0x52, 0x68, 0x3d, 0xa2, 0xd0, 0x44, 0xbe, 0x26, 0xd0, 0x9e, 0x27, 0x30,
	0x84, 0x8e, 0x2c, 0xb0, 0x9c, 0xc3, 0x5e, 0xe7, 0x35, 0x9d, 0x8d, 0x06, 0x3a, 0x9d, 0xa7, 0xe8,
	0x74, 0x57, 0xd1, 0x81, 0x15, 0x74, 0x06, 0xff, 0x46, 0x67, 0x67, 0x99, 0xce, 0x57, 0xf8, 0x9f,
	0x61, 0x31, 0xb9, 0x7f, 0xf4, 0xd9, 0x8c, 0x21, 0x50, 0x9a, 0xeb, 0xa9, 0xb2, 0x8b, 0xec, 0x8d,
	0x87, 0x0b, 0x17, 0xda, 0x13, 0x97, 0xb6, 0x83, 0x55, 0x9d, 0xe4, 0x10, 0xda, 0x22, 0xbf, 0x96,
	0x76, 0xc9, 0xbd, 0xf1, 0x8b, 0x27, 0x7e, 0xa2, 0xb1, 0x67, 0xb6, 0x31, 0xfa, 0xe9, 0x2d, 0x5c,
	0x7d, 0x26, 0x94, 0xfe, 0xab, 0xab, 0x07, 0xb0, 0xae, 0xa5, 0xe6, 0x13, 0x7b, 0x77, 0x9f, 0xb9,
	0xc4, 0xa8, 0x05, 0xbf, 0x41, 0x55, 0x01, 0x76, 0x89, 0x61, 0x63, 0x02, 0x0b, 0xb8, 0xcf, 0x6c,
	0x6c, 0x5e, 0x48, 0x3e, 0xcd, 0xae, 0xd0, 0xd1, 0xed, 0xb3, 0x2a, 0x33, 0x23, 0x4d, 0x84, 0xd2,
	0x74, 0xdb, 0x6e, 0x7d, 0xf5, 0x48, 0xa6, 0x71, 0xfc, 0xcb, 0x7f, 0x78, 0xeb, 0x97, 0x58, 0xda,
	0x8f, 0xfa, 0x03, 0x04, 0x47, 0x69, 0x7a, 0x91, 0x23, 0xd9, 0x5f, 0x1a, 0x65, 0xfe, 0xed, 0x0e,
	0xc3, 0xa6, 0x49, 0xe7, 0x2f, 0x89, 0xd6, 0xc8, 0x3b, 0x08, 0x4e, 0x51, 0x1b, 0xb3, 0xe5, 0xbd,
	0xdc, 0x4e, 0x51, 0x69, 0xd3, 0xf7, 0x47, 0x4e, 0x17, 0xd0, 0x3b, 0x45, 0x7d, 0x7c, 0xff, 0x56,
	0x4c, 0xcc, 0x03, 0xdc, 0x6b, 0xb6, 0x73, 0xd5, 0xd5, 0x86, 0x06, 0x5b, 0xb4, 0x46, 0xce, 0x61,
	0xf3, 0x14, 0xb5, 0x21, 0x23, 0x94, 0x16, 0xc9, 0x33, 0x8e, 0xfb, 0x4f, 0x61, 0xb5, 0x87, 0xa3,
	0x35, 0x72, 0x06, 0x5b, 0x1f, 0x0b, 0xf3, 0x35, 0xd5, 0x3f, 0xf1, 0x60, 0xd9, 0xd0, 0x75, 0x54,
	0x96, 0xb4, 0xc9, 0xb2, 0x9a, 0xf6, 0x04, 0xba, 0x0c, 0x33, 0x79, 0x87, 0xcf, 0xad, 0x6e, 0x85,
	0xc9, 0xf1, 0xe0, 0x33, 0x59, 0xfe, 0x03, 0xba, 0x0a, 0xac, 0xf6, 0xfa, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x20, 0xc0, 0x05, 0x2f, 0xcf, 0x06, 0x00, 0x00,
}
