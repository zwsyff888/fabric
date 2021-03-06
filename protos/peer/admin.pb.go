// Code generated by protoc-gen-go.
// source: peer/admin.proto
// DO NOT EDIT!

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	peer/admin.proto
	peer/chaincodeevent.proto
	peer/chaincode.proto
	peer/chaincodeshim.proto
	peer/configuration.proto
	peer/events.proto
	peer/peer.proto
	peer/processMessage.proto
	peer/proposal.proto
	peer/proposal_response.proto
	peer/server_supervise.proto
	peer/transaction.proto

It has these top-level messages:
	ServerStatus
	LogLevelRequest
	LogLevelResponse
	ChaincodeEvent
	ChaincodeID
	ChaincodeInput
	ChaincodeSpec
	ChaincodeDeploymentSpec
	ChaincodeInvocationSpec
	ChaincodeMessage
	PutStateInfo
	GetStateByRange
	GetQueryResult
	GetHistoryForKey
	QueryStateNext
	QueryStateClose
	QueryStateKeyValue
	QueryStateResponse
	AnchorPeers
	AnchorPeer
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	SignedEvent
	Event
	PeerID
	PeerEndpoint
	QueryBlocks
	ChannelMessage
	MessageInput
	Mblock
	MblockHeader
	MblockData
	TransData
	MblockMetadata
	MessageOutput
	SignedProposal
	Proposal
	ChaincodeHeaderExtension
	ChaincodeProposalPayload
	ChaincodeAction
	ProposalResponse
	Response
	ProposalResponsePayload
	Endorsement
	PeerReply
	ConnectPeer
	PeerInfo
	SignedTransaction
	ProcessedTransaction
	Transaction
	TransactionAction
	ChaincodeActionPayload
	ChaincodeEndorsedAction
*/
package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type ServerStatus_StatusCode int32

const (
	ServerStatus_UNDEFINED ServerStatus_StatusCode = 0
	ServerStatus_STARTED   ServerStatus_StatusCode = 1
	ServerStatus_STOPPED   ServerStatus_StatusCode = 2
	ServerStatus_PAUSED    ServerStatus_StatusCode = 3
	ServerStatus_ERROR     ServerStatus_StatusCode = 4
	ServerStatus_UNKNOWN   ServerStatus_StatusCode = 5
)

var ServerStatus_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "STARTED",
	2: "STOPPED",
	3: "PAUSED",
	4: "ERROR",
	5: "UNKNOWN",
}
var ServerStatus_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"STARTED":   1,
	"STOPPED":   2,
	"PAUSED":    3,
	"ERROR":     4,
	"UNKNOWN":   5,
}

func (x ServerStatus_StatusCode) String() string {
	return proto.EnumName(ServerStatus_StatusCode_name, int32(x))
}
func (ServerStatus_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ServerStatus struct {
	Status ServerStatus_StatusCode `protobuf:"varint,1,opt,name=status,enum=protos.ServerStatus_StatusCode" json:"status,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LogLevelRequest struct {
	LogModule string `protobuf:"bytes,1,opt,name=log_module,json=logModule" json:"log_module,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
}

func (m *LogLevelRequest) Reset()                    { *m = LogLevelRequest{} }
func (m *LogLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*LogLevelRequest) ProtoMessage()               {}
func (*LogLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LogLevelResponse struct {
	LogModule string `protobuf:"bytes,1,opt,name=log_module,json=logModule" json:"log_module,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
}

func (m *LogLevelResponse) Reset()                    { *m = LogLevelResponse{} }
func (m *LogLevelResponse) String() string            { return proto.CompactTextString(m) }
func (*LogLevelResponse) ProtoMessage()               {}
func (*LogLevelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ServerStatus)(nil), "protos.ServerStatus")
	proto.RegisterType((*LogLevelRequest)(nil), "protos.LogLevelRequest")
	proto.RegisterType((*LogLevelResponse)(nil), "protos.LogLevelResponse")
	proto.RegisterEnum("protos.ServerStatus_StatusCode", ServerStatus_StatusCode_name, ServerStatus_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Admin service

type AdminClient interface {
	// Return the serve status.
	GetStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error)
	StartServer(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error)
	StopServer(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error)
	GetModuleLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error)
	SetModuleLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/protos.Admin/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StartServer(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/protos.Admin/StartServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StopServer(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/protos.Admin/StopServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetModuleLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := grpc.Invoke(ctx, "/protos.Admin/GetModuleLogLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetModuleLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := grpc.Invoke(ctx, "/protos.Admin/SetModuleLogLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// Return the serve status.
	GetStatus(context.Context, *google_protobuf.Empty) (*ServerStatus, error)
	StartServer(context.Context, *google_protobuf.Empty) (*ServerStatus, error)
	StopServer(context.Context, *google_protobuf.Empty) (*ServerStatus, error)
	GetModuleLogLevel(context.Context, *LogLevelRequest) (*LogLevelResponse, error)
	SetModuleLogLevel(context.Context, *LogLevelRequest) (*LogLevelResponse, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetStatus(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StartServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StartServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/StartServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StartServer(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StopServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StopServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/StopServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StopServer(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetModuleLogLevel(ctx, req.(*LogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/SetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SetModuleLogLevel(ctx, req.(*LogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Admin_GetStatus_Handler,
		},
		{
			MethodName: "StartServer",
			Handler:    _Admin_StartServer_Handler,
		},
		{
			MethodName: "StopServer",
			Handler:    _Admin_StopServer_Handler,
		},
		{
			MethodName: "GetModuleLogLevel",
			Handler:    _Admin_GetModuleLogLevel_Handler,
		},
		{
			MethodName: "SetModuleLogLevel",
			Handler:    _Admin_SetModuleLogLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("peer/admin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0x29, 0x0a, 0xda, 0x83, 0x7f, 0xc6, 0x89, 0x51, 0x02, 0x31, 0x9a, 0x5e, 0x69, 0x4c,
	0xda, 0x04, 0x2f, 0xbc, 0x50, 0x2f, 0xd0, 0x56, 0x34, 0x42, 0x21, 0x2d, 0xc4, 0xe8, 0x0d, 0x69,
	0xe9, 0xa1, 0x90, 0x4c, 0x99, 0xee, 0xcc, 0x94, 0x84, 0xd7, 0xd9, 0x77, 0xda, 0xf7, 0xd9, 0xb4,
	0x03, 0x61, 0xb3, 0xbb, 0x37, 0xfb, 0xe7, 0xea, 0xf4, 0x9c, 0xf3, 0x9d, 0x2f, 0x9d, 0x5f, 0x3e,
	0x20, 0x39, 0xa2, 0x70, 0xa2, 0x24, 0x5b, 0x6f, 0xec, 0x5c, 0x70, 0xc5, 0x69, 0xb3, 0x2a, 0xb2,
	0xd3, 0x4d, 0x39, 0x4f, 0x19, 0x3a, 0x55, 0x1b, 0x17, 0x4b, 0x07, 0xb3, 0x5c, 0xed, 0xb4, 0xc8,
	0x3a, 0x35, 0xe0, 0x49, 0x88, 0x62, 0x8b, 0x22, 0x54, 0x91, 0x2a, 0x24, 0xfd, 0x0c, 0x4d, 0x59,
	0x7d, 0xb5, 0x8d, 0x77, 0xc6, 0xfb, 0x67, 0xbd, 0xb7, 0x5a, 0x28, 0xed, 0x8b, 0x2a, 0x5b, 0x97,
	0x1f, 0x3c, 0xc1, 0x60, 0x2f, 0xb7, 0xfe, 0x01, 0x1c, 0xa7, 0xf4, 0x29, 0x98, 0x33, 0xdf, 0xf5,
	0x7e, 0xfe, 0xf6, 0x3d, 0x97, 0xd4, 0x68, 0x0b, 0x1e, 0x85, 0xd3, 0x7e, 0x30, 0xf5, 0x5c, 0x62,
	0xe8, 0x66, 0x3c, 0x99, 0x78, 0x2e, 0xa9, 0x53, 0x80, 0xe6, 0xa4, 0x3f, 0x0b, 0x3d, 0x97, 0x3c,
	0xa0, 0x26, 0x34, 0xbc, 0x20, 0x18, 0x07, 0xe4, 0x61, 0xa9, 0x99, 0xf9, 0x7f, 0xfc, 0xf1, 0x5f,
	0x9f, 0x34, 0xac, 0x11, 0x3c, 0x1f, 0xf2, 0x74, 0x88, 0x5b, 0x64, 0x01, 0x9e, 0x14, 0x28, 0x15,
	0x7d, 0x03, 0xc0, 0x78, 0x3a, 0xcf, 0x78, 0x52, 0x30, 0xac, 0x7e, 0xd5, 0x0c, 0x4c, 0xc6, 0xd3,
	0x51, 0x35, 0xa0, 0x5d, 0x28, 0x9b, 0x39, 0x2b, 0x4f, 0xda, 0xf5, 0x6a, 0xfb, 0x98, 0xed, 0x2d,
	0x2c, 0x1f, 0xc8, 0xd1, 0x4e, 0xe6, 0x7c, 0x23, 0xf1, 0x2e, 0x7e, 0xbd, 0xb3, 0x3a, 0x34, 0xfa,
	0x25, 0x78, 0xfa, 0x05, 0xcc, 0x01, 0xaa, 0x3d, 0xc9, 0x57, 0xb6, 0x06, 0x6f, 0x1f, 0xc0, 0xdb,
	0x5e, 0x09, 0xbe, 0xf3, 0xf2, 0x3a, 0xa2, 0x56, 0x8d, 0x7e, 0x83, 0x56, 0xa8, 0x22, 0xa1, 0xf4,
	0xf8, 0xc6, 0xe7, 0x5f, 0x4b, 0xfe, 0x3c, 0xbf, 0xe5, 0xf5, 0x2f, 0x78, 0x31, 0x40, 0xa5, 0x5f,
	0x7b, 0x80, 0x43, 0x5f, 0x1f, 0xc4, 0x97, 0xe8, 0x77, 0xda, 0x57, 0x17, 0x9a, 0xa3, 0x76, 0x0a,
	0xef, 0xc5, 0xe9, 0xfb, 0xc7, 0xff, 0x1f, 0xd2, 0xb5, 0x5a, 0x15, 0xb1, 0xbd, 0xe0, 0x99, 0xb3,
	0xda, 0xe5, 0x28, 0x18, 0x26, 0x29, 0x0a, 0x67, 0x19, 0xc5, 0x62, 0xbd, 0xd0, 0x89, 0x96, 0x4e,
	0x99, 0xfc, 0x58, 0xa7, 0xfd, 0xd3, 0x79, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x24, 0x41, 0xca,
	0x08, 0x03, 0x00, 0x00,
}
