// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatroom/member.proto

package chatroom

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "gitlab.srgow.com/warehouse/proto/common"
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

type Member struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{0}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type MemberResponse struct {
	Result               *common.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Member               *Member        `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MemberResponse) Reset()         { *m = MemberResponse{} }
func (m *MemberResponse) String() string { return proto.CompactTextString(m) }
func (*MemberResponse) ProtoMessage()    {}
func (*MemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{1}
}

func (m *MemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberResponse.Unmarshal(m, b)
}
func (m *MemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberResponse.Marshal(b, m, deterministic)
}
func (m *MemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberResponse.Merge(m, src)
}
func (m *MemberResponse) XXX_Size() int {
	return xxx_messageInfo_MemberResponse.Size(m)
}
func (m *MemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberResponse proto.InternalMessageInfo

func (m *MemberResponse) GetResult() *common.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *MemberResponse) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type SetDataRequest struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SetDataRequest) Reset()         { *m = SetDataRequest{} }
func (m *SetDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataRequest) ProtoMessage()    {}
func (*SetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{2}
}

func (m *SetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataRequest.Unmarshal(m, b)
}
func (m *SetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataRequest.Marshal(b, m, deterministic)
}
func (m *SetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataRequest.Merge(m, src)
}
func (m *SetDataRequest) XXX_Size() int {
	return xxx_messageInfo_SetDataRequest.Size(m)
}
func (m *SetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataRequest proto.InternalMessageInfo

func (m *SetDataRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SetDataRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetMemberRequest struct {
	MemberName           string   `protobuf:"bytes,1,opt,name=memberName,proto3" json:"memberName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemberRequest) Reset()         { *m = GetMemberRequest{} }
func (m *GetMemberRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemberRequest) ProtoMessage()    {}
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{3}
}

func (m *GetMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberRequest.Unmarshal(m, b)
}
func (m *GetMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberRequest.Marshal(b, m, deterministic)
}
func (m *GetMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberRequest.Merge(m, src)
}
func (m *GetMemberRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemberRequest.Size(m)
}
func (m *GetMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberRequest proto.InternalMessageInfo

func (m *GetMemberRequest) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

type LoginResponse struct {
	Result               *common.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Token                string         `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResult() *common.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenRequest struct {
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRequest) Reset()         { *m = TokenRequest{} }
func (m *TokenRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()    {}
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1086526666e68bf, []int{5}
}

func (m *TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRequest.Unmarshal(m, b)
}
func (m *TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRequest.Marshal(b, m, deterministic)
}
func (m *TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRequest.Merge(m, src)
}
func (m *TokenRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRequest.Size(m)
}
func (m *TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRequest proto.InternalMessageInfo

func (m *TokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Member)(nil), "chatroom.Member")
	proto.RegisterMapType((map[string]string)(nil), "chatroom.Member.DataEntry")
	proto.RegisterType((*MemberResponse)(nil), "chatroom.MemberResponse")
	proto.RegisterType((*SetDataRequest)(nil), "chatroom.SetDataRequest")
	proto.RegisterMapType((map[string]string)(nil), "chatroom.SetDataRequest.DataEntry")
	proto.RegisterType((*GetMemberRequest)(nil), "chatroom.GetMemberRequest")
	proto.RegisterType((*LoginResponse)(nil), "chatroom.LoginResponse")
	proto.RegisterType((*TokenRequest)(nil), "chatroom.TokenRequest")
}

func init() { proto.RegisterFile("chatroom/member.proto", fileDescriptor_b1086526666e68bf) }

var fileDescriptor_b1086526666e68bf = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x49, 0xd4, 0x5c, 0x3d, 0x5e, 0x25, 0x9c, 0xfb, 0xa7, 0x21, 0x8b, 0x22, 0xa1, 0x14,
	0x57, 0x11, 0x52, 0xfa, 0x87, 0xd2, 0x4d, 0xa1, 0xa5, 0x1b, 0xed, 0x22, 0x96, 0xee, 0x47, 0x3d,
	0x58, 0xd1, 0x64, 0x6c, 0x32, 0x11, 0x7c, 0x80, 0x3e, 0x43, 0x1f, 0xa3, 0xaf, 0x58, 0x32, 0x93,
	0x98, 0xa8, 0x75, 0x51, 0xe8, 0xca, 0x99, 0x6f, 0xce, 0xf9, 0xfc, 0x9d, 0x6f, 0x26, 0xf0, 0x6f,
	0xfc, 0xc2, 0x44, 0xc4, 0x79, 0xd0, 0x0b, 0x28, 0x18, 0x51, 0xe4, 0x2e, 0x23, 0x2e, 0x38, 0xd6,
	0x73, 0xd9, 0xfe, 0x33, 0xe6, 0x41, 0xc0, 0xc3, 0x9e, 0xfa, 0x51, 0xc7, 0xce, 0x9b, 0x06, 0xc6,
	0x40, 0xd6, 0x23, 0x42, 0x35, 0x64, 0x01, 0x59, 0x5a, 0x47, 0xeb, 0x36, 0x7c, 0xb9, 0x46, 0x17,
	0xaa, 0x13, 0x26, 0x98, 0xa5, 0x77, 0x2a, 0xdd, 0xa6, 0x67, 0xbb, 0xb9, 0x99, 0xab, 0x7a, 0xdc,
	0x3b, 0x26, 0xd8, 0x7d, 0x28, 0xa2, 0xb5, 0x2f, 0xeb, 0xec, 0x4b, 0x68, 0x6c, 0x24, 0x34, 0xa1,
	0x32, 0xa7, 0x75, 0xe6, 0x97, 0x2e, 0xf1, 0x2f, 0xd4, 0x56, 0x6c, 0x91, 0x90, 0xa5, 0x4b, 0x4d,
	0x6d, 0xae, 0xf5, 0x2b, 0xcd, 0x19, 0x41, 0x5b, 0x59, 0xfa, 0x14, 0x2f, 0x79, 0x18, 0x13, 0x9e,
	0x82, 0x11, 0x51, 0x9c, 0x2c, 0x84, 0x34, 0x68, 0x7a, 0x6d, 0x37, 0x03, 0xf7, 0xa5, 0xea, 0x67,
	0xa7, 0xd8, 0x05, 0x43, 0x0d, 0x2c, 0x4d, 0x9b, 0x9e, 0xb9, 0x0b, 0xe9, 0x67, 0xe7, 0xce, 0xbb,
	0x06, 0xed, 0x21, 0x89, 0x14, 0xd0, 0xa7, 0xd7, 0x84, 0x62, 0x91, 0x02, 0x09, 0x3e, 0xa7, 0x30,
	0x83, 0x54, 0x1b, 0xbc, 0xd8, 0x9a, 0xda, 0x29, 0x0c, 0xb7, 0xbb, 0x7f, 0x6e, 0x7a, 0x0f, 0xcc,
	0x07, 0x12, 0x79, 0x00, 0x0a, 0xed, 0x18, 0x40, 0x71, 0x3f, 0x16, 0x97, 0x52, 0x52, 0x9c, 0x01,
	0xb4, 0xfa, 0x7c, 0x3a, 0x0b, 0xbf, 0x1d, 0xd8, 0x66, 0x66, 0xbd, 0x34, 0xb3, 0x73, 0x02, 0xbf,
	0x9f, 0xd2, 0xc5, 0x5e, 0x32, 0xe5, 0x2a, 0xef, 0x43, 0x87, 0x96, 0xc2, 0x1c, 0x52, 0xb4, 0x9a,
	0x8d, 0x09, 0x3d, 0xa8, 0x49, 0x0c, 0xdc, 0xcb, 0xdd, 0x3e, 0x2a, 0x94, 0x6d, 0x52, 0x0f, 0x8c,
	0x3e, 0x9f, 0xf2, 0x44, 0xe0, 0xff, 0xa2, 0xa4, 0xfc, 0xef, 0xb6, 0x59, 0x62, 0x57, 0x3d, 0xe7,
	0xf0, 0x2b, 0x4b, 0x1f, 0xad, 0x43, 0x17, 0xf2, 0x45, 0xdb, 0x2d, 0x34, 0x36, 0xc9, 0x62, 0xe9,
	0xfd, 0xee, 0xc6, 0x6d, 0x5b, 0x7b, 0xcf, 0x26, 0xb7, 0xb8, 0x81, 0xfa, 0x33, 0x5b, 0xcc, 0x26,
	0x4c, 0xd0, 0x41, 0xde, 0x83, 0xdd, 0x23, 0x43, 0x7e, 0x67, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xec, 0xee, 0x97, 0x70, 0x9f, 0x03, 0x00, 0x00,
}
