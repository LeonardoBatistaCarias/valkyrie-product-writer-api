// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka/kafka.proto

package kafka

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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

type ProductCreate struct {
	ProductID            string          `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string          `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Brand                int32           `protobuf:"varint,4,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Price                float32         `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	Quantity             int32           `protobuf:"varint,6,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	CategoryID           string          `protobuf:"bytes,7,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	ProductImages        []*ProductImage `protobuf:"bytes,8,rep,name=productImages,proto3" json:"productImages,omitempty"`
	Active               bool            `protobuf:"varint,9,opt,name=Active,proto3" json:"Active,omitempty"`
	CreatedAt            string          `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string          `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            string          `protobuf:"bytes,12,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProductCreate) ProtoReflect() protoreflect.Message {
	//TODO implement me
	panic("implement me")
}

func (m *ProductCreate) Reset()         { *m = ProductCreate{} }
func (m *ProductCreate) String() string { return proto.CompactTextString(m) }
func (*ProductCreate) ProtoMessage()    {}
func (*ProductCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c92c80ce95b4c7c, []int{0}
}

func (m *ProductCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCreate.Unmarshal(m, b)
}
func (m *ProductCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCreate.Marshal(b, m, deterministic)
}
func (m *ProductCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCreate.Merge(m, src)
}
func (m *ProductCreate) XXX_Size() int {
	return xxx_messageInfo_ProductCreate.Size(m)
}
func (m *ProductCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCreate proto.InternalMessageInfo

func (m *ProductCreate) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *ProductCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductCreate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProductCreate) GetBrand() int32 {
	if m != nil {
		return m.Brand
	}
	return 0
}

func (m *ProductCreate) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductCreate) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *ProductCreate) GetCategoryID() string {
	if m != nil {
		return m.CategoryID
	}
	return ""
}

func (m *ProductCreate) GetProductImages() []*ProductImage {
	if m != nil {
		return m.ProductImages
	}
	return nil
}

func (m *ProductCreate) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ProductCreate) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ProductCreate) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ProductCreate) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type ProductImage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProductID            string   `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductImage) Reset()         { *m = ProductImage{} }
func (m *ProductImage) String() string { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()    {}
func (*ProductImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c92c80ce95b4c7c, []int{1}
}

func (m *ProductImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductImage.Unmarshal(m, b)
}
func (m *ProductImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductImage.Marshal(b, m, deterministic)
}
func (m *ProductImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductImage.Merge(m, src)
}
func (m *ProductImage) XXX_Size() int {
	return xxx_messageInfo_ProductImage.Size(m)
}
func (m *ProductImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductImage.DiscardUnknown(m)
}

var xxx_messageInfo_ProductImage proto.InternalMessageInfo

func (m *ProductImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductImage) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductCreate)(nil), "ProductCreate")
	proto.RegisterType((*ProductImage)(nil), "ProductImage")
}

func init() { proto.RegisterFile("kafka/kafka.proto", fileDescriptor_7c92c80ce95b4c7c) }

var fileDescriptor_7c92c80ce95b4c7c = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x6a, 0x83, 0x40,
	0x10, 0x86, 0xd1, 0x44, 0xab, 0x63, 0xa4, 0x74, 0x29, 0x65, 0x28, 0xa5, 0x2c, 0x39, 0x79, 0x4a,
	0xa1, 0x79, 0x81, 0x26, 0xf1, 0x92, 0x43, 0x8b, 0x15, 0x7a, 0xe9, 0x6d, 0xab, 0xd3, 0x20, 0x69,
	0x54, 0xd6, 0x4d, 0x21, 0xaf, 0xdd, 0x27, 0x28, 0xbb, 0x2b, 0xc6, 0x5c, 0x64, 0xff, 0xef, 0x5b,
	0x74, 0xfc, 0x07, 0x6e, 0xf6, 0xe2, 0x7b, 0x2f, 0x9e, 0xcc, 0x73, 0xd1, 0xca, 0x46, 0x35, 0xf3,
	0x3f, 0x17, 0xe2, 0x4c, 0x36, 0xe5, 0xb1, 0x50, 0x1b, 0x49, 0x42, 0x11, 0x7b, 0x80, 0xb0, 0x07,
	0xdb, 0x14, 0x1d, 0xee, 0x24, 0x61, 0x7e, 0x06, 0x8c, 0xc1, 0xf4, 0x4d, 0x1c, 0x08, 0x5d, 0x23,
	0xcc, 0x99, 0x71, 0x88, 0x52, 0xea, 0x0a, 0x59, 0xb5, 0xaa, 0x6a, 0x6a, 0x9c, 0x18, 0x35, 0x46,
	0xec, 0x16, 0xbc, 0xb5, 0x14, 0x75, 0x89, 0x53, 0xee, 0x24, 0x5e, 0x6e, 0x83, 0xa6, 0x99, 0xac,
	0x0a, 0x42, 0x8f, 0x3b, 0x89, 0x9b, 0xdb, 0xc0, 0xee, 0x21, 0x78, 0x3f, 0x8a, 0x5a, 0x55, 0xea,
	0x84, 0xbe, 0xb9, 0x3e, 0x64, 0xf6, 0x08, 0xb0, 0x11, 0x8a, 0x76, 0x8d, 0x3c, 0x6d, 0x53, 0xbc,
	0x32, 0x1f, 0x1a, 0x11, 0xb6, 0x84, 0xb8, 0xed, 0x47, 0x3d, 0x88, 0x1d, 0x75, 0x18, 0xf0, 0x49,
	0x12, 0x3d, 0xc7, 0x8b, 0x6c, 0x44, 0xf3, 0xcb, 0x3b, 0xec, 0x0e, 0xfc, 0x55, 0xa1, 0xaa, 0x5f,
	0xc2, 0x90, 0x3b, 0x49, 0x90, 0xf7, 0x49, 0x17, 0x61, 0x2b, 0x29, 0x57, 0x0a, 0xc1, 0x16, 0x31,
	0x00, 0x6d, 0x3f, 0xda, 0xb2, 0xb7, 0x91, 0xb5, 0x03, 0xd0, 0x36, 0xa5, 0x1f, 0xb2, 0x76, 0x66,
	0xed, 0x00, 0xe6, 0x2f, 0x30, 0x1b, 0x0f, 0xa4, 0x4b, 0xad, 0x75, 0xa9, 0xb6, 0x6d, 0x73, 0xd6,
	0x6f, 0x68, 0x87, 0x35, 0xd8, 0xb6, 0xcf, 0x60, 0x7d, 0xfd, 0x19, 0x9b, 0x2d, 0xbe, 0x52, 0xd7,
	0xe9, 0x9f, 0xf8, 0xf2, 0xcd, 0x3a, 0x97, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8e, 0x27,
	0x9a, 0xe3, 0x01, 0x00, 0x00,
}