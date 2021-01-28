package handler

import (
	"dotcount/modle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func PayResult(c *gin.Context) {
	var payresultjson PayResultJson
	var paydetails PayDetails
	tablename := "pay_order"
	err := c.ShouldBind(&payresultjson)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"PayResultJson解析错误": err,
		}).Error(c.Request.Method, c.Request.URL)
	} else {
		paydetail := HandleUrlCode(payresultjson.PayDetail)
		payresultjson.PayResultToPayOrder(tablename, paydetail,c)
		err := json.Unmarshal([]byte(paydetail), &paydetails)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"PayDetail解析错误": err,
			}).Error(c.Request.Method, c.Request.URL)
		}
	}
}

func (payresultjson *PayResultJson) PayResultToPayOrder(tablename string, paydetail string , c *gin.Context) {
	var payresultdata modle.SelectPayResoltData
	payresultdata.TransactionId = payresultjson.Tid
	ok, readpayresult := payresultdata.DataExist(tablename)

	if ok {
		readpayresult.PayKey = payresultjson.PayKey
		readpayresult.Detail = paydetail
		rows := readpayresult.GormUpdateRow(tablename)
		if rows == 0 {
			logrus.WithFields(logrus.Fields{
				"订单号":          payresultjson.Tid,
				"详细描述":         readpayresult,
				"RowsAffected": rows,
			}).Warn("数据存储: 不成功：")
			c.JSON(http.StatusOK, gin.H{
				"result":    "error",
				"code": rows,
				"messege":   "payResult",
			})
		} else {
			logrus.WithFields(logrus.Fields{
				"订单号":  payresultjson.Tid,
				"详细描述": readpayresult,
				"RowsAffected": rows,
			}).Info("数据存储: 成功 , RowsAffected: ",rows)
			c.JSON(http.StatusOK, gin.H{
				"result":    "ok",
				"code": rows,
				"messege":   "payResult",
			})
		}

	} else {
		logrus.WithFields(logrus.Fields{
			"订单号":  payresultjson.Tid,
			"详细描述": payresultjson,
		}).Error("数据不存在: 无此订单号：")
		c.JSON(http.StatusOK, gin.H{
			"result":    "error",
			"code": 0,
			"messege":   "payResult",
		})
	}
}

func (payresultjson *PayResultJson) PayResultToUserData()  {

}