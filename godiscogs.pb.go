// Code generated by protoc-gen-go.
// source: godiscogs.proto
// DO NOT EDIT!

/*
Package godiscogs is a generated protocol buffer package.

It is generated from these files:
	godiscogs.proto

It has these top-level messages:
	Label
	Folder
	Artist
	Image
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

type Label struct {
	// The name of the label
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Folder struct {
	// The id number of the folder
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The name of the folder
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Folder) Reset()                    { *m = Folder{} }
func (m *Folder) String() string            { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()               {}
func (*Folder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Artist struct {
	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The name of the artist
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Image struct {
	// The uri to the image
	Uri string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	// The type of image
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Release struct {
	// The id number of the release
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The title of the release
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// Artists associated with the release
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists" json:"artists,omitempty"`
	// The folder in which the record is stored
	FolderId int32 `protobuf:"varint,4,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
	// Images associated with the release
	Images []*Image `protobuf:"bytes,5,rep,name=images" json:"images,omitempty"`
	// The instance id of this release
	InstanceId int32 `protobuf:"varint,6,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	// The labels connected to this release
	Labels []*Label `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Release) GetArtists() []*Artist {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *Release) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Release) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "Label")
	proto.RegisterType((*Folder)(nil), "Folder")
	proto.RegisterType((*Artist)(nil), "Artist")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Release)(nil), "Release")
}

func init() { proto.RegisterFile("godiscogs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x9a, 0xda, 0xa1, 0x5b, 0x09, 0xd0, 0x8a, 0x43, 0xa4, 0x4a, 0x50, 0x72, 0xea, 0x01,
	0x72, 0x80, 0x17, 0x70, 0x41, 0xaa, 0xc4, 0xc9, 0x1f, 0x40, 0x6e, 0xb3, 0x44, 0x96, 0xdc, 0xb8,
	0xca, 0x9a, 0x03, 0xbf, 0xe4, 0x49, 0xd4, 0x9b, 0xa4, 0x12, 0xe2, 0xc2, 0x6d, 0x3c, 0x9e, 0x19,
	0x8f, 0x77, 0xe1, 0xaa, 0x0d, 0x8d, 0xe3, 0x7d, 0x68, 0xb9, 0x3e, 0xf6, 0x21, 0x86, 0x6a, 0x05,
	0xea, 0xcd, 0xee, 0xc8, 0x23, 0xc2, 0xbc, 0xb3, 0x07, 0x2a, 0xb3, 0x75, 0xb6, 0x59, 0x18, 0xc1,
	0xd5, 0x03, 0xe8, 0xd7, 0xe0, 0x1b, 0xea, 0xf1, 0x12, 0x66, 0xae, 0x91, 0x3b, 0x65, 0x4e, 0xe8,
	0xac, 0x9e, 0xfd, 0x56, 0xbf, 0xf4, 0xd1, 0x71, 0xfc, 0x97, 0xfa, 0x11, 0xd4, 0xf6, 0x60, 0x5b,
	0xc2, 0x6b, 0xc8, 0x3f, 0x7b, 0x37, 0xbe, 0x9b, 0x60, 0x92, 0xc7, 0xaf, 0xe3, 0x59, 0x9e, 0x70,
	0xf5, 0x9d, 0x41, 0x61, 0xc8, 0x93, 0x65, 0xfa, 0x13, 0x7f, 0x03, 0x2a, 0xba, 0xe8, 0x27, 0xc3,
	0x70, 0xc0, 0x7b, 0x28, 0xac, 0xd4, 0xe1, 0x32, 0x5f, 0xe7, 0x9b, 0xe5, 0x53, 0x51, 0x0f, 0xf5,
	0xcc, 0xc4, 0xe3, 0x0a, 0x16, 0x1f, 0xf2, 0xbf, 0xf7, 0x53, 0xde, 0x5c, 0xf2, 0x2e, 0x06, 0x62,
	0xdb, 0xe0, 0x2d, 0x68, 0x97, 0x0a, 0x72, 0xa9, 0xc4, 0xae, 0x6b, 0xe9, 0x6b, 0x46, 0x16, 0xef,
	0x60, 0xe9, 0x3a, 0x8e, 0xb6, 0xdb, 0x53, 0xb2, 0x6b, 0xb1, 0xc3, 0x44, 0x0d, 0x01, 0x3e, 0x8d,
	0x96, 0xcb, 0x62, 0x0c, 0x90, 0x49, 0x9b, 0x91, 0xdd, 0x69, 0xd9, 0xc0, 0xf3, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x02, 0x69, 0x22, 0x94, 0x01, 0x00, 0x00,
}
