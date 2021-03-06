// Code generated by protoc-gen-go.
// source: peer/processMessage.proto
// DO NOT EDIT!

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

type QueryBlocks struct {
	BlockIndex uint64 `protobuf:"varint,1,opt,name=BlockIndex" json:"BlockIndex,omitempty"`
	ChannelID  string `protobuf:"bytes,2,opt,name=ChannelID" json:"ChannelID,omitempty"`
}

func (m *QueryBlocks) Reset()                    { *m = QueryBlocks{} }
func (m *QueryBlocks) String() string            { return proto.CompactTextString(m) }
func (*QueryBlocks) ProtoMessage()               {}
func (*QueryBlocks) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type ChannelMessage struct {
	ChannelInput []*MessageInput `protobuf:"bytes,1,rep,name=ChannelInput" json:"ChannelInput,omitempty"`
}

func (m *ChannelMessage) Reset()                    { *m = ChannelMessage{} }
func (m *ChannelMessage) String() string            { return proto.CompactTextString(m) }
func (*ChannelMessage) ProtoMessage()               {}
func (*ChannelMessage) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *ChannelMessage) GetChannelInput() []*MessageInput {
	if m != nil {
		return m.ChannelInput
	}
	return nil
}

type MessageInput struct {
	PeerIp    string    `protobuf:"bytes,1,opt,name=PeerIp" json:"PeerIp,omitempty"`
	Height    uint64    `protobuf:"varint,2,opt,name=Height" json:"Height,omitempty"`
	PeerName  string    `protobuf:"bytes,3,opt,name=PeerName" json:"PeerName,omitempty"`
	Mblocks   []*Mblock `protobuf:"bytes,4,rep,name=Mblocks" json:"Mblocks,omitempty"`
	ChannelID string    `protobuf:"bytes,5,opt,name=ChannelID" json:"ChannelID,omitempty"`
}

func (m *MessageInput) Reset()                    { *m = MessageInput{} }
func (m *MessageInput) String() string            { return proto.CompactTextString(m) }
func (*MessageInput) ProtoMessage()               {}
func (*MessageInput) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *MessageInput) GetMblocks() []*Mblock {
	if m != nil {
		return m.Mblocks
	}
	return nil
}

type Mblock struct {
	Header   *MblockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *MblockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *MblockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Mblock) Reset()                    { *m = Mblock{} }
func (m *Mblock) String() string            { return proto.CompactTextString(m) }
func (*Mblock) ProtoMessage()               {}
func (*Mblock) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *Mblock) GetHeader() *MblockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Mblock) GetData() *MblockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Mblock) GetMetadata() *MblockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MblockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash string `protobuf:"bytes,2,opt,name=PreviousHash" json:"PreviousHash,omitempty"`
	DataHash     string `protobuf:"bytes,3,opt,name=DataHash" json:"DataHash,omitempty"`
	NowHash      string `protobuf:"bytes,4,opt,name=NowHash" json:"NowHash,omitempty"`
}

func (m *MblockHeader) Reset()                    { *m = MblockHeader{} }
func (m *MblockHeader) String() string            { return proto.CompactTextString(m) }
func (*MblockHeader) ProtoMessage()               {}
func (*MblockHeader) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

type MblockData struct {
	Datas []*TransData `protobuf:"bytes,1,rep,name=Datas" json:"Datas,omitempty"`
}

func (m *MblockData) Reset()                    { *m = MblockData{} }
func (m *MblockData) String() string            { return proto.CompactTextString(m) }
func (*MblockData) ProtoMessage()               {}
func (*MblockData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *MblockData) GetDatas() []*TransData {
	if m != nil {
		return m.Datas
	}
	return nil
}

type TransData struct {
	Txid        string                      `protobuf:"bytes,1,opt,name=Txid" json:"Txid,omitempty"`
	ChainID     string                      `protobuf:"bytes,2,opt,name=ChainID" json:"ChainID,omitempty"`
	Time        *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=Time" json:"Time,omitempty"`
	ChainCodeID string                      `protobuf:"bytes,4,opt,name=ChainCodeID" json:"ChainCodeID,omitempty"`
	Payload     string                      `protobuf:"bytes,5,opt,name=Payload" json:"Payload,omitempty"`
	Type        string                      `protobuf:"bytes,6,opt,name=Type" json:"Type,omitempty"`
	Nonce       string                      `protobuf:"bytes,7,opt,name=Nonce" json:"Nonce,omitempty"`
	Signature   string                      `protobuf:"bytes,8,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *TransData) Reset()                    { *m = TransData{} }
func (m *TransData) String() string            { return proto.CompactTextString(m) }
func (*TransData) ProtoMessage()               {}
func (*TransData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *TransData) GetTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type MblockMetadata struct {
	Metadata []string `protobuf:"bytes,1,rep,name=Metadata" json:"Metadata,omitempty"`
}

func (m *MblockMetadata) Reset()                    { *m = MblockMetadata{} }
func (m *MblockMetadata) String() string            { return proto.CompactTextString(m) }
func (*MblockMetadata) ProtoMessage()               {}
func (*MblockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

type MessageOutput struct {
	Output string `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *MessageOutput) Reset()                    { *m = MessageOutput{} }
func (m *MessageOutput) String() string            { return proto.CompactTextString(m) }
func (*MessageOutput) ProtoMessage()               {}
func (*MessageOutput) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func init() {
	proto.RegisterType((*QueryBlocks)(nil), "protos.QueryBlocks")
	proto.RegisterType((*ChannelMessage)(nil), "protos.ChannelMessage")
	proto.RegisterType((*MessageInput)(nil), "protos.MessageInput")
	proto.RegisterType((*Mblock)(nil), "protos.Mblock")
	proto.RegisterType((*MblockHeader)(nil), "protos.MblockHeader")
	proto.RegisterType((*MblockData)(nil), "protos.MblockData")
	proto.RegisterType((*TransData)(nil), "protos.TransData")
	proto.RegisterType((*MblockMetadata)(nil), "protos.MblockMetadata")
	proto.RegisterType((*MessageOutput)(nil), "protos.MessageOutput")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for StatusPeer service

type StatusPeerClient interface {
	ProcessMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*MessageOutput, error)
}

type statusPeerClient struct {
	cc *grpc.ClientConn
}

func NewStatusPeerClient(cc *grpc.ClientConn) StatusPeerClient {
	return &statusPeerClient{cc}
}

func (c *statusPeerClient) ProcessMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*MessageOutput, error) {
	out := new(MessageOutput)
	err := grpc.Invoke(ctx, "/protos.statusPeer/ProcessMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StatusPeer service

type StatusPeerServer interface {
	ProcessMessage(context.Context, *ChannelMessage) (*MessageOutput, error)
}

func RegisterStatusPeerServer(s *grpc.Server, srv StatusPeerServer) {
	s.RegisterService(&_StatusPeer_serviceDesc, srv)
}

func _StatusPeer_ProcessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusPeerServer).ProcessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.statusPeer/ProcessMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusPeerServer).ProcessMessage(ctx, req.(*ChannelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusPeer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.statusPeer",
	HandlerType: (*StatusPeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMessage",
			Handler:    _StatusPeer_ProcessMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor7,
}

// Client API for QueryPeer service

type QueryPeerClient interface {
	QueryMessage(ctx context.Context, in *QueryBlocks, opts ...grpc.CallOption) (*Mblock, error)
}

type queryPeerClient struct {
	cc *grpc.ClientConn
}

func NewQueryPeerClient(cc *grpc.ClientConn) QueryPeerClient {
	return &queryPeerClient{cc}
}

func (c *queryPeerClient) QueryMessage(ctx context.Context, in *QueryBlocks, opts ...grpc.CallOption) (*Mblock, error) {
	out := new(Mblock)
	err := grpc.Invoke(ctx, "/protos.queryPeer/QueryMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QueryPeer service

type QueryPeerServer interface {
	QueryMessage(context.Context, *QueryBlocks) (*Mblock, error)
}

func RegisterQueryPeerServer(s *grpc.Server, srv QueryPeerServer) {
	s.RegisterService(&_QueryPeer_serviceDesc, srv)
}

func _QueryPeer_QueryMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlocks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryPeerServer).QueryMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.queryPeer/QueryMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryPeerServer).QueryMessage(ctx, req.(*QueryBlocks))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryPeer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.queryPeer",
	HandlerType: (*QueryPeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMessage",
			Handler:    _QueryPeer_QueryMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor7,
}

func init() { proto.RegisterFile("peer/processMessage.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xbf, 0xba, 0x69, 0x73, 0x93, 0x2f, 0x12, 0x43, 0xa9, 0x4c, 0x84, 0x20, 0xf2, 0x82,
	0x06, 0x51, 0x39, 0x52, 0x50, 0x25, 0xb6, 0xb4, 0x5d, 0x34, 0xa0, 0xa6, 0xc1, 0x64, 0xc5, 0x6e,
	0x12, 0xdf, 0xda, 0x16, 0x89, 0xc7, 0x78, 0x6c, 0x68, 0x76, 0x3c, 0x03, 0x8f, 0xc0, 0xfb, 0xf1,
	0x0e, 0x68, 0xee, 0xcc, 0xb8, 0x76, 0x57, 0xf6, 0xb9, 0xe7, 0xd8, 0x73, 0xee, 0xdc, 0x1f, 0x78,
	0x9e, 0x23, 0x16, 0x93, 0xbc, 0x10, 0x6b, 0x94, 0xf2, 0x06, 0xa5, 0xe4, 0x31, 0x06, 0x79, 0x21,
	0x4a, 0xc1, 0x3a, 0xf4, 0x90, 0xc3, 0x57, 0xb1, 0x10, 0xf1, 0x06, 0x27, 0x04, 0x57, 0xd5, 0xdd,
	0xa4, 0x4c, 0xb7, 0x28, 0x4b, 0xbe, 0xcd, 0xb5, 0xd0, 0xff, 0x04, 0xbd, 0xcf, 0x15, 0x16, 0xbb,
	0x8b, 0x8d, 0x58, 0x7f, 0x93, 0xec, 0x25, 0x00, 0xbd, 0xcd, 0xb2, 0x08, 0xef, 0x3d, 0x67, 0xe4,
	0x8c, 0xdd, 0xb0, 0x11, 0x61, 0x2f, 0xa0, 0x7b, 0x99, 0xf0, 0x2c, 0xc3, 0xcd, 0xec, 0xca, 0xfb,
	0x6f, 0xe4, 0x8c, 0xbb, 0xe1, 0x43, 0xc0, 0xff, 0x08, 0x03, 0x03, 0x8c, 0x1b, 0xf6, 0x1e, 0xfa,
	0x96, 0xce, 0xf2, 0xaa, 0xf4, 0x9c, 0xd1, 0xfe, 0xb8, 0x37, 0x3d, 0xd6, 0x87, 0xcb, 0xc0, 0xc8,
	0x88, 0x0b, 0x5b, 0x4a, 0xff, 0x8f, 0x03, 0xfd, 0x26, 0xcd, 0x4e, 0xa0, 0xb3, 0x40, 0x2c, 0x66,
	0x39, 0xd9, 0xea, 0x86, 0x06, 0xa9, 0xf8, 0x35, 0xa6, 0x71, 0x52, 0x92, 0x1f, 0x37, 0x34, 0x88,
	0x0d, 0xe1, 0x48, 0x29, 0xe6, 0x7c, 0x8b, 0xde, 0x3e, 0x7d, 0x51, 0x63, 0x36, 0x86, 0xc3, 0x9b,
	0x15, 0x65, 0xec, 0xb9, 0xe4, 0x68, 0x50, 0x3b, 0xa2, 0x70, 0x68, 0xe9, 0x76, 0xc2, 0x07, 0x8f,
	0x13, 0xfe, 0xed, 0x40, 0x47, 0x2b, 0xd9, 0x99, 0xb2, 0xc1, 0x23, 0x2c, 0xc8, 0x5e, 0x33, 0x47,
	0xe2, 0x35, 0x17, 0x1a, 0x0d, 0x7b, 0x0d, 0xee, 0x15, 0x2f, 0x39, 0x59, 0xee, 0x4d, 0x59, 0x5b,
	0xab, 0x98, 0x90, 0x78, 0x36, 0x85, 0xa3, 0x1b, 0x2c, 0x79, 0xa4, 0xb4, 0xfb, 0xa4, 0x3d, 0x69,
	0x6b, 0x2d, 0x1b, 0xd6, 0x3a, 0xff, 0x97, 0xba, 0xb9, 0xc6, 0xa1, 0xea, 0x86, 0xe6, 0xd5, 0x76,
	0x65, 0xac, 0xb9, 0xa1, 0x41, 0xcc, 0x87, 0xfe, 0xa2, 0xc0, 0x1f, 0xa9, 0xa8, 0xe4, 0x35, 0x97,
	0x89, 0xa9, 0x67, 0x2b, 0xa6, 0x6e, 0x51, 0x19, 0x21, 0xde, 0xdc, 0xa2, 0xc5, 0xcc, 0x83, 0xc3,
	0xb9, 0xf8, 0x49, 0x94, 0x4b, 0x94, 0x85, 0xfe, 0x39, 0xc0, 0x43, 0x2a, 0xec, 0x14, 0x0e, 0xd4,
	0x53, 0x9a, 0xea, 0x3f, 0xb1, 0x19, 0x2c, 0x0b, 0x9e, 0x49, 0x4a, 0x56, 0xf3, 0xfe, 0x5f, 0x07,
	0xba, 0x75, 0x90, 0x31, 0x70, 0x97, 0xf7, 0x69, 0x64, 0xca, 0x4d, 0xef, 0xea, 0xc8, 0xcb, 0x84,
	0xa7, 0x59, 0xdd, 0x7d, 0x16, 0xb2, 0x00, 0xdc, 0x65, 0x6a, 0x4a, 0xdd, 0x9b, 0x0e, 0x03, 0xdd,
	0xf8, 0x81, 0x6d, 0xfc, 0x60, 0x69, 0x1b, 0x3f, 0x24, 0x1d, 0x1b, 0x41, 0x8f, 0x3e, 0xbd, 0x14,
	0x11, 0xce, 0xae, 0x4c, 0x02, 0xcd, 0x90, 0x3a, 0x6b, 0xc1, 0x77, 0x1b, 0xc1, 0x23, 0x53, 0x78,
	0x0b, 0xc9, 0xd9, 0x2e, 0x47, 0xaf, 0x63, 0x9c, 0xed, 0x72, 0x64, 0xc7, 0x70, 0x30, 0x17, 0xd9,
	0x1a, 0xbd, 0x43, 0x0a, 0x6a, 0xa0, 0xda, 0xe7, 0x4b, 0x1a, 0x67, 0xbc, 0xac, 0x0a, 0xf4, 0x8e,
	0x74, 0xfb, 0xd4, 0x01, 0xff, 0x0c, 0x06, 0xed, 0x2a, 0xaa, 0xeb, 0xae, 0xeb, 0xad, 0x6e, 0xab,
	0xdb, 0xa8, 0xeb, 0x29, 0xfc, 0x6f, 0x06, 0xe2, 0xb6, 0x2a, 0xcd, 0x44, 0x08, 0x7a, 0xb3, 0x13,
	0xa1, 0xd1, 0xf4, 0x16, 0x40, 0x96, 0xbc, 0xac, 0xa4, 0xea, 0x77, 0xf6, 0x01, 0x06, 0x8b, 0xd6,
	0x8a, 0x60, 0x75, 0x0b, 0xb5, 0x87, 0x75, 0xf8, 0xec, 0xd1, 0x58, 0xea, 0x63, 0xfc, 0xbd, 0xe9,
	0x05, 0x74, 0xbf, 0xab, 0x25, 0x41, 0xff, 0x3b, 0x87, 0x3e, 0x6d, 0x0c, 0xfb, 0xb7, 0xa7, 0xf6,
	0xab, 0xc6, 0x1e, 0x19, 0x3e, 0x9a, 0x27, 0x7f, 0xef, 0xe2, 0xed, 0xd7, 0x37, 0x71, 0x5a, 0x26,
	0xd5, 0x2a, 0x58, 0x8b, 0xed, 0x24, 0xd9, 0xe5, 0x58, 0x6c, 0x30, 0x8a, 0xb1, 0x98, 0xdc, 0xf1,
	0x55, 0x91, 0xae, 0xf5, 0x8a, 0x92, 0x13, 0xb5, 0xd3, 0x56, 0x7a, 0x7d, 0xbd, 0xfb, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x11, 0xe3, 0x7e, 0xef, 0xe2, 0x04, 0x00, 0x00,
}
