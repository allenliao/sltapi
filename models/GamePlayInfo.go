package models

type GamePlayInfo struct {
	Result           [][]string
	PayoutResultList []*PayoutResult
}

type PayoutResult struct {
	TotalCreditPayout uint64
	CreditPayout      uint64
	Multiplier        uint8
}
