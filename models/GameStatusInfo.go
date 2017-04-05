package models

type GameStatusInfo struct {
	GameState         uint8
	TotalCreditPayout uint64 //整個gameRound的totalCreditPayout
	CurrentMode       uint8
	CurrentFreeSpin   uint8
}
