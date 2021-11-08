// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_errors.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	_ "github.com/batchcorp/plumber-schemas/build/go/protos/opts"
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

type ErrorMessage struct {
	Resource             string   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	ResourceId           string   `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMessage) Reset()         { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()    {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f58175ccf95effe0, []int{0}
}

func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMessage.Unmarshal(m, b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ErrorMessage.Size(m)
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ErrorMessage) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ErrorMessage) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type GetErrorsRequest struct {
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetErrorsRequest) Reset()         { *m = GetErrorsRequest{} }
func (m *GetErrorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetErrorsRequest) ProtoMessage()    {}
func (*GetErrorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f58175ccf95effe0, []int{1}
}

func (m *GetErrorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetErrorsRequest.Unmarshal(m, b)
}
func (m *GetErrorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetErrorsRequest.Marshal(b, m, deterministic)
}
func (m *GetErrorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetErrorsRequest.Merge(m, src)
}
func (m *GetErrorsRequest) XXX_Size() int {
	return xxx_messageInfo_GetErrorsRequest.Size(m)
}
func (m *GetErrorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetErrorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetErrorsRequest proto.InternalMessageInfo

func (m *GetErrorsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetErrorsResponse struct {
	Error                *ErrorMessage `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetErrorsResponse) Reset()         { *m = GetErrorsResponse{} }
func (m *GetErrorsResponse) String() string { return proto.CompactTextString(m) }
func (*GetErrorsResponse) ProtoMessage()    {}
func (*GetErrorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f58175ccf95effe0, []int{2}
}

func (m *GetErrorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetErrorsResponse.Unmarshal(m, b)
}
func (m *GetErrorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetErrorsResponse.Marshal(b, m, deterministic)
}
func (m *GetErrorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetErrorsResponse.Merge(m, src)
}
func (m *GetErrorsResponse) XXX_Size() int {
	return xxx_messageInfo_GetErrorsResponse.Size(m)
}
func (m *GetErrorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetErrorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetErrorsResponse proto.InternalMessageInfo

func (m *GetErrorsResponse) GetError() *ErrorMessage {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorMessage)(nil), "protos.ErrorMessage")
	proto.RegisterType((*GetErrorsRequest)(nil), "protos.GetErrorsRequest")
	proto.RegisterType((*GetErrorsResponse)(nil), "protos.GetErrorsResponse")
}

func init() { proto.RegisterFile("ps_errors.proto", fileDescriptor_f58175ccf95effe0) }

var fileDescriptor_f58175ccf95effe0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0xa5, 0x5a, 0x45, 0xb3, 0x82, 0x1a, 0xf7, 0x10, 0xea, 0xc1, 0xa5, 0xa7, 0x22, 0xd8, 0x80,
	0x8a, 0x27, 0x41, 0x14, 0x44, 0x3c, 0xe8, 0xa1, 0x47, 0x2f, 0xa5, 0x69, 0x87, 0xb6, 0xb0, 0xdd,
	0x89, 0x99, 0xe4, 0xe0, 0x57, 0xf8, 0xcb, 0x92, 0x64, 0x2b, 0x7b, 0x9a, 0x37, 0xf3, 0xe6, 0xbd,
	0x79, 0x09, 0x3b, 0xd5, 0x54, 0x83, 0x31, 0x68, 0xa8, 0xd4, 0x06, 0x2d, 0xf2, 0xc3, 0x50, 0x28,
	0x13, 0xa8, 0x2d, 0x49, 0x4d, 0xb5, 0xaf, 0x35, 0xd9, 0xc6, 0x6e, 0x37, 0xb2, 0xcb, 0x16, 0xa7,
	0x09, 0x37, 0x9e, 0x8b, 0xa8, 0x6e, 0x9c, 0x1d, 0x22, 0x99, 0x3b, 0x76, 0xf2, 0xea, 0xed, 0x3e,
	0x80, 0xa8, 0xe9, 0x81, 0x67, 0xec, 0xc8, 0x00, 0xa1, 0x33, 0x2d, 0x88, 0x64, 0x95, 0x14, 0xc7,
	0xd5, 0x7f, 0xcf, 0xaf, 0xd8, 0x62, 0xc6, 0xf5, 0xd8, 0x89, 0xbd, 0x40, 0xb3, 0x79, 0xf4, 0xde,
	0xf1, 0x25, 0x3b, 0x08, 0xd9, 0xc4, 0x7e, 0xa0, 0x62, 0xc3, 0x39, 0x4b, 0x15, 0x76, 0x3f, 0x22,
	0x0d, 0xc3, 0x80, 0xf3, 0x47, 0x76, 0xf6, 0x06, 0x36, 0x5c, 0xa6, 0x0a, 0xbe, 0x1d, 0x90, 0xe5,
	0x05, 0x4b, 0x7d, 0x30, 0xf1, 0xfb, 0xb9, 0x4a, 0x8a, 0xc5, 0xed, 0x45, 0x4c, 0x48, 0x65, 0x0c,
	0x5d, 0x3e, 0x3b, 0x3b, 0x54, 0x61, 0x23, 0x7f, 0x62, 0xe7, 0x3b, 0x6a, 0xd2, 0xb8, 0x21, 0xe0,
	0xd7, 0xf3, 0xf1, 0x24, 0xc8, 0x97, 0xb3, 0x7c, 0xf7, 0x79, 0xdb, 0x48, 0x2f, 0x0f, 0x5f, 0xf7,
	0xfd, 0x68, 0x07, 0xa7, 0xbc, 0xb9, 0x54, 0x8d, 0x6d, 0x87, 0x16, 0x8d, 0x96, 0x7a, 0xed, 0x26,
	0x05, 0xe6, 0x86, 0xda, 0x01, 0xa6, 0x86, 0xa4, 0x72, 0xe3, 0xba, 0x93, 0x3d, 0xca, 0xe8, 0xa5,
	0xe2, 0x67, 0xdf, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x07, 0xc2, 0x2d, 0x86, 0x01, 0x00,
	0x00,
}
