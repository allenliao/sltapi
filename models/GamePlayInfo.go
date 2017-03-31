package models

type GamePlayInfo struct {
	TotalCreditPayout uint64 //這一個play的所有CreditPayout
	SymbolResult      [][]string
	PayoutResultList  []*PayoutResult
}

type PayoutResult struct {
	TotalCreditPayout uint64 //這一個中獎線的結果乘上FREE倍數後的CreditPayout
	CreditPayout      uint64
	Multiplier        uint8
	Paylines          uint8 //幾連線
	SymbolID          string
	Ways              [][]int8 //中幾線
}
