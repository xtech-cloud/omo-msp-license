// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/license/shared.proto

package license

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

// 状态
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f71ef6d47559577, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 空白回复
type BlankResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlankResponse) Reset()         { *m = BlankResponse{} }
func (m *BlankResponse) String() string { return proto.CompactTextString(m) }
func (*BlankResponse) ProtoMessage()    {}
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f71ef6d47559577, []int{1}
}

func (m *BlankResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlankResponse.Unmarshal(m, b)
}
func (m *BlankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlankResponse.Marshal(b, m, deterministic)
}
func (m *BlankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlankResponse.Merge(m, src)
}
func (m *BlankResponse) XXX_Size() int {
	return xxx_messageInfo_BlankResponse.Size(m)
}
func (m *BlankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlankResponse proto.InternalMessageInfo

func (m *BlankResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// 空间
type SpaceEntity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SpaceKey             string   `protobuf:"bytes,2,opt,name=spaceKey,proto3" json:"spaceKey,omitempty"`
	SpaceSecret          string   `protobuf:"bytes,3,opt,name=spaceSecret,proto3" json:"spaceSecret,omitempty"`
	PublicKey            string   `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	PrivateKey           string   `protobuf:"bytes,5,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	Profile              string   `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	CreatedAt            int64    `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpaceEntity) Reset()         { *m = SpaceEntity{} }
func (m *SpaceEntity) String() string { return proto.CompactTextString(m) }
func (*SpaceEntity) ProtoMessage()    {}
func (*SpaceEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f71ef6d47559577, []int{2}
}

func (m *SpaceEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpaceEntity.Unmarshal(m, b)
}
func (m *SpaceEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpaceEntity.Marshal(b, m, deterministic)
}
func (m *SpaceEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceEntity.Merge(m, src)
}
func (m *SpaceEntity) XXX_Size() int {
	return xxx_messageInfo_SpaceEntity.Size(m)
}
func (m *SpaceEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceEntity.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceEntity proto.InternalMessageInfo

func (m *SpaceEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SpaceEntity) GetSpaceKey() string {
	if m != nil {
		return m.SpaceKey
	}
	return ""
}

func (m *SpaceEntity) GetSpaceSecret() string {
	if m != nil {
		return m.SpaceSecret
	}
	return ""
}

func (m *SpaceEntity) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *SpaceEntity) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *SpaceEntity) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *SpaceEntity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *SpaceEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

// 激活码
type KeyEntity struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Space                string   `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Expiry               int32    `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Storage              string   `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty"`
	Profile              string   `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	CreatedAt            int64    `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ActivatedAt          int64    `protobuf:"varint,9,opt,name=activatedAt,proto3" json:"activatedAt,omitempty"`
	Ban                  int32    `protobuf:"varint,10,opt,name=ban,proto3" json:"ban,omitempty"`
	Consumer             []string `protobuf:"bytes,11,rep,name=consumer,proto3" json:"consumer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEntity) Reset()         { *m = KeyEntity{} }
func (m *KeyEntity) String() string { return proto.CompactTextString(m) }
func (*KeyEntity) ProtoMessage()    {}
func (*KeyEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f71ef6d47559577, []int{3}
}

func (m *KeyEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEntity.Unmarshal(m, b)
}
func (m *KeyEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEntity.Marshal(b, m, deterministic)
}
func (m *KeyEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEntity.Merge(m, src)
}
func (m *KeyEntity) XXX_Size() int {
	return xxx_messageInfo_KeyEntity.Size(m)
}
func (m *KeyEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEntity.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEntity proto.InternalMessageInfo

func (m *KeyEntity) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *KeyEntity) GetSpace() string {
	if m != nil {
		return m.Space
	}
	return ""
}

func (m *KeyEntity) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *KeyEntity) GetExpiry() int32 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *KeyEntity) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *KeyEntity) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *KeyEntity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *KeyEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *KeyEntity) GetActivatedAt() int64 {
	if m != nil {
		return m.ActivatedAt
	}
	return 0
}

func (m *KeyEntity) GetBan() int32 {
	if m != nil {
		return m.Ban
	}
	return 0
}

func (m *KeyEntity) GetConsumer() []string {
	if m != nil {
		return m.Consumer
	}
	return nil
}

// 证书
type CertificateEntity struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Space                string   `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	Number               string   `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Consumer             string   `protobuf:"bytes,4,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt            int64    `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateEntity) Reset()         { *m = CertificateEntity{} }
func (m *CertificateEntity) String() string { return proto.CompactTextString(m) }
func (*CertificateEntity) ProtoMessage()    {}
func (*CertificateEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f71ef6d47559577, []int{4}
}

func (m *CertificateEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateEntity.Unmarshal(m, b)
}
func (m *CertificateEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateEntity.Marshal(b, m, deterministic)
}
func (m *CertificateEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateEntity.Merge(m, src)
}
func (m *CertificateEntity) XXX_Size() int {
	return xxx_messageInfo_CertificateEntity.Size(m)
}
func (m *CertificateEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateEntity.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateEntity proto.InternalMessageInfo

func (m *CertificateEntity) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CertificateEntity) GetSpace() string {
	if m != nil {
		return m.Space
	}
	return ""
}

func (m *CertificateEntity) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *CertificateEntity) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *CertificateEntity) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CertificateEntity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "license.Status")
	proto.RegisterType((*BlankResponse)(nil), "license.BlankResponse")
	proto.RegisterType((*SpaceEntity)(nil), "license.SpaceEntity")
	proto.RegisterType((*KeyEntity)(nil), "license.KeyEntity")
	proto.RegisterType((*CertificateEntity)(nil), "license.CertificateEntity")
}

func init() {
	proto.RegisterFile("proto/license/shared.proto", fileDescriptor_2f71ef6d47559577)
}

var fileDescriptor_2f71ef6d47559577 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x25, 0x9b, 0x4d, 0x66, 0x53, 0x83, 0xb8, 0x36, 0x22, 0x61, 0x11, 0x09, 0xb9, 0x38, 0xa7,
	0x59, 0x50, 0x10, 0xaf, 0x2a, 0x9e, 0xf6, 0x96, 0xf9, 0x05, 0x3d, 0x9d, 0x5a, 0x6d, 0x9c, 0xe9,
	0x34, 0xdd, 0x15, 0x71, 0x7e, 0x91, 0x27, 0x7f, 0x9d, 0x7f, 0x40, 0xaa, 0xd2, 0x99, 0x0f, 0xc1,
	0xa3, 0xb7, 0x7a, 0xef, 0x25, 0xaf, 0xf2, 0x5e, 0x77, 0xe0, 0xce, 0x87, 0x81, 0x86, 0xfb, 0x9d,
	0x35, 0xe8, 0x22, 0xde, 0xc7, 0xaf, 0x3a, 0x60, 0xbf, 0x16, 0x52, 0x2d, 0x12, 0xdb, 0xbe, 0x83,
	0x72, 0x43, 0x9a, 0xc6, 0xa8, 0x14, 0x5c, 0x9b, 0xa1, 0xc7, 0x3a, 0x6b, 0xb2, 0x55, 0xd1, 0xc9,
	0xac, 0x6a, 0x58, 0xec, 0x31, 0x46, 0xfd, 0x05, 0xeb, 0xab, 0x26, 0x5b, 0x55, 0xdd, 0x0c, 0xdb,
	0xf7, 0xf0, 0xe4, 0xe3, 0x4e, 0xbb, 0x6f, 0x1d, 0x46, 0x3f, 0xb8, 0x88, 0xea, 0x35, 0x94, 0x51,
	0x8c, 0xc4, 0x60, 0xf9, 0xe6, 0xe9, 0x3a, 0xad, 0x58, 0x4f, 0xfe, 0x5d, 0x92, 0xdb, 0xdf, 0x19,
	0x2c, 0x37, 0x5e, 0x1b, 0xfc, 0xec, 0xc8, 0xd2, 0x81, 0xf7, 0x3a, 0xbd, 0x9f, 0xf6, 0x56, 0x9d,
	0xcc, 0xea, 0x0e, 0x6e, 0x22, 0x3f, 0xf2, 0x80, 0x87, 0xb4, 0xf8, 0x88, 0x55, 0x03, 0x4b, 0x99,
	0x37, 0x68, 0x02, 0x52, 0x9d, 0x8b, 0x7c, 0x4e, 0xa9, 0x97, 0x50, 0xf9, 0x71, 0xbb, 0xb3, 0x86,
	0x5f, 0xbf, 0x16, 0xfd, 0x44, 0xa8, 0x57, 0x00, 0x3e, 0xd8, 0xef, 0x9a, 0xc4, 0xbd, 0x10, 0xf9,
	0x8c, 0xe1, 0xcc, 0x3e, 0x0c, 0x8f, 0x76, 0x87, 0x75, 0x39, 0x65, 0x4e, 0x90, 0x7d, 0x4d, 0x40,
	0x4d, 0xd8, 0x7f, 0xa0, 0x7a, 0xd1, 0x64, 0xab, 0xbc, 0x3b, 0x11, 0xac, 0x8e, 0xbe, 0x4f, 0xea,
	0xcd, 0xa4, 0x1e, 0x89, 0xf6, 0xd7, 0x15, 0x54, 0x0f, 0x78, 0x48, 0x99, 0x5f, 0x40, 0xe9, 0xc6,
	0xfd, 0x16, 0x43, 0x4a, 0x9d, 0x90, 0x7a, 0x0e, 0x85, 0x04, 0x49, 0xa1, 0x27, 0xc0, 0x6d, 0x18,
	0xed, 0xb5, 0xb1, 0x74, 0x90, 0xb8, 0x45, 0x77, 0xc4, 0xec, 0x84, 0x3f, 0xbc, 0x0d, 0x53, 0xd0,
	0xa2, 0x4b, 0x88, 0x53, 0x44, 0x1a, 0x02, 0x9f, 0xdc, 0x14, 0x71, 0x86, 0xff, 0x27, 0x1f, 0x9f,
	0x8a, 0x36, 0x24, 0x25, 0xb2, 0x5e, 0x89, 0x7e, 0x4e, 0xa9, 0x5b, 0xc8, 0xb7, 0xda, 0xd5, 0x20,
	0x9f, 0xc9, 0xa3, 0xe4, 0x1a, 0x5c, 0x1c, 0xf7, 0x18, 0xea, 0x65, 0x93, 0xf3, 0x29, 0xcf, 0xb8,
	0xfd, 0x99, 0xc1, 0xb3, 0x4f, 0x18, 0xc8, 0x3e, 0x5a, 0xa3, 0x69, 0xbe, 0x2b, 0xb7, 0x90, 0x8f,
	0xb6, 0x4f, 0xa5, 0xf1, 0xf8, 0x8f, 0xc6, 0x4e, 0xfd, 0xe6, 0x17, 0xfd, 0x9e, 0x6f, 0x9c, 0x2e,
	0xc6, 0x11, 0x73, 0x2f, 0x66, 0x70, 0x84, 0x8e, 0xe6, 0xc6, 0x12, 0xbc, 0xec, 0xa5, 0xfc, 0xab,
	0x97, 0x6d, 0x29, 0x7f, 0xd4, 0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x63, 0x76, 0x28,
	0x6f, 0x03, 0x00, 0x00,
}
