package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

//追加的方式打开文件
func checkfile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return f
}
func WriteFileAppend(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(data)
	return nil
}
func WriteFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

/*接收请求的body全部存到文件中，建议日志接收服务*/
func main() {
	f := "/Users/zhangyongsheng/data/test/111.txt"
	//f := "/home/ubuntu/111.txt"
	ff := checkfile(f)
	defer ff.Close() //关闭文件
	r := gin.Default()
	r.POST("/log/post", func(c *gin.Context) {
		//c.JSON(200,Ss{"asdfjasldfkj"})
		body, err := ioutil.ReadAll(c.Request.Body) //把请求的Request.Body写入body，body为[]byte格式
		if err != nil {
			fmt.Println(err)
		}
		r, err := ff.Write(body) //把body内容写入到文件
		if err != nil {
			c.String(2001, "Error:%V\n", err)
		} else {
			c.String(200, "%V ok\n", r)
		}
		fmt.Println(string(body))

	})
	r.Run(":9999")
}
