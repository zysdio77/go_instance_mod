package modle

import (
	"fmt"
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
		fmt.Println(err)
		panic(err)
	}
	//Db.AutoMigrate(&variable.UserDetal{})
	GDb.DB().SetMaxOpenConns(10)   //设置最大连接数
	GDb.DB().SetMaxIdleConns(5)    //设置闲置连接数
	//GDb.SingularTable(true)
	//Db.AutoMigrate(&variable.UserData{})
}

func (writeuserdata *WriteUserData)SaveData(tablename string)  {
	readrow := writeuserdata.GormSelectRow(tablename)
	//fmt.Printf("readrow: %v \n",readrow)
	if GDb.NewRecord(readrow) {
		writeuserdata.GormInsertRow(tablename)
		//fmt.Printf("readrow nil : %v \n",readrow)
	} else {
		//fmt.Printf( " Error : writeuserdata  : %v  already exic \n",writeuserdata)
		writeuserdata.GormUpdateRow(tablename)
	}
}

func (writeuserdata *WriteUserData) DataExist(tablename string) bool {
	//数据是否存在，存在true，不存在false
	readrow := writeuserdata.GormSelectRow(tablename)
	if GDb.Table(tablename).NewRecord(readrow) {
		return false
	} else {
		return true
	}
}
func (writeuserdata *WriteUserData)GormInsertRow(tablename string)  {
	//插入数据
		GDb.Table(tablename).Create(&writeuserdata)
}
func(writeuserdata *WriteUserData) GormSelectRow(tablename string) *ReadUserData {
	//查询数据
	var readrow  ReadUserData
	GDb.Table(tablename).Where("id = ?",&writeuserdata.Id).Find(&readrow)
	return &readrow
}

func (writeuserdata *WriteUserData) GormUpdateRow(tablename string)  {
	//更新数据
	GDb.Table(tablename).Save(&writeuserdata)
}

