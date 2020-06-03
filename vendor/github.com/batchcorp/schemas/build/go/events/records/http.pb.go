// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/http.proto

package records

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

type HTTPRecord struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                 string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Source               string            `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Body                 []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Timestamp            int64             `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTPRecord) Reset()         { *m = HTTPRecord{} }
func (m *HTTPRecord) String() string { return proto.CompactTextString(m) }
func (*HTTPRecord) ProtoMessage()    {}
func (*HTTPRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_40dfd77d7b66d03c, []int{0}
}

func (m *HTTPRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPRecord.Unmarshal(m, b)
}
func (m *HTTPRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPRecord.Marshal(b, m, deterministic)
}
func (m *HTTPRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPRecord.Merge(m, src)
}
func (m *HTTPRecord) XXX_Size() int {
	return xxx_messageInfo_HTTPRecord.Size(m)
}
func (m *HTTPRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPRecord.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPRecord proto.InternalMessageInfo

func (m *HTTPRecord) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRecord) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRecord) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPRecord) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *HTTPRecord) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HTTPRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*HTTPRecord)(nil), "records.HTTPRecord")
	proto.RegisterMapType((map[string]string)(nil), "records.HTTPRecord.HeadersEntry")
}

func init() { proto.RegisterFile("events/records/http.proto", fileDescriptor_40dfd77d7b66d03c) }

var fileDescriptor_40dfd77d7b66d03c = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0xa4, 0x4d, 0xe9, 0xd8, 0x83, 0x2c, 0x22, 0xab, 0x78, 0x08, 0x9e, 0x72, 0xca,
	0x82, 0x8a, 0x48, 0x8e, 0x82, 0xd0, 0xa3, 0x84, 0x9e, 0xbc, 0x6d, 0x36, 0x43, 0x37, 0xd8, 0x74,
	0xc3, 0xee, 0xa4, 0x90, 0xbf, 0xee, 0x49, 0xb2, 0x89, 0xd4, 0xde, 0xde, 0xf7, 0x98, 0x99, 0x37,
	0x3c, 0xb8, 0xc3, 0x13, 0x1e, 0xc9, 0x4b, 0x87, 0xda, 0xba, 0xda, 0x4b, 0x43, 0xd4, 0xe5, 0x9d,
	0xb3, 0x64, 0xf9, 0x6a, 0xf6, 0x1e, 0x7f, 0x18, 0xc0, 0x76, 0xb7, 0xfb, 0x2c, 0x03, 0xf3, 0x5b,
	0x48, 0x5a, 0x24, 0x63, 0x6b, 0xc1, 0x52, 0x96, 0xad, 0xcb, 0x99, 0x78, 0x01, 0x2b, 0x83, 0xaa,
	0x46, 0xe7, 0x45, 0x94, 0xc6, 0xd9, 0xd5, 0x53, 0x9a, 0xcf, 0x17, 0xf2, 0xf3, 0x76, 0xbe, 0x9d,
	0x46, 0x3e, 0x8e, 0xe4, 0x86, 0xf2, 0x6f, 0x81, 0x73, 0x58, 0x74, 0x8a, 0x8c, 0x88, 0xc3, 0xc5,
	0xa0, 0xc7, 0x1c, 0x6f, 0x7b, 0xa7, 0x51, 0x2c, 0xa6, 0x9c, 0x89, 0xc6, 0xd9, 0xca, 0xd6, 0x83,
	0x58, 0xa6, 0x2c, 0xdb, 0x94, 0x41, 0xf3, 0x07, 0x58, 0x53, 0xd3, 0xa2, 0x27, 0xd5, 0x76, 0x22,
	0x49, 0x59, 0x16, 0x97, 0x67, 0xe3, 0xbe, 0x80, 0xcd, 0xff, 0x58, 0x7e, 0x0d, 0xf1, 0x37, 0x0e,
	0xf3, 0xfb, 0xa3, 0xe4, 0x37, 0xb0, 0x3c, 0xa9, 0x43, 0x8f, 0x22, 0x0a, 0xde, 0x04, 0x45, 0xf4,
	0xc6, 0xde, 0x5f, 0xbf, 0x5e, 0xf6, 0x0d, 0x99, 0xbe, 0xca, 0xb5, 0x6d, 0x65, 0xa5, 0x48, 0x1b,
	0x6d, 0x5d, 0x27, 0xbd, 0x36, 0xd8, 0x2a, 0x2f, 0xab, 0xbe, 0x39, 0xd4, 0x72, 0x6f, 0xe5, 0x65,
	0x91, 0x55, 0x12, 0x4a, 0x7c, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x21, 0x89, 0xcb, 0x61,
	0x01, 0x00, 0x00,
}
