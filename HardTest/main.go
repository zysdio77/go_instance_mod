//package HardTest
package main

import (
"fmt"
"github.com/go-resty/resty/v2"
"os"
"strconv"
"sync"
"time"
)

func G(s string) {
	client := resty.New() //建立一次连接
	for j := 1; j < 20; j++ {
		_, err := client.R().Get(s)
		if err != nil {
			fmt.Println(err)
			//return
		}
		time.Sleep(time.Second) //延时一秒，每次访问间隔一秒
		//ti := resq.Request.TraceInfo() //返回的状态信息
		//fmt.Println(resq.Status())
	}
	w.Done()
}
var w sync.WaitGroup
func main() {
	l := os.Args
	if len(l) != 3 {
		fmt.Println("Usage: You need 2 Args ,First is thread count ,Second is visit addrass")
		return
	}
	ll := l[1] //命令行参数第一个输入并发次数
	lll, err := strconv.Atoi(ll)
	if err != nil {
		fmt.Println("请输入数字")
	}
	s := l[2] //命令行参数第二个，输入地址

	for i := 0; i < lll; i++ {
		w.Add(1)
		go G(s)
	}
	w.Wait()
}
