package handler

import (
	"dotcount/modle"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SlotDetail(c *gin.Context)  {
	var slotdetailjson SlotDetailJson
	tablenmame := "dc_slot_detail"
	err := c.ShouldBind(&slotdetailjson)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"SlotDetail 解析json错误:":err,
		}).Error()
		c.JSON(500,gin.H{
			"result":  "",
			"message": "freecoindetail 解析错误",
			"status":  "faild",
		})
	} else {
		slist := HandleSlotDetails(slotdetailjson.SlotDetails)
		Circleslot(tablenmame,&slotdetailjson,slist)

		logrus.WithFields(logrus.Fields{
			"uid":slotdetailjson.UID,
			"channel":slotdetailjson.Channel,
			"platform":slotdetailjson.Platform,
			"version":slotdetailjson.Version,
			"slotDetails":slist,
		}).Info(c.Request.Method,c.Request.URL)

		c.JSON(http.StatusOK,gin.H{
			"result":  "ok",
			"message": "slotdetail",
			"status":  "succuss",
		})
	}

}

func HandleSlotDetails(slotdetailjson string) []SlotDetails  {
	var slotdetails  []SlotDetails
	slotdetailstr :=HandleUrlCode(slotdetailjson)
	json.Unmarshal([]byte(slotdetailstr),&slotdetails)
	return slotdetails
}
func Circleslot(tablename string,slotdetailjson *SlotDetailJson,slist []SlotDetails)  {
	for i,j := range slist{
		fmt.Println(i,j)
		writedata := NewSlotDtail(slotdetailjson,j)
		ok,readrow :=writedata.DataExist(tablename)
		if ok {
			//修改
			UpdataSlotDtail(slotdetailjson,j,readrow)
			readrow.GormUpdateRow(tablename)
		} else {
			//添加
			writedata.GormInsertRow(tablename)
		}
		fmt.Printf("ok:%v,readrow:%v\n",ok,readrow)

	}
}

func UniqSlotid(slotdetail *SlotDetails,) {
	//slotdetail.SlotID
}
func NewSlotDtail(slotdetailjson *SlotDetailJson,slotdetails  SlotDetails) *modle.WriteSlotDetailData {
	var writedata modle.WriteSlotDetailData
	writedata.Uid = slotdetailjson.UID
	writedata.Channel = slotdetailjson.Channel
	writedata.Platform = slotdetailjson.Platform
	writedata.SlotId  = slotdetails.SlotID
	writedata.BetNum = SplitYi(IntToString(slotdetails.BetNum))
	writedata.WinNum = SplitYi(IntToString(slotdetails.WinNum))
	writedata.DateStr=TimeStampToDate(slotdetailjson.Time)
	writedata.Extend = HandleNewVersion(slotdetailjson.Version)
	//fmt.Printf("Jsontodata:%v\n",writedata)
	return &writedata
}

func UpdataSlotDtail(slotdetailjson *SlotDetailJson,slotdetails  SlotDetails,readdata *modle.ReadSlotDetailData) {
	//var writedata modle.WriteSlotDetailData
	readdata.BetNum = SplitYi(Updata(IntToString(slotdetails.BetNum),readdata.BetNum))
	readdata.WinNum = SplitYi(Updata(IntToString(slotdetails.WinNum),readdata.WinNum))
	readdata.Extend = HandleVersion(slotdetailjson.Version,readdata.Extend)
	//fmt.Printf("Jsontodata:%v\n",writedata)
}


