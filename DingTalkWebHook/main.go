//package DingTalkWebHook

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strings"
)

type RequestBody struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}
type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func main() {
	var requestbody RequestBody
	var result Result
	requestbody.Msgtype = "text"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxx724eeee07bxxxxxxxxxx389ee" //在钉钉机器人中找到此地址
	r := gin.Default()
	r.POST("/recieve", func(context *gin.Context) {
		body, err := ioutil.ReadAll(context.Request.Body)//接收post请求，拿到body中的json
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"key": "read request body err",
			}).Error(err)
		}

		status := gjson.Get(string(body), "status").Str //取json中的status的值
		ct := gjson.GetMany(string(body), "alerts.0.annotations.aliyun", //取多个值返回一个切片
			"alerts.0.annotations.grafana",
			"alerts.0.annotations.summary",
			"alerts.0.labels.desc",
			"alerts.0.annotations.value")
		c := []string{"阿里云地址：", "监控地址：", "总结：", "单位：", "状态(报警时的触发状态，恢复后的详细状态定请访问：http://10.0.0.66:9090/)："}
		var bb strings.Builder //字符串拼接
		bb.WriteString(status)
		bb.WriteString("\n")
		for i, j := range ct {
			bb.WriteString(c[i])
			bb.WriteString(j.Str)
			bb.WriteString("\n")
		}
		requestbody.Text.Content = bb.String() //拼接之后的字符串
		client := resty.New()
		//发送给钉钉机器人，详情参看钉钉机器人的接口文档
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			//SetBody(`{"msgtype":"text", "text":{"content":"` + bodystring + `"}}`).
			SetBody(&requestbody).
			SetResult(&result).
			Post(webhook)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error(err)
		}
		logrus.WithFields(logrus.Fields{}).Info(resp)
		//fmt.Println(resp)
	})
	r.Run(":9099")
}
