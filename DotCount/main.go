package main

import (
	"dotcount/modle"
	"dotcount/router"
)

func main()  {
	engine := router.NewRouter()
	dsn := "root:4rfvBGT%@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	modle.GormInitDb(dsn)

	engine.Run(":9090")
}