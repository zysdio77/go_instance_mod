package GormInstance

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

func InitDB() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8&parseTime=True&loc=Local"
	db , err := gorm.Open("msyql",dsn)
	if err != nil {
		panic(err)
	}
	sqlDB,err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return sqlDB
}

