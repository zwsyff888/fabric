package models

import (
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MblockMetadata struct {
	Metadata []string
}

func NewMblockMetadata() *MblockMetadata {
	return &MblockMetadata{}
}

func (m *MblockMetadata) initMblockMetadata(input *pb.MblockMetadata) {
	// mblockMetadata := NewMblockMetadata()
	m.Metadata = make([]string, len(input.Metadata))
	copy(m.Metadata, input.Metadata)

}
