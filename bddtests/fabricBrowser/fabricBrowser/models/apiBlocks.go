package models

import ()

type ApiBlock struct {
	BlockNumber uint64 `json:"blockNumber"`
	Count       uint64 `json:"count"`
}

func GetBlocks(startNum uint64, chainid string) []*ApiBlock {

	peerId, height := GetMaxPeer(chainid)
	length := GetBlocksLength(peerId, chainid)
	if startNum > height-1 {
		return []*ApiBlock{}
	} else if startNum < height-length {
		blockindex := uint64(0)
		return MakeBlock(peerId, blockindex, length, chainid)
	} else {
		blockindex := startNum - (height - length)
		return MakeBlock(peerId, blockindex, length, chainid)
	}

}

func MakeBlock(id string, blockindex uint64, length uint64, chainid string) []*ApiBlock {
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	blocks := []*ApiBlock{}

	for i := blockindex; i < length; i++ {
		ans := &ApiBlock{}

		ans.BlockNumber = PeerStatusMap[id].Mblocks[i].Header.Number
		ans.Count = GetPeerTransNumber(id, i, chainid)

		blocks = append(blocks, ans)

	}

	return blocks
}
