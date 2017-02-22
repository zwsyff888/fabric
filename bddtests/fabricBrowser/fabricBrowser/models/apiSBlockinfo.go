package models

import ()

type SBlockinfo struct {
	Method string       `json:"method"`
	Data   []*Blockinfo `json:"data"`
}

func NilSBlockinfos() *SBlockinfo {
	return &SBlockinfo{Method: "blockinfo"}
}

func APISCheckBlockinfos(before []*Blockinfo, chainid string) (bool, []*Blockinfo) {
	var ans []*Blockinfo

	if !BaseCheck(chainid) {
		ans = []*Blockinfo{}
	} else {
		peerip, _ := GetMaxPeer(chainid)
		startNum := uint64(0)
		ans = GetBlockinfos(peerip, startNum, chainid)
	}

	if len(before) != len(ans) {
		return true, ans
	} else {
		for i := 0; i < len(before); i++ {
			if before[i].CurrentBlockHash != ans[i].CurrentBlockHash {
				return true, ans
			}
		}
		return false, before
	}

}
