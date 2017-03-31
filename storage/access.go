package storage

import (
	"database/sql"
	"log"
	"sltapi/models"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}
var err error

func init() {
	//db, err = sql.Open("mysql", "root:y0701003@tcp(localhost:3306)/slt") //公司
	db, err = sql.Open("mysql", "allenslt:y0701003@tcp(allen.com:3306)/slt")
	log.Println("Hello!!!")
}
func checkErr(err error) {
	if err != nil {
		panic(err)
		//log.Println("DB Error:", err.Error())
	}
}
func DB_GetBUInfo(bucode string) {
	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	checkErr(err)
	dbQueryStr := `
	SELECT login_url,
	placebet_url,
	settlebet_url,
	getbalance_url,
	cancelbet_url 
	FROM bu WHERE bucode=?
	`
	rows, err := db.Query(dbQueryStr, bucode)
	defer rows.Close()
	bu := new(models.BUInfo)
	for rows.Next() { //有下一筆就會一直true下去
		err = rows.Scan(&bu.Login_url, &bu.Placebet_url, &bu.Settlebet_url, &bu.Getbalance_url, &bu.Cancelbet_url)
		checkErr(err)
		break
	}
	//最后释放tx内部的连接
	err = tx.Commit()

	log.Println("bu.Login_url:", bu.Login_url)
	checkErr(err)
}

func Insert() {

	//方式4 insert

	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	checkErr(err)

	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	_, err = tx.Exec("INSERT INTO wager(gamesn,bucode,membercode,stake_c,stake_m,payout_c,payout_m) values(?,?,?,?,?,?,?)", 1, "BU001", "AllenLiao", 1000.5, 1000.5, 1000.5, 1000.5)
	checkErr(err)
	//最后释放tx内部的连接
	err = tx.Commit()
	checkErr(err)

}
