package main

import (
	"dotcount/modle"
	"dotcount/router"
)

func main()  {
	engine := router.NewRouter()
	dsn := "root:4rfvBGT%@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	modle.GormInitDb(dsn)
	defer modle.GDb.Close()
	//handler.SplitYi("11234567812345678123456781234567812345678")
	//handler.StringToBigInt("11234567812345678123456781234567812345678")
	engine.Run(":9090")
}