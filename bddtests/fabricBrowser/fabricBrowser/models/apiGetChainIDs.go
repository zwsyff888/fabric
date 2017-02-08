package models

import ()

func GetChainIDs() []string {
	ans := []string{}
	for key := range AllChannelPeerStatusMap {
		ans = append(ans, key)
	}
	return ans
}
