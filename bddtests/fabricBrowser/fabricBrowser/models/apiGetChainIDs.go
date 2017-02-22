package models

import (
	"sort"
)

type APIChainsIDs []string

func (list APIChainsIDs) Less(i, j int) bool {
	if list[i] < list[j] {
		return true
	} else {
		return false
	}
}

func (list APIChainsIDs) Swap(i, j int) {
	var temp string = list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list APIChainsIDs) Len() int {
	return len(list)
}

func GetChainIDs() APIChainsIDs {
	ans := APIChainsIDs{}
	for key := range AllChannelPeerStatusMap {
		ans = append(ans, key)
	}
	sort.Sort(ans)
	return ans
}
