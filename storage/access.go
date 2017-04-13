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
	db, err = sql.Open("mysql", "allenslt:y0701003@tcp(allen.com:3306)/slt") //>>公司的 VPN 要關掉才能連
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
	dbQueryStr := `
	SELECT login_url,
	placebet_url,
	settlebet_url,
	getbalance_url,
	cancelbet_url 
	FROM bu WHERE bucode=?
	`
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(bucode)
	goutils.CheckErr(err)
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
	dbQueryStr := `
	SELECT min_multiplier,
	max_multiplier,
	basecredit,
	engineSN
	FROM game WHERE gamesn=?
	`

	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(gamesn)
	goutils.CheckErr(err)
	defer rows.Close()

	gameInfo := new(models.GameInfo)
	for rows.Next() { //有下一筆就會一直true下去

		err = rows.Scan(&gameInfo.MinMultiplier,
			&gameInfo.MaxMultiplier,
			&gameInfo.BaseCredit,
			&gameInfo.EngineSN)
		goutils.CheckErr(err)
		break
	}
	log.Printf("gamesn:%v,  MinMultiplier:%v, MaxMultiplier:%v, BaseCredit:%v, EngineSN:%v,",
		gamesn,
		gameInfo.MinMultiplier,
		gameInfo.MaxMultiplier,
		gameInfo.BaseCredit,
		gameInfo.EngineSN)
	return gameInfo
}

/*
var BUGameCoinSizeList map[string][uint8]*[]float32 = make(map[string][uint8]*[]float32)

func (s *someStruct) Set(i int, k, v string) {
    child, ok := s.nestedMap[i]
    if !ok {
        child = map[uint8]*[]float32 {}
        s.nestedMap[i] = child
    }
    child[k] = v
}
*/

func GetBUGameCoinSizeList(bucode string, gamesn uint8) *[]float32 {
	//CoinSizeList  []float32
	coinSizeList := make([]float32, 10)

	return &coinSizeList
}
