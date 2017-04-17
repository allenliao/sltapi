package models

type APIGameLoginInput struct {
	Token  string
	BUCode string
	GameSN uint8
}

type APIPartnerLoginSuccessOutput struct {
	Statuscode string
	Membercode string
	Balance    uint64
}
type APIPartnerLoginFailOutput struct {
	Statuscode string
	Msg        string
}

type APIGameLoginOutput struct {
	Body     *APILoginOutputBody
	Messages *APILoginOutputMessages
}

type APILoginOutputBody struct {
	CoinSizeList  *[]float32
	MaxMultiplier uint8
	MinMultiplier uint8
	Status        *GameStatusInfo
}

type APILoginOutputMessages struct {
	Args *[]string
	ID   string
	Type int8
}

func init() {

}
