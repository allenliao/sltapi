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
	log.Println("CurrentBU.Login_url:", models.CurrentBU.Login_url)

	//尋找 Partner 驗證API URL 打過去驗證
	requestJsonStr := fmt.Sprintf(`
	{
		"Token":"%v"
	}
	`, inputObj.Token)

	resp, err := http.Post(models.CurrentBU.Login_url,
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

		//models.GetGameInfo(GameSN)

	}

	//var loginUrl:=
	//if inputObj.BUCode
}
