// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/device.proto

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

type DeviceInfo struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string    `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string    `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string    `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Remark               string    `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string    `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Sn                   string    `protobuf:"bytes,10,opt,name=sn,proto3" json:"sn,omitempty"`
	Quote                string    `protobuf:"bytes,11,opt,name=quote,proto3" json:"quote,omitempty"`
	Os                   string    `protobuf:"bytes,12,opt,name=os,proto3" json:"os,omitempty"`
	Expiry               uint32    `protobuf:"varint,13,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Certificate          string    `protobuf:"bytes,14,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Type                 uint32    `protobuf:"varint,15,opt,name=type,proto3" json:"type,omitempty"`
	Activated            int64     `protobuf:"varint,16,opt,name=activated,proto3" json:"activated,omitempty"`
	Status               uint32    `protobuf:"varint,17,opt,name=status,proto3" json:"status,omitempty"`
	Meta                 string    `protobuf:"bytes,18,opt,name=meta,proto3" json:"meta,omitempty"`
	Auto                 *PairInfo `protobuf:"bytes,19,opt,name=auto,proto3" json:"auto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{0}
}

func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceInfo.Unmarshal(m, b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceInfo.Size(m)
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeviceInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeviceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *DeviceInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *DeviceInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DeviceInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *DeviceInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *DeviceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DeviceInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *DeviceInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *DeviceInfo) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *DeviceInfo) GetExpiry() uint32 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *DeviceInfo) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *DeviceInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DeviceInfo) GetActivated() int64 {
	if m != nil {
		return m.Activated
	}
	return 0
}

func (m *DeviceInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DeviceInfo) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *DeviceInfo) GetAuto() *PairInfo {
	if m != nil {
		return m.Auto
	}
	return nil
}

type ReqDeviceAdd struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Sn                   string   `protobuf:"bytes,5,opt,name=sn,proto3" json:"sn,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDeviceAdd) Reset()         { *m = ReqDeviceAdd{} }
func (m *ReqDeviceAdd) String() string { return proto.CompactTextString(m) }
func (*ReqDeviceAdd) ProtoMessage()    {}
func (*ReqDeviceAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{1}
}

func (m *ReqDeviceAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDeviceAdd.Unmarshal(m, b)
}
func (m *ReqDeviceAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDeviceAdd.Marshal(b, m, deterministic)
}
func (m *ReqDeviceAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDeviceAdd.Merge(m, src)
}
func (m *ReqDeviceAdd) XXX_Size() int {
	return xxx_messageInfo_ReqDeviceAdd.Size(m)
}
func (m *ReqDeviceAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDeviceAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDeviceAdd proto.InternalMessageInfo

func (m *ReqDeviceAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqDeviceAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqDeviceAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqDeviceAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqDeviceAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqDeviceAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqDeviceBase struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDeviceBase) Reset()         { *m = ReqDeviceBase{} }
func (m *ReqDeviceBase) String() string { return proto.CompactTextString(m) }
func (*ReqDeviceBase) ProtoMessage()    {}
func (*ReqDeviceBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{2}
}

func (m *ReqDeviceBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDeviceBase.Unmarshal(m, b)
}
func (m *ReqDeviceBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDeviceBase.Marshal(b, m, deterministic)
}
func (m *ReqDeviceBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDeviceBase.Merge(m, src)
}
func (m *ReqDeviceBase) XXX_Size() int {
	return xxx_messageInfo_ReqDeviceBase.Size(m)
}
func (m *ReqDeviceBase) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDeviceBase.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDeviceBase proto.InternalMessageInfo

func (m *ReqDeviceBase) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqDeviceBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqDeviceBase) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqDeviceBase) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqDeviceBind struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Quote                string   `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	Os                   string   `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Expiry               uint32   `protobuf:"varint,5,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Activated            uint64   `protobuf:"varint,6,opt,name=activated,proto3" json:"activated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDeviceBind) Reset()         { *m = ReqDeviceBind{} }
func (m *ReqDeviceBind) String() string { return proto.CompactTextString(m) }
func (*ReqDeviceBind) ProtoMessage()    {}
func (*ReqDeviceBind) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{3}
}

func (m *ReqDeviceBind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDeviceBind.Unmarshal(m, b)
}
func (m *ReqDeviceBind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDeviceBind.Marshal(b, m, deterministic)
}
func (m *ReqDeviceBind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDeviceBind.Merge(m, src)
}
func (m *ReqDeviceBind) XXX_Size() int {
	return xxx_messageInfo_ReqDeviceBind.Size(m)
}
func (m *ReqDeviceBind) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDeviceBind.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDeviceBind proto.InternalMessageInfo

func (m *ReqDeviceBind) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqDeviceBind) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqDeviceBind) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqDeviceBind) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *ReqDeviceBind) GetExpiry() uint32 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *ReqDeviceBind) GetActivated() uint64 {
	if m != nil {
		return m.Activated
	}
	return 0
}

type ReplyDeviceInfo struct {
	Info                 *DeviceInfo  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyDeviceInfo) Reset()         { *m = ReplyDeviceInfo{} }
func (m *ReplyDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyDeviceInfo) ProtoMessage()    {}
func (*ReplyDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{4}
}

func (m *ReplyDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyDeviceInfo.Unmarshal(m, b)
}
func (m *ReplyDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyDeviceInfo.Marshal(b, m, deterministic)
}
func (m *ReplyDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyDeviceInfo.Merge(m, src)
}
func (m *ReplyDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyDeviceInfo.Size(m)
}
func (m *ReplyDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyDeviceInfo proto.InternalMessageInfo

func (m *ReplyDeviceInfo) GetInfo() *DeviceInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyDeviceInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyDeviceList struct {
	Total                uint32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32        `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pages                uint32        `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Number               uint32        `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*DeviceInfo `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyDeviceList) Reset()         { *m = ReplyDeviceList{} }
func (m *ReplyDeviceList) String() string { return proto.CompactTextString(m) }
func (*ReplyDeviceList) ProtoMessage()    {}
func (*ReplyDeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{5}
}

func (m *ReplyDeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyDeviceList.Unmarshal(m, b)
}
func (m *ReplyDeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyDeviceList.Marshal(b, m, deterministic)
}
func (m *ReplyDeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyDeviceList.Merge(m, src)
}
func (m *ReplyDeviceList) XXX_Size() int {
	return xxx_messageInfo_ReplyDeviceList.Size(m)
}
func (m *ReplyDeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyDeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyDeviceList proto.InternalMessageInfo

func (m *ReplyDeviceList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyDeviceList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyDeviceList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyDeviceList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyDeviceList) GetList() []*DeviceInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyDeviceList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceInfo)(nil), "omo.msp.scene.DeviceInfo")
	proto.RegisterType((*ReqDeviceAdd)(nil), "omo.msp.scene.ReqDeviceAdd")
	proto.RegisterType((*ReqDeviceBase)(nil), "omo.msp.scene.ReqDeviceBase")
	proto.RegisterType((*ReqDeviceBind)(nil), "omo.msp.scene.ReqDeviceBind")
	proto.RegisterType((*ReplyDeviceInfo)(nil), "omo.msp.scene.ReplyDeviceInfo")
	proto.RegisterType((*ReplyDeviceList)(nil), "omo.msp.scene.ReplyDeviceList")
}

func init() {
	proto.RegisterFile("proto/organization/device.proto", fileDescriptor_68b6e6e2e28953de)
}

var fileDescriptor_68b6e6e2e28953de = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0xae, 0x13, 0xc7, 0x6d, 0x36, 0x75, 0xdb, 0xb3, 0xa7, 0x3a, 0x67, 0x09, 0x05, 0x22, 0x5f,
	0x45, 0x42, 0xa4, 0x52, 0x78, 0x01, 0x5a, 0x50, 0x23, 0xa4, 0x22, 0xc0, 0x15, 0x37, 0xdc, 0x6d,
	0xed, 0x69, 0xb5, 0x22, 0xde, 0x75, 0xbc, 0x9b, 0x42, 0x78, 0x0a, 0x6e, 0x79, 0x27, 0x6e, 0xb9,
	0xe6, 0x55, 0xd0, 0xce, 0xe6, 0xcf, 0xb5, 0x9b, 0xf6, 0xca, 0x3b, 0x3b, 0xdf, 0x7c, 0x33, 0xf3,
	0xcd, 0x6c, 0x42, 0x9e, 0xe5, 0x85, 0x32, 0xea, 0x58, 0x15, 0xd7, 0x5c, 0x8a, 0xef, 0xdc, 0x08,
	0x25, 0x8f, 0x53, 0xb8, 0x11, 0x09, 0x0c, 0xd0, 0x43, 0x43, 0x95, 0xa9, 0x41, 0xa6, 0xf3, 0x81,
	0x4e, 0x40, 0x42, 0xb7, 0x0e, 0x9f, 0xa8, 0x2c, 0x53, 0xd2, 0xe1, 0xa3, 0xdf, 0x4d, 0x42, 0xde,
	0x20, 0xc1, 0x5b, 0x79, 0xa5, 0xe8, 0x1e, 0x69, 0x88, 0x94, 0x79, 0x3d, 0xaf, 0xef, 0xc7, 0x0d,
	0x91, 0xd2, 0x03, 0xd2, 0x9c, 0x8a, 0x94, 0x35, 0x7a, 0x5e, 0xbf, 0x1d, 0xdb, 0x23, 0xa5, 0xc4,
	0x97, 0x3c, 0x03, 0xd6, 0xc4, 0x2b, 0x3c, 0x53, 0x46, 0xb6, 0x93, 0x02, 0xb8, 0x81, 0x94, 0xf9,
	0x18, 0xba, 0x30, 0xad, 0x67, 0x9a, 0xa7, 0xe8, 0x69, 0x39, 0xcf, 0xdc, 0x5c, 0xc6, 0xa8, 0x82,
	0x05, 0x48, 0xb5, 0x30, 0x69, 0x97, 0xec, 0xa8, 0x1c, 0x0a, 0x74, 0x6d, 0xa3, 0x6b, 0x69, 0xd3,
	0xff, 0x48, 0x50, 0x40, 0xc6, 0x8b, 0x2f, 0x6c, 0x07, 0x3d, 0x73, 0x8b, 0x1e, 0x92, 0x96, 0xfa,
	0x2a, 0xa1, 0x60, 0x6d, 0xbc, 0x76, 0x86, 0xed, 0x46, 0x4b, 0x46, 0xf0, 0xaa, 0xa1, 0xa5, 0x45,
	0x4d, 0xa6, 0xca, 0x00, 0xeb, 0x38, 0x14, 0x1a, 0x16, 0xa5, 0x34, 0xdb, 0x75, 0x28, 0xa5, 0x6d,
	0x0e, 0xf8, 0x96, 0x8b, 0x62, 0xc6, 0xc2, 0x9e, 0xd7, 0x0f, 0xe3, 0xb9, 0x45, 0x7b, 0xa4, 0x93,
	0x40, 0x61, 0xc4, 0x95, 0x48, 0xb8, 0x01, 0xb6, 0x87, 0x01, 0xeb, 0x57, 0x56, 0x1b, 0x33, 0xcb,
	0x81, 0xed, 0x63, 0x1c, 0x9e, 0xe9, 0x11, 0x69, 0xf3, 0xc4, 0x88, 0x1b, 0xd4, 0xe0, 0xa0, 0xe7,
	0xf5, 0x9b, 0xf1, 0xea, 0xc2, 0xe6, 0xd2, 0x86, 0x9b, 0xa9, 0x66, 0xff, 0xb8, 0x5c, 0xce, 0xb2,
	0x4c, 0x19, 0x18, 0xce, 0xa8, 0x53, 0xd9, 0x9e, 0xe9, 0x73, 0xe2, 0xf3, 0xa9, 0x51, 0xec, 0xdf,
	0x9e, 0xd7, 0xef, 0x0c, 0xff, 0x1f, 0x94, 0x26, 0x3d, 0xf8, 0xc0, 0x45, 0x61, 0x47, 0x18, 0x23,
	0x28, 0xfa, 0xe1, 0x91, 0xdd, 0x18, 0x26, 0x6e, 0xb4, 0x27, 0x69, 0x5a, 0x52, 0xd5, 0xbb, 0xa5,
	0xea, 0x52, 0xbd, 0xc6, 0xba, 0x7a, 0x75, 0x93, 0x5e, 0xe9, 0xef, 0x97, 0xf4, 0x77, 0x4a, 0xb7,
	0x96, 0x4a, 0x2f, 0x94, 0x08, 0x56, 0x4a, 0x44, 0x82, 0x84, 0xcb, 0x8a, 0x4e, 0xb9, 0x86, 0xc5,
	0x72, 0x79, 0xd5, 0xe5, 0x6a, 0xd4, 0xa6, 0x6c, 0x96, 0x52, 0xae, 0x37, 0xe4, 0x97, 0x1b, 0x8a,
	0x7e, 0x7a, 0xeb, 0xb9, 0x84, 0xdc, 0xdc, 0x7e, 0x75, 0xc9, 0x97, 0x8b, 0xd2, 0xac, 0x2e, 0x8a,
	0x5f, 0xb3, 0x28, 0xad, 0xd2, 0xa2, 0x94, 0x46, 0x1e, 0xe0, 0xda, 0xaf, 0x2e, 0x22, 0x43, 0xf6,
	0x63, 0xc8, 0xc7, 0xb3, 0xb5, 0x57, 0xf7, 0x82, 0xf8, 0x42, 0x5e, 0x29, 0x2c, 0xac, 0x33, 0x7c,
	0x74, 0x6b, 0xb2, 0x2b, 0x60, 0x8c, 0x30, 0x3a, 0x5c, 0x2e, 0x4d, 0x03, 0x03, 0xba, 0xb7, 0x02,
	0x90, 0xfe, 0x02, 0x11, 0x8b, 0x85, 0x8a, 0x7e, 0x79, 0xa5, 0xb4, 0xe7, 0x42, 0x1b, 0xdb, 0xa5,
	0x51, 0x86, 0x8f, 0x31, 0x6f, 0x18, 0x3b, 0xc3, 0xce, 0x20, 0xe7, 0xd7, 0x6e, 0x06, 0x61, 0x8c,
	0x67, 0x8b, 0xb4, 0x5f, 0x8d, 0x7a, 0x84, 0xb1, 0x33, 0x6c, 0xff, 0x72, 0x9a, 0x5d, 0x82, 0xd3,
	0x3f, 0x8c, 0xe7, 0x96, 0x6d, 0x67, 0x2c, 0xb4, 0x61, 0xad, 0x5e, 0xf3, 0x9e, 0x76, 0x2c, 0x6c,
	0xad, 0x9d, 0xe0, 0xa1, 0xed, 0x0c, 0xff, 0xf8, 0x24, 0x74, 0x44, 0x17, 0x50, 0xd8, 0x0f, 0x1d,
	0x91, 0xe0, 0x24, 0x4d, 0xdf, 0x4b, 0xa0, 0x8f, 0x2b, 0xf1, 0xab, 0x67, 0xd0, 0x7d, 0x5a, 0x47,
	0xbe, 0x2a, 0x29, 0xda, 0xa2, 0x67, 0x24, 0x18, 0x81, 0xb1, 0x44, 0xd5, 0x42, 0x26, 0x53, 0xd0,
	0xc6, 0xe2, 0x1e, 0xc0, 0xf3, 0x9a, 0xb4, 0x63, 0xc8, 0xd4, 0x0d, 0xdc, 0x47, 0xc5, 0xea, 0xa8,
	0xe6, 0x24, 0x1f, 0xc9, 0xfe, 0x08, 0x8c, 0x9d, 0xd6, 0xe9, 0xec, 0x4c, 0x8c, 0x0d, 0x14, 0xf4,
	0xa8, 0x9e, 0xca, 0x79, 0x37, 0xd5, 0x65, 0x59, 0xa2, 0x2d, 0xfa, 0x8e, 0xec, 0x8e, 0xc0, 0x58,
	0x3d, 0x85, 0x36, 0x22, 0xb9, 0x87, 0xef, 0xc9, 0x5d, 0xc3, 0xc0, 0x60, 0x94, 0x8b, 0x7c, 0xc2,
	0x9f, 0x74, 0x7c, 0xd2, 0x47, 0x77, 0x69, 0x6f, 0xbd, 0x1b, 0x3b, 0x3d, 0x27, 0x7b, 0x73, 0x9e,
	0x45, 0xa3, 0xd5, 0x56, 0x26, 0x0e, 0x31, 0x2f, 0x6d, 0x13, 0xdb, 0x2b, 0xe2, 0xe3, 0xb3, 0xbf,
	0xbb, 0x1e, 0x21, 0xd3, 0x4d, 0x0c, 0xa7, 0x87, 0x9f, 0x69, 0xf5, 0xbf, 0xf3, 0x32, 0xc0, 0xbb,
	0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x47, 0xdf, 0xa9, 0x88, 0x07, 0x00, 0x00,
}
