// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/scene.proto

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

type SceneInfo struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64          `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64          `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string         `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32          `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Status               int32          `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Cover                string         `protobuf:"bytes,8,opt,name=cover,proto3" json:"cover,omitempty"`
	Master               string         `protobuf:"bytes,9,opt,name=master,proto3" json:"master,omitempty"`
	Remark               string         `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string         `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	Entity               string         `protobuf:"bytes,12,opt,name=entity,proto3" json:"entity,omitempty"`
	Creator              string         `protobuf:"bytes,13,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,14,opt,name=operator,proto3" json:"operator,omitempty"`
	Address              *AddressInfo   `protobuf:"bytes,15,opt,name=address,proto3" json:"address,omitempty"`
	Members              []string       `protobuf:"bytes,16,rep,name=members,proto3" json:"members,omitempty"`
	Exhibitions          []*ExhibitInfo `protobuf:"bytes,17,rep,name=exhibitions,proto3" json:"exhibitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SceneInfo) Reset()         { *m = SceneInfo{} }
func (m *SceneInfo) String() string { return proto.CompactTextString(m) }
func (*SceneInfo) ProtoMessage()    {}
func (*SceneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{0}
}

func (m *SceneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SceneInfo.Unmarshal(m, b)
}
func (m *SceneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SceneInfo.Marshal(b, m, deterministic)
}
func (m *SceneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SceneInfo.Merge(m, src)
}
func (m *SceneInfo) XXX_Size() int {
	return xxx_messageInfo_SceneInfo.Size(m)
}
func (m *SceneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SceneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SceneInfo proto.InternalMessageInfo

func (m *SceneInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SceneInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SceneInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SceneInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *SceneInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SceneInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SceneInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SceneInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *SceneInfo) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *SceneInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *SceneInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SceneInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *SceneInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SceneInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *SceneInfo) GetAddress() *AddressInfo {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *SceneInfo) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *SceneInfo) GetExhibitions() []*ExhibitInfo {
	if m != nil {
		return m.Exhibitions
	}
	return nil
}

type ExhibitInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Effect               string   `protobuf:"bytes,2,opt,name=effect,proto3" json:"effect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExhibitInfo) Reset()         { *m = ExhibitInfo{} }
func (m *ExhibitInfo) String() string { return proto.CompactTextString(m) }
func (*ExhibitInfo) ProtoMessage()    {}
func (*ExhibitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{1}
}

func (m *ExhibitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExhibitInfo.Unmarshal(m, b)
}
func (m *ExhibitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExhibitInfo.Marshal(b, m, deterministic)
}
func (m *ExhibitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExhibitInfo.Merge(m, src)
}
func (m *ExhibitInfo) XXX_Size() int {
	return xxx_messageInfo_ExhibitInfo.Size(m)
}
func (m *ExhibitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExhibitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExhibitInfo proto.InternalMessageInfo

func (m *ExhibitInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ExhibitInfo) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

type ReqSceneAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Master               string   `protobuf:"bytes,4,opt,name=master,proto3" json:"master,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Entity               string   `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneAdd) Reset()         { *m = ReqSceneAdd{} }
func (m *ReqSceneAdd) String() string { return proto.CompactTextString(m) }
func (*ReqSceneAdd) ProtoMessage()    {}
func (*ReqSceneAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{2}
}

func (m *ReqSceneAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneAdd.Unmarshal(m, b)
}
func (m *ReqSceneAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneAdd.Marshal(b, m, deterministic)
}
func (m *ReqSceneAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneAdd.Merge(m, src)
}
func (m *ReqSceneAdd) XXX_Size() int {
	return xxx_messageInfo_ReqSceneAdd.Size(m)
}
func (m *ReqSceneAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneAdd proto.InternalMessageInfo

func (m *ReqSceneAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSceneAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqSceneAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqSceneAdd) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqSceneAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSceneAdd) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqSceneAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqSceneAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplySceneOne struct {
	Info                 *SceneInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneOne) Reset()         { *m = ReplySceneOne{} }
func (m *ReplySceneOne) String() string { return proto.CompactTextString(m) }
func (*ReplySceneOne) ProtoMessage()    {}
func (*ReplySceneOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{3}
}

func (m *ReplySceneOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneOne.Unmarshal(m, b)
}
func (m *ReplySceneOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneOne.Marshal(b, m, deterministic)
}
func (m *ReplySceneOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneOne.Merge(m, src)
}
func (m *ReplySceneOne) XXX_Size() int {
	return xxx_messageInfo_ReplySceneOne.Size(m)
}
func (m *ReplySceneOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneOne proto.InternalMessageInfo

func (m *ReplySceneOne) GetInfo() *SceneInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplySceneOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplySceneList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32       `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32       `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*SceneInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneList) Reset()         { *m = ReplySceneList{} }
func (m *ReplySceneList) String() string { return proto.CompactTextString(m) }
func (*ReplySceneList) ProtoMessage()    {}
func (*ReplySceneList) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{4}
}

func (m *ReplySceneList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneList.Unmarshal(m, b)
}
func (m *ReplySceneList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneList.Marshal(b, m, deterministic)
}
func (m *ReplySceneList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneList.Merge(m, src)
}
func (m *ReplySceneList) XXX_Size() int {
	return xxx_messageInfo_ReplySceneList.Size(m)
}
func (m *ReplySceneList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneList proto.InternalMessageInfo

func (m *ReplySceneList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplySceneList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplySceneList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplySceneList) GetList() []*SceneInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplySceneList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqSceneUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Master               string   `protobuf:"bytes,6,opt,name=master,proto3" json:"master,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneUpdate) Reset()         { *m = ReqSceneUpdate{} }
func (m *ReqSceneUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqSceneUpdate) ProtoMessage()    {}
func (*ReqSceneUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{5}
}

func (m *ReqSceneUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneUpdate.Unmarshal(m, b)
}
func (m *ReqSceneUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneUpdate.Marshal(b, m, deterministic)
}
func (m *ReqSceneUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneUpdate.Merge(m, src)
}
func (m *ReqSceneUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqSceneUpdate.Size(m)
}
func (m *ReqSceneUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneUpdate proto.InternalMessageInfo

func (m *ReqSceneUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSceneUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqSceneUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSceneUpdate) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqSceneUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyMasterUsed struct {
	Master               string       `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Used                 bool         `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMasterUsed) Reset()         { *m = ReplyMasterUsed{} }
func (m *ReplyMasterUsed) String() string { return proto.CompactTextString(m) }
func (*ReplyMasterUsed) ProtoMessage()    {}
func (*ReplyMasterUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{6}
}

func (m *ReplyMasterUsed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMasterUsed.Unmarshal(m, b)
}
func (m *ReplyMasterUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMasterUsed.Marshal(b, m, deterministic)
}
func (m *ReplyMasterUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMasterUsed.Merge(m, src)
}
func (m *ReplyMasterUsed) XXX_Size() int {
	return xxx_messageInfo_ReplyMasterUsed.Size(m)
}
func (m *ReplyMasterUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMasterUsed.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMasterUsed proto.InternalMessageInfo

func (m *ReplyMasterUsed) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReplyMasterUsed) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

func (m *ReplyMasterUsed) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqSceneStatus struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneStatus) Reset()         { *m = ReqSceneStatus{} }
func (m *ReqSceneStatus) String() string { return proto.CompactTextString(m) }
func (*ReqSceneStatus) ProtoMessage()    {}
func (*ReqSceneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{7}
}

func (m *ReqSceneStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneStatus.Unmarshal(m, b)
}
func (m *ReqSceneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneStatus.Marshal(b, m, deterministic)
}
func (m *ReqSceneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneStatus.Merge(m, src)
}
func (m *ReqSceneStatus) XXX_Size() int {
	return xxx_messageInfo_ReqSceneStatus.Size(m)
}
func (m *ReqSceneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneStatus proto.InternalMessageInfo

func (m *ReqSceneStatus) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqSceneStatus) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqSceneDisplay struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneDisplay) Reset()         { *m = ReqSceneDisplay{} }
func (m *ReqSceneDisplay) String() string { return proto.CompactTextString(m) }
func (*ReqSceneDisplay) ProtoMessage()    {}
func (*ReqSceneDisplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{8}
}

func (m *ReqSceneDisplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneDisplay.Unmarshal(m, b)
}
func (m *ReqSceneDisplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneDisplay.Marshal(b, m, deterministic)
}
func (m *ReqSceneDisplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneDisplay.Merge(m, src)
}
func (m *ReqSceneDisplay) XXX_Size() int {
	return xxx_messageInfo_ReqSceneDisplay.Size(m)
}
func (m *ReqSceneDisplay) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneDisplay.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneDisplay proto.InternalMessageInfo

func (m *ReqSceneDisplay) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqSceneDisplay) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneDisplay) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqSceneDisplay) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplySceneDisplays struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Uid                  string         `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []*ExhibitInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplySceneDisplays) Reset()         { *m = ReplySceneDisplays{} }
func (m *ReplySceneDisplays) String() string { return proto.CompactTextString(m) }
func (*ReplySceneDisplays) ProtoMessage()    {}
func (*ReplySceneDisplays) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{9}
}

func (m *ReplySceneDisplays) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneDisplays.Unmarshal(m, b)
}
func (m *ReplySceneDisplays) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneDisplays.Marshal(b, m, deterministic)
}
func (m *ReplySceneDisplays) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneDisplays.Merge(m, src)
}
func (m *ReplySceneDisplays) XXX_Size() int {
	return xxx_messageInfo_ReplySceneDisplays.Size(m)
}
func (m *ReplySceneDisplays) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneDisplays.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneDisplays proto.InternalMessageInfo

func (m *ReplySceneDisplays) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplySceneDisplays) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplySceneDisplays) GetList() []*ExhibitInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SceneInfo)(nil), "omo.msp.scene.SceneInfo")
	proto.RegisterType((*ExhibitInfo)(nil), "omo.msp.scene.ExhibitInfo")
	proto.RegisterType((*ReqSceneAdd)(nil), "omo.msp.scene.ReqSceneAdd")
	proto.RegisterType((*ReplySceneOne)(nil), "omo.msp.scene.ReplySceneOne")
	proto.RegisterType((*ReplySceneList)(nil), "omo.msp.scene.ReplySceneList")
	proto.RegisterType((*ReqSceneUpdate)(nil), "omo.msp.scene.ReqSceneUpdate")
	proto.RegisterType((*ReplyMasterUsed)(nil), "omo.msp.scene.ReplyMasterUsed")
	proto.RegisterType((*ReqSceneStatus)(nil), "omo.msp.scene.ReqSceneStatus")
	proto.RegisterType((*ReqSceneDisplay)(nil), "omo.msp.scene.ReqSceneDisplay")
	proto.RegisterType((*ReplySceneDisplays)(nil), "omo.msp.scene.ReplySceneDisplays")
}

func init() {
	proto.RegisterFile("proto/organization/scene.proto", fileDescriptor_059c9a53c993103b)
}

var fileDescriptor_059c9a53c993103b = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x45, 0xfd, 0x79, 0xf4, 0x13, 0x77, 0x11, 0x14, 0x0b, 0xa1, 0x71, 0x55, 0x9e, 0x74,
	0x28, 0x14, 0xc0, 0x2d, 0xd0, 0x4b, 0x2f, 0x76, 0x52, 0x3b, 0x06, 0xfc, 0x07, 0xa6, 0xe9, 0xa1,
	0x37, 0x9a, 0x1c, 0xab, 0x44, 0x44, 0x2e, 0xc3, 0x5d, 0x39, 0x51, 0x1f, 0xa1, 0x4f, 0xd0, 0x87,
	0xe9, 0x1b, 0xf4, 0x21, 0xfa, 0x22, 0x3d, 0x14, 0x3b, 0xbb, 0x94, 0x28, 0x9b, 0x94, 0x5d, 0x38,
	0xb7, 0x9d, 0x9d, 0xe1, 0x37, 0x3f, 0xdf, 0xb7, 0x23, 0xc1, 0x7e, 0x96, 0x0b, 0x25, 0x5e, 0x8a,
	0x7c, 0x16, 0xa4, 0xf1, 0xef, 0x81, 0x8a, 0x45, 0xfa, 0x52, 0x86, 0x98, 0xe2, 0x94, 0x1c, 0x6c,
	0x20, 0x12, 0x31, 0x4d, 0x64, 0x36, 0xa5, 0xcb, 0xd1, 0xd7, 0x15, 0xe1, 0xa1, 0x48, 0x12, 0x91,
	0x9a, 0x78, 0xef, 0x1f, 0x17, 0x76, 0xdf, 0xea, 0xd0, 0xd3, 0xf4, 0x46, 0xb0, 0x3d, 0x70, 0x17,
	0x71, 0xc4, 0x9d, 0xb1, 0x33, 0xd9, 0xf5, 0xf5, 0x91, 0x0d, 0xa1, 0x11, 0x47, 0xbc, 0x31, 0x76,
	0x26, 0x4d, 0xbf, 0x11, 0x47, 0x8c, 0x43, 0x27, 0xcc, 0x31, 0x50, 0x18, 0x71, 0x77, 0xec, 0x4c,
	0x5c, 0xbf, 0x30, 0xb5, 0x67, 0x91, 0x45, 0xe4, 0x69, 0x1a, 0x8f, 0x35, 0x19, 0x83, 0x66, 0x1a,
	0x24, 0xc8, 0x5b, 0x04, 0x4b, 0x67, 0x7d, 0xa7, 0x96, 0x19, 0xf2, 0xf6, 0xd8, 0x99, 0xb4, 0x7c,
	0x3a, 0xb3, 0x2f, 0xa1, 0x2d, 0x55, 0xa0, 0x16, 0x92, 0x77, 0xe8, 0xd6, 0x5a, 0xec, 0x39, 0xb4,
	0x42, 0x71, 0x8b, 0x39, 0xef, 0x12, 0x80, 0x31, 0x74, 0x74, 0x12, 0x48, 0x85, 0x39, 0xdf, 0xa5,
	0x6b, 0x6b, 0xe9, 0xfb, 0x1c, 0x93, 0x20, 0x7f, 0xcf, 0xc1, 0xdc, 0x1b, 0x8b, 0x8d, 0xa0, 0x3b,
	0x17, 0x21, 0x8d, 0x80, 0xf7, 0xc8, 0xb3, 0xb2, 0xf5, 0x37, 0x98, 0xaa, 0x58, 0x2d, 0x79, 0xdf,
	0x7c, 0x63, 0xac, 0x55, 0xb7, 0x22, 0xe7, 0x03, 0x72, 0x14, 0xa6, 0x46, 0x13, 0x19, 0xe6, 0xe4,
	0x1a, 0x1a, 0xb4, 0xc2, 0x66, 0xdf, 0x43, 0x27, 0x88, 0xa2, 0x1c, 0xa5, 0xe4, 0xcf, 0xc6, 0xce,
	0xa4, 0x77, 0x30, 0x9a, 0x6e, 0xb0, 0x32, 0x3d, 0x34, 0x5e, 0x3d, 0x72, 0xbf, 0x08, 0xd5, 0xb9,
	0x12, 0x4c, 0xae, 0x31, 0x97, 0x7c, 0x6f, 0xec, 0xea, 0x5c, 0xd6, 0x64, 0x3f, 0x42, 0x0f, 0x3f,
	0xfd, 0x16, 0x5f, 0xc7, 0xba, 0x56, 0xc9, 0xbf, 0x18, 0xbb, 0x15, 0x98, 0x3f, 0x99, 0x08, 0xc2,
	0x2c, 0x87, 0x7b, 0x3f, 0x40, 0xaf, 0xe4, 0xab, 0xa0, 0x58, 0x37, 0x7f, 0x73, 0x83, 0xa1, 0x22,
	0x9a, 0x75, 0xf3, 0x64, 0x79, 0x7f, 0x3b, 0xd0, 0xf3, 0xf1, 0x03, 0xa9, 0xe3, 0x30, 0x5a, 0xd3,
	0xe8, 0x54, 0xd0, 0xd8, 0x28, 0xd1, 0xb8, 0xa2, 0xcb, 0xad, 0xa6, 0xab, 0x59, 0x43, 0x57, 0xab,
	0x96, 0xae, 0x76, 0x2d, 0x5d, 0x9d, 0x0d, 0xba, 0xca, 0xa4, 0x74, 0x37, 0x49, 0xf1, 0x3e, 0xc0,
	0xc0, 0xc7, 0x6c, 0xbe, 0xa4, 0x76, 0x2e, 0x53, 0x64, 0xdf, 0x42, 0x33, 0x4e, 0x6f, 0x04, 0xb5,
	0xd3, 0x3b, 0xe0, 0x77, 0xc6, 0xb9, 0x7a, 0x13, 0x3e, 0x45, 0xb1, 0x83, 0x95, 0x36, 0x1b, 0x95,
	0x94, 0x1a, 0x6c, 0x8a, 0x28, 0x74, 0xeb, 0xfd, 0xe5, 0xc0, 0x70, 0x9d, 0xf3, 0x2c, 0x96, 0x4a,
	0xcf, 0x46, 0x09, 0x15, 0xcc, 0x29, 0xeb, 0xc0, 0x37, 0x86, 0xa6, 0x3e, 0x0b, 0x66, 0x78, 0x1e,
	0x7c, 0x22, 0xf4, 0x81, 0x5f, 0x98, 0x85, 0xe7, 0x42, 0x7c, 0xa4, 0x69, 0x5a, 0xcf, 0x85, 0xf8,
	0xa8, 0xcb, 0x9f, 0xc7, 0x52, 0xf1, 0x26, 0xa9, 0x61, 0x4b, 0xf9, 0x3a, 0xaa, 0x54, 0x7e, 0xeb,
	0xd1, 0xe5, 0xff, 0x49, 0xe5, 0x1b, 0xfe, 0xdf, 0xd1, 0x53, 0xae, 0x10, 0x4f, 0x21, 0x8a, 0x46,
	0x49, 0x14, 0xb5, 0x02, 0xb0, 0x44, 0x37, 0x37, 0x88, 0x5e, 0x0b, 0xa3, 0xbd, 0x21, 0x8c, 0x32,
	0x99, 0x9d, 0x7b, 0x64, 0x3e, 0xa3, 0x8a, 0xcf, 0x29, 0xf4, 0x9d, 0xc4, 0xa8, 0x04, 0xe3, 0x6c,
	0xc0, 0x30, 0x68, 0x2e, 0x24, 0x9a, 0x15, 0xd6, 0xf5, 0xe9, 0x5c, 0x9a, 0x86, 0xfb, 0xe8, 0x69,
	0xfc, 0xb2, 0x1e, 0x86, 0xf1, 0x54, 0xbf, 0xa4, 0x92, 0x48, 0xd6, 0x0b, 0xac, 0xdc, 0x8a, 0x7b,
	0xa7, 0x95, 0x99, 0x6e, 0xc5, 0xe0, 0xbe, 0x8e, 0x65, 0x36, 0x0f, 0x96, 0x7a, 0x7e, 0x54, 0x87,
	0x85, 0x36, 0x46, 0x91, 0xae, 0xb1, 0x4e, 0xb7, 0x07, 0xee, 0x7b, 0x5c, 0x5a, 0x44, 0x7d, 0xdc,
	0x48, 0xd4, 0xbc, 0x93, 0xe8, 0x0f, 0x07, 0xd8, 0x5a, 0x8d, 0x36, 0x97, 0x2c, 0xcd, 0xc2, 0x79,
	0xec, 0x2c, 0x2a, 0x4a, 0x99, 0x5a, 0x35, 0xba, 0x0f, 0xee, 0x26, 0x8a, 0x3b, 0xf8, 0xb7, 0x0b,
	0x7d, 0x33, 0x4b, 0xcc, 0x6f, 0xe3, 0x10, 0xd9, 0x6b, 0x68, 0x1f, 0x46, 0x91, 0x7e, 0x97, 0xf7,
	0x0b, 0x58, 0xad, 0xa0, 0xd1, 0x57, 0x95, 0xc5, 0xd9, 0x17, 0xed, 0xed, 0x68, 0x94, 0x13, 0x54,
	0x35, 0x28, 0x0b, 0x94, 0x54, 0xc2, 0x83, 0x28, 0x67, 0x30, 0x34, 0x28, 0x47, 0x56, 0x60, 0x4f,
	0x42, 0x7b, 0x05, 0xbb, 0x3e, 0x26, 0xe2, 0x16, 0x1f, 0x2a, 0x8b, 0x57, 0x01, 0x69, 0x0f, 0x95,
	0xd4, 0x3f, 0x95, 0x25, 0xb5, 0x6f, 0xc3, 0xd9, 0xaf, 0xc2, 0x59, 0x7f, 0xeb, 0xed, 0xb0, 0x63,
	0xe8, 0x9c, 0xa0, 0xa2, 0x85, 0x54, 0x03, 0x74, 0x15, 0xcc, 0x70, 0xf4, 0xa2, 0xb6, 0x33, 0xfd,
	0xa9, 0xb7, 0xc3, 0x4e, 0x00, 0xcc, 0x62, 0x38, 0x0a, 0x24, 0xb2, 0x17, 0x35, 0xc4, 0x99, 0x90,
	0xad, 0xed, 0x9d, 0x42, 0xdf, 0x44, 0xd9, 0xa7, 0x55, 0x07, 0x65, 0xdc, 0x5b, 0xa1, 0x2e, 0x60,
	0x60, 0xa0, 0xec, 0x8f, 0x6c, 0x15, 0x96, 0xee, 0xd0, 0xba, 0x1f, 0xa4, 0xef, 0x0d, 0x0c, 0x0d,
	0xde, 0x59, 0xf1, 0xeb, 0x53, 0x33, 0xb2, 0xe3, 0x79, 0x30, 0xdb, 0x5a, 0xd9, 0x31, 0xf4, 0x0f,
	0xb3, 0x0c, 0xd3, 0xe8, 0x9c, 0x7e, 0xd7, 0xff, 0xbf, 0x16, 0xec, 0xd4, 0xdf, 0xc0, 0xf0, 0xed,
	0xe2, 0x5a, 0xe5, 0x41, 0xa8, 0x9e, 0x88, 0x74, 0x09, 0xfd, 0xab, 0x85, 0xba, 0x4c, 0x8b, 0xc5,
	0xb3, 0x0d, 0xe7, 0x9b, 0xda, 0x39, 0x15, 0xab, 0xc4, 0xdb, 0x61, 0x57, 0x30, 0x78, 0x15, 0xa4,
	0x21, 0xce, 0x3f, 0x1b, 0xe2, 0xcf, 0x05, 0x9d, 0x05, 0xe2, 0x7e, 0x8d, 0x34, 0xac, 0xff, 0x51,
	0xa8, 0x47, 0xcf, 0x7f, 0x65, 0xf7, 0xff, 0x18, 0x5f, 0xb7, 0xe9, 0xee, 0xbb, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x2f, 0x3e, 0x65, 0x64, 0x0b, 0x00, 0x00,
}
