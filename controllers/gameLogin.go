package controllers

import (
	"encoding/json"
	"fmt"
	"goutils"
	"io/ioutil"
	"log"
	"strings"

	"sltapi/models"
	"sltapi/storage"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"

	"net/http"
)

// Operations about object
//改了Routing 和controler的名稱 要跑過Bee Run在launch才會生效
type GameLoginController struct {
	beego.Controller
}

//
// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *GameLoginController) Post() {
	var inputObj models.APIGameLoginInput
	log.Println("GameLoginController RequestBody:", string(o.Ctx.Input.RequestBody))
	json.Unmarshal(o.Ctx.Input.RequestBody, &inputObj) //把JSON值塞進Object去
	//Login(驗證)
	verifyLogin(&inputObj)
	//o.Data["json"] = inputObj                          //輸出JSON
	//o.ServeJSON()                                      //輸出JSON
}

func verifyLogin(inputObj *models.APIGameLoginInput) {
	//取得BUInfo
	storage.DB_GetBUInfo(inputObj.BUCode)
	currentBU := models.GetBUInfoByBuCode(inputObj.BUCode)
	log.Println("currentBU.Login_url:", currentBU.Login_url)

	//尋找 Partner 驗證API URL 打過去驗證
	requestJsonStr := fmt.Sprintf(`
	{
		"Token":"%v"
	}
	`, inputObj.Token)

	resp, err := http.Post(currentBU.Login_url,
		"application/x-www-form-urlencoded",
		strings.NewReader(requestJsonStr))

	goutils.CheckErr(err)

	//得到回應
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	goutils.CheckErr(err)

	var outputSuccessObj models.APIPartnerLoginSuccessOutput
	var outputFailObj models.APIPartnerLoginFailOutput
	err = json.Unmarshal(body, &outputSuccessObj) //把JSON值塞進Object去
	goutils.CheckErr(err)
	if outputSuccessObj.Statuscode != "000000" {
		//驗證失敗
		err = json.Unmarshal(body, &outputFailObj) //把JSON值塞進Object去
		log.Printf("verify fail rawdata: %v, Statuscode:%v, Msg:%v", string(body), outputFailObj.Statuscode, outputFailObj.Msg)
		goutils.CheckErr(err)
	} else {
		//驗證成功
		log.Printf("verify success rawdata: %v, Statuscode:%v, Membercode:%v, Balance:%v", string(body), outputSuccessObj.Statuscode, outputSuccessObj.Membercode, outputSuccessObj.Balance)
		//GetGameInfo

		//回傳參考：
		//{"Body":{"CoinSizeList":[0.01,0.02,0.05,0.10,0.25,0.50,1.0,2.0,5.0],"Currency":{"Code":"RMB","ISOCode":"CNY"},"GameToken":"e0314384-8f9a-43f3-b337-e5f8318fabd0","MaxMultiplier":10,"MinMultiplier":1,"Status":{"BaseCredit":50,"CoinSize":0.01,"CurrentMode":0,"GameState":5,"MaxPayout":0.0,"MeetMaxPayout":false,"MemberBalance":1128055.20,"MemberCredits":112805520.0,"Modes":[{"CurrentPlay":0,"CurrentPoint":0,"CurrentRound":0,"GamePlay":{"Entities":[["N006","N009","N010","N008"],["N010","N010","N003","N008"],["N005","N009","N006","N004"],["N006","N007","N003","N007"],["N006","N008","N004","N001"]],"TotalCreditPayout":0.0},"TotalPlay":0,"TotalPoint":0,"TotalRound":0,"Type":0},{"CurrentPlay":0,"CurrentPoint":0,"CurrentRound":0,"TotalPlay":0,"TotalPoint":0,"TotalRound":0,"Type":1}],"Multiplier":1,"TotalCreditPayout":0.0}},"Messages":[{"Args":["160901020300001446","5206650"],"ID":"SLT0205","Type":2}],"Status":"Success"}
		storage.DB_GetGameInfo(inputObj.BUCode, inputObj.GameSN)

	}

	//var loginUrl:=
	//if inputObj.BUCode
}
