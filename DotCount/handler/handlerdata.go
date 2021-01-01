package handler

import (
	"dotcount/variable"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserDetail(c *gin.Context)  {
//处理UserDetail
	var userdetail variable.UserDtail
	err := c.ShouldBind(&userdetail)
	if err != nil {
		fmt.Println(err)
		c.JSON(500,gin.H{
			"result": "",
			"message": "",
			"status": "faild",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"result": "",
			"message": "",
			"status": "success",
		})
	}


}

func FreeCoinDetail(c *gin.Context)  {
//处理FreeCoinDetail
}