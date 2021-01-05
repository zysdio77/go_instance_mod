package handler

import (
	"dotcount/modle"
	"dotcount/variable"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"

	//"github.com/jinzhu/gorm"
	"net/http"
)

func UserDetail(c *gin.Context) {
	//处理UserDetail
	var db = modle.GDb
	var userdetail variable.UserData

	err := c.ShouldBind(&userdetail)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"result":  "",
			"message": "",
			"status":  "faild",
		})
	} else {
		if db.NewRecord(&userdetail) {
			db.Create(&userdetail)
		}

		c.JSON(http.StatusOK, gin.H{
			"result":  "",
			"message": "",
			"status":  "success",
		})
	}
}
func UserDetail2(c *gin.Context) {
	//处理UserDetail
	//var db = modle.GDb
	//var db = modle.XDb
	var userdatadb variable.UserData
	uid := c.PostForm("uid")
	lv := c.PostForm("lv")
	coinNum := c.PostForm("coinNum")
	freeCoin := c.PostForm("freeCoin")
	totalBet := c.PostForm("totalBet")
	totalWin := c.PostForm("totalWin")
	vip := c.PostForm("vip")
	userGambling := c.PostForm("userGambling")
	wildStamp := c.PostForm("wildStamp")
	userValu := c.PostForm("userValu")
	logoutTime := c.PostForm("logoutTime")
	guideF := c.PostForm("guideF")
	guideNF := c.PostForm("guideNF")
	platform := c.PostForm("platform")
	channel := c.PostForm("channel")
	bind_fb := c.PostForm("bind_fb")
	bind_apple := c.PostForm("bind_apple")
	pay_num := c.PostForm("pay_num")
	buy_coin := c.PostForm("buy_coin")
	user_gambling := c.PostForm("user_gambling")
	version := c.PostForm("version")
	id := String2Int(uid)
	id = id - 10000
	fmt.Printf("id:%v\n", id)
	fmt.Printf("uid:%v,%T,lv:%v,coinNum:%v,freeCoin:%v,totalBet:%v,totalWin:%v,vip:%v,userGambling:%v,wildStamp:%v,userValu:%v,logoutTime:%v,guideF:%v,guideNF:%v,platform:%v\n", uid, uid, lv, coinNum, freeCoin, totalBet, totalWin, vip, userGambling, wildStamp, userValu, logoutTime, guideF, guideNF, platform)

	ua := c.GetHeader("User-Agent")
	fmt.Println(ua)
	lvtemp := String2Int(lv)
	bindfb := String2Int(bind_fb)
	bindapple := String2Int(bind_apple)
	paynum := String2Int(pay_num)
	vip2 := String2Int(vip)
	logout := String2Int(logoutTime)
	datetime := TimeStamp2Date(int64(logout))

	var versionlist []string

	versionlist = append(versionlist, version)
	versionstr :=strings.Join(versionlist,",")
	userdatadb.Id = int64(id)
	userdatadb.Channel = channel
	userdatadb.Platform = platform
	userdatadb.DevModel = ua
	userdatadb.BindFb = bindfb
	userdatadb.BindApple = bindapple
	userdatadb.Lv = lvtemp
	userdatadb.CoinNum = coinNum
	userdatadb.PayNum = int64(paynum)
	userdatadb.BuyCoin = buy_coin
	userdatadb.FreeCoin = freeCoin
	userdatadb.TotalBet = totalBet
	userdatadb.TotalWin = totalWin
	userdatadb.Vip = vip2
	userdatadb.UserGambling = user_gambling
	userdatadb.WildStamp = wildStamp
	userdatadb.UserValue = userValu
	userdatadb.LogoutTime = datetime
	userdatadb.GuideF = guideF
	userdatadb.GuideNf = guideNF
	userdatadb.InboxDetail = ""
	userdatadb.Extend = versionstr


	modle.GormInsertRow(&userdatadb)


}
func String2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
func TimeStamp2Date(timestamp int64) string {
	datetime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return datetime
}
func FreeCoinDetail(c *gin.Context) {
	//处理FreeCoinDetail
}
