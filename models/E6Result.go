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
	result.GamePlayInfoList = append(result.GamePlayInfoList, gamePlayInfo)
	payoutResult := new(PayoutResult)
	gamePlayInfo.PayoutResultList = append(gamePlayInfo.PayoutResultList, payoutResult)
	payoutResult.CreditPayout = 99999999999
	payoutResult.TotalCreditPayout = 999999999990
	payoutResult.Multiplier = 10

}

func GetResult() *E6Result {
	return result
}
