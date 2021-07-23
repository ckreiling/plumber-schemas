// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay.proto

package protos

import (
	fmt "fmt"
	args "github.com/batchcorp/plumber-schemas/build/go/protos/args"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
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

type Relay struct {
	// Required
	BatchCollectionToken string `protobuf:"bytes,1,opt,name=batch_collection_token,json=batchCollectionToken,proto3" json:"batch_collection_token,omitempty"`
	// Optional; how many messages to send in a single batch (default: 1000)
	BatchSize int32 `protobuf:"varint,2,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// Optional; how many times plumber will try re-sending a batch (default: 3)
	BatchMaxRetry int32  `protobuf:"varint,3,opt,name=batch_max_retry,json=batchMaxRetry,proto3" json:"batch_max_retry,omitempty"`
	ConnectionId  string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Optional; where to send events to (default: grpc-collector.batch.sh:9000)
	BatchshGrpcAddress string `protobuf:"bytes,5,opt,name=batchsh_grpc_address,json=batchshGrpcAddress,proto3" json:"batchsh_grpc_address,omitempty"`
	// Optional; whether to use TLS for gRPC (default: true)
	BatchshGrpcDisableTls bool `protobuf:"varint,6,opt,name=batchsh_grpc_disable_tls,json=batchshGrpcDisableTls,proto3" json:"batchsh_grpc_disable_tls,omitempty"`
	// Optional: how long to wait before giving up talking to the gRPC collector (default: 10s)
	BatchshGrpcTimeout bool `protobuf:"varint,7,opt,name=batchsh_grpc_timeout,json=batchshGrpcTimeout,proto3" json:"batchsh_grpc_timeout,omitempty"`
	// Set appropriate args based on what connection is specified
	// ie. If connection_id is for kafka - specify Kafka args
	//
	// Types that are valid to be assigned to Args:
	//	*Relay_Kafka
	Args                 isRelay_Args `protobuf_oneof:"Args"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Relay) Reset()         { *m = Relay{} }
func (m *Relay) String() string { return proto.CompactTextString(m) }
func (*Relay) ProtoMessage()    {}
func (*Relay) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{0}
}

func (m *Relay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relay.Unmarshal(m, b)
}
func (m *Relay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relay.Marshal(b, m, deterministic)
}
func (m *Relay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relay.Merge(m, src)
}
func (m *Relay) XXX_Size() int {
	return xxx_messageInfo_Relay.Size(m)
}
func (m *Relay) XXX_DiscardUnknown() {
	xxx_messageInfo_Relay.DiscardUnknown(m)
}

var xxx_messageInfo_Relay proto.InternalMessageInfo

func (m *Relay) GetBatchCollectionToken() string {
	if m != nil {
		return m.BatchCollectionToken
	}
	return ""
}

func (m *Relay) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *Relay) GetBatchMaxRetry() int32 {
	if m != nil {
		return m.BatchMaxRetry
	}
	return 0
}

func (m *Relay) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Relay) GetBatchshGrpcAddress() string {
	if m != nil {
		return m.BatchshGrpcAddress
	}
	return ""
}

func (m *Relay) GetBatchshGrpcDisableTls() bool {
	if m != nil {
		return m.BatchshGrpcDisableTls
	}
	return false
}

func (m *Relay) GetBatchshGrpcTimeout() bool {
	if m != nil {
		return m.BatchshGrpcTimeout
	}
	return false
}

type isRelay_Args interface {
	isRelay_Args()
}

type Relay_Kafka struct {
	Kafka *args.Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

func (*Relay_Kafka) isRelay_Args() {}

func (m *Relay) GetArgs() isRelay_Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Relay) GetKafka() *args.Kafka {
	if x, ok := m.GetArgs().(*Relay_Kafka); ok {
		return x.Kafka
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Relay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Relay_Kafka)(nil),
	}
}

type GetAllRelaysRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllRelaysRequest) Reset()         { *m = GetAllRelaysRequest{} }
func (m *GetAllRelaysRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRelaysRequest) ProtoMessage()    {}
func (*GetAllRelaysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{1}
}

func (m *GetAllRelaysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRelaysRequest.Unmarshal(m, b)
}
func (m *GetAllRelaysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRelaysRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRelaysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRelaysRequest.Merge(m, src)
}
func (m *GetAllRelaysRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRelaysRequest.Size(m)
}
func (m *GetAllRelaysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRelaysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRelaysRequest proto.InternalMessageInfo

func (m *GetAllRelaysRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllRelaysResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Will be set as empty []Relay if no relays are configured
	Relays               []*Relay `protobuf:"bytes,1,rep,name=relays,proto3" json:"relays,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRelaysResponse) Reset()         { *m = GetAllRelaysResponse{} }
func (m *GetAllRelaysResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllRelaysResponse) ProtoMessage()    {}
func (*GetAllRelaysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{2}
}

func (m *GetAllRelaysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRelaysResponse.Unmarshal(m, b)
}
func (m *GetAllRelaysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRelaysResponse.Marshal(b, m, deterministic)
}
func (m *GetAllRelaysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRelaysResponse.Merge(m, src)
}
func (m *GetAllRelaysResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllRelaysResponse.Size(m)
}
func (m *GetAllRelaysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRelaysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRelaysResponse proto.InternalMessageInfo

func (m *GetAllRelaysResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetAllRelaysResponse) GetRelays() []*Relay {
	if m != nil {
		return m.Relays
	}
	return nil
}

type GetRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	RelayId              string       `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetRelayRequest) Reset()         { *m = GetRelayRequest{} }
func (m *GetRelayRequest) String() string { return proto.CompactTextString(m) }
func (*GetRelayRequest) ProtoMessage()    {}
func (*GetRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{3}
}

func (m *GetRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRelayRequest.Unmarshal(m, b)
}
func (m *GetRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRelayRequest.Marshal(b, m, deterministic)
}
func (m *GetRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRelayRequest.Merge(m, src)
}
func (m *GetRelayRequest) XXX_Size() int {
	return xxx_messageInfo_GetRelayRequest.Size(m)
}
func (m *GetRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRelayRequest proto.InternalMessageInfo

func (m *GetRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetRelayRequest) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

type GetRelayResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Set only if status is OK
	Relay                *Relay   `protobuf:"bytes,1,opt,name=relay,proto3" json:"relay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRelayResponse) Reset()         { *m = GetRelayResponse{} }
func (m *GetRelayResponse) String() string { return proto.CompactTextString(m) }
func (*GetRelayResponse) ProtoMessage()    {}
func (*GetRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{4}
}

func (m *GetRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRelayResponse.Unmarshal(m, b)
}
func (m *GetRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRelayResponse.Marshal(b, m, deterministic)
}
func (m *GetRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRelayResponse.Merge(m, src)
}
func (m *GetRelayResponse) XXX_Size() int {
	return xxx_messageInfo_GetRelayResponse.Size(m)
}
func (m *GetRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRelayResponse proto.InternalMessageInfo

func (m *GetRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetRelayResponse) GetRelay() *Relay {
	if m != nil {
		return m.Relay
	}
	return nil
}

type CreateRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth  *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Relay *Relay       `protobuf:"bytes,1,opt,name=relay,proto3" json:"relay,omitempty"`
	// Types that are valid to be assigned to Args:
	//	*CreateRelayRequest_Kafka
	Args                 isCreateRelayRequest_Args `protobuf_oneof:"Args"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateRelayRequest) Reset()         { *m = CreateRelayRequest{} }
func (m *CreateRelayRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRelayRequest) ProtoMessage()    {}
func (*CreateRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{5}
}

func (m *CreateRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRelayRequest.Unmarshal(m, b)
}
func (m *CreateRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRelayRequest.Marshal(b, m, deterministic)
}
func (m *CreateRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRelayRequest.Merge(m, src)
}
func (m *CreateRelayRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRelayRequest.Size(m)
}
func (m *CreateRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRelayRequest proto.InternalMessageInfo

func (m *CreateRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateRelayRequest) GetRelay() *Relay {
	if m != nil {
		return m.Relay
	}
	return nil
}

type isCreateRelayRequest_Args interface {
	isCreateRelayRequest_Args()
}

type CreateRelayRequest_Kafka struct {
	Kafka *args.Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

func (*CreateRelayRequest_Kafka) isCreateRelayRequest_Args() {}

func (m *CreateRelayRequest) GetArgs() isCreateRelayRequest_Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CreateRelayRequest) GetKafka() *args.Kafka {
	if x, ok := m.GetArgs().(*CreateRelayRequest_Kafka); ok {
		return x.Kafka
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateRelayRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateRelayRequest_Kafka)(nil),
	}
}

type CreateRelayResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// ID of the created relay entry
	RelayId              string   `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRelayResponse) Reset()         { *m = CreateRelayResponse{} }
func (m *CreateRelayResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRelayResponse) ProtoMessage()    {}
func (*CreateRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{6}
}

func (m *CreateRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRelayResponse.Unmarshal(m, b)
}
func (m *CreateRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRelayResponse.Marshal(b, m, deterministic)
}
func (m *CreateRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRelayResponse.Merge(m, src)
}
func (m *CreateRelayResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRelayResponse.Size(m)
}
func (m *CreateRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRelayResponse proto.InternalMessageInfo

func (m *CreateRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateRelayResponse) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

// WARNING: Any in-progress relay will be interrupted/restarted
type UpdateRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	RelayId              string       `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	Relay                *Relay       `protobuf:"bytes,2,opt,name=relay,proto3" json:"relay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRelayRequest) Reset()         { *m = UpdateRelayRequest{} }
func (m *UpdateRelayRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRelayRequest) ProtoMessage()    {}
func (*UpdateRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{7}
}

func (m *UpdateRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRelayRequest.Unmarshal(m, b)
}
func (m *UpdateRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRelayRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRelayRequest.Merge(m, src)
}
func (m *UpdateRelayRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRelayRequest.Size(m)
}
func (m *UpdateRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRelayRequest proto.InternalMessageInfo

func (m *UpdateRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateRelayRequest) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

func (m *UpdateRelayRequest) GetRelay() *Relay {
	if m != nil {
		return m.Relay
	}
	return nil
}

type UpdateRelayResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRelayResponse) Reset()         { *m = UpdateRelayResponse{} }
func (m *UpdateRelayResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateRelayResponse) ProtoMessage()    {}
func (*UpdateRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{8}
}

func (m *UpdateRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRelayResponse.Unmarshal(m, b)
}
func (m *UpdateRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRelayResponse.Marshal(b, m, deterministic)
}
func (m *UpdateRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRelayResponse.Merge(m, src)
}
func (m *UpdateRelayResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateRelayResponse.Size(m)
}
func (m *UpdateRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRelayResponse proto.InternalMessageInfo

func (m *UpdateRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Resume a paused relay
type ResumeRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	RelayId              string       `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResumeRelayRequest) Reset()         { *m = ResumeRelayRequest{} }
func (m *ResumeRelayRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRelayRequest) ProtoMessage()    {}
func (*ResumeRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{9}
}

func (m *ResumeRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRelayRequest.Unmarshal(m, b)
}
func (m *ResumeRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRelayRequest.Marshal(b, m, deterministic)
}
func (m *ResumeRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRelayRequest.Merge(m, src)
}
func (m *ResumeRelayRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRelayRequest.Size(m)
}
func (m *ResumeRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRelayRequest proto.InternalMessageInfo

func (m *ResumeRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ResumeRelayRequest) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

type ResumeRelayResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResumeRelayResponse) Reset()         { *m = ResumeRelayResponse{} }
func (m *ResumeRelayResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeRelayResponse) ProtoMessage()    {}
func (*ResumeRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{10}
}

func (m *ResumeRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRelayResponse.Unmarshal(m, b)
}
func (m *ResumeRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRelayResponse.Marshal(b, m, deterministic)
}
func (m *ResumeRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRelayResponse.Merge(m, src)
}
func (m *ResumeRelayResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeRelayResponse.Size(m)
}
func (m *ResumeRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRelayResponse proto.InternalMessageInfo

func (m *ResumeRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Temporarily stop/pause a relay
type StopRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	RelayId              string       `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StopRelayRequest) Reset()         { *m = StopRelayRequest{} }
func (m *StopRelayRequest) String() string { return proto.CompactTextString(m) }
func (*StopRelayRequest) ProtoMessage()    {}
func (*StopRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{11}
}

func (m *StopRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRelayRequest.Unmarshal(m, b)
}
func (m *StopRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRelayRequest.Marshal(b, m, deterministic)
}
func (m *StopRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRelayRequest.Merge(m, src)
}
func (m *StopRelayRequest) XXX_Size() int {
	return xxx_messageInfo_StopRelayRequest.Size(m)
}
func (m *StopRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRelayRequest proto.InternalMessageInfo

func (m *StopRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StopRelayRequest) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

type StopRelayResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StopRelayResponse) Reset()         { *m = StopRelayResponse{} }
func (m *StopRelayResponse) String() string { return proto.CompactTextString(m) }
func (*StopRelayResponse) ProtoMessage()    {}
func (*StopRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{12}
}

func (m *StopRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRelayResponse.Unmarshal(m, b)
}
func (m *StopRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRelayResponse.Marshal(b, m, deterministic)
}
func (m *StopRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRelayResponse.Merge(m, src)
}
func (m *StopRelayResponse) XXX_Size() int {
	return xxx_messageInfo_StopRelayResponse.Size(m)
}
func (m *StopRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopRelayResponse proto.InternalMessageInfo

func (m *StopRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteRelayRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	RelayId              string       `protobuf:"bytes,1,opt,name=relay_id,json=relayId,proto3" json:"relay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteRelayRequest) Reset()         { *m = DeleteRelayRequest{} }
func (m *DeleteRelayRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRelayRequest) ProtoMessage()    {}
func (*DeleteRelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{13}
}

func (m *DeleteRelayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRelayRequest.Unmarshal(m, b)
}
func (m *DeleteRelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRelayRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRelayRequest.Merge(m, src)
}
func (m *DeleteRelayRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRelayRequest.Size(m)
}
func (m *DeleteRelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRelayRequest proto.InternalMessageInfo

func (m *DeleteRelayRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteRelayRequest) GetRelayId() string {
	if m != nil {
		return m.RelayId
	}
	return ""
}

type DeleteRelayResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteRelayResponse) Reset()         { *m = DeleteRelayResponse{} }
func (m *DeleteRelayResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRelayResponse) ProtoMessage()    {}
func (*DeleteRelayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f69a7d5a802d584, []int{14}
}

func (m *DeleteRelayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRelayResponse.Unmarshal(m, b)
}
func (m *DeleteRelayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRelayResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRelayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRelayResponse.Merge(m, src)
}
func (m *DeleteRelayResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRelayResponse.Size(m)
}
func (m *DeleteRelayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRelayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRelayResponse proto.InternalMessageInfo

func (m *DeleteRelayResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Relay)(nil), "protos.Relay")
	proto.RegisterType((*GetAllRelaysRequest)(nil), "protos.GetAllRelaysRequest")
	proto.RegisterType((*GetAllRelaysResponse)(nil), "protos.GetAllRelaysResponse")
	proto.RegisterType((*GetRelayRequest)(nil), "protos.GetRelayRequest")
	proto.RegisterType((*GetRelayResponse)(nil), "protos.GetRelayResponse")
	proto.RegisterType((*CreateRelayRequest)(nil), "protos.CreateRelayRequest")
	proto.RegisterType((*CreateRelayResponse)(nil), "protos.CreateRelayResponse")
	proto.RegisterType((*UpdateRelayRequest)(nil), "protos.UpdateRelayRequest")
	proto.RegisterType((*UpdateRelayResponse)(nil), "protos.UpdateRelayResponse")
	proto.RegisterType((*ResumeRelayRequest)(nil), "protos.ResumeRelayRequest")
	proto.RegisterType((*ResumeRelayResponse)(nil), "protos.ResumeRelayResponse")
	proto.RegisterType((*StopRelayRequest)(nil), "protos.StopRelayRequest")
	proto.RegisterType((*StopRelayResponse)(nil), "protos.StopRelayResponse")
	proto.RegisterType((*DeleteRelayRequest)(nil), "protos.DeleteRelayRequest")
	proto.RegisterType((*DeleteRelayResponse)(nil), "protos.DeleteRelayResponse")
}

func init() { proto.RegisterFile("relay.proto", fileDescriptor_9f69a7d5a802d584) }

var fileDescriptor_9f69a7d5a802d584 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0x3f, 0xb7, 0x4d, 0xd2, 0x4e, 0xbe, 0xa8, 0xe9, 0xa6, 0x45, 0x4b, 0x25, 0xa4, 0xc8,
	0x15, 0xc8, 0x42, 0x22, 0x46, 0xa5, 0x82, 0x23, 0x4a, 0x53, 0x14, 0x2a, 0x04, 0x07, 0x27, 0x80,
	0xe0, 0x62, 0xd6, 0xf6, 0x62, 0x5b, 0xb1, 0xbd, 0x66, 0x77, 0x2d, 0xb5, 0x3d, 0xf0, 0x0a, 0x1c,
	0x79, 0x45, 0x9e, 0x81, 0x13, 0xf2, 0xae, 0x4b, 0x62, 0x2a, 0x2a, 0xb0, 0xca, 0x29, 0xca, 0xfc,
	0x67, 0x7e, 0x33, 0xf3, 0x9f, 0x55, 0x02, 0x5d, 0x4e, 0x13, 0x72, 0x3e, 0xca, 0x39, 0x93, 0x0c,
	0xb5, 0xd5, 0x87, 0xd8, 0xef, 0x13, 0x1e, 0x0a, 0x7b, 0x41, 0x3e, 0x2e, 0x88, 0x56, 0xf6, 0x77,
	0x7c, 0x96, 0xa6, 0x2c, 0xb3, 0x49, 0x21, 0xa3, 0x2a, 0x34, 0xa8, 0x42, 0x42, 0x12, 0x59, 0x08,
	0x1d, 0x34, 0xbf, 0xaf, 0x41, 0xcb, 0x29, 0x89, 0xe8, 0x08, 0x6e, 0x79, 0x44, 0xfa, 0x91, 0xeb,
	0xb3, 0x24, 0xa1, 0xbe, 0x8c, 0x59, 0xe6, 0x4a, 0xb6, 0xa0, 0x19, 0x36, 0x86, 0x86, 0xb5, 0xe5,
	0xec, 0x2a, 0x75, 0xf2, 0x53, 0x9c, 0x97, 0x1a, 0xba, 0x03, 0xa0, 0xab, 0x44, 0x7c, 0x41, 0xf1,
	0xda, 0xd0, 0xb0, 0x5a, 0xce, 0x96, 0x8a, 0xcc, 0xe2, 0x0b, 0x8a, 0xee, 0xc1, 0xb6, 0x96, 0x53,
	0x72, 0xe6, 0x72, 0x2a, 0xf9, 0x39, 0x5e, 0x57, 0x39, 0x3d, 0x15, 0x7e, 0x49, 0xce, 0x9c, 0x32,
	0x88, 0x0e, 0xa0, 0xe7, 0xb3, 0x2c, 0xab, 0xda, 0xc6, 0x01, 0xde, 0x50, 0x3d, 0xff, 0x5f, 0x06,
	0x4f, 0x03, 0xf4, 0x10, 0xf4, 0x0c, 0x22, 0x72, 0x43, 0x9e, 0xfb, 0x2e, 0x09, 0x02, 0x4e, 0x85,
	0xc0, 0x2d, 0x95, 0x8b, 0x2a, 0x6d, 0xca, 0x73, 0x7f, 0xac, 0x15, 0xf4, 0x04, 0x70, 0xad, 0x22,
	0x88, 0x05, 0xf1, 0x12, 0xea, 0xca, 0x44, 0xe0, 0xf6, 0xd0, 0xb0, 0x36, 0x9d, 0xbd, 0x95, 0xaa,
	0x13, 0xad, 0xce, 0x13, 0x71, 0xa5, 0x95, 0x8c, 0x53, 0xca, 0x0a, 0x89, 0x3b, 0xaa, 0x68, 0xb5,
	0xd5, 0x5c, 0x2b, 0xe8, 0x3e, 0xb4, 0x94, 0xff, 0x38, 0x18, 0x1a, 0x56, 0xf7, 0x10, 0x69, 0x7f,
	0xc5, 0xa8, 0xbc, 0xcc, 0xe8, 0x45, 0xa9, 0x3c, 0xff, 0xcf, 0xd1, 0x29, 0xc7, 0x6d, 0xd8, 0x18,
	0xf3, 0x50, 0x98, 0x4f, 0x61, 0x30, 0xa5, 0x72, 0x9c, 0x24, 0xea, 0x02, 0xc2, 0xa1, 0x9f, 0x0a,
	0x2a, 0x24, 0xb2, 0x60, 0xa3, 0x3c, 0x1b, 0xfe, 0xf2, 0x4a, 0xa1, 0x06, 0x97, 0x28, 0x7d, 0xbf,
	0xd1, 0xb8, 0x90, 0x91, 0xa3, 0x32, 0xcc, 0x14, 0x76, 0xeb, 0x00, 0x91, 0xb3, 0x4c, 0x50, 0x34,
	0x82, 0xb6, 0xbe, 0x32, 0xfe, 0xd6, 0x51, 0x8c, 0xbd, 0x5f, 0x18, 0x33, 0xa5, 0x3a, 0x55, 0x16,
	0xba, 0x0b, 0x6d, 0xf5, 0xac, 0x04, 0x36, 0x86, 0xeb, 0x56, 0xf7, 0xb0, 0x77, 0x99, 0xae, 0xb8,
	0x4e, 0x25, 0x9a, 0x6f, 0x60, 0x7b, 0x4a, 0xa5, 0x8e, 0xfd, 0xed, 0xac, 0xe8, 0x36, 0x6c, 0x2a,
	0x4c, 0x79, 0x5d, 0xfd, 0xa2, 0x3a, 0xea, 0xfb, 0x69, 0x60, 0x86, 0xd0, 0x5f, 0x72, 0x1b, 0xae,
	0x70, 0x00, 0x2d, 0x85, 0x53, 0xec, 0x2b, 0x1b, 0x68, 0xcd, 0xfc, 0x6a, 0x00, 0x9a, 0x70, 0x4a,
	0x24, 0x6d, 0xb8, 0xc4, 0x9f, 0x74, 0x69, 0xf4, 0x14, 0x3e, 0xc0, 0xa0, 0x36, 0x58, 0x43, 0x17,
	0xae, 0x31, 0xf9, 0x33, 0xa0, 0xd7, 0x79, 0xd0, 0x7c, 0xf5, 0xdf, 0xa3, 0x97, 0xae, 0xac, 0x5d,
	0xe3, 0xfd, 0x33, 0x18, 0xd4, 0xfa, 0x37, 0xdb, 0xd0, 0x7c, 0x07, 0xc8, 0xa1, 0xa2, 0x48, 0x6f,
	0x7e, 0x8d, 0x72, 0xc2, 0x1a, 0xba, 0xe1, 0x84, 0x6f, 0xa1, 0x3f, 0x93, 0x2c, 0xbf, 0xf9, 0xf9,
	0x26, 0xb0, 0xb3, 0x02, 0x6e, 0xee, 0xdf, 0x09, 0x4d, 0xa8, 0xfc, 0x37, 0xfe, 0xd5, 0xd0, 0xcd,
	0x26, 0x3c, 0x7e, 0xfc, 0xfe, 0x28, 0x8c, 0x65, 0x54, 0x78, 0xa5, 0x6e, 0xab, 0x9f, 0x5a, 0x9f,
	0xf1, 0xdc, 0xce, 0x93, 0x22, 0xf5, 0x28, 0x7f, 0x20, 0xfc, 0x88, 0xa6, 0x44, 0xd8, 0x5e, 0x11,
	0x27, 0x81, 0x1d, 0x32, 0x5b, 0xd3, 0x3c, 0xfd, 0x67, 0xf8, 0xe8, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbc, 0xf4, 0x74, 0x05, 0x22, 0x07, 0x00, 0x00,
}
