// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replay.proto

package events

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

// Replay destination type
type Replay_Type int32

const (
	Replay_HTTP Replay_Type = 0
)

var Replay_Type_name = map[int32]string{
	0: "HTTP",
}

var Replay_Type_value = map[string]int32{
	"HTTP": 0,
}

func (x Replay_Type) String() string {
	return proto.EnumName(Replay_Type_name, int32(x))
}

func (Replay_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eed9461330ccfc03, []int{0, 0}
}

type Replay struct {
	// Replay id in DB
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Filters determine which events qualify for replay (think WHERE clause in SQL)
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	// How the replayer will replay the data
	Type Replay_Type `protobuf:"varint,3,opt,name=type,proto3,enum=events.Replay_Type" json:"type,omitempty"`
	// Where the replayer will "write" the event to
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	// Token that will be sent to destination URL along with event
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Replay) Reset()         { *m = Replay{} }
func (m *Replay) String() string { return proto.CompactTextString(m) }
func (*Replay) ProtoMessage()    {}
func (*Replay) Descriptor() ([]byte, []int) {
	return fileDescriptor_eed9461330ccfc03, []int{0}
}

func (m *Replay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Replay.Unmarshal(m, b)
}
func (m *Replay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Replay.Marshal(b, m, deterministic)
}
func (m *Replay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Replay.Merge(m, src)
}
func (m *Replay) XXX_Size() int {
	return xxx_messageInfo_Replay.Size(m)
}
func (m *Replay) XXX_DiscardUnknown() {
	xxx_messageInfo_Replay.DiscardUnknown(m)
}

var xxx_messageInfo_Replay proto.InternalMessageInfo

func (m *Replay) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Replay) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Replay) GetType() Replay_Type {
	if m != nil {
		return m.Type
	}
	return Replay_HTTP
}

func (m *Replay) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Replay) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.Replay_Type", Replay_Type_name, Replay_Type_value)
	proto.RegisterType((*Replay)(nil), "events.Replay")
}

func init() { proto.RegisterFile("replay.proto", fileDescriptor_eed9461330ccfc03) }

var fileDescriptor_eed9461330ccfc03 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xc9, 0x9f, 0x06, 0x38, 0xaa, 0x28, 0x32, 0x0c, 0x16, 0x53, 0xd4, 0x85, 0x0c, 0xc8,
	0x96, 0xca, 0x1b, 0x30, 0x20, 0x46, 0x14, 0x65, 0x62, 0x8b, 0x9d, 0xa3, 0xb1, 0x70, 0x6b, 0xcb,
	0x76, 0x90, 0xf2, 0x40, 0xbc, 0x67, 0x55, 0x5b, 0xdd, 0xec, 0xdf, 0xfd, 0xee, 0xfb, 0x74, 0xb0,
	0x75, 0x68, 0xf5, 0xb8, 0x32, 0xeb, 0x4c, 0x30, 0xa4, 0xc2, 0x3f, 0x3c, 0x05, 0xff, 0xbc, 0xfd,
	0x51, 0x3a, 0xa0, 0x4b, 0x74, 0xf7, 0x9f, 0x41, 0xd5, 0x47, 0x8d, 0xd4, 0x90, 0xab, 0x89, 0x66,
	0x6d, 0xd6, 0xdd, 0xf7, 0xb9, 0x9a, 0x48, 0x07, 0xb7, 0x49, 0xf5, 0x34, 0x6f, 0x8b, 0xee, 0x61,
	0x5f, 0xb3, 0x14, 0xc1, 0x3e, 0x22, 0xee, 0xaf, 0x63, 0xf2, 0x02, 0x65, 0x58, 0x2d, 0xd2, 0xa2,
	0xcd, 0xba, 0x7a, 0xff, 0x78, 0xd5, 0x52, 0x2e, 0x1b, 0x56, 0x8b, 0x7d, 0x14, 0x48, 0x03, 0xc5,
	0xe2, 0x34, 0x2d, 0x63, 0xc7, 0xe5, 0x49, 0x9e, 0x60, 0x13, 0xcc, 0x2f, 0x9e, 0xe8, 0x26, 0xb2,
	0xf4, 0xd9, 0x35, 0x50, 0x5e, 0xb6, 0xc8, 0x1d, 0x94, 0x9f, 0xc3, 0xf0, 0xd5, 0xdc, 0xbc, 0xb3,
	0xef, 0xd7, 0x83, 0x0a, 0xf3, 0x22, 0x98, 0x34, 0x47, 0x2e, 0xc6, 0x20, 0x67, 0x69, 0x9c, 0xe5,
	0x5e, 0xce, 0x78, 0x1c, 0x3d, 0x17, 0x8b, 0xd2, 0x13, 0x3f, 0x18, 0x9e, 0xba, 0x45, 0x15, 0xcf,
	0x7b, 0x3b, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xd3, 0x5e, 0x06, 0x04, 0x01, 0x00, 0x00,
}
