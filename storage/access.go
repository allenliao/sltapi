package storage

import (
	"bytes"
	"database/sql"
	"log"
	"sltapi/models"
	"strconv"

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
	}

	return bUInfo

}

func DB_GetGameInfo(bucode string, gameSN uint8) *models.GameInfo {
	if models.GameInfoList[gameSN] == nil {
		//取回來快取
		models.GameInfoList[gameSN] = GetGameInfo(bucode, gameSN)
	}
	return models.GameInfoList[gameSN]
}

func GetGameInfo(bucode string, gamesn uint8) *models.GameInfo {
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
	}

	gameInfo.CoinSizeList = DB_GetCoinSizeListByBUGame(bucode, gamesn)

	log.Printf("gamesn:%v,  MinMultiplier:%v, MaxMultiplier:%v, BaseCredit:%v, EngineSN:%v,",
		gamesn,
		gameInfo.MinMultiplier,
		gameInfo.MaxMultiplier,
		gameInfo.BaseCredit,
		gameInfo.EngineSN)

	return gameInfo
}

var BUGameCoinSizeList map[string]*[]float32 = make(map[string]*[]float32) //map[BU001_1]*[]float32

func DB_GetCoinSizeListByBUGame(bucode string, gamesn uint8) *[]float32 {

	var bucode_gamesn bytes.Buffer
	bucode_gamesn.WriteString(bucode)
	bucode_gamesn.WriteString("_")
	bucode_gamesn.WriteString(strconv.Itoa(int(gamesn))) //接字串

	bucode_gamesnStr := bucode_gamesn.String()
	log.Println("bucode_gamesnStr:", bucode_gamesnStr)

	if BUGameCoinSizeList[bucode_gamesnStr] == nil {
		//取回來快取
		BUGameCoinSizeList[bucode_gamesnStr] = GetCoinSizeListByBUGame(bucode, gamesn)
	}
	return BUGameCoinSizeList[bucode_gamesnStr]
}

func GetCoinSizeListByBUGame(bucode string, gamesn uint8) *[]float32 {

	dbQueryStr := `
	SELECT coinsize
	FROM coinsize as c inner join coinsizegroup as cg on c.coinsize_groupid=cg.coinsize_groupid 
    inner join bucoinsize as bc on bc.coinsize_groupid=cg.coinsize_groupid 
    WHERE bc.bucode=? and bc.gamesn=?
	`

	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(bucode, gamesn)
	goutils.CheckErr(err)
	defer rows.Close()

	colNames, err := rows.Columns()
	log.Println("GetCoinSizeListByBUGame rows.colNames:", colNames)
	//CoinSizeList  []float32
	coinSizeList := make([]float32, 10)
	readCols := make([]interface{}, len(colNames))
	writeCols := make([]float32, len(colNames))
	for i, _ := range writeCols {
		readCols[i] = &writeCols[i]
	}
	for rows.Next() { //有下一筆就會一直true下去
		err = rows.Scan(readCols...)
		goutils.CheckErr(err)

		idx := 0
		for idy := range writeCols {
			coinSizeList[idx] = writeCols[idy]
			log.Println("coinsize:", coinSizeList[idx])
			idx++
		}

	}

	return &coinSizeList
}
