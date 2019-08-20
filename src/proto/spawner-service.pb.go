// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spawner-service.proto

package proto

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

type Status int32

const (
	Status_OK Status = 0
)

var Status_name = map[int32]string{
	0: "OK",
}

var Status_value = map[string]int32{
	"OK": 0,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6e612d6c76e05af, []int{0}
}

type TestRequest struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6e612d6c76e05af, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type TestResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6e612d6c76e05af, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func init() {
	proto.RegisterEnum("proto.Status", Status_name, Status_value)
	proto.RegisterType((*TestRequest)(nil), "proto.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "proto.TestResponse")
}

func init() { proto.RegisterFile("spawner-service.proto", fileDescriptor_d6e612d6c76e05af) }

var fileDescriptor_d6e612d6c76e05af = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0x48, 0x2c,
	0xcf, 0x4b, 0x2d, 0xd2, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x8a, 0x5c, 0xdc, 0x21, 0xa9, 0xc5, 0x25, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x49, 0xf9, 0x29, 0x95, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x29, 0x17, 0x0f, 0x44, 0x49, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x2a, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x58, 0x15, 0x9f, 0x11, 0x2f,
	0xc4, 0x44, 0xbd, 0x60, 0xb0, 0x60, 0x10, 0x54, 0x52, 0x4b, 0x80, 0x8b, 0x0d, 0x22, 0x22, 0xc4,
	0xc6, 0xc5, 0xe4, 0xef, 0x2d, 0xc0, 0x60, 0xe4, 0xcc, 0xc5, 0x17, 0x0c, 0x71, 0x4b, 0x30, 0xc4,
	0x29, 0x42, 0x86, 0x5c, 0x2c, 0x21, 0x60, 0x6b, 0xa1, 0x46, 0x20, 0x39, 0x45, 0x4a, 0x18, 0x45,
	0x0c, 0x62, 0xb7, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x36,
	0x48, 0xd7, 0xa6, 0xd7, 0x00, 0x00, 0x00,
}
