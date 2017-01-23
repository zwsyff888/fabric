package models

import ()

func GetBlockDetail(peerId string, blockNum uint64) *Mblock {
	height := PeerStatusMap[peerId].Height

	if blockNum > height {
		return nil
	} else if blockNum > height-GetBlocksLength(peerId) {
		blockindex := blockNum - (height - GetBlocksLength(peerId))
		//直接从内存中拿
		return PeerStatusMap[peerId].Mblocks[blockindex]
	} else {
		//去rpc中拿
		pbBlocks := QueryClient(blockNum)
		mblock := NewMblock()
		mblock.initMblock(pbBlocks)
		return mblock
	}

}
