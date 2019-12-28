// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pbevents

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

type Message_Type int32

const (
	Message_START_COLLECTION Message_Type = 0
	Message_STOP_COLLECTION  Message_Type = 1
)

var Message_Type_name = map[int32]string{
	0: "START_COLLECTION",
	1: "STOP_COLLECTION",
}

var Message_Type_value = map[string]int32{
	"START_COLLECTION": 0,
	"STOP_COLLECTION":  1,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type Message struct {
	// What kind of a message is this?
	Type Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=pbevents.Message_Type" json:"type,omitempty"`
	// Indicates which client this event is tied to
	Client *Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	// Info used by collectors to manage event collection
	Collect *Collect `protobuf:"bytes,3,opt,name=collect,proto3" json:"collect,omitempty"`
	// Contains event debug info
	Info                 *Info    `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_START_COLLECTION
}

func (m *Message) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Message) GetCollect() *Collect {
	if m != nil {
		return m.Collect
	}
	return nil
}

func (m *Message) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterEnum("pbevents.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterType((*Message)(nil), "pbevents.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x48, 0x4a, 0x2d, 0x4b,
	0xcd, 0x2b, 0x29, 0x96, 0xe2, 0x49, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x81, 0x88, 0x4b, 0xf1, 0x26,
	0xe7, 0xe7, 0xe4, 0xa4, 0x26, 0xc3, 0xb8, 0x5c, 0x99, 0x79, 0x69, 0xf9, 0x10, 0xb6, 0xd2, 0x63,
	0x46, 0x2e, 0x76, 0x5f, 0x88, 0x21, 0x42, 0x5a, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x62, 0x7a, 0x30, 0xd3, 0xf4, 0xa0, 0x0a, 0xf4, 0x42, 0x2a, 0x0b,
	0x52, 0x83, 0xc0, 0x6a, 0x84, 0x34, 0xb8, 0xd8, 0x20, 0x56, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x1b, 0x09, 0x20, 0x54, 0x3b, 0x83, 0xc5, 0x83, 0xa0, 0xf2, 0x42, 0xda, 0x5c, 0xec, 0x50, 0xeb,
	0x25, 0x98, 0xc1, 0x4a, 0x05, 0x91, 0x94, 0x42, 0x24, 0x82, 0x60, 0x2a, 0x84, 0x94, 0xb8, 0x58,
	0x40, 0x8e, 0x93, 0x60, 0x01, 0xab, 0xe4, 0x43, 0xa8, 0xf4, 0xcc, 0x4b, 0xcb, 0x0f, 0x02, 0xcb,
	0x29, 0x19, 0x72, 0xb1, 0x80, 0x1c, 0x22, 0x24, 0xc2, 0x25, 0x10, 0x1c, 0xe2, 0x18, 0x14, 0x12,
	0xef, 0xec, 0xef, 0xe3, 0xe3, 0xea, 0x1c, 0xe2, 0xe9, 0xef, 0x27, 0xc0, 0x20, 0x24, 0xcc, 0xc5,
	0x1f, 0x1c, 0xe2, 0x1f, 0x80, 0x2c, 0xc8, 0x98, 0xc4, 0x06, 0xf6, 0xac, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xb2, 0x6e, 0xa4, 0x07, 0x30, 0x01, 0x00, 0x00,
}
