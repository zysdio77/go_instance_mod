package modle

import (
	"dotcount/variable"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"strings"
)

var GDb *gorm.DB
var XDb *sqlx.DB
var err error
func GormInitDb(dsn string)  {
	//dsn := "root:4rfvBGT%@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	//db,err :=gorm.Open("mysql",dsn)
	GDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//Db.AutoMigrate(&variable.UserDetal{})
	GDb.DB().SetMaxOpenConns(50)   //设置最大连接数
	GDb.DB().SetMaxIdleConns(5)    //设置闲置连接数
	GDb.SingularTable(true)
	//Db.AutoMigrate(&variable.UserData{})
}

func SqlxInitDb (dsn string) {
	db,err := sqlx.Connect("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
}
func ConnectTable()  {

}
func GormInsertRow(userdata *variable.UserData)  {
	var versionlist []string
	//var userdata variable.UserData
	if GDb.NewRecord(&userdata){
		GDb.Create(&userdata)
	} else {
		var userdata2 variable.UserDataDb
		GDb.Table("user_data").Where(&userdata).Find(&userdata2)
		versionlist = append(versionlist, userdata2.Extend)
		versionstr :=strings.Join(versionlist,",")
		fmt.Println(userdata2,versionstr,versionlist)
		userdata.Extend = versionstr
		GDb.Save(&userdata)
	}
	//db.NewRecord()

}

func SqlxInsertRow()  {
	//sqlstr := "insert into user_data(id,channel,platform,dev_model,bind_fb,bind_apple,lv,coin_num,pay_num,buy_coin,free_coin,total_bet,total_win,vip,user_gambling,wild_stamp,user_value,logout_time,guide_f,guide_nf,inbox_detail,extend) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	//ret,err := XDb.Exec(sqlstr,)

}

func DevModel() {
}

func FreeCoin()  {

}

func TotalBet()  {

}

func TotalWin()  {

}
func Extend()  {
	
}
