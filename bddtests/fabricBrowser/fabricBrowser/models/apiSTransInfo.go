package models

import ()

type APISTransInfo struct {
	Method string          `json:"method"`
	Data   []*APITransInfo `json:"data"`
}

func NilSAPITransInfo() *APISTransInfo {
	return &APISTransInfo{Method: "transinfo"}
}

func equalAPITransInfo(before *APITransInfo, after *APITransInfo) bool {
	if before.BelongToBlock == after.BelongToBlock && before.ChainCodeID == after.ChainCodeID && before.TransactionNo == after.TransactionNo && before.Type == after.Type {
		return true
	}
	return false
}

func APISCheckTransInfo(before []*APITransInfo, chainid string) (bool, []*APITransInfo) {
	var ans []*APITransInfo
	if !BaseCheck(chainid) {
		ans = []*APITransInfo{}
	} else {
		peerid, _ := GetMaxPeer(chainid)

		blockNum := uint64(0)
		ans = GetTransinfo(peerid, blockNum, chainid)
	}
	if len(before) != len(ans) {
		return true, ans
	} else {
		for i := 0; i < len(ans); i++ {
			if !equalAPITransInfo(before[i], ans[i]) {
				return true, ans
			}
		}
		return false, before
	}
}
