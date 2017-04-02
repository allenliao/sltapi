package storage

import (
	"database/sql"
	"log"
	"sltapi/models"

	"goutils"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}
var err error

func init() {
	//db, err = sql.Open("mysql", "root:y0701003@tcp(localhost:3306)/slt") //公司
	db, err = sql.Open("mysql", "allenslt:y0701003@tcp(allen.com:3306)/slt")
	log.Println("Hello!!!")
}

func DB_GetBUInfo(bucode string) {
	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	goutils.CheckErr(err)

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
	models.CurrentBU = new(models.BUInfo)
	//bu := *models.CurrentBU
	for rows.Next() { //有下一筆就會一直true下去
		//err = rows.Scan(&bu.Login_url, &bu.Placebet_url, &bu.Settlebet_url, &bu.Getbalance_url, &bu.Cancelbet_url)

		err = rows.Scan(&models.CurrentBU.Login_url,
			&models.CurrentBU.Placebet_url,
			&models.CurrentBU.Settlebet_url,
			&models.CurrentBU.Getbalance_url,
			&models.CurrentBU.Cancelbet_url)
		goutils.CheckErr(err)
		break
	}
	//最后释放tx内部的连接
	err = tx.Commit()
	goutils.CheckErr(err)
}

func Insert() {

	//方式4 insert

	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	goutils.CheckErr(err)

	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	_, err = tx.Exec("INSERT INTO wager(gamesn,bucode,membercode,stake_c,stake_m,payout_c,payout_m) values(?,?,?,?,?,?,?)", 1, "BU001", "AllenLiao", 1000.5, 1000.5, 1000.5, 1000.5)
	goutils.CheckErr(err)
	//最后释放tx内部的连接
	err = tx.Commit()
	goutils.CheckErr(err)

}
