package models

import (
	"fmt"
	"github.com/astaxie/beego"
	// "strconv"
	// "encoding/json"
	"sync"
	"time"
)

var MapMutex sync.Mutex
var AllChannelPeerStatusMap map[string](map[string]*PeerMessage)
var FabricTimeCycle, _ = beego.AppConfig.Int64("fabricTimeCycle")
var CheckTimeCycle, _ = beego.AppConfig.Int64("checkTimeCycle")
var BufferTime, _ = beego.AppConfig.Int64("bufferTime")

func IsChange(before *PeerMessage, after *PeerMessage) bool {
	if before.Status != after.Status || before.Name != after.Name || before.PeerIp != after.PeerIp || before.Height != after.Height {
		return true
	} else { //找到最高块的hash
		highest := int(before.Height - 1)
		if before.Mblocks[highest].Header.CurrentHash == after.Mblocks[highest].Header.CurrentHash {
			return false
		} else {
			return true
		}

	}
}

func BaseCheck(args ...interface{}) bool {
	if len(args) == 0 {
		if len(AllChannelPeerStatusMap) == 0 {
			// fmt.Println("the map is nil!!")
			return false
		}
		// fmt.Println("the map is not nil ", PeerStatusMap)
		return true
	} else {
		_, ok := AllChannelPeerStatusMap[args[0].(string)]
		if ok {
			return true
		}
		return false
	}

}

func GetTransByTxID(peerid string, txid string, chainid string) *TransData {
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	for blockindex := 0; blockindex < len(PeerStatusMap[peerid].Mblocks); blockindex++ {
		tmptrans := PeerStatusMap[peerid].Mblocks[blockindex].Data.Data
		length := len(tmptrans)
		for i := 0; i < length; i++ {
			if txid == tmptrans[i].Txid {
				return tmptrans[i]
			}
		}
	}
	return nil
}

func GetPeerTransNumber(id string, blockindex uint64, chainid string) uint64 {
	ans := uint64(0)
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	length := len(PeerStatusMap[id].Mblocks[blockindex].Data.Data)
	for i := 0; i < length; i++ {
		if len(PeerStatusMap[id].Mblocks[blockindex].Data.Data[i].Txid) > 0 {
			ans++
		}
	}
	return ans
}

func GetBlocksLength(peerId string, chainid string) uint64 {
	// peerID := GetMaxPeer()
	return uint64(len(AllChannelPeerStatusMap[chainid][peerId].Mblocks))
}

func GetMaxPeer(chainid string) (string, uint64) {
	var maxPeerId string
	var maxHeight uint64
	maxHeight = 0
	PeerStatusMap := AllChannelPeerStatusMap[chainid]
	for key := range PeerStatusMap {
		if PeerStatusMap[key].Height > maxHeight {
			maxHeight = PeerStatusMap[key].Height
			maxPeerId = key
		} else if PeerStatusMap[key].Height == maxHeight {
			if maxPeerId > key {
				maxPeerId = key
			}
		}
	}
	return maxPeerId, maxHeight
}

func CheckEnable() {
	for {
		now := time.Now().Unix()
		MapMutex.Lock()
		for key, _ := range AllChannelPeerStatusMap {
			PeerStatusMap := AllChannelPeerStatusMap[key]
			for _, value := range PeerStatusMap {
				// fmt.Println("@@@@chenqiao Check Time: ", now-value.Time)
				if now-value.Time > (FabricTimeCycle + BufferTime) {
					value.Status = 0

					//peer状态需要更新
					fmt.Println("it's time to send DATA!!")
					SocketsProperty.Update()
					fmt.Println("TTTTTT", SocketsProperty)
				}
			}

		}
		MapMutex.Unlock()
		time.Sleep(time.Duration(CheckTimeCycle) * 1e9)
	}

}
