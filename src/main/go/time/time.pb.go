// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devim/protobuf/time.proto

package time // import "github.com/devimteam/proto-utils/src/main/go/time"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Period struct {
	Start                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_44da16b8362697d2, []int{0}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Period.Unmarshal(m, b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Period.Marshal(b, m, deterministic)
}
func (dst *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(dst, src)
}
func (m *Period) XXX_Size() int {
	return xxx_messageInfo_Period.Size(m)
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Period) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*Period)(nil), "devim.protobuf.Period")
}

func init() { proto.RegisterFile("devim/protobuf/time.proto", fileDescriptor_time_44da16b8362697d2) }

var fileDescriptor_time_44da16b8362697d2 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0x2d, 0xcb,
	0xcc, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d, 0xd5,
	0x03, 0xf3, 0x84, 0xf8, 0xc0, 0x52, 0x7a, 0x30, 0x29, 0x29, 0xf9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c,
	0x54, 0x54, 0xb5, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0x35, 0x4a, 0x19, 0x5c, 0x6c, 0x01, 0xa9,
	0x45, 0x99, 0xf9, 0x29, 0x42, 0x06, 0x5c, 0xac, 0xc5, 0x25, 0x89, 0x45, 0x25, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0xad, 0x70, 0xb3, 0xf4, 0x42, 0x60, 0x5a, 0x83, 0x20,
	0x0a, 0x85, 0x74, 0xb8, 0x98, 0x53, 0xf3, 0x52, 0x24, 0x98, 0x08, 0xaa, 0x07, 0x29, 0x73, 0x6a,
	0x60, 0xe4, 0x12, 0x4a, 0xce, 0xcf, 0xd5, 0x43, 0x75, 0xa1, 0x13, 0x27, 0x48, 0x59, 0x00, 0x88,
	0x17, 0xc0, 0x18, 0x65, 0x98, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0x56, 0x57, 0x92, 0x9a, 0x08, 0xf5, 0xa8, 0x6e, 0x69, 0x49, 0x66, 0x4e, 0xb1, 0x7e, 0x71, 0x51,
	0xb2, 0x7e, 0x6e, 0x62, 0x66, 0x9e, 0x7e, 0x7a, 0x3e, 0xd8, 0x2f, 0x0b, 0x18, 0x19, 0x7f, 0x30,
	0x32, 0x2e, 0x62, 0x62, 0x76, 0x0f, 0x70, 0x5a, 0xc5, 0x24, 0xeb, 0x02, 0x36, 0x3c, 0x00, 0xe6,
	0x84, 0xf0, 0xd4, 0x9c, 0x1c, 0xef, 0xbc, 0xfc, 0xf2, 0xbc, 0x90, 0xca, 0x82, 0xd4, 0xe2, 0x24,
	0x36, 0xb0, 0x49, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xa5, 0xfc, 0x38, 0x41, 0x01,
	0x00, 0x00,
}
