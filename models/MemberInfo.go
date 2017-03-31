package models

type MemberInfo struct {
	Token         uint8
	MemberCode    string
	CoinSize      float32
	Multiplier    uint8
	MemberBalance uint64 //真錢
	MemberCredit  uint64 //計算過coin size之後的 member credits.
	BUCode        string
}
