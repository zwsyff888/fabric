package models

type APISChainIDs struct {
	Method string       `json:"method"`
	Data   APIChainsIDs `json:"data"`
}

func NilAPISChainIDs() *APISChainIDs {
	return &APISChainIDs{Method: "chainids"}
}

func APISCheckChainIDs(before APIChainsIDs) (bool, APIChainsIDs) {
	var ans APIChainsIDs

	if !BaseCheck() {
		ans = APIChainsIDs{}
	} else {
		ans = GetChainIDs()
	}

	if len(before) != len(ans) {
		return true, ans
	} else {
		for i := 0; i < len(before); i++ {
			if before[i] != ans[i] {
				return true, ans
			}
		}
		return false, before
	}

}
