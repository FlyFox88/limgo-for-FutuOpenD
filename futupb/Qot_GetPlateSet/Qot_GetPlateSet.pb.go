// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetPlateSet.proto

package Qot_GetPlateSet

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
	Market               *int32   `protobuf:"varint,1,req,name=market" json:"market,omitempty"`
	PlateSetType         *int32   `protobuf:"varint,2,req,name=plateSetType" json:"plateSetType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_866019980cce8e3a, []int{0}
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

func (m *C2S) GetMarket() int32 {
	if m != nil && m.Market != nil {
		return *m.Market
	}
	return 0
}

func (m *C2S) GetPlateSetType() int32 {
	if m != nil && m.PlateSetType != nil {
		return *m.PlateSetType
	}
	return 0
}

type PlateInfo struct {
	Plate                *Security `protobuf:"bytes,1,req,name=plate" json:"plate,omitempty"`
	Name                 *string   `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlateInfo) Reset()         { *m = PlateInfo{} }
func (m *PlateInfo) String() string { return proto.CompactTextString(m) }
func (*PlateInfo) ProtoMessage()    {}
func (*PlateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_866019980cce8e3a, []int{1}
}
func (m *PlateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlateInfo.Unmarshal(m, b)
}
func (m *PlateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlateInfo.Marshal(b, m, deterministic)
}
func (dst *PlateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlateInfo.Merge(dst, src)
}
func (m *PlateInfo) XXX_Size() int {
	return xxx_messageInfo_PlateInfo.Size(m)
}
func (m *PlateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlateInfo proto.InternalMessageInfo

func (m *PlateInfo) GetPlate() *Security {
	if m != nil {
		return m.Plate
	}
	return nil
}

func (m *PlateInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type S2C struct {
	PlateInfoList        []*PlateInfo `protobuf:"bytes,1,rep,name=plateInfoList" json:"plateInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_866019980cce8e3a, []int{2}
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

func (m *S2C) GetPlateInfoList() []*PlateInfo {
	if m != nil {
		return m.PlateInfoList
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
	return fileDescriptor_Qot_GetPlateSet_866019980cce8e3a, []int{3}
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
	return fileDescriptor_Qot_GetPlateSet_866019980cce8e3a, []int{4}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetPlateSet.C2S")
	proto.RegisterType((*PlateInfo)(nil), "Qot_GetPlateSet.PlateInfo")
	proto.RegisterType((*S2C)(nil), "Qot_GetPlateSet.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetPlateSet.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetPlateSet.Response")
}

func init() {
	proto.RegisterFile("Qot_GetPlateSet.proto", fileDescriptor_Qot_GetPlateSet_866019980cce8e3a)
}

var fileDescriptor_Qot_GetPlateSet_866019980cce8e3a = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4b, 0xfb, 0x30,
	0x18, 0xa5, 0xeb, 0xf6, 0xdb, 0x6f, 0xdf, 0x26, 0x4a, 0x50, 0x09, 0x3b, 0x48, 0xc9, 0x41, 0x8a,
	0xe0, 0x98, 0xc1, 0x93, 0x27, 0x25, 0x87, 0x21, 0x2a, 0x68, 0xe2, 0x5d, 0x4a, 0xfd, 0x94, 0xa1,
	0x6d, 0x6a, 0x92, 0x1d, 0x76, 0xf7, 0x0f, 0x97, 0xa4, 0xad, 0xb4, 0xc3, 0x5b, 0xbe, 0xf7, 0x5e,
	0xde, 0x7b, 0xdf, 0x07, 0x47, 0x4f, 0xda, 0xbd, 0xac, 0xd0, 0x3d, 0x7e, 0x66, 0x0e, 0x15, 0xba,
	0x45, 0x65, 0xb4, 0xd3, 0x64, 0x7f, 0x07, 0x9e, 0xcf, 0x84, 0x2e, 0x0a, 0x5d, 0xd6, 0xf4, 0xfc,
	0xc0, 0xd3, 0x5d, 0x84, 0xdd, 0x40, 0x2c, 0xb8, 0x22, 0xc7, 0xf0, 0xaf, 0xc8, 0xcc, 0x07, 0x3a,
	0x1a, 0x25, 0x83, 0x74, 0x24, 0x9b, 0x89, 0x30, 0x98, 0x55, 0x8d, 0xd5, 0xf3, 0xb6, 0x42, 0x3a,
	0x08, 0x6c, 0x0f, 0x63, 0x77, 0x30, 0x09, 0x71, 0xb7, 0xe5, 0x9b, 0x26, 0x67, 0x30, 0x0a, 0x64,
	0xf0, 0x99, 0xf2, 0xc3, 0x45, 0x27, 0x51, 0x61, 0xbe, 0x31, 0x6b, 0xb7, 0x95, 0xb5, 0x84, 0x10,
	0x18, 0x96, 0x59, 0x51, 0x9b, 0x4e, 0x64, 0x78, 0xb3, 0x15, 0xc4, 0x8a, 0x0b, 0x72, 0x0d, 0x7b,
	0x55, 0xeb, 0x79, 0xbf, 0xb6, 0xbe, 0x56, 0x9c, 0x4e, 0xf9, 0x7c, 0xb1, 0xbb, 0xf6, 0x6f, 0xb2,
	0xec, 0x7f, 0x60, 0x17, 0x30, 0x96, 0xf8, 0xb5, 0x41, 0xeb, 0xc8, 0x29, 0xc4, 0x39, 0xb7, 0xbd,
	0x46, 0x5d, 0x0b, 0xc1, 0x95, 0xf4, 0x02, 0xf6, 0x1d, 0xc1, 0x7f, 0x89, 0xb6, 0xd2, 0xa5, 0x45,
	0x72, 0x02, 0x63, 0xd3, 0x2c, 0x1d, 0x4e, 0x72, 0x35, 0x3c, 0xbf, 0x5c, 0x2e, 0x65, 0x0b, 0xfa,
	0x8b, 0x19, 0x74, 0x0f, 0xf6, 0x9d, 0x0e, 0x92, 0x28, 0x9d, 0xc8, 0x66, 0x22, 0x14, 0xc6, 0x68,
	0x8c, 0xd0, 0xaf, 0x48, 0xe3, 0x24, 0x4a, 0x47, 0xb2, 0x1d, 0x7d, 0x0d, 0xcb, 0x73, 0x3a, 0x4c,
	0xa2, 0x3f, 0x6b, 0x28, 0x2e, 0xa4, 0x17, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x49, 0x22, 0xda,
	0x2f, 0xdb, 0x01, 0x00, 0x00,
}
