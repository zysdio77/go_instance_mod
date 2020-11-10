//package ReceiveLogs
package main

import (
	"example.com/m/v2/ConfigOption"
	"example.com/m/v2/WriteFile"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
)

/*接收请求的body全部存到文件中，建议日志接收服务*/
func main() {

	log.SetFormatter(&log.JSONFormatter{})    //设置log格式
	log.SetLevel(log.InfoLevel)               //设置log等级
	v := ConfigOption.InitConif()             //初始化配置文件
	c, err := ConfigOption.Unmarshalconfig(v) //把配置文件内容解析到结构提
	if err != nil {
		log.WithFields(log.Fields{
			"key": "ConfigOption.Unmarshalconfig",
		}).Error(err)
	}
	//f := "/Users/zhangyongsheng/data/test/111.txt"
	f := c.LogFile
	port := ":" + strconv.Itoa(c.Port)
	var ff *os.File
	if WriteFile.CheckFile(f) { //检测文件是否存在
		ff = WriteFile.AppendWrite(f) //存在就追加写入文件
	} else {
		ff = WriteFile.GAppendWrite(f) //不存在就创建追加写入文件
	}
	defer ff.Close() //关闭文件
	//gin.Default()
	r := gin.Default()
	r.POST("/log/receive", func(c *gin.Context) {
		//c.JSON(200,Ss{"asdfjasldfkj"})
		job := c.Query("jobname")
		//job := c.Param("name")
		//fmt.Println(name,job,age)
		body, err := ioutil.ReadAll(c.Request.Body) //把请求的Request.Body写入body，body为[]byte格式
		if err != nil {
			log.WithFields(log.Fields{
				"key": "ioutil.ReadAll",
			}).Error(err)
		}
		b := job + ":" + string(body) + "\n"
		//_, err = ff.Write(body) //把body内容写入到文件
		resolt, err := ff.WriteString(b) //把字符串写入文件
		//fmt.Println(r)
		if err != nil {
			c.String(2001, "Error:%V\n", err)
			log.WithFields(log.Fields{
				"key": "ff.Write(body)",
			}).Error(err)
		} else {
			c.String(200, "%v ok\n", resolt)
			log.WithFields(log.Fields{
				"key": "main.c.String",
			}).Info("Write log file succussful\n")
		}

	})

	r.Run(port)
}
