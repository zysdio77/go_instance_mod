package modle

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var GDb *gorm.DB
var err error
func GormInitDb(dsn string)  {
	//dsn := "root:4rfvBGT%@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	//db,err :=gorm.Open("mysql",dsn)
	GDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//Db.AutoMigrate(&variable.UserDetal{})
	GDb.DB().SetMaxOpenConns(10)   //设置最大连接数
	GDb.DB().SetMaxIdleConns(5)    //设置闲置连接数
	//GDb.SingularTable(true)
	//Db.AutoMigrate(&variable.UserData{})
}

