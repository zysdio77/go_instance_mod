package gincode

import (
	"example.com/m/v2/ctldb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Data struct {
	PostUserInfo
	Time string
}

func GetInfo(context *gin.Context) {
	id := context.Query("id")	//获取参数
	i, err := strconv.Atoi(id)	//字符串转换为int
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Querystring": err,
		}).Error("Query err don,t strconv.AtoI")
	}
	var d Data
	d.Name, d.Gender, d.Hobby, d.Time = ctldb.Select(uint(i))	//查询数据库获取想要的结果
	context.JSON(http.StatusOK, &d)

}
