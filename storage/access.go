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

func DB_GetBUInfo(bucode string) *models.BUInfo {
	if models.BUInfoList[bucode] == nil {
		//取回來快取
		models.BUInfoList[bucode] = GetBUInfo(bucode)
	}
	return models.BUInfoList[bucode]
}

func GetBUInfo(bucode string) *models.BUInfo {
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
	bUInfo := new(models.BUInfo)

	for rows.Next() { //有下一筆就會一直true下去
		err = rows.Scan(&bUInfo.Login_url,
			&bUInfo.Placebet_url,
			&bUInfo.Settlebet_url,
			&bUInfo.Getbalance_url,
			&bUInfo.Cancelbet_url)
		goutils.CheckErr(err)
		break
	}
	//最后释放tx内部的连接
	err = tx.Commit()
	goutils.CheckErr(err)
	return bUInfo

}

func DB_GetGameInfo(gameSN uint8) *models.GameInfo {
	if models.GameInfoList[gameSN] == nil {
		//取回來快取
		models.GameInfoList[gameSN] = GetGameInfo(gameSN)
	}
	return models.GameInfoList[gameSN]
}

func GetGameInfo(gamesn uint8) *models.GameInfo {
	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	goutils.CheckErr(err)

	dbQueryStr := `
	SELECT min_multiplier,
	max_multiplier,
	basecredit,
	engineSN
	FROM game WHERE gamesn=?
	`
	rows, err := db.Query(dbQueryStr, gamesn)
	defer rows.Close()
	//這裡要處理 CoinSizeList 看DB怎麼 回傳兩個 result

	gameInfo := new(models.GameInfo)
	/*
		for rows.Next() { //有下一筆就會一直true下去

			err = rows.Scan(&gameInfo.Login_url,
				&gameInfo.Placebet_url,
				&gameInfo.Settlebet_url,
				&gameInfo.Getbalance_url,
				&gameInfo.Cancelbet_url)
			goutils.CheckErr(err)
			break
		}
	*/
	//最后释放tx内部的连接
	err = tx.Commit()
	goutils.CheckErr(err)

	return gameInfo
}
