package models

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "strconv"
	// "encoding/json"
	"sync"
	"time"
)

var MapMutex sync.Mutex
var PeerStatusMap map[string]*PeerMessage
var AllChannelPeerStatusMap map[string](map[string]*PeerMessage)
var FabricTimeCycle, _ = beego.AppConfig.Int64("fabricTimeCycle")
var CheckTimeCycle, _ = beego.AppConfig.Int64("checkTimeCycle")
var BufferTime, _ = beego.AppConfig.Int64("bufferTime")

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
					value.Status = 2
				}
			}
			// fmt.Println("@@@@chenqiao Check : ", PeerStatusMap["172.18.0.3"])

		}
		// fmt.Println("@@@chenqiao Check Peer: ", PeerStatusMap)
		// fmt.Println("@@@chenqiao Time: ", CheckTimeCycle)
		MapMutex.Unlock()
		time.Sleep(time.Duration(CheckTimeCycle) * 1e9)
	}

}
