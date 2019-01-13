// Code generated by protoc-gen-go. DO NOT EDIT.
// source: godiscogs.proto

package godiscogs

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SaleState int32

const (
	SaleState_NOT_FOR_SALE SaleState = 0
	SaleState_FOR_SALE     SaleState = 1
	SaleState_SOLD         SaleState = 2
)

var SaleState_name = map[int32]string{
	0: "NOT_FOR_SALE",
	1: "FOR_SALE",
	2: "SOLD",
}

var SaleState_value = map[string]int32{
	"NOT_FOR_SALE": 0,
	"FOR_SALE":     1,
	"SOLD":         2,
}

func (x SaleState) String() string {
	return proto.EnumName(SaleState_name, int32(x))
}

func (SaleState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

type Label struct {
	// The name of the label
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The catalogue number
	Catno string `protobuf:"bytes,2,opt,name=catno,proto3" json:"catno,omitempty"`
	// The id of the label
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetCatno() string {
	if m != nil {
		return m.Catno
	}
	return ""
}

func (m *Label) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Folder struct {
	//The id number of the folder
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the folder
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Folder) Reset()         { *m = Folder{} }
func (m *Folder) String() string { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()    {}
func (*Folder) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{1}
}

func (m *Folder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Folder.Unmarshal(m, b)
}
func (m *Folder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Folder.Marshal(b, m, deterministic)
}
func (m *Folder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Folder.Merge(m, src)
}
func (m *Folder) XXX_Size() int {
	return xxx_messageInfo_Folder.Size(m)
}
func (m *Folder) XXX_DiscardUnknown() {
	xxx_messageInfo_Folder.DiscardUnknown(m)
}

var xxx_messageInfo_Folder proto.InternalMessageInfo

func (m *Folder) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Folder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Artist struct {
	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the artist
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artist) Reset()         { *m = Artist{} }
func (m *Artist) String() string { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()    {}
func (*Artist) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{2}
}

func (m *Artist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artist.Unmarshal(m, b)
}
func (m *Artist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artist.Marshal(b, m, deterministic)
}
func (m *Artist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artist.Merge(m, src)
}
func (m *Artist) XXX_Size() int {
	return xxx_messageInfo_Artist.Size(m)
}
func (m *Artist) XXX_DiscardUnknown() {
	xxx_messageInfo_Artist.DiscardUnknown(m)
}

var xxx_messageInfo_Artist proto.InternalMessageInfo

func (m *Artist) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Artist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Image struct {
	// The uri to the image
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The type of image
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Format struct {
	// The descriptions of the Format
	Descriptions []string `protobuf:"bytes,1,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
	// The name of the Format
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The number of the format
	Qty string `protobuf:"bytes,3,opt,name=qty,proto3" json:"qty,omitempty"`
	// Text associated with the Format
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Format) Reset()         { *m = Format{} }
func (m *Format) String() string { return proto.CompactTextString(m) }
func (*Format) ProtoMessage()    {}
func (*Format) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{4}
}

func (m *Format) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Format.Unmarshal(m, b)
}
func (m *Format) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Format.Marshal(b, m, deterministic)
}
func (m *Format) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Format.Merge(m, src)
}
func (m *Format) XXX_Size() int {
	return xxx_messageInfo_Format.Size(m)
}
func (m *Format) XXX_DiscardUnknown() {
	xxx_messageInfo_Format.DiscardUnknown(m)
}

var xxx_messageInfo_Format proto.InternalMessageInfo

func (m *Format) GetDescriptions() []string {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

func (m *Format) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Format) GetQty() string {
	if m != nil {
		return m.Qty
	}
	return ""
}

func (m *Format) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Release struct {
	// The id number of the release
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The title of the release
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Artists associated with the release
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists,proto3" json:"artists,omitempty"`
	// The folder in which the record is stored
	FolderId int32 `protobuf:"varint,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Images associated with the release
	Images []*Image `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
	// The instance id of this release
	InstanceId int32 `protobuf:"varint,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// The labels connected to this release
	Labels []*Label `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty"`
	// The number of discs in the release
	FormatQuantity int32 `protobuf:"varint,8,opt,name=format_quantity,json=formatQuantity,proto3" json:"format_quantity,omitempty"`
	// The rating given to this release
	Rating int32 `protobuf:"varint,9,opt,name=rating,proto3" json:"rating,omitempty"`
	// The earliest release date of this record
	EarliestReleaseDate int64 `protobuf:"varint,10,opt,name=earliest_release_date,json=earliestReleaseDate,proto3" json:"earliest_release_date,omitempty"`
	// The master ID of this release
	MasterId int32 `protobuf:"varint,11,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The release date of this release
	Released string `protobuf:"bytes,12,opt,name=released,proto3" json:"released,omitempty"`
	// The formats of the release
	Formats []*Format `protobuf:"bytes,13,rep,name=formats,proto3" json:"formats,omitempty"`
	// Is this a gatefold?
	Gatefold bool `protobuf:"varint,14,opt,name=gatefold,proto3" json:"gatefold,omitempty"`
	// Is this a boxset
	Boxset               bool     `protobuf:"varint,15,opt,name=boxset,proto3" json:"boxset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{5}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Release) GetArtists() []*Artist {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *Release) GetFolderId() int32 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

func (m *Release) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Release) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *Release) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Release) GetFormatQuantity() int32 {
	if m != nil {
		return m.FormatQuantity
	}
	return 0
}

func (m *Release) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Release) GetEarliestReleaseDate() int64 {
	if m != nil {
		return m.EarliestReleaseDate
	}
	return 0
}

func (m *Release) GetMasterId() int32 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *Release) GetReleased() string {
	if m != nil {
		return m.Released
	}
	return ""
}

func (m *Release) GetFormats() []*Format {
	if m != nil {
		return m.Formats
	}
	return nil
}

func (m *Release) GetGatefold() bool {
	if m != nil {
		return m.Gatefold
	}
	return false
}

func (m *Release) GetBoxset() bool {
	if m != nil {
		return m.Boxset
	}
	return false
}

func init() {
	proto.RegisterEnum("godiscogs.SaleState", SaleState_name, SaleState_value)
	proto.RegisterType((*Label)(nil), "godiscogs.Label")
	proto.RegisterType((*Folder)(nil), "godiscogs.Folder")
	proto.RegisterType((*Artist)(nil), "godiscogs.Artist")
	proto.RegisterType((*Image)(nil), "godiscogs.Image")
	proto.RegisterType((*Format)(nil), "godiscogs.Format")
	proto.RegisterType((*Release)(nil), "godiscogs.Release")
}

func init() { proto.RegisterFile("godiscogs.proto", fileDescriptor_579be35d7413271f) }

var fileDescriptor_579be35d7413271f = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x4d, 0x73, 0x69, 0x93, 0x69, 0x6d, 0xeb, 0x7a, 0xca, 0xa2, 0x0f, 0x86, 0xbc, 0x18,
	0xfc, 0x71, 0x0f, 0x27, 0xfe, 0x01, 0x85, 0xf3, 0xa0, 0x50, 0x2c, 0x6e, 0x7d, 0x0f, 0xdb, 0xee,
	0x36, 0x2c, 0xa4, 0x49, 0x2f, 0x3b, 0x07, 0xed, 0x1f, 0x2f, 0xc8, 0xfe, 0x48, 0x4e, 0xa9, 0x0f,
	0xbe, 0xcd, 0x7c, 0x67, 0xf8, 0xce, 0xcc, 0x27, 0x1b, 0x98, 0x95, 0x8d, 0x50, 0x7a, 0xd7, 0x94,
	0xfa, 0xe6, 0xd8, 0x36, 0xd8, 0x90, 0xa4, 0x17, 0xb2, 0x05, 0x44, 0x2b, 0xbe, 0x95, 0x15, 0x21,
	0x70, 0x55, 0xf3, 0x83, 0xa4, 0x41, 0x1a, 0xe4, 0x09, 0xb3, 0x31, 0xb9, 0x86, 0x68, 0xc7, 0xb1,
	0x6e, 0xe8, 0xc0, 0x8a, 0x2e, 0x21, 0x53, 0x18, 0x28, 0x41, 0xc3, 0x34, 0xc8, 0x23, 0x36, 0x50,
	0x22, 0xfb, 0x04, 0xc3, 0xfb, 0xa6, 0x12, 0xb2, 0xf5, 0x95, 0xa0, 0xab, 0xf4, 0x9e, 0x83, 0x27,
	0x4f, 0xd3, 0xbd, 0x68, 0x51, 0x69, 0xfc, 0xaf, 0xee, 0xcf, 0x10, 0x2d, 0x0f, 0xbc, 0x94, 0x64,
	0x0e, 0xe1, 0x63, 0xab, 0xfc, 0x76, 0x26, 0x34, 0xed, 0x78, 0x3e, 0xf6, 0xed, 0x26, 0xce, 0xf6,
	0x66, 0x95, 0xf6, 0xc0, 0x91, 0x64, 0x30, 0x11, 0x52, 0xef, 0x5a, 0x75, 0x44, 0xd5, 0xd4, 0x9a,
	0x06, 0x69, 0x98, 0x27, 0xec, 0x2f, 0xed, 0x5f, 0x03, 0xcd, 0x9c, 0x07, 0x3c, 0xdb, 0xeb, 0x12,
	0x66, 0x42, 0x3b, 0x47, 0x9e, 0x90, 0x5e, 0xf9, 0x39, 0xf2, 0x84, 0xd9, 0xaf, 0x10, 0x46, 0x4c,
	0x56, 0x92, 0x6b, 0x79, 0x71, 0xc6, 0x35, 0x44, 0xa8, 0xb0, 0xea, 0x6c, 0x5d, 0x42, 0x3e, 0xc2,
	0x88, 0xdb, 0xb3, 0x35, 0x0d, 0xd3, 0x30, 0x1f, 0xdf, 0xbe, 0xb8, 0x79, 0xfa, 0x2a, 0x0e, 0x08,
	0xeb, 0x3a, 0xc8, 0x5b, 0x48, 0xf6, 0x96, 0x68, 0xa1, 0x84, 0x9d, 0x1b, 0xb1, 0xd8, 0x09, 0x4b,
	0x41, 0x72, 0x18, 0x2a, 0x83, 0x44, 0xd3, 0xc8, 0x1a, 0xcd, 0xff, 0x30, 0xb2, 0xac, 0x98, 0xaf,
	0x93, 0x77, 0x30, 0x56, 0xb5, 0x46, 0x5e, 0xef, 0xa4, 0x31, 0x1a, 0x5a, 0x23, 0xe8, 0x24, 0x67,
	0x55, 0x99, 0x8f, 0xaf, 0xe9, 0xe8, 0xc2, 0xca, 0xbe, 0x0a, 0xe6, 0xeb, 0xe4, 0x3d, 0xcc, 0xf6,
	0x16, 0x6c, 0xf1, 0xf0, 0xc8, 0x6b, 0x54, 0x78, 0xa6, 0xb1, 0xb5, 0x9b, 0x3a, 0xf9, 0x87, 0x57,
	0xc9, 0x6b, 0x18, 0xb6, 0x1c, 0x55, 0x5d, 0xd2, 0xc4, 0xd6, 0x7d, 0x46, 0x6e, 0xe1, 0x95, 0xe4,
	0x6d, 0xa5, 0xa4, 0xc6, 0xa2, 0x75, 0xe4, 0x0a, 0xc1, 0x51, 0x52, 0x48, 0x83, 0x3c, 0x64, 0x2f,
	0xbb, 0xa2, 0xa7, 0x7a, 0xc7, 0x51, 0x1a, 0x0c, 0x07, 0xae, 0xd1, 0x61, 0x18, 0x3b, 0x0c, 0x4e,
	0x58, 0x0a, 0xf2, 0x06, 0x62, 0xef, 0x23, 0xe8, 0xc4, 0x92, 0xee, 0x73, 0x03, 0xdb, 0xad, 0xa5,
	0xe9, 0xf3, 0x0b, 0xd8, 0xee, 0x81, 0xb0, 0xae, 0xc3, 0x18, 0x95, 0x1c, 0xa5, 0xe1, 0x4b, 0xa7,
	0x69, 0x90, 0xc7, 0xac, 0xcf, 0xcd, 0x35, 0xdb, 0xe6, 0xa4, 0x25, 0xd2, 0x99, 0xad, 0xf8, 0xec,
	0xc3, 0x57, 0x48, 0x36, 0xbc, 0x92, 0x1b, 0x34, 0x6b, 0xce, 0x61, 0xf2, 0x7d, 0xfd, 0xb3, 0xb8,
	0x5f, 0xb3, 0x62, 0xb3, 0x58, 0x7d, 0x9b, 0x3f, 0x23, 0x13, 0x88, 0xfb, 0x2c, 0x20, 0x31, 0x5c,
	0x6d, 0xd6, 0xab, 0xbb, 0xf9, 0x60, 0x3b, 0xb4, 0xbf, 0xdf, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x6f, 0xf6, 0xd7, 0xb1, 0x91, 0x03, 0x00, 0x00,
}
