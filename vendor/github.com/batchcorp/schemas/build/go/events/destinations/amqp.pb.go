// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/destinations/amqp.proto

package destinations

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

type AMQP struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	RoutingKey           string   `protobuf:"bytes,3,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AMQP) Reset()         { *m = AMQP{} }
func (m *AMQP) String() string { return proto.CompactTextString(m) }
func (*AMQP) ProtoMessage()    {}
func (*AMQP) Descriptor() ([]byte, []int) {
	return fileDescriptor_b13c1ce86a1f6ce7, []int{0}
}

func (m *AMQP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AMQP.Unmarshal(m, b)
}
func (m *AMQP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AMQP.Marshal(b, m, deterministic)
}
func (m *AMQP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AMQP.Merge(m, src)
}
func (m *AMQP) XXX_Size() int {
	return xxx_messageInfo_AMQP.Size(m)
}
func (m *AMQP) XXX_DiscardUnknown() {
	xxx_messageInfo_AMQP.DiscardUnknown(m)
}

var xxx_messageInfo_AMQP proto.InternalMessageInfo

func (m *AMQP) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AMQP) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *AMQP) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func init() {
	proto.RegisterType((*AMQP)(nil), "destinations.AMQP")
}

func init() { proto.RegisterFile("events/destinations/amqp.proto", fileDescriptor_b13c1ce86a1f6ce7) }

var fileDescriptor_b13c1ce86a1f6ce7 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x4f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6,
	0x4f, 0xcc, 0x2d, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x96, 0x50, 0x0a,
	0xe5, 0x62, 0x71, 0xf4, 0x0d, 0x0c, 0x10, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xa4, 0xb8, 0x38, 0x52, 0x2b, 0x92, 0x33, 0x12, 0xf3,
	0xd2, 0x53, 0x25, 0x98, 0xc0, 0xc2, 0x70, 0xbe, 0x90, 0x3c, 0x17, 0x77, 0x51, 0x7e, 0x69, 0x49,
	0x66, 0x5e, 0x7a, 0x7c, 0x76, 0x6a, 0xa5, 0x04, 0x33, 0x58, 0x9a, 0x0b, 0x2a, 0xe4, 0x9d, 0x5a,
	0xe9, 0x64, 0x1d, 0x65, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x94, 0x58, 0x92, 0x9c, 0x91, 0x9c, 0x5f, 0x54, 0xa0, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0x58,
	0xac, 0x9f, 0x54, 0x9a, 0x99, 0x93, 0xa2, 0x9f, 0x9e, 0xaf, 0x8f, 0xc5, 0xb1, 0x49, 0x6c, 0x60,
	0x87, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0x4a, 0x08, 0xa3, 0xca, 0x00, 0x00, 0x00,
}
