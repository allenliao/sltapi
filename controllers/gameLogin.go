package controllers

import (
	"database/sql"
	"encoding/json"
	"log"

	"sltapi/models"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}
var err error

func init() {
	//db, err = sql.Open("mysql", "root:y0701003@tcp(localhost:3306)/slt")//公司
	db, err = sql.Open("mysql", "allenslt:y0701003@tcp(allen.com:3306)/slt")
}

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
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *GameLoginController) Get() {
	insert()
	obb := models.GetResult()

	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		_, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			//ob.ObjectId = objectId
			o.Data["json"] = obb
		}
	}
	o.ServeJSON()
}

func insert() {

	//方式4 insert

	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	if err != nil {
		log.Println("DB Error:", err.Error())
	}

	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	_, err3 := tx.Exec("INSERT INTO wager(gamesn,bucode,membercode,stake_c,stake_m,payout_c,payout_m) values(?,?,?,?,?,?,?)", 1, "BU001", "AllenLiao", 1000.5, 1000.5, 1000.5, 1000.5)
	if err3 != nil {
		log.Println("DB Error:", err3.Error())
	}
	//最后释放tx内部的连接
	err2 := tx.Commit()
	if err2 != nil {
		log.Println("DB Error:", err2.Error())
	}

}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *GameLoginController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *GameLoginController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *GameLoginController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
