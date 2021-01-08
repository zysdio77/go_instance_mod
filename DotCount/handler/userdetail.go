package handler

import (
	"dotcount/modle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserDetail(c *gin.Context) {
	//处理UserDetail
	//var db = modle.GDb
	var userdetail UserDetailJson
	//var userdetaildb modle.UserDataDb
	tablename := "user_data"

	err := c.ShouldBind(&userdetail)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"result":  "",
			"message": "",
			"status":  "faild",
		})
	} else {
		userdata := userdetail.NewDetailToData()
		if !userdata.DataExist(tablename) {
			userdata.GormInsertRow(tablename)
			c.JSON(http.StatusOK, gin.H{
				"result":  "",
				"message": "instert",
				"status":  "success",
			})
		} else {
			readuserdata := userdata.GormSelectRow(tablename)
			userdetail.UpdateDetailToData(readuserdata, userdata)
			userdata.SaveData(tablename)
			c.JSON(http.StatusOK, gin.H{
				"result":  "",
				"message": "update",
				"status":  "success",
			})
		}
	}
	fmt.Printf("userdetail: %v \n", userdetail)

}
func (userdetail *UserDetailJson) HandleNewVersion() string {
	var extend string
	if userdetail.Version == "" {
		extend = fmt.Sprintf("%v|%v", 0, 0)
	} else {
		extend = fmt.Sprintf("%v|%v", userdetail.Version, 0)
	}
	return extend
}
func (userdetail *UserDetailJson) HandleVersion(userdata *modle.ReadUserData) string {
	var extend string
	versionlist := strings.Split(userdata.Extend, "|")
	fmt.Printf("versionlist:%v\n", versionlist)
	versionlist1 := strings.Split(versionlist[1], ",")
	fmt.Printf("versionlist1:%v\n", versionlist1)
	if len(versionlist1) >= 100 {
		versionlist1 = versionlist1[1:]
	}
	versionlist1 = append(versionlist1, userdetail.Version)
	//a := versionlist[1] + "," + versionlist[0]
	a := strings.Join(versionlist1, ",")
	extend = fmt.Sprintf("%v|%v", userdetail.Version, a)
	fmt.Println(extend)

	return extend
}

func (userdetail *UserDetailJson) NewDetailToData() *modle.WriteUserData {
	var userdata modle.WriteUserData
	userdata.Id = userdetail.UID - 10000
	userdata.Channel = userdetail.Channel
	userdata.Platform = userdetail.Platform
	userdata.DevModel = userdetail.DevModel
	userdata.Lv = userdetail.Lv
	userdata.CoinNum = SplitYi(userdetail.CoinNum)
	userdata.FreeCoin = SplitYi(userdetail.FreeCoin)
	userdata.TotalBet = SplitYi(userdetail.TotalBet)
	userdata.TotalWin = SplitYi(userdetail.TotalWin)
	userdata.Vip = userdetail.Vip
	userdata.UserGambling = userdetail.UserGambling
	userdata.WildStamp = userdetail.WildStamp
	userdata.UserValue = userdetail.UserValue
	userdata.LogoutTime = TimeStamp2Date(userdetail.Time)
	userdata.GuideF = userdetail.GuideF
	userdata.GuideNf = userdetail.GuideNF
	userdata.Extend = userdetail.HandleNewVersion()
	fmt.Printf("userdata: %v \n", userdata)
	return &userdata
}

func (userdetail *UserDetailJson) UpdateDetailToData(readuserdata *modle.ReadUserData, userdata *modle.WriteUserData) {
	//var userdata modle.WriteUserData

	if userdetail.UID != 0 {
		userdata.Id = userdetail.UID - 10000
	}
	if userdetail.Channel != "" {
		userdata.Channel = userdetail.Channel
	} else {
		userdata.Channel = readuserdata.Channel
	}
	if userdetail.Platform != "" {
		userdata.Platform = userdetail.Platform
	} else {
		userdata.Platform = readuserdata.Platform
	}
	if userdetail.DevModel != "" {
		userdata.DevModel = userdetail.DevModel
	} else {
		userdata.DevModel = readuserdata.DevModel
	}
	if userdetail.Lv != 0 {
		userdata.Lv = userdetail.Lv
	} else {
		userdata.Lv = readuserdata.Lv
	}
	if userdetail.CoinNum != "" {
		userdata.CoinNum = SplitYi(userdetail.CoinNum)
	} else {
		userdata.CoinNum = readuserdata.CoinNum
	}
	if userdetail.FreeCoin != "" {
		freecoin := Updata(userdetail.FreeCoin, readuserdata.FreeCoin)
		userdata.FreeCoin = SplitYi(freecoin)
	} else {
		userdata.FreeCoin = readuserdata.FreeCoin
	}
	if userdetail.TotalBet != "" {
		totalbet := Updata(userdetail.TotalBet, readuserdata.TotalBet)
		userdata.TotalBet = SplitYi(totalbet)
	} else {
		userdata.TotalBet = readuserdata.TotalBet
	}
	if userdetail.TotalWin != "" {
		totalwin := Updata(userdetail.TotalWin, readuserdata.TotalWin)
		userdata.TotalWin = SplitYi(totalwin)
	} else {
		userdata.TotalWin = readuserdata.TotalWin
	}
	if userdetail.Vip != 0 {
		userdata.Vip = userdetail.Vip
	} else {
		userdata.Vip = readuserdata.Vip
	}
	if userdetail.UserGambling != "" {
		userdata.UserGambling = userdetail.UserGambling
	} else {
		userdata.UserGambling = readuserdata.UserGambling
	}
	if userdetail.WildStamp != "" {
		userdata.WildStamp = userdetail.WildStamp
	} else {
		userdata.WildStamp = readuserdata.WildStamp
	}
	if userdetail.UserValue != "" {
		userdata.UserValue = userdetail.UserValue
	} else {
		userdata.UserValue = readuserdata.UserValue
	}
	if userdetail.Time != 0 {
		userdata.LogoutTime = TimeStamp2Date(userdetail.Time)
	} else {
		userdata.LogoutTime = readuserdata.LogoutTime
	}
	if userdetail.GuideF != "" {
		userdata.GuideF = userdetail.GuideF
	} else {
		userdata.GuideF = readuserdata.GuideF
	}
	if userdetail.GuideNF != "" {
		userdata.GuideNf = userdetail.GuideNF
	} else {
		userdata.GuideNf = readuserdata.GuideNf
	}
	if userdetail.Version != "" {
		userdata.Extend = userdetail.HandleVersion(readuserdata)
	} else {
		userdata.Extend = readuserdata.Extend
	}
	fmt.Printf("userdata: %v \n", userdata)
}
