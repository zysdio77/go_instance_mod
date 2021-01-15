package router

import (
	"dotcount/handler"
	"fmt"
	//"dotcount/handler"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func NewRouter() *gin.Engine {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	job := router.Group("/Cash-hoard-slots-job/dc2")
	{
		job.POST("/userDetail",handler.UserDetail)
		job.POST("/freeCoinDetail",handler.FreeCoinDetail)

	}



	//test
	router.POST("/Cash-hoard-slots-job/dc2/slotDetail", func(c *gin.Context) {
		r := c.Request
		body,_ := ioutil.ReadAll(r.Body)
		//log.Logger.Printf("Method:%v,URL:%v,PostForm:%v,Body:%v,Header:%v\n",r.Method,r.URL,r.PostForm,string(body),r.Header)
		//logrus.Info("Method: %v, URL: %v, Body: %v, Header: %v\n",r.Method,r.URL,string(body),r.Header)
		//log.Printf("Method:%v,URL:%v,PostForm:%v,Body:%v,Header:%v\n",r.Method,r.URL,r.PostForm,string(body),r.Header)
		fmt.Printf("Method: %v, URL: %v, Body: %v, Header: %v\n",r.Method,r.URL,string(body),r.Header)
	})
	router.POST("/Cash-hoard-slots-job/dc2/payResult", func(c *gin.Context) {
		r := c.Request
		body,_ := ioutil.ReadAll(r.Body)
		//log.Logger.Printf("Method:%v,URL:%v,PostForm:%v,Body:%v,Header:%v\n",r.Method,r.URL,r.PostForm,string(body),r.Header)
		//logrus.Info("Method: %v, URL: %v, Body: %v, Header: %v\n",r.Method,r.URL,string(body),r.Header)
		//log.Printf("Method:%v,URL:%v,PostForm:%v,Body:%v,Header:%v\n",r.Method,r.URL,r.PostForm,string(body),r.Header)
		fmt.Printf("Method: %v, URL: %v, Body: %v, Header: %v\n",r.Method,r.URL,string(body),r.Header)
	})
	return router
}