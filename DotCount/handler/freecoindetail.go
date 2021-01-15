package handler

import (
	"dotcount/modle"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func FreeCoinDetail(c *gin.Context) {
	var freecoindetail FreeCoinDataJson
	//接收传来的json
	//var readdata *modle.ReadFreeData
	var data modle.WriteFreeData
	tablename := "dc_free_coin_detail"
	err := c.ShouldBind(&freecoindetail)
	if err != nil {
		fmt.Println()
		logrus.WithFields(logrus.Fields{
			"freecoindetail 解析错误: ": err,
			//"修改的数据 userdata":   userdata,
		}).Error(c.Request.Method, c.Request.URL)
		c.JSON(500, gin.H{
			"result":  "",
			"message": "freecoindetail 解析错误",
			"status":  "faild",
		})
	} else {
		var freecoindetails  FreeCoinDetailJson
		freecoindetailstr := HandleUrlCode(freecoindetail.FreeCoinDetails)
		//url解码
		if freecoindetailstr == "" {
			return
		}
		err := json.Unmarshal([]byte(freecoindetailstr),&freecoindetails) //解码之后的数据存入freecoindetails
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"freecoindetailstr 解析错误: ": err,
			}).Error(c.Request.Method, c.Request.URL)
		}
		logrus.WithFields(logrus.Fields{
			"freecoindetail : ": freecoindetails,
		}).Info(c.Request.Method, c.Request.URL)

		freemap := HandleuFreeCoinDetails(freecoindetails) //freecoindetail字段的处理存入map。
		freedetail :=Dbstr(freemap)//需要存储的数据

		for k,v := range freedetail{
			d ,err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
			}
			NewFreeJsonToData(&freecoindetail,&data,k,string(d))
			//吧需要入库的数据映射到对应的结构提
			ok,readdata := data.DataExist(tablename)
			if ok {
				//修改
				UpdateFreeJsonToData(&freecoindetail,&data,readdata)
				readdata.GormUpdateRow(tablename,readdata)
				c.JSON(http.StatusOK, gin.H{
					"result":  "修改",
					"message": "freecoindetail",
					"status":  "succuss",
				})
			} else {
				//添加
				data.GormInsertRow(tablename)
				c.JSON(http.StatusOK, gin.H{
					"result":  "添加",
					"message": "freecoindetail",
					"status":  "succuss",
				})
			}
		}
	}
}

func HandleuFreeCoinDetails(freedetail FreeCoinDetailJson)  map[string]map[string]int{
	//freecoindetail字段的处理
	var freemap map[string]map[string]int
	freemap = make(map[string]map[string]int)
	for _,j := range freedetail{
		//初始化key:key:value的map
		datestr := Datestr(j.TimeStamp)
		freemap[datestr] = map[string]int{}
	}
	for _,j := range freedetail{
		datestr := Datestr(j.TimeStamp)
		freemap[datestr][j.CoinName]= freemap[datestr][j.CoinName] +j.CoinNum
	}
 return freemap
}

func Dbstr(freemap map[string]map[string]int) map[string][]FreeCoinDetailJson2 {
	freedetail := make(map[string][]FreeCoinDetailJson2)
	var f2 FreeCoinDetailJson2
	for i,j := range freemap{
		for k, v := range j {
			f2.CoinName = k
			f2.CoinNum = v
			f2.TimeStamp =  DateToTimeStamp(i)
			freedetail[i] = append(freedetail[i], f2)
		}
	}
	return freedetail
}

func Datestr(timestamp int64) string {
	t := TimeStampToDate(timestamp)
	tlist  := strings.Fields(t)
	return tlist[0]
}
func NewFreeJsonToData(freecoinjson *FreeCoinDataJson,data *modle.WriteFreeData,datestr string,freedetail string)  {
	data.Uid = freecoinjson.Uid
	data.Platform =freecoinjson.Platform
	data.Channel = freecoinjson.Channel
	data.Version = HandleNewVersion(freecoinjson.Version)
	data.DateStr = datestr
	data.FreeCoinDetail = freedetail
}
func UpdateFreeJsonToData(freecoinjson *FreeCoinDataJson,writedata *modle.WriteFreeData,readdata *modle.ReadFreeData)  {
	readdata.Version = HandleVersion(freecoinjson.Version,readdata.Version)
	readdata.FreeCoinDetail = modle.FreeCoinDetails(readdata.FreeCoinDetail,writedata.FreeCoinDetail)
}
