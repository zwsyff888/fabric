package models

import (
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MblockHeader struct {
	Number       uint64
	PreviousHash string
	DataHash     string
	CurrentHash  string
}

func NewMblockHeader() *MblockHeader {
	return &MblockHeader{}
}

func (m *MblockHeader) initMblockHeader(input *pb.MblockHeader) {
	m.Number = uint64(input.Number)
	m.PreviousHash = string(input.PreviousHash)
	m.DataHash = string(input.DataHash)
	m.CurrentHash = string(input.NowHash)
}
