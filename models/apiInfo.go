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

func init() {

}
