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
	SaleState_EXPIRED      SaleState = 3
)

var SaleState_name = map[int32]string{
	0: "NOT_FOR_SALE",
	1: "FOR_SALE",
	2: "SOLD",
	3: "EXPIRED",
}

var SaleState_value = map[string]int32{
	"NOT_FOR_SALE": 0,
	"FOR_SALE":     1,
	"SOLD":         2,
	"EXPIRED":      3,
}

func (x SaleState) String() string {
	return proto.EnumName(SaleState_name, int32(x))
}

func (SaleState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

type Track_TrackType int32

const (
	Track_UNKNOWN Track_TrackType = 0
	Track_TRACK   Track_TrackType = 1
	Track_HEADING Track_TrackType = 2
)

var Track_TrackType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TRACK",
	2: "HEADING",
}

var Track_TrackType_value = map[string]int32{
	"UNKNOWN": 0,
	"TRACK":   1,
	"HEADING": 2,
}

func (x Track_TrackType) String() string {
	return proto.EnumName(Track_TrackType_name, int32(x))
}

func (Track_TrackType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{8, 0}
}

type ForSale struct {
	// The id of the record
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the sale
	SaleId int32 `protobuf:"varint,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	// The current price
	SalePrice            int32    `protobuf:"varint,3,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForSale) Reset()         { *m = ForSale{} }
func (m *ForSale) String() string { return proto.CompactTextString(m) }
func (*ForSale) ProtoMessage()    {}
func (*ForSale) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

func (m *ForSale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForSale.Unmarshal(m, b)
}
func (m *ForSale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForSale.Marshal(b, m, deterministic)
}
func (m *ForSale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForSale.Merge(m, src)
}
func (m *ForSale) XXX_Size() int {
	return xxx_messageInfo_ForSale.Size(m)
}
func (m *ForSale) XXX_DiscardUnknown() {
	xxx_messageInfo_ForSale.DiscardUnknown(m)
}

var xxx_messageInfo_ForSale proto.InternalMessageInfo

func (m *ForSale) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ForSale) GetSaleId() int32 {
	if m != nil {
		return m.SaleId
	}
	return 0
}

func (m *ForSale) GetSalePrice() int32 {
	if m != nil {
		return m.SalePrice
	}
	return 0
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
	return fileDescriptor_579be35d7413271f, []int{1}
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
	return fileDescriptor_579be35d7413271f, []int{2}
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
	return fileDescriptor_579be35d7413271f, []int{3}
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
	return fileDescriptor_579be35d7413271f, []int{4}
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
	return fileDescriptor_579be35d7413271f, []int{5}
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
	Boxset bool `protobuf:"varint,15,opt,name=boxset,proto3" json:"boxset,omitempty"`
	// The tracks for this release
	Tracklist            []*Track `protobuf:"bytes,16,rep,name=tracklist,proto3" json:"tracklist,omitempty"`
	InstanceNotes        []*Note  `protobuf:"bytes,19,rep,name=instance_notes,json=instanceNotes,proto3" json:"instance_notes,omitempty"`
	RecordCondition      string   `protobuf:"bytes,17,opt,name=record_condition,json=recordCondition,proto3" json:"record_condition,omitempty"`
	SleeveCondition      string   `protobuf:"bytes,18,opt,name=sleeve_condition,json=sleeveCondition,proto3" json:"sleeve_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{6}
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

func (m *Release) GetTracklist() []*Track {
	if m != nil {
		return m.Tracklist
	}
	return nil
}

func (m *Release) GetInstanceNotes() []*Note {
	if m != nil {
		return m.InstanceNotes
	}
	return nil
}

func (m *Release) GetRecordCondition() string {
	if m != nil {
		return m.RecordCondition
	}
	return ""
}

func (m *Release) GetSleeveCondition() string {
	if m != nil {
		return m.SleeveCondition
	}
	return ""
}

type Note struct {
	FieldId              int32    `protobuf:"varint,1,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{7}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetFieldId() int32 {
	if m != nil {
		return m.FieldId
	}
	return 0
}

func (m *Note) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Track struct {
	Position             string          `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Title                string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type_                string          `protobuf:"bytes,3,opt,name=type_,json=type,proto3" json:"type_,omitempty"`
	TrackType            Track_TrackType `protobuf:"varint,4,opt,name=track_type,json=trackType,proto3,enum=godiscogs.Track_TrackType" json:"track_type,omitempty"`
	SubTracks            []*Track        `protobuf:"bytes,5,rep,name=sub_tracks,json=subTracks,proto3" json:"sub_tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Track) Reset()         { *m = Track{} }
func (m *Track) String() string { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()    {}
func (*Track) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{8}
}

func (m *Track) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Track.Unmarshal(m, b)
}
func (m *Track) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Track.Marshal(b, m, deterministic)
}
func (m *Track) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Track.Merge(m, src)
}
func (m *Track) XXX_Size() int {
	return xxx_messageInfo_Track.Size(m)
}
func (m *Track) XXX_DiscardUnknown() {
	xxx_messageInfo_Track.DiscardUnknown(m)
}

var xxx_messageInfo_Track proto.InternalMessageInfo

func (m *Track) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Track) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Track) GetType_() string {
	if m != nil {
		return m.Type_
	}
	return ""
}

func (m *Track) GetTrackType() Track_TrackType {
	if m != nil {
		return m.TrackType
	}
	return Track_UNKNOWN
}

func (m *Track) GetSubTracks() []*Track {
	if m != nil {
		return m.SubTracks
	}
	return nil
}

func init() {
	proto.RegisterEnum("godiscogs.SaleState", SaleState_name, SaleState_value)
	proto.RegisterEnum("godiscogs.Track_TrackType", Track_TrackType_name, Track_TrackType_value)
	proto.RegisterType((*ForSale)(nil), "godiscogs.ForSale")
	proto.RegisterType((*Label)(nil), "godiscogs.Label")
	proto.RegisterType((*Folder)(nil), "godiscogs.Folder")
	proto.RegisterType((*Artist)(nil), "godiscogs.Artist")
	proto.RegisterType((*Image)(nil), "godiscogs.Image")
	proto.RegisterType((*Format)(nil), "godiscogs.Format")
	proto.RegisterType((*Release)(nil), "godiscogs.Release")
	proto.RegisterType((*Note)(nil), "godiscogs.Note")
	proto.RegisterType((*Track)(nil), "godiscogs.Track")
}

func init() { proto.RegisterFile("godiscogs.proto", fileDescriptor_579be35d7413271f) }

var fileDescriptor_579be35d7413271f = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xea, 0x46,
	0x10, 0xbd, 0xc6, 0x18, 0xf0, 0xc0, 0x05, 0xdf, 0xcd, 0x6d, 0xbb, 0xbd, 0x55, 0x55, 0xe4, 0x97,
	0xd2, 0xaf, 0xb4, 0x4a, 0xa5, 0x56, 0x7d, 0xa4, 0x81, 0xb4, 0x28, 0x11, 0x24, 0x0b, 0x55, 0xfb,
	0x66, 0x2d, 0xf6, 0x82, 0x56, 0x35, 0x36, 0xf1, 0x2e, 0x51, 0xf8, 0x6d, 0xfd, 0x4b, 0xfd, 0x11,
	0xd5, 0xce, 0xda, 0x4e, 0x5a, 0x52, 0xe9, 0xbe, 0xa0, 0x9d, 0x33, 0x87, 0x33, 0xbb, 0x67, 0x66,
	0x0c, 0x83, 0x6d, 0x9e, 0x48, 0x15, 0xe7, 0x5b, 0x75, 0xbe, 0x2f, 0x72, 0x9d, 0x13, 0xbf, 0x06,
	0xc2, 0x3b, 0x68, 0x5f, 0xe5, 0xc5, 0x92, 0xa7, 0x82, 0xf4, 0xa1, 0x21, 0x13, 0xea, 0x0c, 0x9d,
	0x91, 0xc7, 0x1a, 0x32, 0x21, 0x1f, 0x41, 0x5b, 0xf1, 0x54, 0x44, 0x32, 0xa1, 0x0d, 0x04, 0x5b,
	0x26, 0x9c, 0x25, 0xe4, 0x53, 0x00, 0x4c, 0xec, 0x0b, 0x19, 0x0b, 0xea, 0x62, 0xce, 0x37, 0xc8,
	0xad, 0x01, 0xc2, 0x31, 0x78, 0x37, 0x7c, 0x2d, 0x52, 0x42, 0xa0, 0x99, 0xf1, 0x9d, 0x40, 0x49,
	0x9f, 0xe1, 0x99, 0xbc, 0x05, 0x2f, 0xe6, 0x3a, 0xcb, 0x51, 0xd2, 0x67, 0x36, 0x28, 0x4b, 0xbb,
	0x55, 0xe9, 0xf0, 0x6b, 0x68, 0x5d, 0xe5, 0x69, 0x22, 0x8a, 0x93, 0x4b, 0x55, 0x9a, 0x8d, 0x27,
	0x4d, 0xc3, 0x1e, 0x17, 0x5a, 0x2a, 0xfd, 0x5e, 0xec, 0x6f, 0xc0, 0x9b, 0xed, 0xf8, 0x56, 0x90,
	0x00, 0xdc, 0x43, 0x21, 0xcb, 0xdb, 0x99, 0xa3, 0xa1, 0xeb, 0xe3, 0xbe, 0xa6, 0x9b, 0x73, 0xb8,
	0x31, 0x57, 0x29, 0x76, 0x5c, 0x93, 0x10, 0x7a, 0x89, 0x50, 0x71, 0x21, 0xf7, 0x5a, 0xe6, 0x99,
	0xa2, 0xce, 0xd0, 0x1d, 0xf9, 0xec, 0x5f, 0xd8, 0x4b, 0x05, 0x4d, 0x9d, 0x7b, 0x7d, 0xc4, 0xd7,
	0xf9, 0xcc, 0x1c, 0xb1, 0x8e, 0x78, 0xd4, 0xb4, 0x59, 0xd6, 0x11, 0x8f, 0x3a, 0xfc, 0xcb, 0x83,
	0x36, 0x13, 0xa9, 0xe0, 0xea, 0xb4, 0x13, 0x6f, 0xc1, 0xd3, 0x52, 0xa7, 0x95, 0xac, 0x0d, 0xc8,
	0x57, 0xd0, 0xe6, 0xf8, 0x6c, 0x45, 0xdd, 0xa1, 0x3b, 0xea, 0x5e, 0xbc, 0x39, 0x7f, 0x6a, 0xb4,
	0x35, 0x84, 0x55, 0x0c, 0xf2, 0x09, 0xf8, 0x1b, 0x74, 0xd4, 0xb4, 0xb3, 0x89, 0xca, 0x1d, 0x0b,
	0xcc, 0x12, 0x32, 0x82, 0x96, 0x34, 0x96, 0x28, 0xea, 0xa1, 0x50, 0xf0, 0x4c, 0x08, 0xbd, 0x62,
	0x65, 0x9e, 0x7c, 0x06, 0x5d, 0x99, 0x29, 0xcd, 0xb3, 0x18, 0xe7, 0xa2, 0x85, 0x42, 0x50, 0x41,
	0x56, 0x2a, 0x35, 0xcd, 0x57, 0xb4, 0x7d, 0x22, 0x85, 0x53, 0xc1, 0xca, 0x3c, 0xf9, 0x1c, 0x06,
	0x1b, 0x34, 0x36, 0xba, 0x3f, 0xf0, 0x4c, 0x4b, 0x7d, 0xa4, 0x1d, 0x94, 0xeb, 0x5b, 0xf8, 0xae,
	0x44, 0xc9, 0x87, 0xd0, 0x2a, 0xb8, 0x96, 0xd9, 0x96, 0xfa, 0x76, 0x0c, 0x6d, 0x44, 0x2e, 0xe0,
	0x03, 0xc1, 0x8b, 0x54, 0x0a, 0xa5, 0xa3, 0xc2, 0x3a, 0x17, 0x25, 0x5c, 0x0b, 0x0a, 0x43, 0x67,
	0xe4, 0xb2, 0xb3, 0x2a, 0x59, 0xba, 0x3a, 0xe1, 0x5a, 0x18, 0x1b, 0x76, 0x5c, 0x69, 0x6b, 0x43,
	0xd7, 0xda, 0x60, 0x81, 0x59, 0x42, 0xde, 0x41, 0xa7, 0xd4, 0x49, 0x68, 0x0f, 0x9d, 0xae, 0x63,
	0x63, 0xb6, 0xbd, 0x96, 0xa2, 0xaf, 0x4f, 0xcc, 0xb6, 0x03, 0xc2, 0x2a, 0x86, 0x11, 0xda, 0x72,
	0x2d, 0x8c, 0xbf, 0xb4, 0x3f, 0x74, 0x46, 0x1d, 0x56, 0xc7, 0xe6, 0x35, 0xeb, 0xfc, 0x51, 0x09,
	0x4d, 0x07, 0x98, 0x29, 0x23, 0x72, 0x0e, 0xbe, 0x2e, 0x78, 0xfc, 0x67, 0x2a, 0x95, 0xa6, 0xc1,
	0x89, 0x77, 0x2b, 0x93, 0x63, 0x4f, 0x14, 0xf2, 0x03, 0xf4, 0xeb, 0x4e, 0x64, 0xb9, 0x16, 0x8a,
	0x9e, 0xe1, 0x9f, 0x06, 0xcf, 0xfe, 0x34, 0xcf, 0xb5, 0x60, 0xaf, 0x2b, 0x9a, 0x89, 0x14, 0xf9,
	0x02, 0x82, 0x42, 0xc4, 0x79, 0x91, 0x44, 0x71, 0x9e, 0x25, 0xd2, 0x8c, 0x2d, 0x7d, 0x83, 0x8f,
	0x1d, 0x58, 0xfc, 0xb2, 0x82, 0x0d, 0x55, 0xa5, 0x42, 0x3c, 0x88, 0x67, 0x54, 0x62, 0xa9, 0x16,
	0xaf, 0xa9, 0xe1, 0x8f, 0xd0, 0x34, 0xf2, 0xe4, 0x63, 0xe8, 0x6c, 0xa4, 0x48, 0x93, 0xa8, 0x9e,
	0xdf, 0x36, 0xc6, 0x33, 0x1c, 0xe2, 0x07, 0x9e, 0x1e, 0xea, 0x21, 0xc6, 0x20, 0xfc, 0xdb, 0x01,
	0x0f, 0xdf, 0x66, 0x4c, 0xdb, 0xe7, 0xca, 0x56, 0xb1, 0x3b, 0x59, 0xc7, 0xff, 0xb3, 0x00, 0x67,
	0xe0, 0x99, 0x15, 0x8d, 0xca, 0xd5, 0xc2, 0x7d, 0x25, 0x3f, 0x01, 0xa0, 0x49, 0x11, 0x6e, 0xb2,
	0x99, 0xf4, 0xfe, 0xc5, 0xbb, 0xff, 0x1a, 0x69, 0x7f, 0x57, 0xc7, 0xbd, 0x28, 0x2d, 0x35, 0x47,
	0xf2, 0x2d, 0x80, 0x3a, 0xac, 0x23, 0x04, 0x5e, 0x5a, 0x85, 0xb2, 0x07, 0xea, 0xb0, 0xc6, 0x93,
	0x0a, 0xbf, 0x03, 0xbf, 0x16, 0x22, 0x5d, 0x68, 0xff, 0x36, 0xbf, 0x9e, 0x2f, 0x7e, 0x9f, 0x07,
	0xaf, 0x88, 0x0f, 0xde, 0x8a, 0x8d, 0x2f, 0xaf, 0x03, 0xc7, 0xe0, 0xbf, 0x4e, 0xc7, 0x93, 0xd9,
	0xfc, 0x97, 0xa0, 0xf1, 0xe5, 0xcf, 0xe0, 0x9b, 0x6f, 0xed, 0x52, 0x9b, 0x61, 0x0c, 0xa0, 0x37,
	0x5f, 0xac, 0xa2, 0xab, 0x05, 0x8b, 0x96, 0xe3, 0x9b, 0x69, 0xf0, 0x8a, 0xf4, 0xa0, 0x53, 0x47,
	0x0e, 0xe9, 0x40, 0x73, 0xb9, 0xb8, 0x99, 0x04, 0x0d, 0xa3, 0x31, 0xfd, 0xe3, 0x76, 0xc6, 0xa6,
	0x93, 0xc0, 0x5d, 0xb7, 0xf0, 0x23, 0xfe, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa8,
	0xbf, 0x51, 0xd7, 0x05, 0x00, 0x00,
}
