// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetHistoryKL.proto

package Qot_GetHistoryKL

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
	RehabType            *int32    `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`
	KlType               *int32    `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`
	Security             *Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`
	BeginTime            *string   `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string   `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`
	MaxAckKLNum          *int32    `protobuf:"varint,6,opt,name=maxAckKLNum" json:"maxAckKLNum,omitempty"`
	NeedKLFieldsFlag     *int64    `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584, []int{0}
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

func (m *C2S) GetRehabType() int32 {
	if m != nil && m.RehabType != nil {
		return *m.RehabType
	}
	return 0
}

func (m *C2S) GetKlType() int32 {
	if m != nil && m.KlType != nil {
		return *m.KlType
	}
	return 0
}

func (m *C2S) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *C2S) GetMaxAckKLNum() int32 {
	if m != nil && m.MaxAckKLNum != nil {
		return *m.MaxAckKLNum
	}
	return 0
}

func (m *C2S) GetNeedKLFieldsFlag() int64 {
	if m != nil && m.NeedKLFieldsFlag != nil {
		return *m.NeedKLFieldsFlag
	}
	return 0
}

type S2C struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	KlList               []*KLine  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`
	NextKLTime           *string   `protobuf:"bytes,3,opt,name=nextKLTime" json:"nextKLTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584, []int{1}
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

func (m *S2C) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetKlList() []*KLine {
	if m != nil {
		return m.KlList
	}
	return nil
}

func (m *S2C) GetNextKLTime() string {
	if m != nil && m.NextKLTime != nil {
		return *m.NextKLTime
	}
	return ""
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
	return fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584, []int{2}
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
	return fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetHistoryKL.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetHistoryKL.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetHistoryKL.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetHistoryKL.Response")
}

func init() {
	proto.RegisterFile("Qot_GetHistoryKL.proto", fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584)
}

var fileDescriptor_Qot_GetHistoryKL_ed0ba2747a479584 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x99, 0xc4, 0x18, 0x1d, 0x7b, 0xb0, 0x43, 0x2b, 0x83, 0x14, 0x19, 0x72, 0x69, 0x5a,
	0xa8, 0xc8, 0xd0, 0x53, 0x6f, 0x25, 0x60, 0x0b, 0x49, 0x0b, 0x9d, 0x78, 0x2f, 0x9a, 0x3c, 0xdc,
	0xa0, 0xc9, 0xb8, 0x33, 0x23, 0xe8, 0x75, 0x61, 0x3f, 0xf1, 0x7e, 0x81, 0x25, 0x63, 0x5c, 0xb3,
	0xeb, 0x1e, 0xf6, 0xf8, 0x7e, 0xbf, 0x3f, 0xbc, 0xff, 0xcc, 0xc3, 0xa3, 0x7f, 0xd2, 0xfc, 0xff,
	0x05, 0xe6, 0x77, 0xa1, 0x8d, 0x54, 0xc7, 0x38, 0x99, 0xee, 0x94, 0x34, 0x92, 0x0c, 0x5f, 0xf2,
	0xf1, 0xbb, 0x48, 0x96, 0xa5, 0xac, 0x4e, 0x7e, 0x6c, 0x7d, 0x9b, 0x04, 0x0f, 0x08, 0xbb, 0x11,
	0x4f, 0xc9, 0x27, 0xdc, 0x57, 0x70, 0xb3, 0x5c, 0x2d, 0x8e, 0x3b, 0xa0, 0x88, 0x39, 0xa1, 0x27,
	0x2e, 0x80, 0x8c, 0x70, 0x77, 0xb3, 0xb5, 0xca, 0xb1, 0xaa, 0x99, 0xc8, 0x0c, 0xf7, 0x34, 0x64,
	0x7b, 0x55, 0x98, 0x23, 0x75, 0x99, 0x13, 0x0e, 0xf8, 0x87, 0x69, 0x6b, 0x45, 0xda, 0x38, 0xf1,
	0x94, 0xaa, 0xf7, 0xac, 0x60, 0x5d, 0x54, 0x8b, 0xa2, 0x04, 0xda, 0x61, 0x4e, 0xd8, 0x17, 0x17,
	0x40, 0x28, 0xf6, 0xa1, 0xca, 0xad, 0xf3, 0xac, 0x3b, 0x8f, 0x84, 0xe1, 0x41, 0xb9, 0x3c, 0xfc,
	0xcc, 0x36, 0x71, 0xf2, 0x77, 0x5f, 0xd2, 0x2e, 0x43, 0xa1, 0x27, 0xda, 0x88, 0x7c, 0xc5, 0xc3,
	0x0a, 0x20, 0x8f, 0x93, 0x79, 0x01, 0xdb, 0x5c, 0xcf, 0xb7, 0xcb, 0x35, 0xf5, 0x19, 0x0a, 0x5d,
	0x71, 0xc5, 0x83, 0x3b, 0x84, 0xdd, 0x94, 0x47, 0xcf, 0xfa, 0xa3, 0x37, 0xf5, 0xff, 0x52, 0xff,
	0x44, 0x52, 0x68, 0x43, 0x1d, 0xe6, 0x86, 0x03, 0xfe, 0xbe, 0x9d, 0x8f, 0x93, 0xa2, 0x02, 0xd1,
	0x04, 0xc8, 0x04, 0xe3, 0x0a, 0x0e, 0x26, 0x4e, 0xec, 0x7b, 0x5c, 0x86, 0xc2, 0xbe, 0x68, 0x91,
	0x80, 0x63, 0x5f, 0xc0, 0xed, 0x1e, 0xb4, 0x21, 0x9f, 0xb1, 0x9b, 0x71, 0xdd, 0x54, 0xf8, 0x38,
	0xbd, 0xba, 0x6e, 0xc4, 0x53, 0x51, 0x27, 0x82, 0x7b, 0x84, 0x7b, 0x02, 0xf4, 0x4e, 0x56, 0x1a,
	0xc8, 0x04, 0xfb, 0x0a, 0xcc, 0xe5, 0x62, 0x3f, 0x3a, 0xdf, 0xbe, 0xcf, 0x66, 0xe2, 0x0c, 0xeb,
	0xab, 0x29, 0x30, 0x7f, 0xf4, 0x9a, 0x3a, 0x76, 0x79, 0x33, 0xd9, 0x5f, 0x56, 0x2a, 0x92, 0xf9,
	0xa9, 0x95, 0x27, 0xce, 0x63, 0xdd, 0x43, 0xf3, 0x8c, 0x76, 0x18, 0x7a, 0xbd, 0x47, 0xca, 0x23,
	0x51, 0x27, 0x1e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x28, 0x24, 0xed, 0x81, 0x02, 0x00, 0x00,
}
