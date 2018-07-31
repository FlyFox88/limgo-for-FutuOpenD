// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_GetFunds.proto

package Trd_GetFunds

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	Header               *TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetHeader() *TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type S2C struct {
	Header               *TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Funds                *Funds     `protobuf:"bytes,2,opt,name=funds" json:"funds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetHeader() *TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetFunds() *Funds {
	if m != nil {
		return m.Funds
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Trd_GetFunds.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_GetFunds.S2C")
	proto.RegisterType((*Request)(nil), "Trd_GetFunds.Request")
	proto.RegisterType((*Response)(nil), "Trd_GetFunds.Response")
}

func init() { proto.RegisterFile("Trd_GetFunds.proto", fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40) }

var fileDescriptor_Trd_GetFunds_bd3693e5a7f82f40 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xa6, 0x69, 0xe0, 0xda, 0x01, 0x2c, 0x81, 0xac, 0x0e, 0x28, 0x0a, 0x03, 0x59,
	0x1a, 0x55, 0x56, 0x27, 0x56, 0x4b, 0xc0, 0xc2, 0x72, 0xc9, 0x8a, 0x10, 0xaa, 0x0f, 0x58, 0x1a,
	0x07, 0xdb, 0x1d, 0x18, 0xf9, 0xe7, 0xc8, 0x4e, 0x1a, 0x79, 0xed, 0x78, 0xef, 0x7b, 0xcf, 0xfa,
	0x0c, 0xac, 0x35, 0xea, 0xfd, 0x99, 0xdc, 0xd3, 0xb1, 0x53, 0xb6, 0xee, 0x8d, 0x76, 0x9a, 0xad,
	0xe2, 0x6c, 0xbd, 0x92, 0xfa, 0x70, 0xd0, 0xdd, 0xc0, 0xd6, 0x57, 0x9e, 0xc5, 0x49, 0xb9, 0x83,
	0x54, 0x8a, 0x86, 0x6d, 0x60, 0xf1, 0x4d, 0x1f, 0x8a, 0x0c, 0x4f, 0x8a, 0x59, 0xb5, 0x14, 0x37,
	0x75, 0xd4, 0x6c, 0x8d, 0x7a, 0x09, 0x10, 0xc7, 0x52, 0xf9, 0x06, 0x69, 0x23, 0xe4, 0x99, 0x2b,
	0xf6, 0x00, 0xd9, 0xa7, 0x97, 0xe2, 0xb3, 0x22, 0xa9, 0x96, 0xe2, 0x3a, 0x6e, 0x07, 0x5b, 0x1c,
	0x78, 0x59, 0x43, 0x8e, 0xf4, 0x73, 0x24, 0xeb, 0xd8, 0x3d, 0xa4, 0x7b, 0x61, 0xc7, 0xf7, 0x87,
	0xc5, 0xf4, 0x5f, 0x29, 0x1a, 0xf4, 0xb4, 0xfc, 0x4b, 0xe0, 0x02, 0xc9, 0xf6, 0xba, 0xb3, 0xc4,
	0xee, 0x20, 0x37, 0xe4, 0xda, 0xdf, 0x9e, 0xc2, 0x2a, 0x7b, 0x9c, 0x6f, 0x76, 0xdb, 0x2d, 0x9e,
	0x42, 0x76, 0x0b, 0x0b, 0x43, 0xee, 0xd5, 0x7e, 0x05, 0x8d, 0x4b, 0x1c, 0x2f, 0xc6, 0x21, 0x27,
	0x63, 0xa4, 0x56, 0xc4, 0xd3, 0x22, 0xa9, 0x32, 0x3c, 0x9d, 0xde, 0xc1, 0x8a, 0x3d, 0x9f, 0x47,
	0xd6, 0x93, 0x43, 0x23, 0x24, 0x7a, 0xfa, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x71, 0x10, 0xa2, 0x0c,
	0x8b, 0x01, 0x00, 0x00,
}
