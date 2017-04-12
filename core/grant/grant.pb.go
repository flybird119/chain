// Code generated by protoc-gen-go.
// source: grant.proto
// DO NOT EDIT!

/*
Package grant is a generated protocol buffer package.

It is generated from these files:
	grant.proto

It has these top-level messages:
	GrantList
	Grant
*/
package grant

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

type GrantList struct {
	Grants []*Grant `protobuf:"bytes,1,rep,name=grants" json:"grants,omitempty"`
}

func (m *GrantList) Reset()                    { *m = GrantList{} }
func (m *GrantList) String() string            { return proto.CompactTextString(m) }
func (*GrantList) ProtoMessage()               {}
func (*GrantList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GrantList) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

type Grant struct {
	GuardType string `protobuf:"bytes,1,opt,name=guard_type,json=guardType" json:"guard_type,omitempty"`
	GuardData []byte `protobuf:"bytes,2,opt,name=guard_data,json=guardData,proto3" json:"guard_data,omitempty"`
	Policy    string `protobuf:"bytes,3,opt,name=policy" json:"policy,omitempty"`
}

func (m *Grant) Reset()                    { *m = Grant{} }
func (m *Grant) String() string            { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()               {}
func (*Grant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Grant) GetGuardType() string {
	if m != nil {
		return m.GuardType
	}
	return ""
}

func (m *Grant) GetGuardData() []byte {
	if m != nil {
		return m.GuardData
	}
	return nil
}

func (m *Grant) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func init() {
	proto.RegisterType((*GrantList)(nil), "grant.GrantList")
	proto.RegisterType((*Grant)(nil), "grant.Grant")
}

func init() { proto.RegisterFile("grant.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0xcc,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x0c, 0xb9, 0x38, 0xdd,
	0x41, 0x0c, 0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x15, 0x2e, 0x36, 0xb0, 0x68, 0xb1, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x11, 0x8f, 0x1e, 0x44, 0x07, 0x58, 0x45, 0x10, 0x54, 0x4e, 0x29, 0x96, 0x8b,
	0x15, 0x2c, 0x20, 0x24, 0xcb, 0xc5, 0x95, 0x5e, 0x9a, 0x58, 0x94, 0x12, 0x5f, 0x52, 0x59, 0x90,
	0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09, 0x16, 0x09, 0xa9, 0x2c, 0x48, 0x45, 0x48,
	0xa7, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x40, 0xa5, 0x5d, 0x12, 0x4b, 0x12,
	0x85, 0xc4, 0xb8, 0xd8, 0x0a, 0xf2, 0x73, 0x32, 0x93, 0x2b, 0x25, 0x98, 0xc1, 0x3a, 0xa1, 0xbc,
	0x24, 0x36, 0xb0, 0xfb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x5e, 0x92, 0x42, 0xae,
	0x00, 0x00, 0x00,
}
