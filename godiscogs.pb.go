// Code generated by protoc-gen-go.
// source: godiscogs.proto
// DO NOT EDIT!

/*
Package godiscogs is a generated protocol buffer package.

It is generated from these files:
	godiscogs.proto

It has these top-level messages:
	Artist
	Release
*/
package godiscogs

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

type Artist struct {
	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The name of the artist
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Release struct {
	// The id number of the release
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The title of the release
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// Artists associated with the release
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists" json:"artists,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Release) GetArtists() []*Artist {
	if m != nil {
		return m.Artists
	}
	return nil
}

func init() {
	proto.RegisterType((*Artist)(nil), "Artist")
	proto.RegisterType((*Release)(nil), "Release")
}

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcf, 0x4f, 0xc9,
	0x2c, 0x4e, 0xce, 0x4f, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe1, 0x62, 0x73,
	0x2c, 0x2a, 0xc9, 0x2c, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x02, 0xb2, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x80, 0x22,
	0x9c, 0x41, 0x60, 0xb6, 0x52, 0x10, 0x17, 0x7b, 0x50, 0x6a, 0x4e, 0x6a, 0x62, 0x71, 0x2a, 0x86,
	0x72, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x1c, 0x98, 0x7a, 0x08, 0x47, 0x48, 0x91, 0x8b, 0x3d,
	0x11, 0x6c, 0x7c, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0xbb, 0x1e, 0xc4, 0xba, 0x20,
	0x98, 0x78, 0x12, 0x1b, 0xd8, 0x21, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xb1, 0x26,
	0xb3, 0x9b, 0x00, 0x00, 0x00,
}
