package models

import ()

func GetTransDetail(peerid string, txid string, chainid string) *TransData {
	return GetTransByTxID(peerid, txid, chainid)
}
