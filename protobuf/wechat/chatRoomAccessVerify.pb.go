// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatRoomAccessVerify.proto

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

type ModChatRoomAccessVerifyRequest struct {
	ChatRoomName         *string  `protobuf:"bytes,1,opt,name=chatRoomName" json:"chatRoomName,omitempty"`
	Status               *uint32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModChatRoomAccessVerifyRequest) Reset()         { *m = ModChatRoomAccessVerifyRequest{} }
func (m *ModChatRoomAccessVerifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModChatRoomAccessVerifyRequest) ProtoMessage()    {}
func (*ModChatRoomAccessVerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a2e75d48969b67, []int{0}
}

func (m *ModChatRoomAccessVerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModChatRoomAccessVerifyRequest.Unmarshal(m, b)
}
func (m *ModChatRoomAccessVerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModChatRoomAccessVerifyRequest.Marshal(b, m, deterministic)
}
func (m *ModChatRoomAccessVerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModChatRoomAccessVerifyRequest.Merge(m, src)
}
func (m *ModChatRoomAccessVerifyRequest) XXX_Size() int {
	return xxx_messageInfo_ModChatRoomAccessVerifyRequest.Size(m)
}
func (m *ModChatRoomAccessVerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModChatRoomAccessVerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModChatRoomAccessVerifyRequest proto.InternalMessageInfo

func (m *ModChatRoomAccessVerifyRequest) GetChatRoomName() string {
	if m != nil && m.ChatRoomName != nil {
		return *m.ChatRoomName
	}
	return ""
}

func (m *ModChatRoomAccessVerifyRequest) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

type AddChatRoomAdminRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,opt,name=baseRequest" json:"baseRequest,omitempty"`
	ChatRoomName         *string      `protobuf:"bytes,2,opt,name=chatRoomName" json:"chatRoomName,omitempty"`
	UserNameList         []string     `protobuf:"bytes,3,rep,name=userNameList" json:"userNameList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AddChatRoomAdminRequest) Reset()         { *m = AddChatRoomAdminRequest{} }
func (m *AddChatRoomAdminRequest) String() string { return proto.CompactTextString(m) }
func (*AddChatRoomAdminRequest) ProtoMessage()    {}
func (*AddChatRoomAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a2e75d48969b67, []int{1}
}

func (m *AddChatRoomAdminRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddChatRoomAdminRequest.Unmarshal(m, b)
}
func (m *AddChatRoomAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddChatRoomAdminRequest.Marshal(b, m, deterministic)
}
func (m *AddChatRoomAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddChatRoomAdminRequest.Merge(m, src)
}
func (m *AddChatRoomAdminRequest) XXX_Size() int {
	return xxx_messageInfo_AddChatRoomAdminRequest.Size(m)
}
func (m *AddChatRoomAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddChatRoomAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddChatRoomAdminRequest proto.InternalMessageInfo

func (m *AddChatRoomAdminRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *AddChatRoomAdminRequest) GetChatRoomName() string {
	if m != nil && m.ChatRoomName != nil {
		return *m.ChatRoomName
	}
	return ""
}

func (m *AddChatRoomAdminRequest) GetUserNameList() []string {
	if m != nil {
		return m.UserNameList
	}
	return nil
}

type AddChatRoomAdminResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,opt,name=baseResponse" json:"baseResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddChatRoomAdminResponse) Reset()         { *m = AddChatRoomAdminResponse{} }
func (m *AddChatRoomAdminResponse) String() string { return proto.CompactTextString(m) }
func (*AddChatRoomAdminResponse) ProtoMessage()    {}
func (*AddChatRoomAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a2e75d48969b67, []int{2}
}

func (m *AddChatRoomAdminResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddChatRoomAdminResponse.Unmarshal(m, b)
}
func (m *AddChatRoomAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddChatRoomAdminResponse.Marshal(b, m, deterministic)
}
func (m *AddChatRoomAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddChatRoomAdminResponse.Merge(m, src)
}
func (m *AddChatRoomAdminResponse) XXX_Size() int {
	return xxx_messageInfo_AddChatRoomAdminResponse.Size(m)
}
func (m *AddChatRoomAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddChatRoomAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddChatRoomAdminResponse proto.InternalMessageInfo

func (m *AddChatRoomAdminResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

type DelChatRoomAdminRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,opt,name=baseRequest" json:"baseRequest,omitempty"`
	ChatRoomName         *string      `protobuf:"bytes,2,opt,name=chatRoomName" json:"chatRoomName,omitempty"`
	UserNameList         []string     `protobuf:"bytes,3,rep,name=userNameList" json:"userNameList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DelChatRoomAdminRequest) Reset()         { *m = DelChatRoomAdminRequest{} }
func (m *DelChatRoomAdminRequest) String() string { return proto.CompactTextString(m) }
func (*DelChatRoomAdminRequest) ProtoMessage()    {}
func (*DelChatRoomAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a2e75d48969b67, []int{3}
}

func (m *DelChatRoomAdminRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelChatRoomAdminRequest.Unmarshal(m, b)
}
func (m *DelChatRoomAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelChatRoomAdminRequest.Marshal(b, m, deterministic)
}
func (m *DelChatRoomAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelChatRoomAdminRequest.Merge(m, src)
}
func (m *DelChatRoomAdminRequest) XXX_Size() int {
	return xxx_messageInfo_DelChatRoomAdminRequest.Size(m)
}
func (m *DelChatRoomAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelChatRoomAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelChatRoomAdminRequest proto.InternalMessageInfo

func (m *DelChatRoomAdminRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *DelChatRoomAdminRequest) GetChatRoomName() string {
	if m != nil && m.ChatRoomName != nil {
		return *m.ChatRoomName
	}
	return ""
}

func (m *DelChatRoomAdminRequest) GetUserNameList() []string {
	if m != nil {
		return m.UserNameList
	}
	return nil
}

type DelChatRoomAdminResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,opt,name=baseResponse" json:"baseResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DelChatRoomAdminResponse) Reset()         { *m = DelChatRoomAdminResponse{} }
func (m *DelChatRoomAdminResponse) String() string { return proto.CompactTextString(m) }
func (*DelChatRoomAdminResponse) ProtoMessage()    {}
func (*DelChatRoomAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a2e75d48969b67, []int{4}
}

func (m *DelChatRoomAdminResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelChatRoomAdminResponse.Unmarshal(m, b)
}
func (m *DelChatRoomAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelChatRoomAdminResponse.Marshal(b, m, deterministic)
}
func (m *DelChatRoomAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelChatRoomAdminResponse.Merge(m, src)
}
func (m *DelChatRoomAdminResponse) XXX_Size() int {
	return xxx_messageInfo_DelChatRoomAdminResponse.Size(m)
}
func (m *DelChatRoomAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelChatRoomAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelChatRoomAdminResponse proto.InternalMessageInfo

func (m *DelChatRoomAdminResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*ModChatRoomAccessVerifyRequest)(nil), "wechat_proto.ModChatRoomAccessVerifyRequest")
	proto.RegisterType((*AddChatRoomAdminRequest)(nil), "wechat_proto.AddChatRoomAdminRequest")
	proto.RegisterType((*AddChatRoomAdminResponse)(nil), "wechat_proto.AddChatRoomAdminResponse")
	proto.RegisterType((*DelChatRoomAdminRequest)(nil), "wechat_proto.DelChatRoomAdminRequest")
	proto.RegisterType((*DelChatRoomAdminResponse)(nil), "wechat_proto.DelChatRoomAdminResponse")
}

func init() { proto.RegisterFile("chatRoomAccessVerify.proto", fileDescriptor_a4a2e75d48969b67) }

var fileDescriptor_a4a2e75d48969b67 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcd, 0x4f, 0x71,
	0xce, 0x48, 0x2c, 0x09, 0xca, 0xcf, 0xcf, 0x75, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0x0e, 0x4b, 0x2d,
	0xca, 0x4c, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x4f, 0x4d, 0xce, 0x48,
	0x2c, 0x89, 0x07, 0xf3, 0xa4, 0xa0, 0x3c, 0x88, 0x9c, 0x52, 0x0c, 0x97, 0x9c, 0x2f, 0x76, 0xcd,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x4a, 0x5c, 0x3c, 0xc9, 0x50, 0x69, 0xbf, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x14, 0x31, 0x21, 0x31, 0x2e, 0xb6, 0xe2,
	0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x28, 0x4f, 0x69, 0x1e,
	0x23, 0x97, 0xb8, 0x63, 0x0a, 0xc2, 0xf8, 0x94, 0xdc, 0xcc, 0x3c, 0x98, 0xb9, 0xd6, 0x5c, 0xdc,
	0x49, 0x89, 0xc5, 0xa9, 0x50, 0x2e, 0xd8, 0x58, 0x6e, 0x23, 0x49, 0x3d, 0x64, 0xb7, 0xea, 0x39,
	0x21, 0x14, 0x04, 0x21, 0xab, 0xc6, 0x70, 0x14, 0x13, 0x16, 0x47, 0x29, 0x71, 0xf1, 0x94, 0x16,
	0xa7, 0x16, 0x81, 0xd8, 0x3e, 0x99, 0xc5, 0x25, 0x12, 0xcc, 0x0a, 0xcc, 0x20, 0x35, 0xc8, 0x62,
	0x4a, 0x51, 0x5c, 0x12, 0x98, 0xee, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe3, 0xe2,
	0x81, 0x58, 0x09, 0xe1, 0x43, 0x5d, 0x28, 0x85, 0xcd, 0x85, 0x10, 0x15, 0x41, 0x28, 0xea, 0xc1,
	0x9e, 0x77, 0x49, 0xcd, 0x19, 0xd4, 0x9e, 0xc7, 0x74, 0x1f, 0x75, 0x3c, 0xef, 0xc4, 0x1d, 0xc5,
	0xa9, 0xa7, 0xa7, 0x0f, 0x51, 0x0d, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x89, 0xb3, 0x26, 0xa0,
	0x02, 0x00, 0x00,
}