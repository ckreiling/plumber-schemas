// Code generated by protoc-gen-go. DO NOT EDIT.
// source: azure-service-bus.proto

package args

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

type AzureServiceBusConn struct {
	// NOTE: This is an azure-defined env var
	// @gotags: kong:"help='Connection string',env='SERVICEBUS_CONNECTION_STRING',required"
	ConnectionString     string   `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty" kong:"help='Connection string',env='SERVICEBUS_CONNECTION_STRING',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureServiceBusConn) Reset()         { *m = AzureServiceBusConn{} }
func (m *AzureServiceBusConn) String() string { return proto.CompactTextString(m) }
func (*AzureServiceBusConn) ProtoMessage()    {}
func (*AzureServiceBusConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_454428be5f04e702, []int{0}
}

func (m *AzureServiceBusConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureServiceBusConn.Unmarshal(m, b)
}
func (m *AzureServiceBusConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureServiceBusConn.Marshal(b, m, deterministic)
}
func (m *AzureServiceBusConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureServiceBusConn.Merge(m, src)
}
func (m *AzureServiceBusConn) XXX_Size() int {
	return xxx_messageInfo_AzureServiceBusConn.Size(m)
}
func (m *AzureServiceBusConn) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureServiceBusConn.DiscardUnknown(m)
}

var xxx_messageInfo_AzureServiceBusConn proto.InternalMessageInfo

func (m *AzureServiceBusConn) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type AzureServiceBusReadArgs struct {
	// @gotags: kong:"help='Queue name',env='PLUMBER_RELAY_AZURE_QUEUE_NAME',required"
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty" kong:"help='Queue name',env='PLUMBER_RELAY_AZURE_QUEUE_NAME',required"`
	// @gotags: kong:"help='Topic name',env='PLUMBER_RELAY_AZURE_TOPIC_NAME',required"
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic name',env='PLUMBER_RELAY_AZURE_TOPIC_NAME',required"`
	// @gotags: kong:"help='Subscription name',env='PLUMBER_RELAY_AZURE_SUBSCRIPTION',required"
	SubscriptionName     string   `protobuf:"bytes,3,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty" kong:"help='Subscription name',env='PLUMBER_RELAY_AZURE_SUBSCRIPTION',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureServiceBusReadArgs) Reset()         { *m = AzureServiceBusReadArgs{} }
func (m *AzureServiceBusReadArgs) String() string { return proto.CompactTextString(m) }
func (*AzureServiceBusReadArgs) ProtoMessage()    {}
func (*AzureServiceBusReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_454428be5f04e702, []int{1}
}

func (m *AzureServiceBusReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureServiceBusReadArgs.Unmarshal(m, b)
}
func (m *AzureServiceBusReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureServiceBusReadArgs.Marshal(b, m, deterministic)
}
func (m *AzureServiceBusReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureServiceBusReadArgs.Merge(m, src)
}
func (m *AzureServiceBusReadArgs) XXX_Size() int {
	return xxx_messageInfo_AzureServiceBusReadArgs.Size(m)
}
func (m *AzureServiceBusReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureServiceBusReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AzureServiceBusReadArgs proto.InternalMessageInfo

func (m *AzureServiceBusReadArgs) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *AzureServiceBusReadArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *AzureServiceBusReadArgs) GetSubscriptionName() string {
	if m != nil {
		return m.SubscriptionName
	}
	return ""
}

type AzureServiceBusWriteArgs struct {
	// @gotags: kong:"help='Queue name',required"
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty" kong:"help='Queue name',required"`
	// @gotags: kong:"help='Topic name',required"
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic name',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureServiceBusWriteArgs) Reset()         { *m = AzureServiceBusWriteArgs{} }
func (m *AzureServiceBusWriteArgs) String() string { return proto.CompactTextString(m) }
func (*AzureServiceBusWriteArgs) ProtoMessage()    {}
func (*AzureServiceBusWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_454428be5f04e702, []int{2}
}

func (m *AzureServiceBusWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureServiceBusWriteArgs.Unmarshal(m, b)
}
func (m *AzureServiceBusWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureServiceBusWriteArgs.Marshal(b, m, deterministic)
}
func (m *AzureServiceBusWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureServiceBusWriteArgs.Merge(m, src)
}
func (m *AzureServiceBusWriteArgs) XXX_Size() int {
	return xxx_messageInfo_AzureServiceBusWriteArgs.Size(m)
}
func (m *AzureServiceBusWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureServiceBusWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AzureServiceBusWriteArgs proto.InternalMessageInfo

func (m *AzureServiceBusWriteArgs) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *AzureServiceBusWriteArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*AzureServiceBusConn)(nil), "protos.args.AzureServiceBusConn")
	proto.RegisterType((*AzureServiceBusReadArgs)(nil), "protos.args.AzureServiceBusReadArgs")
	proto.RegisterType((*AzureServiceBusWriteArgs)(nil), "protos.args.AzureServiceBusWriteArgs")
}

func init() { proto.RegisterFile("azure-service-bus.proto", fileDescriptor_454428be5f04e702) }

var fileDescriptor_454428be5f04e702 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbb, 0x4b, 0x04, 0x31,
	0x10, 0xc6, 0x39, 0x45, 0xc1, 0xd8, 0xe8, 0x2a, 0xdc, 0x96, 0xb2, 0x95, 0x20, 0xbb, 0x29, 0xac,
	0xc4, 0xea, 0x56, 0xb0, 0xb4, 0xb8, 0x2b, 0x04, 0x1b, 0x49, 0x72, 0x43, 0x36, 0x70, 0x79, 0x38,
	0x93, 0xb1, 0xf0, 0xaf, 0x97, 0x4d, 0x04, 0x65, 0x3b, 0xab, 0xf0, 0x3d, 0x32, 0x3f, 0xf8, 0xc4,
	0x5a, 0x7d, 0x31, 0x42, 0x4f, 0x80, 0x9f, 0xce, 0x40, 0xaf, 0x99, 0x86, 0x84, 0x31, 0xc7, 0xe6,
	0xbc, 0x3c, 0x34, 0x28, 0xb4, 0xd4, 0x8d, 0xe2, 0x6a, 0x33, 0xf7, 0x76, 0xb5, 0x36, 0x32, 0x3d,
	0xc5, 0x10, 0x9a, 0x3b, 0x71, 0x69, 0x62, 0x08, 0x60, 0xb2, 0x8b, 0xe1, 0x9d, 0x32, 0xba, 0x60,
	0xdb, 0xd5, 0xcd, 0xea, 0xf6, 0x6c, 0x7b, 0xf1, 0x1b, 0xec, 0x8a, 0xdf, 0xa1, 0x58, 0x2f, 0x6e,
	0x6c, 0x41, 0xed, 0x37, 0x68, 0xa9, 0xb9, 0x16, 0x27, 0x1f, 0x0c, 0x0c, 0x3f, 0x7f, 0xab, 0x98,
	0xdd, 0x1c, 0x93, 0x33, 0xed, 0x51, 0x75, 0x8b, 0x98, 0x99, 0xc4, 0x9a, 0x0c, 0xba, 0x54, 0xa8,
	0x41, 0x79, 0x68, 0x8f, 0x2b, 0xf3, 0x6f, 0xf0, 0xa2, 0x3c, 0x74, 0xcf, 0xa2, 0x5d, 0x30, 0x5f,
	0xd1, 0x65, 0xf8, 0x2f, 0x74, 0x7c, 0x7c, 0x7b, 0xb0, 0x2e, 0x4f, 0xac, 0x07, 0x13, 0xbd, 0xd4,
	0x2a, 0x9b, 0xc9, 0x44, 0x4c, 0x32, 0x1d, 0xd8, 0x6b, 0xc0, 0x9e, 0xcc, 0x04, 0x5e, 0x91, 0xd4,
	0xec, 0x0e, 0x7b, 0x69, 0xa3, 0xac, 0xe3, 0xc9, 0x79, 0x3c, 0x7d, 0x5a, 0xc4, 0xfd, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x13, 0x06, 0xc5, 0x6b, 0x01, 0x00, 0x00,
}
