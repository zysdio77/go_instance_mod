package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name string
	Age int

}


func (u User)InitDB()*gorm.DB  {
	dsn := "root:4rfvBGT%@tcp(10.0.0.66:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db , err := gorm.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	sqlDB := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println(u.Name)
	//db.CreateTable()
	return db }

func main ()  {
	//db := InitDB()
	var user User
	db := user.InitDB()
	defer db.Close()
	//db.CreateTable(&user)


}
