package MysqlInstance

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB
var err error

var uinfo DBUserInfo

type DBUserInfo struct {
	Id     int    `json:"id"  form:"id" db:"id"`
	Name   string `json:"name"  form:"name" db:"name"`
	Gender string `json:"gender" form:"gender" db:"gender"`
	Hobby  string `json:"hobby" form:"hobby" db:"hobby"`
}

func InitDB() (db *sqlx.DB) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		panic(err)
		//return nil,err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db
}

// 查询多条数据示例
func QueryMultiRowDemo() {
	sqlStr := "select * from user_info2"
	var us []DBUserInfo
	//查询多行数据用select
	err := db.Select(&us, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", us)
	for _, j := range us {
		fmt.Printf("id:%v,name:%v,gender:%v,hobby:%v\n", j.Id, j.Name, j.Gender, j.Hobby)

	}

}

// 查询单条数据示例
func QueryRowDemo(tablename string, uid int) {
	sqlStr := "select * from ? where id=?"
	var u DBUserInfo
	//查询一行数据用get
	err := db.Get(&u, sqlStr, tablename, uid)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(u.Id, u.Name, u.Gender, u.Hobby)
	//fmt.Printf("id:%d name:%s age:%d\n", , u.Name, u.Age)
}

//插入数据
func InsertRow(tablename string, ui *DBUserInfo) {
	sqlStr := "insert into ? (id,name,gender,hobby) values (?,?,?,?)"
	result, err := db.Exec(sqlStr, tablename, ui.Id, ui.Name, ui.Gender, ui.Hobby)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error()
		return
	}
	insertId, err := result.LastInsertId() //自增ID的id号
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("insert data success,id:%v\n", insertId)
}

//更改数据
func UpdateRow(tablename string, ui DBUserInfo) {
	sqlStr := "update ? set name = ? where id = ?"
	result, err := db.Exec(sqlStr, tablename, ui.Name, ui.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	affecteRows, err := result.RowsAffected() //影响的行数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("update data success,affected rows:%d\n", affecteRows)
}

//删除一行
func DeleteRow() {
	sqlStr := "delete from user_info2 where id =?"
	result, err := db.Exec(sqlStr, 14)
	if err != nil {
		fmt.Println(err)
		return
	}
	affectdRows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("delete data success,affected rows:%d\n", affectdRows)
}

/*
func main() {
	err = initDB()
	if err != nil {
		fmt.Println(err)
	}
	queryRowDemo()
	queryMultiRowDemo()
	insertRow()
	updateRow()
	deleteRow()
}*/
