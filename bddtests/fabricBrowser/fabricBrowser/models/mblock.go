package models

import (
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Mblock struct {
	Header   MblockHeader
	Data     MblockData
	Metadata MblockMetadata
}

func NewMblock() *Mblock {
	return &Mblock{}
}

func (m *Mblock) initMblock(block *pb.Mblock) {
	//header
	header := NewMblockHeader()
	header.initMblockHeader(block.Header)
	// fmt.Println("@@@chenqiao: prehash 0:", inputMessage.Mblocks[index].Data)

	//mblockMetadata
	mblockMetadata := NewMblockMetadata()
	mblockMetadata.initMblockMetadata(block.Metadata)

	//mblockData
	mblockData := NewMblockData()
	mblockData.initMblockData(block.Data)

	m.Header = *header
	m.Data = *mblockData
	m.Metadata = *mblockMetadata
}
