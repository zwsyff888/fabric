// Code generated by protoc-gen-go.
// source: peer/admin.proto
// DO NOT EDIT!

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	peer/admin.proto
	peer/chaincode.proto
	peer/chaincodeevent.proto
	peer/configuration.proto
	peer/events.proto
	peer/peer.proto
	peer/proposal.proto
	peer/proposal_response.proto
	peer/server_supervise.proto
	peer/transaction.proto

It has these top-level messages:
	ServerStatus
	LogLevelRequest
	LogLevelResponse
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
	ChaincodeEvent
	AnchorPeers
	AnchorPeer
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	Event
	PeerID
	PeerEndpoint
	BlockchainInfo
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
	InvalidTransaction
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
	LogModule string `protobuf:"bytes,1,opt,name=logModule" json:"logModule,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=logLevel" json:"logLevel,omitempty"`
}

func (m *LogLevelRequest) Reset()                    { *m = LogLevelRequest{} }
func (m *LogLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*LogLevelRequest) ProtoMessage()               {}
func (*LogLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LogLevelResponse struct {
	LogModule string `protobuf:"bytes,1,opt,name=logModule" json:"logModule,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=logLevel" json:"logLevel,omitempty"`
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
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x31, 0x2d, 0xb4, 0x1e, 0xfa, 0x67, 0xbb, 0xaa, 0x5a, 0xe4, 0x56, 0x6a, 0xe5, 0x53,
	0xab, 0x4a, 0xb6, 0x44, 0x0f, 0x39, 0x24, 0x39, 0x90, 0xd8, 0x21, 0x11, 0xc4, 0x20, 0x1b, 0x14,
	0x25, 0x37, 0x1b, 0x0f, 0x06, 0xc9, 0xb0, 0xce, 0xee, 0x1a, 0x89, 0xd7, 0xc9, 0x3b, 0xe5, 0x7d,
	0x22, 0x7b, 0x21, 0x44, 0x49, 0x2e, 0x41, 0x39, 0x8d, 0x67, 0xe6, 0x9b, 0x4f, 0xde, 0x9f, 0x3e,
	0x20, 0x19, 0x22, 0xb7, 0xc3, 0x78, 0x3e, 0x5b, 0x58, 0x19, 0x67, 0x92, 0xd1, 0x7a, 0x59, 0x84,
	0xf1, 0x23, 0x61, 0x2c, 0x49, 0xd1, 0x2e, 0xdb, 0x28, 0x9f, 0xd8, 0x38, 0xcf, 0xe4, 0x4a, 0x89,
	0xcc, 0x1b, 0x0d, 0x3e, 0x04, 0xc8, 0x97, 0xc8, 0x03, 0x19, 0xca, 0x5c, 0xd0, 0x3d, 0xa8, 0x8b,
	0xf2, 0xab, 0xa9, 0xfd, 0xd6, 0xfe, 0x7c, 0x6a, 0xfd, 0x52, 0x42, 0x61, 0x3d, 0x54, 0x59, 0xaa,
	0x1c, 0xb3, 0x18, 0xfd, 0xb5, 0xdc, 0xbc, 0x04, 0xd8, 0x4e, 0xe9, 0x47, 0xd0, 0x47, 0x9e, 0xe3,
	0x9e, 0x9c, 0x79, 0xae, 0x43, 0x2a, 0xb4, 0x01, 0xef, 0x82, 0x61, 0xdb, 0x1f, 0xba, 0x0e, 0xd1,
	0x54, 0xd3, 0x1f, 0x0c, 0x5c, 0x87, 0x54, 0x29, 0x40, 0x7d, 0xd0, 0x1e, 0x05, 0xae, 0x43, 0xde,
	0x50, 0x1d, 0x6a, 0xae, 0xef, 0xf7, 0x7d, 0xf2, 0xb6, 0xd0, 0x8c, 0xbc, 0xae, 0xd7, 0xbf, 0xf0,
	0x48, 0xcd, 0xec, 0xc2, 0xe7, 0x1e, 0x4b, 0x7a, 0xb8, 0xc4, 0xd4, 0xc7, 0xeb, 0x1c, 0x85, 0xa4,
	0x3f, 0x41, 0x4f, 0x59, 0x72, 0xce, 0xe2, 0x3c, 0xc5, 0xf2, 0x4f, 0x75, 0x7f, 0x3b, 0xa0, 0x06,
	0xbc, 0x4f, 0xd7, 0x07, 0xcd, 0x6a, 0xb9, 0xbc, 0xef, 0xcd, 0x1e, 0x90, 0xad, 0x99, 0xc8, 0xd8,
	0x42, 0xe0, 0xee, 0x6e, 0xad, 0xdb, 0x2a, 0xd4, 0xda, 0x05, 0x74, 0xba, 0x0f, 0x7a, 0x07, 0xe5,
	0x9a, 0xe2, 0x37, 0x4b, 0x41, 0xb7, 0x36, 0xd0, 0x2d, 0xb7, 0x80, 0x6e, 0x7c, 0x7d, 0x8e, 0xa6,
	0x59, 0xa1, 0x87, 0xd0, 0x08, 0x64, 0xc8, 0xa5, 0x1a, 0xbf, 0xf8, 0xfc, 0xa0, 0x60, 0xcf, 0xb2,
	0x1d, 0xaf, 0x4f, 0xe1, 0x4b, 0x07, 0xa5, 0x7a, 0xec, 0x06, 0x0d, 0xfd, 0xbe, 0x11, 0x3f, 0x22,
	0x6f, 0x34, 0x9f, 0x2e, 0x14, 0x45, 0xe5, 0x14, 0xbc, 0x8a, 0xd3, 0xd1, 0xbf, 0xab, 0xbf, 0xc9,
	0x4c, 0x4e, 0xf3, 0xc8, 0x1a, 0xb3, 0xb9, 0x3d, 0x5d, 0x65, 0xc8, 0x53, 0x8c, 0x13, 0xe4, 0xf6,
	0x24, 0x8c, 0xf8, 0x6c, 0xac, 0xd2, 0x2c, 0xec, 0x22, 0xf5, 0x91, 0x4a, 0xfa, 0xff, 0xbb, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0x1f, 0x0c, 0x90, 0x04, 0x03, 0x00, 0x00,
}
