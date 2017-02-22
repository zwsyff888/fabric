package models

import (
	"github.com/astaxie/beego/logs"
)

func GetBlockDetail(peerId string, blockNum uint64, chainid string) *Mblock {
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	height := PeerStatusMap[peerId].Height
	// height := uint64(1)

	if blockNum >= height {
		return nil
	} else if blockNum >= height-GetBlocksLength(peerId, chainid) {
		// fmt.Println("hehe, come here!")
		logs.Debug("hehe, come here!")
		blockindex := blockNum - (height - GetBlocksLength(peerId, chainid))
		//直接从内存中拿
		return PeerStatusMap[peerId].Mblocks[blockindex]
		// return nil
	} else {
		//去rpc中拿
		// fmt.Println("hehe, come rpc!!")
		//暂时不考虑
		/*
			logs.Debug("hehe, come rpc!!!")
			// return nil
			pbBlocks := QueryClient(blockNum)
			mblock := NewMblock()
			mblock.initMblock(pbBlocks)
			return mblock*/
		return nil

	}

}
