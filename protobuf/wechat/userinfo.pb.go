// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userinfo.proto

package wechat

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

type OplogUserNameRequest struct {
	CmId                 *uint32   `protobuf:"varint,1,opt,name=CmId" json:"CmId,omitempty"`
	UserInfo             *UserInfo `protobuf:"bytes,2,opt,name=userInfo" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OplogUserNameRequest) Reset()         { *m = OplogUserNameRequest{} }
func (m *OplogUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*OplogUserNameRequest) ProtoMessage()    {}
func (*OplogUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{0}
}

func (m *OplogUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OplogUserNameRequest.Unmarshal(m, b)
}
func (m *OplogUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OplogUserNameRequest.Marshal(b, m, deterministic)
}
func (m *OplogUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OplogUserNameRequest.Merge(m, src)
}
func (m *OplogUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_OplogUserNameRequest.Size(m)
}
func (m *OplogUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OplogUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OplogUserNameRequest proto.InternalMessageInfo

func (m *OplogUserNameRequest) GetCmId() uint32 {
	if m != nil && m.CmId != nil {
		return *m.CmId
	}
	return 0
}

func (m *OplogUserNameRequest) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type UserInfo struct {
	Code                 *uint32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	InfoNew              *InfoNew `protobuf:"bytes,2,opt,name=infoNew" json:"infoNew,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *UserInfo) GetInfoNew() *InfoNew {
	if m != nil {
		return m.InfoNew
	}
	return nil
}

type InfoNew struct {
	InfoCode             *uint32  `protobuf:"varint,1,opt,name=infoCode" json:"infoCode,omitempty"`
	NickName             *string  `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoNew) Reset()         { *m = InfoNew{} }
func (m *InfoNew) String() string { return proto.CompactTextString(m) }
func (*InfoNew) ProtoMessage()    {}
func (*InfoNew) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{2}
}

func (m *InfoNew) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoNew.Unmarshal(m, b)
}
func (m *InfoNew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoNew.Marshal(b, m, deterministic)
}
func (m *InfoNew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoNew.Merge(m, src)
}
func (m *InfoNew) XXX_Size() int {
	return xxx_messageInfo_InfoNew.Size(m)
}
func (m *InfoNew) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoNew.DiscardUnknown(m)
}

var xxx_messageInfo_InfoNew proto.InternalMessageInfo

func (m *InfoNew) GetInfoCode() uint32 {
	if m != nil && m.InfoCode != nil {
		return *m.InfoCode
	}
	return 0
}

func (m *InfoNew) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func init() {
	proto.RegisterType((*OplogUserNameRequest)(nil), "wechat_proto.oplogUserNameRequest")
	proto.RegisterType((*UserInfo)(nil), "wechat_proto.userInfo")
	proto.RegisterType((*InfoNew)(nil), "wechat_proto.infoNew")
}

func init() { proto.RegisterFile("userinfo.proto", fileDescriptor_785a78c34699a93d) }

var fileDescriptor_785a78c34699a93d = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x4f, 0x4d, 0xce,
	0x48, 0x2c, 0x89, 0x07, 0xf3, 0xa4, 0xa0, 0x3c, 0x88, 0x9c, 0x52, 0x1c, 0x97, 0x48, 0x7e, 0x41,
	0x4e, 0x7e, 0x7a, 0x68, 0x71, 0x6a, 0x91, 0x5f, 0x62, 0x6e, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x90, 0x10, 0x17, 0x8b, 0x73, 0xae, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6f,
	0x10, 0x98, 0x2d, 0x64, 0xc4, 0xc5, 0x01, 0x32, 0xd9, 0x33, 0x2f, 0x2d, 0x5f, 0x82, 0x49, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x4c, 0x0f, 0xd9, 0x68, 0x3d, 0x98, 0x6c, 0x10, 0x5c, 0x9d, 0x92, 0x3f,
	0x42, 0x0f, 0xc8, 0xcc, 0xe4, 0xfc, 0x94, 0x54, 0x98, 0x99, 0x20, 0xb6, 0x90, 0x3e, 0x17, 0x3b,
	0xc8, 0xa5, 0x7e, 0xa9, 0xe5, 0x50, 0x23, 0x45, 0x51, 0x8d, 0x84, 0x4a, 0x06, 0xc1, 0x54, 0x29,
	0x39, 0xc2, 0x35, 0x08, 0x49, 0x71, 0x71, 0x80, 0x98, 0xce, 0x08, 0x33, 0xe1, 0x7c, 0x90, 0x5c,
	0x5e, 0x66, 0x72, 0x36, 0xc8, 0x4b, 0x60, 0x83, 0x39, 0x83, 0xe0, 0x7c, 0x27, 0xee, 0x28, 0x4e,
	0x3d, 0x3d, 0x7d, 0x88, 0x35, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xc0, 0xa6, 0x84, 0x2d,
	0x01, 0x00, 0x00,
}