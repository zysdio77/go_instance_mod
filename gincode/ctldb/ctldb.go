package ctldb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type UserInfo struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}

var ui UserInfo
var db *gorm.DB
var err error

func InitData() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(172.20.0.1:3306)/test1?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		//logrus.WithFields(logrus.Fields{}).Info(err)
		logrus.Error(err)
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)          //最大连接
	db.DB().SetMaxIdleConns(50)           //最大空闲连接
	db.DB().SetConnMaxLifetime(time.Hour) //最大存活时间

	if db.HasTable(&UserInfo{}) { // 检查模型`UserInfo`表是否存在
		logrus.Info("tales exiesed")
	} else {
		db.CreateTable(&UserInfo{}) // 为模型`UserInfo`创建表
	}
}
func Insert(name, gender, hobby string) error {
	ui = UserInfo{Name: name, Gender: gender, Hobby: hobby}
	err := db.Create(&ui).Error //把ui中的数据查到数据库中
	if err != nil {
		//logrus.Error(err)
		return err
	}
	return nil

}

func Select(id uint) (name, gender, hobby, times string) { //根据ID查询姓名，性别，爱好，创建时间 ，返回json

	//us := new(UserInfo)
	//var us *UserInfo
	//db.First(&ui)
	db.Where("id = ?", id).Find(&ui)  //查询数据库
	var t int64 = ui.CreatedAt.Unix() //把时间格式数据转换为时间戳
	//fmt.Println(t)
	tt := time.Unix(t, 0).Format("2006-01-02 15:04:05 PM") //把时间戳转换为字符串
	//fmt.Println(tt)

	return ui.Name, ui.Gender, ui.Hobby, tt
	//fmt.Printf("%v,%T\n",ui,ui)
}
