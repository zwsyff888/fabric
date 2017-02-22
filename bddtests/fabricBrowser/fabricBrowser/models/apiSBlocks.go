package models

import ()

type APISBlock struct {
	Method string      `json:"method"`
	Data   []*ApiBlock `json:"data"`
}

func NilSBlock() *APISBlock {
	return &APISBlock{Method: "blocks"}
}

func APISCheckSBlock(before []*ApiBlock, chainid string) (bool, []*ApiBlock) {
	var ans []*ApiBlock
	if !BaseCheck(chainid) {
		ans = []*ApiBlock{}
	} else {
		startnum := uint64(0)
		ans = GetBlocks(startnum, chainid)
	}
	if len(before) != len(ans) {
		return true, ans
	} else {
		for i := 0; i < len(before); i++ {
			if before[i].BlockNumber != ans[i].BlockNumber || before[i].Count != ans[i].Count {
				return true, ans
			}
		}
		return false, before
	}

}
