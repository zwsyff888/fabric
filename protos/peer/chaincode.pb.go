// Code generated by protoc-gen-go.
// source: peer/chaincode.proto
// DO NOT EDIT!

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	peer/chaincode.proto
	peer/chaincode_proposal.proto
	peer/chaincode_transaction.proto
	peer/chaincodeevent.proto
	peer/configuration.proto
	peer/events.proto
	peer/fabric.proto
	peer/fabric_message.proto
	peer/fabric_proposal.proto
	peer/fabric_proposal_response.proto
	peer/fabric_service.proto
	peer/fabric_transaction.proto
	peer/server_admin.proto
	peer/server_supervise.proto

It has these top-level messages:
	ChaincodeID
	ChaincodeInput
	ChaincodeSpec
	ChaincodeDeploymentSpec
	ChaincodeInvocationSpec
	ChaincodeProposalContext
	ChaincodeMessage
	PutStateInfo
	RangeQueryState
	RangeQueryStateNext
	RangeQueryStateClose
	RangeQueryStateKeyValue
	RangeQueryStateResponse
	ChaincodeHeaderExtension
	ChaincodeProposalPayload
	ChaincodeAction
	ChaincodeActionPayload
	ChaincodeEndorsedAction
	ChaincodeEvent
	AnchorPeers
	AnchorPeer
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	Event
	PeerAddress
	PeerID
	PeerEndpoint
	PeersMessage
	PeersAddresses
	BlockchainInfo
	Message
	SignedProposal
	Proposal
	ProposalResponse
	Response
	ProposalResponsePayload
	Endorsement
	SignedTransaction
	InvalidTransaction
	Transaction
	TransactionAction
	ServerStatus
	LogLevelRequest
	LogLevelResponse
	PeerReply
	ConnectPeer
	PeerInfo
*/
package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED               ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER                ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED              ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                    ChaincodeMessage_Type = 3
	ChaincodeMessage_READY                   ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION             ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED               ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR                   ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE               ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE               ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE               ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE        ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE                ChaincodeMessage_Type = 13
	ChaincodeMessage_RANGE_QUERY_STATE       ChaincodeMessage_Type = 14
	ChaincodeMessage_RANGE_QUERY_STATE_NEXT  ChaincodeMessage_Type = 15
	ChaincodeMessage_RANGE_QUERY_STATE_CLOSE ChaincodeMessage_Type = 16
	ChaincodeMessage_KEEPALIVE               ChaincodeMessage_Type = 17
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "RANGE_QUERY_STATE",
	15: "RANGE_QUERY_STATE_NEXT",
	16: "RANGE_QUERY_STATE_CLOSE",
	17: "KEEPALIVE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":               0,
	"REGISTER":                1,
	"REGISTERED":              2,
	"INIT":                    3,
	"READY":                   4,
	"TRANSACTION":             5,
	"COMPLETED":               6,
	"ERROR":                   7,
	"GET_STATE":               8,
	"PUT_STATE":               9,
	"DEL_STATE":               10,
	"INVOKE_CHAINCODE":        11,
	"RESPONSE":                13,
	"RANGE_QUERY_STATE":       14,
	"RANGE_QUERY_STATE_NEXT":  15,
	"RANGE_QUERY_STATE_CLOSE": 16,
	"KEEPALIVE":               17,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeID *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetChaincodeID() *ChaincodeID {
	if m != nil {
		return m.ChaincodeID
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincodeSpec" json:"chaincodeSpec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effectiveDate" json:"effectiveDate,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=codePackage,proto3" json:"codePackage,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"execEnv,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincodeSpec" json:"chaincodeSpec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=idGenerationAlg" json:"idGenerationAlg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// ChaincodeProposalContext contains proposal data that we send to the chaincode
// container shim and allow the chaincode to access through the shim interface.
type ChaincodeProposalContext struct {
	// Creator corresponds to SignatureHeader.Creator
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Transient corresponds to ChaincodeProposalPayload.Transient
	// TODO: The transient field is supposed to carry application-specific
	// data. They might be realted to access-control, encryption and so on.
	// To simply access to this data, replacing bytes with a map
	// is the next step to be carried.
	Transient []byte `protobuf:"bytes,2,opt,name=transient,proto3" json:"transient,omitempty"`
}

func (m *ChaincodeProposalContext) Reset()                    { *m = ChaincodeProposalContext{} }
func (m *ChaincodeProposalContext) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeProposalContext) ProtoMessage()               {}
func (*ChaincodeProposalContext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChaincodeMessage struct {
	Type            ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp       *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload         []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid            string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	ProposalContext *ChaincodeProposalContext   `protobuf:"bytes,5,opt,name=proposalContext" json:"proposalContext,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincodeEvent" json:"chaincodeEvent,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetProposalContext() *ChaincodeProposalContext {
	if m != nil {
		return m.ProposalContext
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type RangeQueryState struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *RangeQueryState) Reset()                    { *m = RangeQueryState{} }
func (m *RangeQueryState) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryState) ProtoMessage()               {}
func (*RangeQueryState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type RangeQueryStateNext struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateNext) Reset()                    { *m = RangeQueryStateNext{} }
func (m *RangeQueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateNext) ProtoMessage()               {}
func (*RangeQueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type RangeQueryStateClose struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateClose) Reset()                    { *m = RangeQueryStateClose{} }
func (m *RangeQueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateClose) ProtoMessage()               {}
func (*RangeQueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

type RangeQueryStateKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RangeQueryStateKeyValue) Reset()                    { *m = RangeQueryStateKeyValue{} }
func (m *RangeQueryStateKeyValue) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateKeyValue) ProtoMessage()               {}
func (*RangeQueryStateKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type RangeQueryStateResponse struct {
	KeysAndValues []*RangeQueryStateKeyValue `protobuf:"bytes,1,rep,name=keysAndValues" json:"keysAndValues,omitempty"`
	HasMore       bool                       `protobuf:"varint,2,opt,name=hasMore" json:"hasMore,omitempty"`
	ID            string                     `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *RangeQueryStateResponse) Reset()                    { *m = RangeQueryStateResponse{} }
func (m *RangeQueryStateResponse) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryStateResponse) ProtoMessage()               {}
func (*RangeQueryStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *RangeQueryStateResponse) GetKeysAndValues() []*RangeQueryStateKeyValue {
	if m != nil {
		return m.KeysAndValues
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*ChaincodeProposalContext)(nil), "protos.ChaincodeProposalContext")
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*RangeQueryState)(nil), "protos.RangeQueryState")
	proto.RegisterType((*RangeQueryStateNext)(nil), "protos.RangeQueryStateNext")
	proto.RegisterType((*RangeQueryStateClose)(nil), "protos.RangeQueryStateClose")
	proto.RegisterType((*RangeQueryStateKeyValue)(nil), "protos.RangeQueryStateKeyValue")
	proto.RegisterType((*RangeQueryStateResponse)(nil), "protos.RangeQueryStateResponse")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x0e, 0x75, 0xb0, 0xa5, 0xd1, 0xc1, 0x9b, 0x8d, 0xe2, 0xe8, 0xd7, 0xdf, 0x22, 0x06, 0xd1,
	0x16, 0x6e, 0x51, 0x48, 0xad, 0x9a, 0x14, 0x05, 0x0a, 0x04, 0x65, 0xc8, 0x8d, 0xca, 0x58, 0xa6,
	0x94, 0x15, 0x6d, 0xc4, 0xbd, 0x31, 0x68, 0x6a, 0x25, 0x13, 0x96, 0xb9, 0x04, 0xb9, 0x12, 0xac,
	0xbb, 0x5e, 0xf7, 0x35, 0xfa, 0x16, 0x7d, 0x99, 0xbe, 0x49, 0x51, 0x2c, 0x0f, 0xb2, 0x0e, 0x36,
	0x10, 0xa0, 0x57, 0xda, 0xd9, 0xf9, 0xbe, 0xd9, 0x99, 0x6f, 0x76, 0x56, 0x84, 0x46, 0xc0, 0x58,
	0xd8, 0x71, 0xaf, 0x1d, 0xcf, 0x77, 0xf9, 0x98, 0xb5, 0x83, 0x90, 0x0b, 0x8e, 0xf7, 0xe2, 0x9f,
	0xa8, 0xf5, 0xbf, 0x4d, 0x2f, 0x5b, 0x30, 0x5f, 0x24, 0x90, 0xd6, 0xcb, 0x29, 0xe7, 0xd3, 0x19,
	0xeb, 0xc4, 0xd6, 0xd5, 0x7c, 0xd2, 0x11, 0xde, 0x2d, 0x8b, 0x84, 0x73, 0x1b, 0x24, 0x00, 0xf5,
	0x35, 0x54, 0xf4, 0x8c, 0x68, 0x1a, 0x18, 0x43, 0x21, 0x70, 0xc4, 0x75, 0x53, 0x39, 0x52, 0x8e,
	0xcb, 0x34, 0x5e, 0xcb, 0x3d, 0xdf, 0xb9, 0x65, 0xcd, 0x5c, 0xb2, 0x27, 0xd7, 0xea, 0x17, 0x50,
	0xbf, 0xa7, 0xf9, 0xc1, 0x5c, 0x48, 0x94, 0x13, 0x4e, 0xa3, 0xa6, 0x72, 0x94, 0x3f, 0xae, 0xd2,
	0x78, 0xad, 0xfe, 0xa3, 0x40, 0x6d, 0x05, 0x1b, 0x05, 0xcc, 0xc5, 0x6d, 0x28, 0x88, 0x65, 0xc0,
	0xe2, 0xf8, 0xf5, 0x6e, 0x2b, 0x49, 0x22, 0x6a, 0x6f, 0x80, 0xda, 0xf6, 0x32, 0x60, 0x34, 0xc6,
	0xe1, 0xd7, 0x50, 0x71, 0xef, 0xd3, 0x8b, 0x53, 0xa8, 0x74, 0x9f, 0xed, 0xd0, 0x4c, 0x83, 0xae,
	0xe3, 0xf0, 0xb7, 0x50, 0xf4, 0x64, 0x56, 0xcd, 0x7c, 0x4c, 0x38, 0xdc, 0x25, 0x48, 0x2f, 0x4d,
	0x40, 0xb8, 0x09, 0xfb, 0x52, 0x16, 0x3e, 0x17, 0xcd, 0xc2, 0x91, 0x72, 0x5c, 0xa4, 0x99, 0xa9,
	0xbe, 0x81, 0x82, 0x4c, 0x06, 0xd7, 0xa0, 0x7c, 0x66, 0x19, 0xe4, 0x9d, 0x69, 0x11, 0x03, 0x3d,
	0xc1, 0x00, 0x7b, 0xbd, 0x41, 0x5f, 0xb3, 0x7a, 0x48, 0xc1, 0x25, 0x28, 0x58, 0x03, 0x83, 0xa0,
	0x1c, 0xde, 0x87, 0xbc, 0xae, 0x51, 0x94, 0x97, 0x5b, 0xef, 0xb5, 0x73, 0x0d, 0x15, 0xd4, 0xbf,
	0x72, 0xf0, 0x62, 0x75, 0xa6, 0xc1, 0x82, 0x19, 0x5f, 0xde, 0x32, 0x5f, 0xc4, 0x52, 0xfc, 0x0c,
	0x35, 0x77, 0xbd, 0xec, 0x58, 0x93, 0x4a, 0xf7, 0xf9, 0x83, 0x9a, 0xd0, 0x4d, 0x2c, 0xfe, 0x05,
	0x6a, 0x6c, 0x32, 0x61, 0xae, 0xf0, 0x16, 0xcc, 0x70, 0x04, 0x4b, 0x95, 0x69, 0xb5, 0x93, 0x7e,
	0xb7, 0xb3, 0x7e, 0xb7, 0xed, 0xac, 0xdf, 0x74, 0x93, 0x80, 0x8f, 0xa0, 0x22, 0xa3, 0x0d, 0x1d,
	0xf7, 0xc6, 0x99, 0xb2, 0x58, 0xa8, 0x2a, 0x5d, 0xdf, 0xc2, 0x16, 0xec, 0xb3, 0x3b, 0xe6, 0x12,
	0x7f, 0x11, 0xcb, 0x52, 0xef, 0xbe, 0xda, 0x49, 0x6d, 0xb3, 0xa4, 0x36, 0xb9, 0x63, 0xee, 0x5c,
	0x78, 0xdc, 0x27, 0xfe, 0xc2, 0x0b, 0xb9, 0x2f, 0x1d, 0x34, 0x0b, 0xa2, 0xb6, 0xa1, 0xf1, 0x10,
	0x40, 0xaa, 0x69, 0x0c, 0xf4, 0x13, 0x42, 0x13, 0x65, 0x47, 0x17, 0x23, 0x9b, 0x9c, 0x22, 0x45,
	0xfd, 0x5d, 0x59, 0x13, 0xcf, 0xf4, 0x17, 0xdc, 0x75, 0x24, 0xf5, 0xbf, 0x8b, 0x77, 0x0c, 0x07,
	0xde, 0xb8, 0xc7, 0x7c, 0x16, 0xc6, 0x01, 0xb5, 0xd9, 0x34, 0xbd, 0xdb, 0xdb, 0xdb, 0x2a, 0x85,
	0xe6, 0x2a, 0xd2, 0x30, 0xe4, 0x01, 0x8f, 0x9c, 0x99, 0xce, 0x7d, 0xc1, 0xee, 0xe2, 0x5b, 0xe3,
	0x86, 0xcc, 0x11, 0x3c, 0x8c, 0x0f, 0xaf, 0xd2, 0xcc, 0xc4, 0x9f, 0x41, 0x59, 0x84, 0x8e, 0x1f,
	0x79, 0xcc, 0x17, 0x71, 0xe4, 0x2a, 0xbd, 0xdf, 0x50, 0xff, 0x2e, 0x00, 0x5a, 0x05, 0x3d, 0x65,
	0x51, 0x24, 0xb5, 0xfe, 0x7e, 0x63, 0x2e, 0x3e, 0xdf, 0x29, 0x23, 0xc5, 0xad, 0x8f, 0xc6, 0x4f,
	0x50, 0x5e, 0x0d, 0xf3, 0x27, 0xb4, 0xff, 0x1e, 0x2c, 0x33, 0x0f, 0x9c, 0xe5, 0x8c, 0x3b, 0xe3,
	0xb4, 0xed, 0x99, 0x29, 0x87, 0x58, 0xdc, 0x79, 0xe3, 0xb8, 0xdf, 0x65, 0x1a, 0xaf, 0xf1, 0x7b,
	0x38, 0x08, 0x36, 0x4b, 0x6f, 0x16, 0xe3, 0xd3, 0x8e, 0x76, 0xb2, 0xdc, 0x92, 0x88, 0x6e, 0x13,
	0xf1, 0x1b, 0xa8, 0xaf, 0x5a, 0x41, 0xe4, 0x33, 0xd5, 0xdc, 0x7b, 0x64, 0x40, 0x63, 0x2f, 0xdd,
	0x42, 0xab, 0x7f, 0xe6, 0x1e, 0x1e, 0xc8, 0x2a, 0x94, 0x28, 0xe9, 0x99, 0x23, 0x9b, 0x50, 0xa4,
	0xe0, 0x3a, 0x40, 0x66, 0x11, 0x03, 0xe5, 0xe4, 0x3c, 0x9a, 0x96, 0x69, 0xa3, 0x3c, 0x2e, 0x43,
	0x91, 0x12, 0xcd, 0xb8, 0x40, 0x05, 0x7c, 0x00, 0x15, 0x9b, 0x6a, 0xd6, 0x48, 0xd3, 0x6d, 0x73,
	0x60, 0xa1, 0xa2, 0x0c, 0xa9, 0x0f, 0x4e, 0x87, 0x7d, 0x62, 0x13, 0x03, 0xed, 0x49, 0x28, 0xa1,
	0x74, 0x40, 0xd1, 0xbe, 0xf4, 0xf4, 0x88, 0x7d, 0x39, 0xb2, 0x35, 0x9b, 0xa0, 0x92, 0x34, 0x87,
	0x67, 0x99, 0x59, 0x96, 0xa6, 0x41, 0xfa, 0xa9, 0x09, 0xb8, 0x01, 0xc8, 0xb4, 0xce, 0x07, 0x27,
	0xe4, 0x52, 0xff, 0x55, 0x33, 0x2d, 0x5d, 0xbe, 0x0d, 0x95, 0x24, 0xc1, 0xd1, 0x70, 0x60, 0x8d,
	0x08, 0xaa, 0xe1, 0xe7, 0xf0, 0x94, 0x6a, 0x56, 0x8f, 0x5c, 0x7e, 0x38, 0x23, 0xf4, 0x22, 0xa5,
	0xd6, 0x71, 0x0b, 0x0e, 0x77, 0xb6, 0x2f, 0x2d, 0xf2, 0xd1, 0x46, 0x07, 0xf8, 0xff, 0xf0, 0x62,
	0xd7, 0xa7, 0xf7, 0x07, 0x23, 0x82, 0x90, 0x4c, 0xe1, 0x84, 0x90, 0xa1, 0xd6, 0x37, 0xcf, 0x09,
	0x7a, 0xaa, 0xfe, 0x08, 0xd5, 0xe1, 0x5c, 0x8c, 0x84, 0x23, 0x98, 0xe9, 0x4f, 0x38, 0x46, 0x90,
	0xbf, 0x61, 0xcb, 0xf4, 0x4d, 0x97, 0x4b, 0xdc, 0x80, 0xe2, 0xc2, 0x99, 0xcd, 0x59, 0x7a, 0x3b,
	0x13, 0x43, 0x25, 0x70, 0x40, 0x1d, 0x7f, 0xca, 0x3e, 0xcc, 0x59, 0xb8, 0x8c, 0xe9, 0xb8, 0x05,
	0xa5, 0x48, 0x38, 0xa1, 0x38, 0x59, 0xf1, 0x57, 0x36, 0x3e, 0x84, 0x3d, 0xe6, 0x8f, 0xa5, 0x27,
	0x99, 0x9e, 0xd4, 0x52, 0xbf, 0x84, 0x67, 0x5b, 0x61, 0x2c, 0xd9, 0xfb, 0x3a, 0xe4, 0x4c, 0x23,
	0x0d, 0x92, 0x33, 0x0d, 0xf5, 0x2b, 0x68, 0x6c, 0xc1, 0xf4, 0x19, 0x8f, 0xd8, 0x0e, 0x4e, 0x83,
	0x17, 0x5b, 0xb8, 0x13, 0xb6, 0x3c, 0x97, 0x09, 0x7f, 0x72, 0x61, 0x7f, 0x28, 0x3b, 0x31, 0x28,
	0x8b, 0x02, 0xee, 0x47, 0x0c, 0x13, 0xa8, 0xdd, 0xb0, 0x65, 0xa4, 0xf9, 0xe3, 0x38, 0x66, 0xf2,
	0x07, 0x56, 0xe9, 0xbe, 0xcc, 0x6e, 0xe4, 0x23, 0x67, 0xd3, 0x4d, 0x96, 0x9c, 0xa9, 0x6b, 0x27,
	0x3a, 0xe5, 0x61, 0x72, 0x74, 0x89, 0x66, 0x66, 0x5a, 0x4f, 0x3e, 0xab, 0xe7, 0x9b, 0x57, 0xd0,
	0xd0, 0xb9, 0x3f, 0xf1, 0xc6, 0xcc, 0x17, 0x9e, 0x33, 0xf3, 0xc4, 0xb2, 0xcf, 0x16, 0x6c, 0x26,
	0x9f, 0xbe, 0xe1, 0xd9, 0xdb, 0xbe, 0xa9, 0xa3, 0x27, 0x18, 0x41, 0x55, 0x1f, 0x58, 0xef, 0x4c,
	0x83, 0x58, 0xb6, 0xa9, 0xf5, 0x91, 0xd2, 0xfd, 0xb8, 0xf6, 0x68, 0x8c, 0xe6, 0x41, 0xc0, 0x43,
	0x81, 0x0d, 0x28, 0x51, 0x36, 0xf5, 0x22, 0xc1, 0x42, 0xdc, 0x7c, 0xec, 0xc9, 0x68, 0x3d, 0xea,
	0x51, 0x9f, 0x1c, 0x2b, 0xdf, 0x29, 0x6f, 0x75, 0x38, 0xe4, 0xe1, 0xb4, 0x7d, 0xbd, 0x0c, 0x58,
	0x38, 0x63, 0xe3, 0x29, 0x0b, 0x53, 0xc2, 0x6f, 0x5f, 0x4f, 0x3d, 0x71, 0x3d, 0xbf, 0x6a, 0xbb,
	0xfc, 0xb6, 0xb3, 0xe6, 0xee, 0x4c, 0x9c, 0xab, 0xd0, 0x73, 0x93, 0x6f, 0x8a, 0xa8, 0x23, 0x3f,
	0x3e, 0xae, 0x92, 0x4f, 0x91, 0x1f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xad, 0x76, 0x5b, 0xa4,
	0xa9, 0x08, 0x00, 0x00,
}
