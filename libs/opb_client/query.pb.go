// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package opb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryOwnerRequest is the request type for the Query/Owner RPC method
type QueryOwnerRequest struct {
	DenomId string `protobuf:"bytes,1,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// pagination defines an optional pagination for the request.
	Pagination           *PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryOwnerRequest) Reset()         { *m = QueryOwnerRequest{} }
func (m *QueryOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOwnerRequest) ProtoMessage()    {}
func (*QueryOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *QueryOwnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOwnerRequest.Unmarshal(m, b)
}
func (m *QueryOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOwnerRequest.Marshal(b, m, deterministic)
}
func (m *QueryOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOwnerRequest.Merge(m, src)
}
func (m *QueryOwnerRequest) XXX_Size() int {
	return xxx_messageInfo_QueryOwnerRequest.Size(m)
}
func (m *QueryOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOwnerRequest proto.InternalMessageInfo

func (m *QueryOwnerRequest) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

func (m *QueryOwnerRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *QueryOwnerRequest) GetPagination() *PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryOwnerResponse is the response type for the Query/Owner RPC method
type QueryOwnerResponse struct {
	Owner                *Owner        `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Pagination           *PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueryOwnerResponse) Reset()         { *m = QueryOwnerResponse{} }
func (m *QueryOwnerResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOwnerResponse) ProtoMessage()    {}
func (*QueryOwnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *QueryOwnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOwnerResponse.Unmarshal(m, b)
}
func (m *QueryOwnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOwnerResponse.Marshal(b, m, deterministic)
}
func (m *QueryOwnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOwnerResponse.Merge(m, src)
}
func (m *QueryOwnerResponse) XXX_Size() int {
	return xxx_messageInfo_QueryOwnerResponse.Size(m)
}
func (m *QueryOwnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOwnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOwnerResponse proto.InternalMessageInfo

func (m *QueryOwnerResponse) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *QueryOwnerResponse) GetPagination() *PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDenomRequest is the request type for the Query/Denom RPC method
type QueryDenomRequest struct {
	DenomId              string   `protobuf:"bytes,1,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDenomRequest) Reset()         { *m = QueryDenomRequest{} }
func (m *QueryDenomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomRequest) ProtoMessage()    {}
func (*QueryDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}
func (m *QueryDenomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDenomRequest.Unmarshal(m, b)
}
func (m *QueryDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDenomRequest.Marshal(b, m, deterministic)
}
func (m *QueryDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomRequest.Merge(m, src)
}
func (m *QueryDenomRequest) XXX_Size() int {
	return xxx_messageInfo_QueryDenomRequest.Size(m)
}
func (m *QueryDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomRequest proto.InternalMessageInfo

func (m *QueryDenomRequest) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

// QueryDenomResponse is the response type for the Query/Denom RPC method
type QueryDenomResponse struct {
	Denom                *Denom   `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDenomResponse) Reset()         { *m = QueryDenomResponse{} }
func (m *QueryDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomResponse) ProtoMessage()    {}
func (*QueryDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}
func (m *QueryDenomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDenomResponse.Unmarshal(m, b)
}
func (m *QueryDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDenomResponse.Marshal(b, m, deterministic)
}
func (m *QueryDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomResponse.Merge(m, src)
}
func (m *QueryDenomResponse) XXX_Size() int {
	return xxx_messageInfo_QueryDenomResponse.Size(m)
}
func (m *QueryDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomResponse proto.InternalMessageInfo

func (m *QueryDenomResponse) GetDenom() *Denom {
	if m != nil {
		return m.Denom
	}
	return nil
}

// QueryNFTRequest is the request type for the Query/NFT RPC method
type QueryNFTRequest struct {
	DenomId              string   `protobuf:"bytes,1,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	TokenId              string   `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty" yaml:"token_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryNFTRequest) Reset()         { *m = QueryNFTRequest{} }
func (m *QueryNFTRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNFTRequest) ProtoMessage()    {}
func (*QueryNFTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}
func (m *QueryNFTRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryNFTRequest.Unmarshal(m, b)
}
func (m *QueryNFTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryNFTRequest.Marshal(b, m, deterministic)
}
func (m *QueryNFTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNFTRequest.Merge(m, src)
}
func (m *QueryNFTRequest) XXX_Size() int {
	return xxx_messageInfo_QueryNFTRequest.Size(m)
}
func (m *QueryNFTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNFTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNFTRequest proto.InternalMessageInfo

func (m *QueryNFTRequest) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

func (m *QueryNFTRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

// QueryNFTResponse is the response type for the Query/NFT RPC method
type QueryNFTResponse struct {
	NFT                  *BaseNFT `protobuf:"bytes,1,opt,name=nft,proto3" json:"nft,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryNFTResponse) Reset()         { *m = QueryNFTResponse{} }
func (m *QueryNFTResponse) String() string { return proto.CompactTextString(m) }
func (*QueryNFTResponse) ProtoMessage()    {}
func (*QueryNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}
func (m *QueryNFTResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryNFTResponse.Unmarshal(m, b)
}
func (m *QueryNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryNFTResponse.Marshal(b, m, deterministic)
}
func (m *QueryNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNFTResponse.Merge(m, src)
}
func (m *QueryNFTResponse) XXX_Size() int {
	return xxx_messageInfo_QueryNFTResponse.Size(m)
}
func (m *QueryNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNFTResponse proto.InternalMessageInfo

func (m *QueryNFTResponse) GetNFT() *BaseNFT {
	if m != nil {
		return m.NFT
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryOwnerRequest)(nil), "proto.QueryOwnerRequest")
	proto.RegisterType((*QueryOwnerResponse)(nil), "proto.QueryOwnerResponse")
	proto.RegisterType((*QueryDenomRequest)(nil), "proto.QueryDenomRequest")
	proto.RegisterType((*QueryDenomResponse)(nil), "proto.QueryDenomResponse")
	proto.RegisterType((*QueryNFTRequest)(nil), "proto.QueryNFTRequest")
	proto.RegisterType((*QueryNFTResponse)(nil), "proto.QueryNFTResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x56, 0x52, 0x85, 0x6e, 0xde, 0xc4, 0x5a, 0x57, 0x82, 0xae, 0x07, 0x5a, 0xf9, 0x30, 0xb6,
	0x03, 0xb1, 0x68, 0x2f, 0x08, 0x89, 0x4b, 0x41, 0x93, 0x76, 0x19, 0x10, 0x4d, 0x42, 0xe2, 0x82,
	0x92, 0xc6, 0x0d, 0x66, 0x8b, 0x9d, 0xd6, 0xae, 0xd0, 0x34, 0xed, 0xc2, 0x2b, 0xf0, 0x00, 0x3c,
	0x02, 0x0f, 0xc3, 0xbd, 0x07, 0xc4, 0x13, 0xec, 0x09, 0x90, 0x7f, 0x3b, 0x4a, 0xb2, 0xee, 0xd4,
	0x43, 0x94, 0x5f, 0xff, 0xff, 0xfd, 0xdf, 0xf7, 0xf9, 0xb3, 0xd1, 0xde, 0x62, 0xc5, 0x96, 0xd7,
	0x61, 0xb1, 0x94, 0x5a, 0xe2, 0x00, 0x7e, 0x03, 0x94, 0xc9, 0x4c, 0xda, 0xd6, 0xa0, 0x1b, 0x0b,
	0x21, 0x75, 0xac, 0xb9, 0x14, 0xca, 0xb5, 0x76, 0xc5, 0x5c, 0xbb, 0xb2, 0x53, 0xc4, 0x19, 0x17,
	0x30, 0xb5, 0x1d, 0xf2, 0xcb, 0x43, 0xdd, 0x8f, 0x86, 0xf2, 0xfd, 0x77, 0xc1, 0x96, 0x11, 0x5b,
	0xac, 0x98, 0xd2, 0x38, 0x44, 0x3b, 0x29, 0x13, 0x32, 0xff, 0xc2, 0xd3, 0xbe, 0x37, 0xf2, 0x8e,
	0x77, 0xa7, 0xbd, 0xbb, 0xf5, 0xf0, 0xe0, 0x3a, 0xce, 0xaf, 0x5e, 0x93, 0x72, 0x42, 0xa2, 0x36,
	0x94, 0x67, 0x29, 0x3e, 0x42, 0x81, 0x34, 0xfb, 0x7d, 0x1f, 0xc0, 0x9d, 0xbb, 0xf5, 0x70, 0xdf,
	0x82, 0xa1, 0x4d, 0x22, 0x3b, 0xc6, 0x63, 0x84, 0x2a, 0x07, 0xfd, 0xd6, 0xc8, 0x3b, 0xde, 0x1b,
	0x63, 0xeb, 0x24, 0xfc, 0x10, 0x67, 0xcc, 0xe9, 0x47, 0x35, 0x14, 0xc9, 0x11, 0xae, 0x1b, 0x54,
	0x85, 0x14, 0x8a, 0x61, 0x52, 0x2a, 0x7a, 0x40, 0xb2, 0xef, 0x48, 0x2c, 0xc8, 0xa9, 0x4d, 0x1a,
	0x6a, 0x3e, 0x00, 0x7b, 0x0d, 0x35, 0x4b, 0xd6, 0x90, 0x7b, 0xeb, 0xf2, 0x78, 0x67, 0x8e, 0xb6,
	0x65, 0x1e, 0xe4, 0x95, 0xf3, 0xec, 0x48, 0x2a, 0xcf, 0x00, 0xb8, 0xe7, 0xd9, 0x82, 0xec, 0x88,
	0x2c, 0xd0, 0x01, 0x6c, 0x9e, 0x9f, 0x5e, 0x6c, 0x7b, 0x19, 0x21, 0xda, 0xd1, 0xf2, 0x92, 0x09,
	0x83, 0xf7, 0xef, 0xe3, 0xcb, 0x09, 0x89, 0xda, 0x50, 0x9e, 0xa5, 0xe4, 0x0d, 0xea, 0x54, 0x92,
	0xce, 0xea, 0x09, 0x6a, 0x89, 0xb9, 0x76, 0x46, 0x1f, 0x3b, 0xa3, 0xd3, 0x58, 0xb1, 0xf3, 0xd3,
	0x8b, 0x69, 0xfb, 0xef, 0x7a, 0xd8, 0x32, 0x68, 0x83, 0x19, 0xff, 0xf6, 0x51, 0x00, 0xfb, 0xf8,
	0x13, 0x0a, 0x20, 0x7f, 0xdc, 0x77, 0x0b, 0x1b, 0x0f, 0x6b, 0x70, 0xf8, 0xc0, 0xc4, 0x4a, 0x92,
	0xc3, 0x1f, 0x7f, 0xfe, 0xfd, 0xf4, 0x7b, 0xb8, 0x4b, 0xf9, 0x92, 0xab, 0x5c, 0xa6, 0x54, 0xcc,
	0xb5, 0xf9, 0x14, 0x4e, 0x51, 0x00, 0x21, 0x35, 0x89, 0xeb, 0x37, 0xd4, 0x24, 0x6e, 0xc4, 0x4e,
	0x8e, 0x80, 0x78, 0x84, 0x9f, 0x35, 0x88, 0x21, 0x2d, 0x45, 0x6f, 0xca, 0x00, 0x6f, 0x71, 0x86,
	0xcc, 0xa1, 0xf0, 0x93, 0x3a, 0x53, 0x75, 0x0d, 0x83, 0xa7, 0x1b, 0x7d, 0xc7, 0x4f, 0x81, 0xff,
	0x04, 0x3f, 0xdf, 0x30, 0x5e, 0x63, 0xa7, 0x37, 0x65, 0xf2, 0xb7, 0xd3, 0xc9, 0xe7, 0x97, 0x19,
	0xd7, 0x5f, 0x57, 0x49, 0x38, 0x93, 0x39, 0x4d, 0x78, 0x2c, 0xbe, 0x71, 0x16, 0x73, 0x9a, 0x5c,
	0xc9, 0xd9, 0xe5, 0x8b, 0x99, 0xcc, 0x0b, 0x26, 0x14, 0xbc, 0x46, 0x0a, 0x92, 0x54, 0x16, 0x49,
	0xf2, 0x08, 0xca, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xb5, 0x97, 0x90, 0x01, 0x04,
	0x00, 0x00,
}
