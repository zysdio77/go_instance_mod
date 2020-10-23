package main

import (
	"example.com/m/v2/ctldb"
	"example.com/m/v2/gincode"
	"github.com/gin-gonic/gin"
)

func main() {
	ctldb.InitData()
	r := gin.Default()
	r.GET("/userinfo", gincode.GetInfo)
	r.POST("/userinfo", gincode.SetInfo)
	r.Run(":9090")


}