package models

var (
	result *E6Result
)

type E6Result struct {
	GameInfo         *GameInfo
	GameStatusInfo   *GameStatusInfo
	GamePlayInfoList []*GamePlayInfo
	//gameLevelInfo    *GameLevelInfo
}

func init() {
	result = new(E6Result)
	result.GameInfo = new(GameInfo)
	result.GameInfo.MaxMultiplier = 1
	result.GameInfo.MinMultiplier = 20

	result.GameStatusInfo = new(GameStatusInfo)
	result.GameStatusInfo.gameState = 1

	gamePlayInfo := new(GamePlayInfo)

	//init gamePlayInfo.Result
	//二維陣列 的宣告方法
	reelNum := 5
	reelSymbolNum := 4
	gamePlayInfo.Result = make([][]string, reelNum)
	for reelIdx := 0; reelIdx < reelNum; reelNum++ {
		gamePlayInfo.Result[reelIdx] = make([]string, reelSymbolNum)
		for reelSymbolIdx := 0; reelSymbolIdx < reelSymbolNum; reelSymbolIdx++ {
			gamePlayInfo.Result[reelIdx][reelSymbolIdx] = "N001"

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

func GetResult() *E6Result {
	return result
}
