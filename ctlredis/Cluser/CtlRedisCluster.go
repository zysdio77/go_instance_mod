package Cluser

import (
	"fmt"
	"github.com/gitstliu/go-redis-cluster"
	//"gopkg.in/redis.v4"
	"time"
)

func Cluster()  {
	c,err := redis.NewCluster( //连接rediscluster集群
		&redis.Options{
			StartNodes: []string{"10.0.0.66:6371","10.0.0.66:6372","10.0.0.66:6373","10.0.0.66:6374","10.0.0.66:6375","10.0.0.66:6376"},
			ConnTimeout: 2 * time.Second,
			ReadTimeout: 2 * time.Second,
			WriteTimeout: 2 * time.Second,
			KeepAlive: 16,
			AliveTime: 60 * time.Second,
		})
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	c.Do("SET","name","zhangyognsheng") //集群的操作方法
	r,_ := c.Do("GET","name")
	//fmt.Println(r,err)
	fmt.Printf("%T,%v",r,r)

}