package models

var GameInfoList map[uint8]*GameInfo = make(map[uint8]*GameInfo)

type GameInfo struct {
	MinMultiplier uint8
	MaxMultiplier uint8
	CoinSizeList  []float32
	BaseCredit    uint8
	EngineSN      uint8
}
