package models

import ()

type DataNode struct {
}

type SoloRate struct {
	Time *[]uint64 `json:"time"`
	Rate *[]int64  `json:"rate"`
}

type BTRate struct {
	Block *[]SoloRate `json:"block"`
	Tx    *[]SoloRate `json:"tx"`
}

func GetRate() {

}
