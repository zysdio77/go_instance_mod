package router

import (
	"dotcount/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	job := router.Group("/Cash-hoard-slots-job/dc2")
	{
		job.POST("/userDetail",handler.UserDetail2)
		job.POST("/freeCoinDetail",handler.FreeCoinDetail)
	}


	return router
}