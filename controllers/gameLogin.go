package controllers

import (
	"encoding/json"
	"log"

	"sltapi/models"
	"sltapi/storage"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
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
	json.Unmarshal(o.Ctx.Input.RequestBody, &inputObj) //把值塞進去
	//Login(驗證)
	verifyLogin(&inputObj)
	//o.Data["json"] = inputObj                          //輸出JSON
	//o.ServeJSON()                                      //輸出JSON
}

func verifyLogin(inputObj *models.APIGameLoginInput) {
	//取得BUInfo
	storage.DB_GetBUInfo(inputObj.BUCode)

	//尋找 Partner 驗證API URL 打過去驗證

	//var loginUrl:=
	//if inputObj.BUCode
}
