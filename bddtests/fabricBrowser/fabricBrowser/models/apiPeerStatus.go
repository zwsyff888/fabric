package models

import (
	// "fmt"
	// "strconv"
	"sort"
)

type ApiPeerStatus struct {
	Address string `json:"address"`
	Height  uint64 `json:"height"`
	Key     int64  `json:"key"`
	Name    string `json:"name"`
	Status  int64  `json:"status"`
}

type ApiPeerList []*ApiPeerStatus

func GetPeerStatusMap() ApiPeerList {

	ans := ApiPeerList{}
	// var index int64
	index := 0
	for key := range PeerStatusMap {
		ans = append(ans, makePeerStatus(key, index))
		index = index + 1
	}

	sort.Sort(ans)
	return ans
}

func (list ApiPeerList) Less(i, j int) bool {
	if list[i].Name < list[j].Name {
		return true
	} else {
		return false
	}
}

func (list ApiPeerList) Swap(i, j int) {
	var temp *ApiPeerStatus = list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list ApiPeerList) Len() int {
	return len(list)
}

func makePeerStatus(id string, index int) *ApiPeerStatus {
	// strindex := strconv.Itoa(index)

	ans := &ApiPeerStatus{}
	ans.Address = PeerStatusMap[id].PeerIp
	ans.Height = PeerStatusMap[id].Height
	ans.Key = int64(index)
	ans.Name = PeerStatusMap[id].Name
	ans.Status = PeerStatusMap[id].Status

	return ans
}
