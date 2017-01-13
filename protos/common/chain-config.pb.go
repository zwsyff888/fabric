// Code generated by protoc-gen-go.
// source: common/chain-config.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/chain-config.proto
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	MSPPrincipal
	OrganizationUnit
	MSPRole
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ByMSPRole MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ByOrganizationUnit MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_ByIdentity MSPPrincipal_Classification = 2
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ByMSPRole",
	1: "ByOrganizationUnit",
	2: "ByIdentity",
}
var MSPPrincipal_Classification_value = map[string]int32{
	"ByMSPRole":          0,
	"ByOrganizationUnit": 1,
	"ByIdentity":         2,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}
func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_Member MSPRole_MSPRoleType = 0
	MSPRole_Admin  MSPRole_MSPRoleType = 1
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "Member",
	1: "Admin",
}
var MSPRole_MSPRoleType_value = map[string]int32{
	"Member": 0,
	"Admin":  1,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}
func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//  - a group of identities that are member of the same MSP
//  - a group of identities that are member of the same organization unit
//    in the same MSP
//  - a group of identities that are administering a specific MSP
//  - a specific identity
// Expressing these groups is done given two fields of the fields below
//  - Classification, that defines the type of classification of identities
//    in an MSP this principal would be defined on; Classification can take
//    three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//          MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//          within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//          identity/certificate; this would mean that the Principal bytes
//          message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=PrincipalClassification,enum=common.MSPPrincipal_Classification" json:"PrincipalClassification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	Principal []byte `protobuf:"bytes,2,opt,name=Principal,proto3" json:"Principal,omitempty"`
}

func (m *MSPPrincipal) Reset()                    { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string            { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()               {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// OrganizationUnitIdentifier defines the organization unit under the
	// MSP identified with MSPIdentifier
	OrganizationUnitIdentifier string `protobuf:"bytes,2,opt,name=OrganizationUnitIdentifier" json:"OrganizationUnitIdentifier,omitempty"`
}

func (m *OrganizationUnit) Reset()                    { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()               {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=Role,enum=common.MSPRole_MSPRoleType" json:"Role,omitempty"`
}

func (m *MSPRole) Reset()                    { *m = MSPRole{} }
func (m *MSPRole) String() string            { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()               {}
func (*MSPRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*MSPPrincipal)(nil), "common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "common.MSPRole")
	proto.RegisterEnum("common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
}

func init() { proto.RegisterFile("common/chain-config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xbb, 0xe5, 0x7d, 0x2b, 0x19, 0xdb, 0x10, 0xf6, 0xa0, 0xf5, 0xcf, 0xa1, 0xc4, 0x1e,
	0x0a, 0xd2, 0x04, 0xea, 0x5d, 0x30, 0x1e, 0xc4, 0x43, 0x30, 0xa4, 0x7a, 0x11, 0x3c, 0x24, 0xe9,
	0x26, 0x1d, 0x48, 0x76, 0xc3, 0x76, 0x05, 0xd7, 0x0f, 0xe0, 0x27, 0xf4, 0x03, 0x49, 0x37, 0x5a,
	0xd3, 0x82, 0xe2, 0x69, 0x99, 0x67, 0x7e, 0xcf, 0x33, 0xbb, 0xcb, 0xc0, 0x51, 0x26, 0xaa, 0x4a,
	0x70, 0x3f, 0x5b, 0x26, 0xc8, 0xa7, 0x99, 0xe0, 0x39, 0x16, 0x5e, 0x2d, 0x85, 0x12, 0xb4, 0xd7,
	0xb4, 0xdc, 0x77, 0x02, 0xfd, 0x70, 0x1e, 0x45, 0x12, 0x79, 0x86, 0x75, 0x52, 0xd2, 0x27, 0x38,
	0xdc, 0x14, 0xd7, 0x65, 0xb2, 0x5a, 0x61, 0x8e, 0x59, 0xa2, 0x50, 0xf0, 0x21, 0x19, 0x91, 0x89,
	0x3d, 0x3b, 0xf3, 0x1a, 0xab, 0xd7, 0xb6, 0x79, 0xdb, 0x68, 0xfc, 0x53, 0x06, 0x3d, 0x05, 0x6b,
	0xd3, 0x1a, 0x76, 0x47, 0x64, 0xd2, 0x8f, 0xbf, 0x05, 0xf7, 0x06, 0xec, 0x1d, 0x7e, 0x00, 0x56,
	0xa0, 0xc3, 0x79, 0x14, 0x8b, 0x92, 0x39, 0x1d, 0x7a, 0x00, 0x34, 0xd0, 0x77, 0xb2, 0x48, 0x38,
	0xbe, 0x1a, 0xe0, 0x81, 0xa3, 0x72, 0x08, 0xb5, 0x01, 0x02, 0x7d, 0xbb, 0x60, 0x5c, 0xa1, 0xd2,
	0x4e, 0xd7, 0x7d, 0x01, 0x67, 0x97, 0xa2, 0x63, 0x18, 0x84, 0xf3, 0xa8, 0x81, 0x72, 0x64, 0xd2,
	0xbc, 0xc7, 0x8a, 0xb7, 0x45, 0x7a, 0x09, 0xc7, 0xbb, 0xce, 0x96, 0xa5, 0x6b, 0x2c, 0xbf, 0x10,
	0xee, 0x1b, 0x81, 0xbd, 0xcf, 0xfb, 0xfe, 0x71, 0xa2, 0x0f, 0xff, 0xd6, 0xb4, 0xc9, 0xb6, 0x67,
	0x27, 0xad, 0xef, 0x5d, 0xcb, 0x5f, 0xe7, 0xbd, 0xae, 0x59, 0x6c, 0x40, 0x77, 0x0c, 0xfb, 0x2d,
	0x91, 0x02, 0xf4, 0x42, 0x56, 0xa5, 0x4c, 0x3a, 0x1d, 0x6a, 0xc1, 0xff, 0xab, 0x45, 0x85, 0xdc,
	0x21, 0xc1, 0xf4, 0xf1, 0xbc, 0x40, 0xb5, 0x7c, 0x4e, 0xd7, 0x81, 0xfe, 0x52, 0xd7, 0x4c, 0x96,
	0x6c, 0x51, 0x30, 0xe9, 0xe7, 0x49, 0x2a, 0x31, 0xf3, 0xcd, 0x22, 0xac, 0xfc, 0x66, 0x5c, 0xda,
	0x33, 0xe5, 0xc5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x19, 0xfd, 0x1d, 0x34, 0x02, 0x00,
	0x00,
}
