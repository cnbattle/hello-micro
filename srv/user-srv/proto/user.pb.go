// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type MiniProgramLoginRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MiniProgramLoginRequest) Reset()         { *m = MiniProgramLoginRequest{} }
func (m *MiniProgramLoginRequest) String() string { return proto.CompactTextString(m) }
func (*MiniProgramLoginRequest) ProtoMessage()    {}
func (*MiniProgramLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *MiniProgramLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiniProgramLoginRequest.Unmarshal(m, b)
}
func (m *MiniProgramLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiniProgramLoginRequest.Marshal(b, m, deterministic)
}
func (m *MiniProgramLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniProgramLoginRequest.Merge(m, src)
}
func (m *MiniProgramLoginRequest) XXX_Size() int {
	return xxx_messageInfo_MiniProgramLoginRequest.Size(m)
}
func (m *MiniProgramLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniProgramLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MiniProgramLoginRequest proto.InternalMessageInfo

func (m *MiniProgramLoginRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *MiniProgramLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PhoneCodeLoginRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneCodeLoginRequest) Reset()         { *m = PhoneCodeLoginRequest{} }
func (m *PhoneCodeLoginRequest) String() string { return proto.CompactTextString(m) }
func (*PhoneCodeLoginRequest) ProtoMessage()    {}
func (*PhoneCodeLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *PhoneCodeLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneCodeLoginRequest.Unmarshal(m, b)
}
func (m *PhoneCodeLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneCodeLoginRequest.Marshal(b, m, deterministic)
}
func (m *PhoneCodeLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneCodeLoginRequest.Merge(m, src)
}
func (m *PhoneCodeLoginRequest) XXX_Size() int {
	return xxx_messageInfo_PhoneCodeLoginRequest.Size(m)
}
func (m *PhoneCodeLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneCodeLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneCodeLoginRequest proto.InternalMessageInfo

func (m *PhoneCodeLoginRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PhoneCodeLoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PhoneCodeLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PasswordLoginRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordLoginRequest) Reset()         { *m = PasswordLoginRequest{} }
func (m *PasswordLoginRequest) String() string { return proto.CompactTextString(m) }
func (*PasswordLoginRequest) ProtoMessage()    {}
func (*PasswordLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *PasswordLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordLoginRequest.Unmarshal(m, b)
}
func (m *PasswordLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordLoginRequest.Marshal(b, m, deterministic)
}
func (m *PasswordLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordLoginRequest.Merge(m, src)
}
func (m *PasswordLoginRequest) XXX_Size() int {
	return xxx_messageInfo_PasswordLoginRequest.Size(m)
}
func (m *PasswordLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordLoginRequest proto.InternalMessageInfo

func (m *PasswordLoginRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PasswordLoginRequest) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *PasswordLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
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

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DefaultRequest struct {
	UID                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultRequest) Reset()         { *m = DefaultRequest{} }
func (m *DefaultRequest) String() string { return proto.CompactTextString(m) }
func (*DefaultRequest) ProtoMessage()    {}
func (*DefaultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *DefaultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultRequest.Unmarshal(m, b)
}
func (m *DefaultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultRequest.Marshal(b, m, deterministic)
}
func (m *DefaultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultRequest.Merge(m, src)
}
func (m *DefaultRequest) XXX_Size() int {
	return xxx_messageInfo_DefaultRequest.Size(m)
}
func (m *DefaultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultRequest proto.InternalMessageInfo

func (m *DefaultRequest) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

type AuthResponse struct {
	UID                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	IdentityType         string   `protobuf:"bytes,2,opt,name=identityType,proto3" json:"identityType,omitempty"`
	Identity             string   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Certificate          string   `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *AuthResponse) GetIdentityType() string {
	if m != nil {
		return m.IdentityType
	}
	return ""
}

func (m *AuthResponse) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *AuthResponse) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type BaseResponse struct {
	UID                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Signature            string   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *BaseResponse) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *BaseResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *BaseResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *BaseResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*MiniProgramLoginRequest)(nil), "MiniProgramLoginRequest")
	proto.RegisterType((*PhoneCodeLoginRequest)(nil), "PhoneCodeLoginRequest")
	proto.RegisterType((*PasswordLoginRequest)(nil), "PasswordLoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*DefaultRequest)(nil), "DefaultRequest")
	proto.RegisterType((*AuthResponse)(nil), "AuthResponse")
	proto.RegisterType((*BaseResponse)(nil), "BaseResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xdd, 0xec, 0x76, 0x97, 0xdd, 0xa1, 0x2d, 0x95, 0xd5, 0x96, 0x28, 0xe2, 0x50, 0x59, 0x42,
	0xe2, 0xe4, 0x03, 0xdc, 0x90, 0x40, 0x02, 0x2a, 0x71, 0x01, 0x54, 0x55, 0xdc, 0x38, 0x99, 0x64,
	0x9a, 0x9a, 0xa6, 0x76, 0xb0, 0x1d, 0x50, 0xaf, 0x1c, 0xf9, 0x0b, 0xfe, 0x14, 0x39, 0x71, 0xd2,
	0x04, 0xd2, 0x55, 0x6f, 0x7e, 0x93, 0x37, 0x6f, 0x9e, 0x66, 0x5e, 0x00, 0x0a, 0x83, 0x9a, 0xe5,
	0x5a, 0x59, 0x45, 0xdf, 0xc3, 0xe3, 0x8f, 0x42, 0x8a, 0x95, 0x56, 0xa9, 0xe6, 0xfb, 0x0f, 0x2a,
	0x15, 0x72, 0x8d, 0xdf, 0x0b, 0x34, 0x96, 0x84, 0xf0, 0x20, 0xde, 0x72, 0x29, 0x31, 0x0b, 0x83,
	0x45, 0xf0, 0xec, 0x6e, 0x5d, 0x43, 0x42, 0x60, 0x10, 0xab, 0x04, 0xc3, 0xcb, 0xb2, 0x5c, 0xbe,
	0xe9, 0x17, 0x98, 0xad, 0xb6, 0x4a, 0xe2, 0x3b, 0x95, 0xe0, 0x99, 0x32, 0x53, 0xb8, 0xce, 0x5d,
	0x8b, 0xd7, 0xa9, 0x40, 0x23, 0x7e, 0xd5, 0x12, 0xff, 0x06, 0xd3, 0x15, 0x37, 0xe6, 0xa7, 0xd2,
	0xc9, 0x99, 0xda, 0x4f, 0xe0, 0x2e, 0x73, 0xcc, 0x4f, 0x7c, 0x5f, 0xeb, 0x1f, 0x0b, 0x24, 0x82,
	0xdb, 0xdc, 0xeb, 0xf9, 0x39, 0x0d, 0xa6, 0x4f, 0x61, 0xe4, 0x67, 0x98, 0x5c, 0x49, 0x83, 0xce,
	0xa6, 0x55, 0x3b, 0x94, 0x7e, 0x44, 0x05, 0x28, 0x85, 0xf1, 0x12, 0x37, 0xbc, 0xc8, 0x6c, 0x6d,
	0x66, 0x02, 0x57, 0x85, 0x48, 0x3c, 0xcb, 0x3d, 0xe9, 0xaf, 0x00, 0x86, 0x6f, 0x0a, 0xbb, 0x6d,
	0xa4, 0xfe, 0xa3, 0x10, 0x0a, 0x43, 0x91, 0xa0, 0xb4, 0xc2, 0x1e, 0x3e, 0x1f, 0xf2, 0xda, 0x6a,
	0xa7, 0xe6, 0xdc, 0xd6, 0xb8, 0x76, 0x5b, 0x63, 0xb2, 0x80, 0x87, 0x31, 0x6a, 0x2b, 0x36, 0x22,
	0xe6, 0x16, 0xc3, 0x41, 0xf9, 0xb9, 0x5d, 0xa2, 0xbf, 0x03, 0x18, 0xbe, 0xe5, 0x06, 0xef, 0x31,
	0x11, 0xc1, 0xad, 0x14, 0xf1, 0x4e, 0x1e, 0x77, 0xd5, 0x60, 0x32, 0x87, 0x1b, 0xfe, 0x83, 0x5b,
	0xae, 0xfd, 0x68, 0x8f, 0x5c, 0x3d, 0x45, 0x99, 0xa0, 0xf6, 0x33, 0x3d, 0x72, 0x8b, 0x37, 0x22,
	0x95, 0xdc, 0x16, 0x1a, 0xc3, 0xeb, 0x6a, 0xf1, 0x4d, 0xe1, 0xf9, 0x9f, 0x4b, 0x18, 0xb8, 0xf4,
	0x91, 0xd7, 0x30, 0xf9, 0x37, 0x77, 0x24, 0x64, 0x27, 0xa2, 0x18, 0x8d, 0x59, 0xe7, 0x24, 0xf4,
	0x82, 0xbc, 0x84, 0x71, 0x37, 0x6e, 0x64, 0xce, 0x7a, 0xf3, 0xd7, 0xd3, 0xfb, 0x0a, 0x48, 0x49,
	0xed, 0x44, 0x8a, 0xcc, 0x58, 0x5f, 0xc4, 0x7a, 0xda, 0x19, 0x80, 0x3b, 0xea, 0x12, 0x2d, 0x17,
	0x19, 0x79, 0xc4, 0xba, 0x31, 0x88, 0x46, 0xac, 0x7d, 0xf2, 0x8a, 0xef, 0xf6, 0x7f, 0x9a, 0xdf,
	0xbe, 0x0e, 0xbd, 0xf8, 0x7a, 0x53, 0xfe, 0x99, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x64,
	0xe5, 0x51, 0x53, 0xa7, 0x03, 0x00, 0x00,
}
