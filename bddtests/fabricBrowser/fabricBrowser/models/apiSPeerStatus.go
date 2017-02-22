package models

import ()

type APISPeerList struct {
	Method string      `json:"method"`
	Data   ApiPeerList `json:"data"`
}

func NilSPeerStatus() *APISPeerList {
	return &APISPeerList{Method: "peers"}
}

func equalApiPeerStatus(before *ApiPeerStatus, after *ApiPeerStatus) bool {
	if before.Address == after.Address && before.Height == after.Height && before.Key == after.Key && before.Name == after.Name && before.Status == after.Status {
		return true
	}
	return false
}

func APISCheckSPeerList(before ApiPeerList, chainid string) (bool, ApiPeerList) {
	var ans ApiPeerList
	if !BaseCheck(chainid) {
		ans = ApiPeerList{}
	} else {
		ans = GetPeerStatusMap(chainid)
	}
	if len(before) != len(ans) {
		return true, ans
	} else {
		for i := 0; i < len(ans); i++ {
			if !equalApiPeerStatus(before[i], ans[i]) {
				return true, ans
			}
		}
		return false, before
	}
}
