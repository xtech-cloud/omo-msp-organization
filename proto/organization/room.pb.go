// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/room.proto

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

type RoomInfo struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              int64          `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64          `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string         `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string         `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string         `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Quotes               []string       `protobuf:"bytes,10,rep,name=quotes,proto3" json:"quotes,omitempty"`
	Devices              []*ProductInfo `protobuf:"bytes,11,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{0}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RoomInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoomInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoomInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RoomInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *RoomInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RoomInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RoomInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RoomInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RoomInfo) GetQuotes() []string {
	if m != nil {
		return m.Quotes
	}
	return nil
}

func (m *RoomInfo) GetDevices() []*ProductInfo {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ReqRoomAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoomAdd) Reset()         { *m = ReqRoomAdd{} }
func (m *ReqRoomAdd) String() string { return proto.CompactTextString(m) }
func (*ReqRoomAdd) ProtoMessage()    {}
func (*ReqRoomAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{1}
}

func (m *ReqRoomAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoomAdd.Unmarshal(m, b)
}
func (m *ReqRoomAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoomAdd.Marshal(b, m, deterministic)
}
func (m *ReqRoomAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoomAdd.Merge(m, src)
}
func (m *ReqRoomAdd) XXX_Size() int {
	return xxx_messageInfo_ReqRoomAdd.Size(m)
}
func (m *ReqRoomAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoomAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoomAdd proto.InternalMessageInfo

func (m *ReqRoomAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqRoomAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRoomAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRoomAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqRoomUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoomUpdate) Reset()         { *m = ReqRoomUpdate{} }
func (m *ReqRoomUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqRoomUpdate) ProtoMessage()    {}
func (*ReqRoomUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{2}
}

func (m *ReqRoomUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoomUpdate.Unmarshal(m, b)
}
func (m *ReqRoomUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoomUpdate.Marshal(b, m, deterministic)
}
func (m *ReqRoomUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoomUpdate.Merge(m, src)
}
func (m *ReqRoomUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqRoomUpdate.Size(m)
}
func (m *ReqRoomUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoomUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoomUpdate proto.InternalMessageInfo

func (m *ReqRoomUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqRoomUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRoomUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRoomUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqRoomQuotes struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Room                 string   `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoomQuotes) Reset()         { *m = ReqRoomQuotes{} }
func (m *ReqRoomQuotes) String() string { return proto.CompactTextString(m) }
func (*ReqRoomQuotes) ProtoMessage()    {}
func (*ReqRoomQuotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{3}
}

func (m *ReqRoomQuotes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoomQuotes.Unmarshal(m, b)
}
func (m *ReqRoomQuotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoomQuotes.Marshal(b, m, deterministic)
}
func (m *ReqRoomQuotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoomQuotes.Merge(m, src)
}
func (m *ReqRoomQuotes) XXX_Size() int {
	return xxx_messageInfo_ReqRoomQuotes.Size(m)
}
func (m *ReqRoomQuotes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoomQuotes.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoomQuotes proto.InternalMessageInfo

func (m *ReqRoomQuotes) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqRoomQuotes) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *ReqRoomQuotes) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoomQuotes) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqRoomDisplays struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Room                 string   `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Sn                   string   `protobuf:"bytes,4,opt,name=sn,proto3" json:"sn,omitempty"`
	List                 []string `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoomDisplays) Reset()         { *m = ReqRoomDisplays{} }
func (m *ReqRoomDisplays) String() string { return proto.CompactTextString(m) }
func (*ReqRoomDisplays) ProtoMessage()    {}
func (*ReqRoomDisplays) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{4}
}

func (m *ReqRoomDisplays) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoomDisplays.Unmarshal(m, b)
}
func (m *ReqRoomDisplays) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoomDisplays.Marshal(b, m, deterministic)
}
func (m *ReqRoomDisplays) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoomDisplays.Merge(m, src)
}
func (m *ReqRoomDisplays) XXX_Size() int {
	return xxx_messageInfo_ReqRoomDisplays.Size(m)
}
func (m *ReqRoomDisplays) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoomDisplays.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoomDisplays proto.InternalMessageInfo

func (m *ReqRoomDisplays) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqRoomDisplays) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *ReqRoomDisplays) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoomDisplays) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqRoomDisplays) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqRoomDevice struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Device               string   `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoomDevice) Reset()         { *m = ReqRoomDevice{} }
func (m *ReqRoomDevice) String() string { return proto.CompactTextString(m) }
func (*ReqRoomDevice) ProtoMessage()    {}
func (*ReqRoomDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{5}
}

func (m *ReqRoomDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoomDevice.Unmarshal(m, b)
}
func (m *ReqRoomDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoomDevice.Marshal(b, m, deterministic)
}
func (m *ReqRoomDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoomDevice.Merge(m, src)
}
func (m *ReqRoomDevice) XXX_Size() int {
	return xxx_messageInfo_ReqRoomDevice.Size(m)
}
func (m *ReqRoomDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoomDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoomDevice proto.InternalMessageInfo

func (m *ReqRoomDevice) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqRoomDevice) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRoomDevice) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ReqRoomDevice) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqRoomDevice) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ReplyRoomDevices struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*ProductInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyRoomDevices) Reset()         { *m = ReplyRoomDevices{} }
func (m *ReplyRoomDevices) String() string { return proto.CompactTextString(m) }
func (*ReplyRoomDevices) ProtoMessage()    {}
func (*ReplyRoomDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{6}
}

func (m *ReplyRoomDevices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoomDevices.Unmarshal(m, b)
}
func (m *ReplyRoomDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoomDevices.Marshal(b, m, deterministic)
}
func (m *ReplyRoomDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoomDevices.Merge(m, src)
}
func (m *ReplyRoomDevices) XXX_Size() int {
	return xxx_messageInfo_ReplyRoomDevices.Size(m)
}
func (m *ReplyRoomDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoomDevices.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoomDevices proto.InternalMessageInfo

func (m *ReplyRoomDevices) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRoomDevices) GetList() []*ProductInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyRoomInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *RoomInfo    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyRoomInfo) Reset()         { *m = ReplyRoomInfo{} }
func (m *ReplyRoomInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyRoomInfo) ProtoMessage()    {}
func (*ReplyRoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{7}
}

func (m *ReplyRoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoomInfo.Unmarshal(m, b)
}
func (m *ReplyRoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoomInfo.Marshal(b, m, deterministic)
}
func (m *ReplyRoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoomInfo.Merge(m, src)
}
func (m *ReplyRoomInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyRoomInfo.Size(m)
}
func (m *ReplyRoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoomInfo proto.InternalMessageInfo

func (m *ReplyRoomInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRoomInfo) GetInfo() *RoomInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyRoomList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*RoomInfo  `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyRoomList) Reset()         { *m = ReplyRoomList{} }
func (m *ReplyRoomList) String() string { return proto.CompactTextString(m) }
func (*ReplyRoomList) ProtoMessage()    {}
func (*ReplyRoomList) Descriptor() ([]byte, []int) {
	return fileDescriptor_002834448c4f3442, []int{8}
}

func (m *ReplyRoomList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRoomList.Unmarshal(m, b)
}
func (m *ReplyRoomList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRoomList.Marshal(b, m, deterministic)
}
func (m *ReplyRoomList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRoomList.Merge(m, src)
}
func (m *ReplyRoomList) XXX_Size() int {
	return xxx_messageInfo_ReplyRoomList.Size(m)
}
func (m *ReplyRoomList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRoomList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRoomList proto.InternalMessageInfo

func (m *ReplyRoomList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRoomList) GetList() []*RoomInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RoomInfo)(nil), "omo.msp.scene.RoomInfo")
	proto.RegisterType((*ReqRoomAdd)(nil), "omo.msp.scene.ReqRoomAdd")
	proto.RegisterType((*ReqRoomUpdate)(nil), "omo.msp.scene.ReqRoomUpdate")
	proto.RegisterType((*ReqRoomQuotes)(nil), "omo.msp.scene.ReqRoomQuotes")
	proto.RegisterType((*ReqRoomDisplays)(nil), "omo.msp.scene.ReqRoomDisplays")
	proto.RegisterType((*ReqRoomDevice)(nil), "omo.msp.scene.ReqRoomDevice")
	proto.RegisterType((*ReplyRoomDevices)(nil), "omo.msp.scene.ReplyRoomDevices")
	proto.RegisterType((*ReplyRoomInfo)(nil), "omo.msp.scene.ReplyRoomInfo")
	proto.RegisterType((*ReplyRoomList)(nil), "omo.msp.scene.ReplyRoomList")
}

func init() {
	proto.RegisterFile("proto/organization/room.proto", fileDescriptor_002834448c4f3442)
}

var fileDescriptor_002834448c4f3442 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x4f, 0xdb, 0x3e,
	0x14, 0x6d, 0x93, 0x34, 0x85, 0x5b, 0xe8, 0x0f, 0x59, 0x88, 0x5f, 0x56, 0xc1, 0xa8, 0xf2, 0x54,
	0x69, 0x52, 0x91, 0xba, 0x7d, 0x01, 0xfe, 0x88, 0x6e, 0x12, 0x13, 0x23, 0x68, 0x2f, 0x7b, 0x0b,
	0x89, 0x99, 0xbc, 0x35, 0x71, 0x88, 0x1d, 0x26, 0xb6, 0xa7, 0x7d, 0x99, 0x7d, 0xc7, 0xbd, 0x4d,
	0xbe, 0x76, 0x4a, 0x4a, 0x13, 0x2a, 0xa1, 0xbe, 0xf9, 0xfa, 0x5e, 0x9f, 0x73, 0xef, 0x39, 0x76,
	0x02, 0x07, 0x59, 0xce, 0x25, 0x3f, 0xe2, 0xf9, 0xd7, 0x30, 0x65, 0x3f, 0x43, 0xc9, 0x78, 0x7a,
	0x94, 0x73, 0x9e, 0x8c, 0x71, 0x9f, 0x6c, 0xf3, 0x84, 0x8f, 0x13, 0x91, 0x8d, 0x45, 0x44, 0x53,
	0x3a, 0x38, 0xac, 0xa9, 0x8e, 0x78, 0x92, 0xf0, 0x54, 0xd7, 0xfb, 0x7f, 0x2c, 0xd8, 0x08, 0x38,
	0x4f, 0x3e, 0xa4, 0xb7, 0x9c, 0xec, 0x80, 0x5d, 0xb0, 0xd8, 0x6b, 0x0f, 0xdb, 0xa3, 0xcd, 0x40,
	0x2d, 0x49, 0x1f, 0x2c, 0x16, 0x7b, 0xd6, 0xb0, 0x3d, 0x72, 0x02, 0x8b, 0xc5, 0x84, 0x80, 0x93,
	0x86, 0x09, 0xf5, 0x6c, 0x2c, 0xc1, 0x35, 0xf1, 0xa0, 0x1b, 0xe5, 0x34, 0x94, 0x34, 0xf6, 0x9c,
	0x61, 0x7b, 0x64, 0x07, 0x65, 0xa8, 0x32, 0x45, 0x16, 0x63, 0xa6, 0xa3, 0x33, 0x26, 0x9c, 0x9f,
	0xe1, 0xb9, 0xe7, 0x22, 0x54, 0x19, 0x92, 0x01, 0x6c, 0xf0, 0x8c, 0xe6, 0x98, 0xea, 0x62, 0x6a,
	0x1e, 0x93, 0x5d, 0xe8, 0xf0, 0x1f, 0x29, 0xcd, 0xbd, 0x0d, 0x4c, 0xe8, 0x80, 0xec, 0x81, 0x9b,
	0xd3, 0x24, 0xcc, 0xbf, 0x7b, 0x9b, 0xb8, 0x6d, 0x22, 0xb5, 0x7f, 0x57, 0x70, 0x49, 0x85, 0x07,
	0x43, 0x5b, 0xed, 0xeb, 0x88, 0xbc, 0x83, 0x6e, 0x4c, 0xef, 0x59, 0x44, 0x85, 0xd7, 0x1b, 0xda,
	0xa3, 0xde, 0x64, 0x30, 0x5e, 0x10, 0x6d, 0xfc, 0x29, 0xe7, 0x71, 0x11, 0x49, 0x25, 0x49, 0x50,
	0x96, 0xfa, 0xdf, 0x00, 0x02, 0x7a, 0xa7, 0xa4, 0x3a, 0x8e, 0xe3, 0xc7, 0x4e, 0xda, 0xd5, 0x4e,
	0x4a, 0x75, 0xac, 0x8a, 0x3a, 0x8f, 0xdd, 0xd9, 0x0b, 0xdd, 0x55, 0xe7, 0x74, 0x16, 0xe7, 0xf4,
	0x19, 0x6c, 0x1b, 0xae, 0xcf, 0xa8, 0x57, 0x8d, 0x31, 0xeb, 0xa7, 0xba, 0xd2, 0xea, 0xec, 0x42,
	0x07, 0x55, 0x28, 0x27, 0xc3, 0x40, 0xd1, 0xa9, 0x4b, 0x56, 0xd2, 0xa9, 0xf5, 0x02, 0xac, 0xfd,
	0xc4, 0x29, 0x02, 0xce, 0x8c, 0x09, 0xe9, 0x39, 0xa8, 0x3c, 0xae, 0xfd, 0x5f, 0xf0, 0x9f, 0xa1,
	0x3a, 0x63, 0x22, 0x9b, 0x85, 0x0f, 0xeb, 0x22, 0xeb, 0x83, 0x25, 0x52, 0x33, 0x99, 0x25, 0xd2,
	0x39, 0x79, 0xa7, 0x42, 0xfe, 0xbb, 0x3d, 0x1f, 0xf4, 0x0c, 0x1d, 0xad, 0xd1, 0xb4, 0xca, 0x61,
	0x3d, 0xe1, 0xd8, 0x03, 0x57, 0xdf, 0x84, 0x52, 0x5b, 0x1d, 0x29, 0x2e, 0xf9, 0x90, 0x51, 0x64,
	0xdf, 0x0e, 0x70, 0x5d, 0xf1, 0xa1, 0x53, 0xf5, 0xc1, 0xbf, 0x87, 0x9d, 0x80, 0x66, 0xb3, 0x87,
	0xc7, 0x26, 0x04, 0x99, 0x80, 0x2b, 0x64, 0x28, 0x0b, 0x81, 0x8d, 0x2c, 0xdf, 0x45, 0x3c, 0x70,
	0x8d, 0x15, 0x81, 0xa9, 0x24, 0x63, 0x33, 0x9f, 0xb5, 0xf2, 0xf6, 0xea, 0xd9, 0x33, 0x35, 0xba,
	0xe1, 0xc5, 0x77, 0xfe, 0x12, 0xd2, 0x37, 0xe0, 0xb0, 0xf4, 0x96, 0xa3, 0x30, 0xbd, 0xc9, 0xff,
	0x4f, 0x4f, 0x18, 0xe8, 0x00, 0x8b, 0x16, 0x18, 0x2f, 0x98, 0x90, 0x2f, 0x65, 0xac, 0x8c, 0xd9,
	0xcc, 0xa8, 0x8a, 0x26, 0x7f, 0x3b, 0xd0, 0x53, 0x5b, 0xd7, 0x34, 0x47, 0x5f, 0x4e, 0xc1, 0x3d,
	0x8e, 0xe3, 0xcb, 0x94, 0x92, 0x57, 0x4b, 0x54, 0xe5, 0x2b, 0x1e, 0xec, 0xd7, 0x75, 0x51, 0x02,
	0xfb, 0x2d, 0x72, 0x0e, 0xa0, 0x1f, 0xe0, 0x49, 0x28, 0x28, 0xd9, 0xaf, 0x07, 0xd2, 0x15, 0x03,
	0xaf, 0x0e, 0xcb, 0xe0, 0x9c, 0xc2, 0x66, 0x40, 0x13, 0x7e, 0x4f, 0x55, 0x3f, 0xcb, 0xa3, 0xdf,
	0x15, 0x54, 0xa0, 0x5f, 0xcf, 0x82, 0x9c, 0x81, 0x3b, 0xa5, 0x72, 0x15, 0xc2, 0xaa, 0x91, 0xa6,
	0xd0, 0x9d, 0x52, 0x89, 0x9e, 0xec, 0xd7, 0xc3, 0x9c, 0xb3, 0x99, 0xa4, 0x79, 0x33, 0x90, 0x3a,
	0xeb, 0xb7, 0xc8, 0x47, 0xd8, 0x9a, 0x52, 0xa9, 0x2c, 0x63, 0x42, 0xb2, 0x68, 0x05, 0xda, 0x41,
	0x93, 0xdf, 0x78, 0xd8, 0x6f, 0x91, 0x4b, 0xd8, 0x3a, 0xce, 0x32, 0x9a, 0xc6, 0xe6, 0x75, 0x36,
	0x88, 0xad, 0xb3, 0x83, 0xc3, 0xa6, 0xe6, 0xcc, 0xb3, 0xf2, 0x5b, 0xe4, 0x0a, 0xfa, 0xd7, 0xc5,
	0x8d, 0xcc, 0xc3, 0x48, 0xae, 0x0b, 0xf2, 0x3d, 0x6c, 0x69, 0xb3, 0xcd, 0xa7, 0xb2, 0x01, 0x50,
	0x67, 0x9f, 0xf5, 0xf2, 0x02, 0xfa, 0x1a, 0x69, 0xfe, 0x25, 0x7c, 0xdd, 0xd0, 0x9c, 0xc9, 0x3f,
	0x87, 0x76, 0xb2, 0xfb, 0x85, 0x2c, 0xff, 0xe6, 0x6f, 0x5c, 0xdc, 0x7b, 0xfb, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x84, 0x3a, 0x6f, 0x75, 0x31, 0x08, 0x00, 0x00,
}
