package models

var CurrentBU *BUInfo

type BUInfo struct {
	Login_url      string
	Placebet_url   string
	Settlebet_url  string
	Getbalance_url string
	Cancelbet_url  string
	BUCode         string
}
