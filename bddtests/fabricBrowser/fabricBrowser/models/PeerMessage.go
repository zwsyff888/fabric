package models

import (
	// "fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
)

type PeerMessage struct {
	Time      int64
	PeerIp    string
	Height    uint64
	Status    int64
	Name      string
	Mblocks   []*Mblock
	ChannelID string
}

func NewPeerMessage() *PeerMessage {
	return &PeerMessage{}
}

func (m *PeerMessage) InitMessage(inputMessage *pb.MessageInput) {
	m.Time = time.Now().Unix()
	m.PeerIp = inputMessage.PeerIp
	m.Height = inputMessage.Height
	m.Name = inputMessage.PeerName
	m.Mblocks = []*Mblock{}
	m.Status = 1
	m.ChannelID = inputMessage.ChannelID

	length := len(inputMessage.Mblocks)
	for index := 0; index < length; index++ {

		//mblock
		mblock := NewMblock()
		mblock.initMblock(inputMessage.Mblocks[index])

		m.Mblocks = append(m.Mblocks, mblock)
	}
}
