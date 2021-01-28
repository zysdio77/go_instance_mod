package router

import (
	"dotcount/handler"
	//"dotcount/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	job := router.Group("/Cash-hoard-slots-job/dc2")
	{
		job.POST("/userDetail",handler.UserDetail)
		job.POST("/freeCoinDetail",handler.FreeCoinDetail)
		job.POST("/slotDetail",handler.SlotDetail)
		job.POST("/payResult",handler.PayResult)
	}

	return router
}