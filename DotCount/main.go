package main

import (
	"dotcount/modle"
	"dotcount/router"
)


func main()  {
	engine := router.NewRouter()
	dsn := "root:123123@tcp(127.0.0.1:3306)/goadmin?charset=utf8&parseTime=True&loc=Local"
	modle.GormInitDb(dsn)
	defer modle.GDb.Close()
	//handler.SplitYi("11234567812345678123456781234567812345678")
	//handler.StringToBigInt("11234567812345678123456781234567812345678")
	//handler.NowTime()
	//fmt.Println(handler.Datestr(1610073016))
	//handler.DateToTimeStamp("2021-01-08")
	engine.Run(":9090")
}