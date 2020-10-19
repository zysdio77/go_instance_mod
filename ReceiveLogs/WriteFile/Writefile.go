package WriteFile

//package main
import (
	"github.com/sirupsen/logrus"
	"os"
)

//检查文件是否存在存在返回ture，不存在返回false
func CheckFile(filenmae string) bool {
	_, err := os.Stat(filenmae)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//追加的方式打开文件，如果文件不存在则创建
func GAppendWrite(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": "Writefile.GAppendWrite",
		}).Error(err)
	}
	return f
}

//追加的方式打开文件
func AppendWrite(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": "Writefile.AppendWrite",
		}).Error(err)
	}
	return f
}

/*
func main()  {
	f := "/Users/zhangyongsheng/Desktop/appconfiga"
	if CheckFile(f) {
		ff := AppendWrite(f)
		ff.WriteString("dsafdsfasdfasdf")
	}
	ff :=GAppendWrite(f)
	ff.WriteString("a ")
}
*/
