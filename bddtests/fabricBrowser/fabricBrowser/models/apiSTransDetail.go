package models

import ()

type STransData struct {
	Method string     `json:"method"`
	Data   *TransData `json:"data"`
}

func NilSTransData() *STransData {
	return &STransData{Method: "transdetail"}
}

func equalTransData(before *TransData, after *TransData) bool {
	if before.ChainCodeID == after.ChainCodeID && before.ChainID == after.ChainID && before.Nonce == after.Nonce && before.Input == after.Input && before.Result == after.Result && before.Signature == after.Signature && before.Txid == after.Txid && before.Type == after.Type {
		return true
	}
	return false
}

func APISCheckTransData(before *TransData, chainid string, txid string) (bool, *TransData) {
	var ans *TransData
	if !BaseCheck(chainid) {
		ans = &TransData{}
	} else {
		peerid, _ := GetMaxPeer(chainid)
		ans = GetTransDetail(peerid, txid, chainid)
	}
	if before == nil && ans == nil {
		return false, before
	}
	if before == nil || ans == nil {
		return true, ans
	}

	if !equalTransData(before, ans) {
		return true, ans
	}
	return false, before

}
