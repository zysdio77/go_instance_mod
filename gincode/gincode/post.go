package gincode

import (
	"example.com/m/v2/ctldb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostUserInfo struct {
	Name   string `json:"name" form:"name"`
	Gender string `json:"gender" form:"gender"`
	Hobby  string `json:"hobby" form:"hobby"`
}

var pui PostUserInfo
var err error

func SetInfo(context *gin.Context) {
	err = context.ShouldBind(&pui)
	if err != nil {
		logrus.Error(err)
	}
	err = ctldb.Insert(pui.Name, pui.Gender, pui.Hobby)
	if err != nil {
		logrus.Error(err)
	}
}
