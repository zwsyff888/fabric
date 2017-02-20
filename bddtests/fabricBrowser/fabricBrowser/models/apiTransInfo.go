package models

import (
	"strconv"
)

type APITransInfo struct {
	TransactionNo string
	BelongToBlock string
	ChainCodeID   string
	Type          string
}

func GetTransinfo(peerip string, startNum uint64, chainid string) []*APITransInfo {
	PeerStatusMap := AllChannelPeerStatusMap[chainid]

	height := PeerStatusMap[peerip].Height

	length := GetBlocksLength(peerip, chainid)

	if startNum > height-1 {
		return []*APITransInfo{}
	} else if startNum < height-length {
		blockindex := uint64(0)
		return MakeTransinfo(peerip, blockindex, length, chainid)
	} else {
		blockindex := startNum - (height - length)
		return MakeTransinfo(peerip, blockindex, length, chainid)
	}

}

func MakeTransinfo(peerid string, blockindex uint64, length uint64, chainid string) []*APITransInfo {
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	transinfos := []*APITransInfo{}

	for i := blockindex; i < length; i++ {

		dataLength := len(PeerStatusMap[peerid].Mblocks[i].Data.Data)
		for j := 0; j < dataLength; j++ {
			tmpData := PeerStatusMap[peerid].Mblocks[i].Data.Data[j]
			if len([]rune(tmpData.Txid)) > 0 {
				tmptransinfo := &APITransInfo{}
				tmptransinfo.BelongToBlock = strconv.Itoa(int(i))
				tmptransinfo.ChainCodeID = tmpData.ChainCodeID
				tmptransinfo.TransactionNo = tmpData.Txid
				tmptransinfo.Type = tmpData.Type

				transinfos = append(transinfos, tmptransinfo)
			} else {
				continue
			}

		}
	}
	return transinfos
}
