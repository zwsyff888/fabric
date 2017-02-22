package models

import ()

type SMblock struct {
	Method string  `json:"method"`
	Data   *Mblock `json:"data"`
}

func NilSMblock() *SMblock {
	return &SMblock{Method: "blockdetail"}
}

func equalHeader(before MblockHeader, ans MblockHeader) bool {
	if before.CurrentHash == ans.CurrentHash && before.DataHash == ans.DataHash && before.Number == ans.Number && before.PreviousHash == ans.PreviousHash {
		return true
	}
	return false
}

func APISCheckSMblock(before *Mblock, chainid string, blockNum uint64) (bool, *Mblock) {
	var ans *Mblock
	if !BaseCheck(chainid) {
		ans = &Mblock{}
	} else {
		peerip, _ := GetMaxPeer(chainid)
		ans = GetBlockDetail(peerip, blockNum, chainid)
	}

	if before == nil && ans == nil {
		return false, before
	}
	if before == nil || ans == nil {
		return true, ans
	}

	if !equalHeader(ans.Header, before.Header) {
		return true, ans
	}

	return false, before
}
