package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

var db *sqlx.DB
var err error

type DBUserInfo struct {
	//Id     int    `json:"id"  form:"id" db:"id"`
	UserId    string `json:"name"  form:"name" db:"user_id"`
	PlatForm  string `json:"gender" form:"gender" db:"platform"`
	CreatTime string `json:"hobby" form:"hobby" db:"create_time"`
}

func InitDB() (db *sqlx.DB) {

	dsn := "pitch:Pitch1009@tcp(rm-0xi3tbytj2cl87fsq.mysql.rds.aliyuncs.com:3306)/cash-hoard-slots?charset=utf8"
	// 也可以使用MustConnect连接不成功就panic
	//dsn := "root:4rfvBGT%@tcp(127.0.0.1:3306)/test?charset=utf8"

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
	sqlStr := "select * from dau_log limit 100"
	var us []DBUserInfo
	//查询多行数据用select
	err := db.Select(&us, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", us)
	for _, j := range us {
		fmt.Printf("name:%v,gender:%v,hobby:%v\n", j.UserId, j.PlatForm, j.CreatTime)

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
	fmt.Println(u.UserId, &u.PlatForm, u.CreatTime)
	//fmt.Printf("id:%d name:%s age:%d\n", , u.Name, u.Age)
}

//插入数据
func InsertRow(ui *DBUserInfo) {
	//sqlStr := "insert into dau_log (user_id,platform,create_time) values (?,?,?)"
	//result, err := db.Exec(sqlStr,ui.UserId,ui.PlatForm, ui.CreatTime)
	result, err := db.Exec("insert into dau_log (user_id,platform,create_time) values (?,?,?)", ui.UserId, ui.PlatForm,ui.CreatTime)
	if err != nil {
		//fmt.Println(err)
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
	result, err := db.Exec(sqlStr, tablename, ui.UserId)
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

func ReadLine(filename string) {

	var uinfo DBUserInfo
	f, err := os.Open(filename)
	if err != nil {
		println(err)
		panic(err)
	}
	defer f.Close()
	bfrd := bufio.NewReader(f)
	for {
		if err != nil {
			if err == io.EOF {
				println("read all ok")
				return
			} else {
				println(err)
			}
		}
		fmt.Printf("ui:%v\n", uinfo)
		line, err := bfrd.ReadBytes('\n')

		uinfo.UserId = strings.Split(string(line), ",")[1]
		uinfo.PlatForm = strings.Split(string(line), ",")[2]
		uinfo.CreatTime = strings.Split(strings.Split(string(line), ",")[3], ")")[0]
		//transtime(uinfo.CreatTime)
		uid := strings.Split(uinfo.UserId,"'")
		uinfo.UserId = uid[1]
		pl :=  strings.Split(uinfo.PlatForm,"'")
		uinfo.PlatForm = pl[1]
		tt := strings.Split(uinfo.CreatTime,"'")
		uinfo.CreatTime = tt[1]

		//fmt.Println(tt[1])
		//t,err  := time.Parse("2006-01-02 15:04:05",tt[1])
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("t:%v,%T\n",t.Unix(),t.Unix())
		//fmt.Println(uinfo)
		InsertRow(&uinfo)
	}

}

func transtime(timeString string)  {

	//timeString := "2019-07-29 00:00:00"
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(layout, timeString, loc)
	fmt.Println("trancetime : ```````",tmp,timeString)
}
func main() {
	//fmt.Println(time.Now().Unix())
	//return
	starttime := time.Now().UnixNano()
	defer func(){

		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
			endtime := time.Now().UnixNano()
			spendtime := endtime - starttime
			fmt.Printf("用时：%v\n",spendtime/1e9)
		}
	}()
	db = InitDB()
	defer db.Close()
	filename := "dau_log.sql"
	ReadLine(filename)


}

