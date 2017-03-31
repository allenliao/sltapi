package models

var (
	result *ResultE6
)

type ResultE6 struct {
	GameInfo         *GameInfo
	GameStatusInfo   *GameStatusInfo
	GamePlayInfoList []*GamePlayInfo
	//gameLevelInfo    *GameLevelInfo
}

func init() {
	result = new(ResultE6)
	result.GameInfo = new(GameInfo)
	result.GameInfo.MaxMultiplier = 1
	result.GameInfo.MinMultiplier = 20

	result.GameStatusInfo = new(GameStatusInfo)
	result.GameStatusInfo.GameState = 1

	gamePlayInfo := new(GamePlayInfo)

	//init gamePlayInfo.Result
	//二維陣列 的宣告方法
	//測試英文commit
	reelNum := 5
	reelSymbolNum := 4

	gamePlayInfo.SymbolResult = make([][]string, reelNum)
	for reelIdx := range gamePlayInfo.SymbolResult {
		gamePlayInfo.SymbolResult[reelIdx] = make([]string, reelSymbolNum)
		for reelSymbolIdx := range gamePlayInfo.SymbolResult[reelIdx] {
			gamePlayInfo.SymbolResult[reelIdx][reelSymbolIdx] = "N001"

		}
	}

	//init gamePlayInfo.PayoutResultList
	payoutResult := new(PayoutResult)
	payoutResult.CreditPayout = 99999999999
	payoutResult.TotalCreditPayout = 999999999990
	payoutResult.Multiplier = 10
	gamePlayInfo.PayoutResultList = append(gamePlayInfo.PayoutResultList, payoutResult)

	result.GamePlayInfoList = append(result.GamePlayInfoList, gamePlayInfo)

}

func GetResult() *ResultE6 {
	return result
}
