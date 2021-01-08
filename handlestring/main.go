package main

import (
	"fmt"
	"strings"
)

func joinstrings() {
	//切片的每个元素连接成字符串
	a := []string{"1", "2", "3", "a", "b", "asdf"}
	alist := strings.Join(a, ",")
	fmt.Printf("alist: %v,%T\n", alist, alist)
}

func splitstrings() {
	//按 ， 分割字符串，返回一个字符串切片
	a := "zh,asd,yohg,sheo,asdf,aniug"
	alist := strings.Split(a, ",")
	fmt.Printf("alist:%v,%T\n", alist, alist)
}

func addsting() {
	a := "11234567812345678123456781234567812345678"
	fmt.Println(a[0:1]) //1,从第0位开始到第1位结束不包括第1位

	//每间隔8位添加一个"｜"的字符串
	var bb string
	tag := 8
	c := len(a) / tag
	var t int
	fmt.Println(c)
	for i := 0; i < c; i++ {
		fmt.Println(bb)
		if i == 0 {
			bb = fmt.Sprintf(a[:len(a)-tag] + "|" + a[len(a)-tag:])
			t = tag
			continue
		}
		t = t + tag + 1
		bb = fmt.Sprintf(bb[:len(bb)-t] + "|" + bb[len(bb)-t:])
	}

	fmt.Println(bb)
}
func main() {
	joinstrings()
	splitstrings()
	addsting()
}
