package main

import (
	"fmt"
	"math"
	"math/big"
)

func count()  {
	maxint64 := math.MaxInt64
	a := int64(maxint64)
	b:= big.NewInt(a)	//int64转bigint
	b.Add(b,b)	//加法
	fmt.Printf("加法: a值: %v,a类型：%T,b值：%v,b类型：%T\n",a,a,b,b)

	b = big.NewInt(a)
	b.Sub(b,b)	//减法
	fmt.Printf("减法：a值：%v,a类型：%T,b值：%v,b类型：%T\n",a,a,b,b)

	b = big.NewInt(a)
	b.Mul(b,b)	//乘法
	fmt.Printf("乘法：a值：%v,a类型：%T,b值：%v,b类型：%T\n",a,a,b,b)

	b = big.NewInt(a)
	b.Div(b,b)	//除法
	fmt.Printf("乘法：a值：%v,a类型：%T,b值：%v,b类型：%T\n",a,a,b,b)

	big1 := new(big.Int).SetUint64(uint64(10000000))
	fmt.Println("big1 is: ", big1)

	big2, _ := new(big.Int).SetString("123123123123123123123123123123123123123123", 10) //字符串转bigint
	fmt.Printf("big2:%v %T\n", big2,big2)

	big2str :=big2.String()	//bigint转字符串
	fmt.Printf("big2str: %v %T\n",big2str,big2str)
}

func main	()  {
	count()
}
