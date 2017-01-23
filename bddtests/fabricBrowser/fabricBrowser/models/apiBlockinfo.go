package models

import (
	"fmt"
	// "strconv"
)

type Blockinfo struct {
	BlockNumber       uint64 `json:"BlockNumber"`
	TransactionsCount uint64 `json:"TransactionsCount"`
	PreviousHash      string `json:"PreviousHash"`
	CurrentBlockHash  string `json:"CurrentBlockHash"`
	DataHash          string `json:"DataHash"`
}

func GetBlockinfos(peerip string, startNum uint64) []*Blockinfo {

	fmt.Println("@@@@chenqiao :", PeerStatusMap[peerip])
	height := PeerStatusMap[peerip].Height

	length := GetBlocksLength(peerip)

	if startNum > height-1 {
		return []*Blockinfo{}
	} else if startNum < height-length {
		blockindex := uint64(0)
		return MakeBlockinfo(peerip, blockindex, length)
	} else {
		blockindex := startNum - (height - length)
		return MakeBlockinfo(peerip, blockindex, length)
	}

}

func MakeBlockinfo(id string, blockindex uint64, length uint64) []*Blockinfo {

	blockinfos := []*Blockinfo{}

	for i := blockindex; i < length; i++ {
		blockinfo := &Blockinfo{}
		blockinfo.BlockNumber = PeerStatusMap[id].Mblocks[i].Header.Number
		blockinfo.TransactionsCount = GetPeerTransNumber(id, i)
		blockinfo.PreviousHash = PeerStatusMap[id].Mblocks[i].Header.PreviousHash
		blockinfo.DataHash = PeerStatusMap[id].Mblocks[i].Header.DataHash
		blockinfo.CurrentBlockHash = PeerStatusMap[id].Mblocks[i].Header.CurrentHash

		blockinfos = append(blockinfos, blockinfo)
	}

	return blockinfos

}
