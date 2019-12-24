// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package article

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

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ViewNum              int64    `protobuf:"varint,5,opt,name=view_num,json=viewNum,proto3" json:"view_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Response) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Response) GetViewNum() int64 {
	if m != nil {
		return m.ViewNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x86, 0x05, 0x84, 0xc2, 0x25, 0x1a, 0xd3, 0x38, 0x14, 0x17, 0x09, 0x13, 0x71, 0x60, 0xd0,
	0x57, 0x70, 0x76, 0xe8, 0x0b, 0x18, 0x84, 0x1b, 0x9a, 0x40, 0x8b, 0x70, 0x95, 0xd7, 0x37, 0x54,
	0x58, 0xdc, 0xee, 0xfb, 0x6e, 0xf9, 0x7e, 0xd8, 0x55, 0x03, 0xa9, 0xba, 0xc5, 0xb2, 0x1f, 0x0c,
	0x99, 0x3c, 0x05, 0x26, 0xf1, 0x6d, 0x71, 0x24, 0xbe, 0x07, 0x5f, 0x35, 0xc2, 0xcb, 0xbc, 0x22,
	0x91, 0xbe, 0x6a, 0xf2, 0x09, 0x62, 0x89, 0x63, 0x6f, 0xf4, 0x88, 0xff, 0x3f, 0x7e, 0x80, 0xc0,
	0xaa, 0x46, 0xf8, 0x4e, 0xcc, 0x27, 0x3f, 0x42, 0x48, 0x8a, 0x5a, 0x14, 0x81, 0x73, 0x3f, 0xe0,
	0x02, 0x58, 0x6d, 0x34, 0xa1, 0x26, 0xb1, 0x75, 0x7e, 0x45, 0x9e, 0x42, 0xfc, 0x51, 0x38, 0x3d,
	0xb5, 0xed, 0x44, 0x98, 0x79, 0x45, 0x20, 0xd9, 0xcc, 0x0f, 0xdb, 0x5d, 0x2f, 0xc0, 0x96, 0x48,
	0x7e, 0x86, 0xe8, 0x8e, 0x54, 0xa9, 0x96, 0xc7, 0xe5, 0xd2, 0x79, 0x4a, 0xca, 0x35, 0x2b, 0xdf,
	0xbc, 0x22, 0x37, 0xe3, 0xf6, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x01, 0x9c, 0xaf, 0xd3, 0xd7, 0x00,
	0x00, 0x00,
}
