package models

var BUInfoList map[string]*BUInfo = make(map[string]*BUInfo)

type BUInfo struct {
	Login_url      string
	Placebet_url   string
	Settlebet_url  string
	Getbalance_url string
	Cancelbet_url  string
	BUCode         string
}

func GetBUInfoByBuCode(bucode string) *BUInfo {
	return BUInfoList[bucode]
}
