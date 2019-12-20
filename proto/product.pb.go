// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

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

type DetailRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

type Detailesponse struct {
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Detailesponse) Reset()         { *m = Detailesponse{} }
func (m *Detailesponse) String() string { return proto.CompactTextString(m) }
func (*Detailesponse) ProtoMessage()    {}
func (*Detailesponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *Detailesponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Detailesponse.Unmarshal(m, b)
}
func (m *Detailesponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Detailesponse.Marshal(b, m, deterministic)
}
func (m *Detailesponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Detailesponse.Merge(m, src)
}
func (m *Detailesponse) XXX_Size() int {
	return xxx_messageInfo_Detailesponse.Size(m)
}
func (m *Detailesponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Detailesponse.DiscardUnknown(m)
}

var xxx_messageInfo_Detailesponse proto.InternalMessageInfo

func (m *Detailesponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*DetailRequest)(nil), "DetailRequest")
	proto.RegisterType((*Detailesponse)(nil), "Detailesponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xe2, 0xe7, 0xe2, 0x75, 0x49, 0x2d,
	0x49, 0xcc, 0xcc, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x52, 0x85, 0x09, 0xa4, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x97, 0x24, 0xe6, 0x48, 0x30,
	0x29, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x46, 0xa6, 0x5c, 0xec, 0x01, 0x10, 0x83, 0x84, 0xb4,
	0xb8, 0xd8, 0x20, 0x3a, 0x84, 0xf8, 0xf4, 0x50, 0xcc, 0x92, 0x82, 0xf1, 0xa1, 0x46, 0x29, 0x31,
	0x24, 0xb1, 0x81, 0x6d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x11, 0xd1, 0x85, 0x86,
	0x00, 0x00, 0x00,
}
