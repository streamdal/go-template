// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/metrics/counter.proto

package metrics

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

type Counter_Type int32

const (
	Counter_UNSET        Counter_Type = 0
	Counter_COLLECTION   Counter_Type = 1
	Counter_REPLAY       Counter_Type = 2
	Counter_STORAGE      Counter_Type = 3
	Counter_REPLAY_BYTES Counter_Type = 4
)

var Counter_Type_name = map[int32]string{
	0: "UNSET",
	1: "COLLECTION",
	2: "REPLAY",
	3: "STORAGE",
	4: "REPLAY_BYTES",
}

var Counter_Type_value = map[string]int32{
	"UNSET":        0,
	"COLLECTION":   1,
	"REPLAY":       2,
	"STORAGE":      3,
	"REPLAY_BYTES": 4,
}

func (x Counter_Type) String() string {
	return proto.EnumName(Counter_Type_name, int32(x))
}

func (Counter_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b78758d2d8a4fc21, []int{0, 0}
}

// Used for internal metrics collection with "mlib" library and metrics service.
type Counter struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               string       `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Type                 Counter_Type `protobuf:"varint,3,opt,name=type,proto3,enum=metrics.Counter_Type" json:"type,omitempty"`
	Source               string       `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Value                int64        `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64        `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b78758d2d8a4fc21, []int{0}
}

func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Counter) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Counter) GetType() Counter_Type {
	if m != nil {
		return m.Type
	}
	return Counter_UNSET
}

func (m *Counter) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Counter) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Counter) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("metrics.Counter_Type", Counter_Type_name, Counter_Type_value)
	proto.RegisterType((*Counter)(nil), "metrics.Counter")
}

func init() { proto.RegisterFile("events/metrics/counter.proto", fileDescriptor_b78758d2d8a4fc21) }

var fileDescriptor_b78758d2d8a4fc21 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0x47, 0x9b, 0x18, 0x13, 0x9c, 0x16, 0x59, 0x86, 0xfe, 0xd9, 0x83, 0x07, 0xf1, 0x64, 0x2f,
	0x59, 0x68, 0x4b, 0xef, 0x2a, 0xa1, 0x08, 0x41, 0x4b, 0x4c, 0x0f, 0xf6, 0x22, 0xc9, 0x66, 0xd1,
	0x05, 0xd7, 0x0d, 0xd9, 0x8d, 0xe0, 0x27, 0xef, 0xb5, 0x98, 0x04, 0x4a, 0x8f, 0xbf, 0xf7, 0x86,
	0x77, 0x18, 0x18, 0x89, 0xb3, 0x38, 0x59, 0xc3, 0x94, 0xb0, 0x95, 0xe4, 0x86, 0x71, 0x5d, 0x9f,
	0xac, 0xa8, 0xc2, 0xb2, 0xd2, 0x56, 0x63, 0xd0, 0xe1, 0xc9, 0x8f, 0x03, 0xc1, 0xa2, 0x55, 0x38,
	0x04, 0x57, 0x16, 0xd4, 0x19, 0x3b, 0xd3, 0x41, 0xe2, 0xca, 0x02, 0x9f, 0x20, 0xb0, 0x22, 0x53,
	0x3b, 0x59, 0x50, 0xb7, 0x81, 0xfe, 0x75, 0x2e, 0x0b, 0x7c, 0x06, 0xcf, 0x5e, 0x4a, 0x41, 0x7b,
	0x63, 0x67, 0x3a, 0x7c, 0x79, 0x08, 0xbb, 0x58, 0xd8, 0x85, 0xc2, 0xf4, 0x52, 0x8a, 0xa4, 0x39,
	0xc1, 0x47, 0xf0, 0x8d, 0xae, 0x2b, 0x2e, 0xa8, 0xd7, 0x26, 0xda, 0x85, 0xf7, 0xd0, 0x3f, 0x67,
	0xc7, 0x5a, 0xd0, 0xfe, 0xd8, 0x99, 0xf6, 0x92, 0x76, 0xe0, 0x08, 0x06, 0x56, 0x2a, 0x61, 0x6c,
	0xa6, 0x4a, 0xea, 0x37, 0xe6, 0x0f, 0x4c, 0x62, 0xf0, 0xae, 0x65, 0x1c, 0x40, 0xff, 0x6b, 0xb5,
	0x89, 0x52, 0x72, 0x83, 0x43, 0x80, 0xc5, 0x3a, 0x8e, 0xa3, 0x45, 0xba, 0x5c, 0xaf, 0x88, 0x83,
	0x00, 0x7e, 0x12, 0x7d, 0xc6, 0xb3, 0x2d, 0x71, 0xf1, 0x16, 0x82, 0x4d, 0xba, 0x4e, 0x66, 0x1f,
	0x11, 0xe9, 0x21, 0x81, 0xbb, 0x56, 0xec, 0xe6, 0xdb, 0x34, 0xda, 0x10, 0x6f, 0xfe, 0xfe, 0xfd,
	0xb6, 0x97, 0xf6, 0x50, 0xe7, 0x21, 0xd7, 0x8a, 0xe5, 0x99, 0xe5, 0x07, 0xae, 0xab, 0x92, 0x19,
	0x7e, 0x10, 0x2a, 0x33, 0x2c, 0xaf, 0xe5, 0xb1, 0x60, 0x7b, 0xcd, 0xfe, 0x3f, 0x32, 0xf7, 0x9b,
	0x0f, 0xbe, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x75, 0x1e, 0xc7, 0x61, 0x01, 0x00, 0x00,
}
