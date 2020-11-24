package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"xuanxuanshuxue/jiajianfa"
)

type Conf struct {
	Duoshaodaoti int
}
func main() {
	viper.SetConfigFile("./ti.yaml")
	var conf Conf
	err := viper.ReadInConfig()	//读取配置文件
	if err != nil {
		fmt.Println(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(conf)
	alist := []int{}
	count := 0
	var randnum int
	//////////初始化，随机插入第一个数//////
	ti := jiajianfa.JiaAndJianfa()
	rand.Seed(time.Now().UnixNano())
	randnum = rand.Intn(len(ti))
	alist = append(alist, randnum)
	//////////初始化，随机插入第一个数//////

	var ok bool
	for i := 0; i < len(ti); i++ {
		ok = false
		rand.Seed(time.Now().UnixNano())
		randnum = rand.Intn(len(ti))
		for _, j := range alist {
			//fmt.Println(alist,randnum, j,count)
			if randnum == j {	//随机的数子在alist里跳出
				//fmt.Println("break")
				break
			}else {
				//fmt.Println("ok")
				ok = true
			}
		}
		if ok {
			count++
			alist = append(alist, randnum)
			//fmt.Println(alist)
		}
		if count == conf.Duoshaodaoti {
			break
		}
	}
	var b strings.Builder
	var data  string
	for _, j := range alist {
		b.WriteString(ti[j])
		b.WriteString("\n")
		fmt.Println(ti[j])

		data =b.String()

	}

	ioutil.WriteFile("out.txt",[]byte(data),0644)
	//fmt.Println(alist)
}
