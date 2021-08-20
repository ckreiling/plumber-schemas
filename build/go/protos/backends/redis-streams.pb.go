// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redis-streams.proto

package backends

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

type CreateConsumerConfig_OffsetStart int32

const (
	CreateConsumerConfig_Latest CreateConsumerConfig_OffsetStart = 0
	CreateConsumerConfig_Oldest CreateConsumerConfig_OffsetStart = 1
)

var CreateConsumerConfig_OffsetStart_name = map[int32]string{
	0: "Latest",
	1: "Oldest",
}

var CreateConsumerConfig_OffsetStart_value = map[string]int32{
	"Latest": 0,
	"Oldest": 1,
}

func (x CreateConsumerConfig_OffsetStart) String() string {
	return proto.EnumName(CreateConsumerConfig_OffsetStart_name, int32(x))
}

func (CreateConsumerConfig_OffsetStart) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55f03e487ac12019, []int{1, 0}
}

type RedisStreams struct {
	// Required
	// Database (0-16)
	Database uint32 `protobuf:"varint,1,opt,name=database,proto3" json:"database,omitempty"`
	// Required
	// Stream(s) to read from/write to
	Stream []string `protobuf:"bytes,2,rep,name=stream,proto3" json:"stream,omitempty"`
	// Optional for Writes
	// Ignored for reads
	// Default is "*", which means redis will auto generate it
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Required for writes
	// Ignored for reads
	Key           string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	ConsumerGroup string `protobuf:"bytes,5,opt,name=consumer_group,json=consumerGroup,proto3" json:"consumer_group,omitempty"`
	ConsumerName  string `protobuf:"bytes,6,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	// Optional for reads
	// Ignored for writes
	// Specify if you the user needs to create/recreate a consumer group
	CreateConsumerConfig *CreateConsumerConfig `protobuf:"bytes,7,opt,name=create_consumer_config,json=createConsumerConfig,proto3" json:"create_consumer_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RedisStreams) Reset()         { *m = RedisStreams{} }
func (m *RedisStreams) String() string { return proto.CompactTextString(m) }
func (*RedisStreams) ProtoMessage()    {}
func (*RedisStreams) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f03e487ac12019, []int{0}
}

func (m *RedisStreams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisStreams.Unmarshal(m, b)
}
func (m *RedisStreams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisStreams.Marshal(b, m, deterministic)
}
func (m *RedisStreams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisStreams.Merge(m, src)
}
func (m *RedisStreams) XXX_Size() int {
	return xxx_messageInfo_RedisStreams.Size(m)
}
func (m *RedisStreams) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisStreams.DiscardUnknown(m)
}

var xxx_messageInfo_RedisStreams proto.InternalMessageInfo

func (m *RedisStreams) GetDatabase() uint32 {
	if m != nil {
		return m.Database
	}
	return 0
}

func (m *RedisStreams) GetStream() []string {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *RedisStreams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RedisStreams) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RedisStreams) GetConsumerGroup() string {
	if m != nil {
		return m.ConsumerGroup
	}
	return ""
}

func (m *RedisStreams) GetConsumerName() string {
	if m != nil {
		return m.ConsumerName
	}
	return ""
}

func (m *RedisStreams) GetCreateConsumerConfig() *CreateConsumerConfig {
	if m != nil {
		return m.CreateConsumerConfig
	}
	return nil
}

type CreateConsumerConfig struct {
	// Create the streams if we're creating a new consumer group
	CreateStreams bool `protobuf:"varint,3,opt,name=create_streams,json=createStreams,proto3" json:"create_streams,omitempty"`
	// Recreate this consumer group if it doesn't exist
	RecreateConsumerGroup bool `protobuf:"varint,4,opt,name=recreate_consumer_group,json=recreateConsumerGroup,proto3" json:"recreate_consumer_group,omitempty"`
	// Required
	OffsetStart          CreateConsumerConfig_OffsetStart `protobuf:"varint,5,opt,name=offset_start,json=offsetStart,proto3,enum=protos.backends.CreateConsumerConfig_OffsetStart" json:"offset_start,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateConsumerConfig) Reset()         { *m = CreateConsumerConfig{} }
func (m *CreateConsumerConfig) String() string { return proto.CompactTextString(m) }
func (*CreateConsumerConfig) ProtoMessage()    {}
func (*CreateConsumerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f03e487ac12019, []int{1}
}

func (m *CreateConsumerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsumerConfig.Unmarshal(m, b)
}
func (m *CreateConsumerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsumerConfig.Marshal(b, m, deterministic)
}
func (m *CreateConsumerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsumerConfig.Merge(m, src)
}
func (m *CreateConsumerConfig) XXX_Size() int {
	return xxx_messageInfo_CreateConsumerConfig.Size(m)
}
func (m *CreateConsumerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsumerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsumerConfig proto.InternalMessageInfo

func (m *CreateConsumerConfig) GetCreateStreams() bool {
	if m != nil {
		return m.CreateStreams
	}
	return false
}

func (m *CreateConsumerConfig) GetRecreateConsumerGroup() bool {
	if m != nil {
		return m.RecreateConsumerGroup
	}
	return false
}

func (m *CreateConsumerConfig) GetOffsetStart() CreateConsumerConfig_OffsetStart {
	if m != nil {
		return m.OffsetStart
	}
	return CreateConsumerConfig_Latest
}

func init() {
	proto.RegisterEnum("protos.backends.CreateConsumerConfig_OffsetStart", CreateConsumerConfig_OffsetStart_name, CreateConsumerConfig_OffsetStart_value)
	proto.RegisterType((*RedisStreams)(nil), "protos.backends.RedisStreams")
	proto.RegisterType((*CreateConsumerConfig)(nil), "protos.backends.CreateConsumerConfig")
}

func init() { proto.RegisterFile("redis-streams.proto", fileDescriptor_55f03e487ac12019) }

var fileDescriptor_55f03e487ac12019 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0x87, 0x6b, 0x43, 0x5d, 0x18, 0xfe, 0x14, 0x6d, 0x29, 0xb5, 0x7a, 0xb2, 0xa8, 0x90, 0x7c,
	0xc1, 0x56, 0xa9, 0xd4, 0x5b, 0x55, 0x29, 0x3e, 0xe4, 0x12, 0x05, 0xc9, 0xe4, 0x94, 0x1c, 0xac,
	0xf5, 0x7a, 0x30, 0x16, 0xd8, 0x6b, 0xed, 0xae, 0x0f, 0x79, 0x88, 0xbc, 0x6c, 0x9e, 0x20, 0xf2,
	0x1a, 0x93, 0x04, 0x71, 0xc8, 0x89, 0x99, 0x6f, 0xbe, 0xc5, 0xf3, 0x1b, 0xf8, 0x26, 0x30, 0xc9,
	0xe4, 0x52, 0x2a, 0x81, 0x34, 0x97, 0x5e, 0x29, 0xb8, 0xe2, 0xe4, 0xab, 0xfe, 0x91, 0x5e, 0x4c,
	0xd9, 0x1e, 0x8b, 0x44, 0xce, 0x9f, 0x4c, 0x18, 0x86, 0xb5, 0xb8, 0x69, 0x3c, 0xf2, 0x13, 0x7a,
	0x09, 0x55, 0x34, 0xa6, 0x12, 0x6d, 0xc3, 0x31, 0xdc, 0x51, 0x78, 0xea, 0xc9, 0x0c, 0xac, 0xe6,
	0xef, 0x6c, 0xd3, 0xe9, 0xb8, 0xfd, 0xf0, 0xd8, 0x91, 0x31, 0x98, 0x59, 0x62, 0x77, 0x1c, 0xc3,
	0xed, 0x87, 0x66, 0x96, 0x90, 0x09, 0x74, 0xf6, 0xf8, 0x68, 0x77, 0x35, 0xa8, 0x4b, 0xb2, 0x80,
	0x31, 0xe3, 0x85, 0xac, 0x72, 0x14, 0x51, 0x2a, 0x78, 0x55, 0xda, 0x9f, 0xf5, 0x70, 0xd4, 0xd2,
	0xeb, 0x1a, 0x92, 0x5f, 0x70, 0x02, 0x51, 0x41, 0x73, 0xb4, 0x2d, 0x6d, 0x0d, 0x5b, 0x78, 0x4b,
	0x73, 0x24, 0x0f, 0x30, 0x63, 0x02, 0xa9, 0xc2, 0xe8, 0xe4, 0x32, 0x5e, 0x6c, 0xb3, 0xd4, 0xfe,
	0xe2, 0x18, 0xee, 0x60, 0xb5, 0xf0, 0xce, 0x42, 0x7a, 0x81, 0xd6, 0x83, 0xa3, 0x1d, 0x68, 0x39,
	0x9c, 0xb2, 0x0b, 0x74, 0xfe, 0x6c, 0xc0, 0xf4, 0x92, 0xae, 0x13, 0x34, 0x5f, 0x3d, 0x5e, 0x54,
	0xe7, 0xed, 0x85, 0xa3, 0x86, 0xb6, 0xe7, 0xfb, 0x0b, 0x3f, 0x04, 0x9e, 0xaf, 0xd7, 0x24, 0xee,
	0x6a, 0xff, 0x7b, 0x3b, 0x0e, 0xde, 0x25, 0xbf, 0x83, 0x21, 0xdf, 0x6e, 0x25, 0xaa, 0x48, 0x2a,
	0x2a, 0x94, 0x3e, 0xcf, 0x78, 0xf5, 0xfb, 0x43, 0x51, 0xbc, 0xb5, 0x7e, 0xb9, 0xa9, 0x1f, 0x86,
	0x03, 0xfe, 0xda, 0xcc, 0x17, 0x30, 0x78, 0x33, 0x23, 0x00, 0xd6, 0x0d, 0x55, 0x28, 0xd5, 0xe4,
	0x53, 0x5d, 0xaf, 0x0f, 0x49, 0x5d, 0x1b, 0x57, 0xff, 0xef, 0xff, 0xa5, 0x99, 0xda, 0x55, 0xb1,
	0xc7, 0x78, 0xee, 0xc7, 0x54, 0xb1, 0x1d, 0xe3, 0xa2, 0xf4, 0xcb, 0x43, 0x95, 0xc7, 0x28, 0x96,
	0x92, 0xed, 0x30, 0xa7, 0xd2, 0x8f, 0xab, 0xec, 0x90, 0xf8, 0x29, 0xf7, 0x9b, 0xad, 0xfc, 0x76,
	0xab, 0xd8, 0xd2, 0xe0, 0xcf, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x6e, 0xd3, 0x62, 0x74,
	0x02, 0x00, 0x00,
}
