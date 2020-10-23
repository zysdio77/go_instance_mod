package ctlredis

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/redis.v4"
)

func CliRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{Addr: "redis:6379", Password: "", DB: 3}) //redis:6379 里的redis就是docker-compose。yml里的redis
	return c
}

func GetData(c *redis.Client,string2 string) string { //查找
	r,err := c.Get(string2).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"SetData":err,
		}).Error()
	}
	return r
}
func SetData(c *redis.Client,k,v string)  { //插入
	err := c.Set(k, v, 0).Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"SetData":err,
		}).Error()
	}

}