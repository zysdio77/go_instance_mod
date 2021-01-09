package handler

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
func Init64ToInt (int642 int64) int{
	int2 :=int(int642)
	return int2
}
func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeStamp2Date(timestamp int64) string {
	haomiao := timestamp/1000
	datetime := time.Unix(haomiao, 0).Format("2006-01-02 15:04:05")
	return datetime
}

func HitSplitYi(str string )bool{
	tag := 8
	c := len(str) / tag
	if c == 0 {
		return false
	} else {
		return true
	}
}
func SplitYi(str string) string {
	//每间隔8位添加一个"｜"的字符串
	var strtemp string
	tag := 8 //每隔 8位
	c := len(str) / tag
	var t int
	if HitSplitYi(str){
		for i := 0; i < c; i++ {
			//fmt.Printf("strtemp:%v\n",strtemp)
			if i == 0 {
				strtemp = fmt.Sprintf(str[:len(str)-tag] + "|" + str[len(str)-tag:])
				t = tag
				continue
			}
			t = t + tag + 1
			strtemp = fmt.Sprintf(strtemp[:len(strtemp)-t] + "|" + strtemp[len(strtemp)-t:])
		}
		//fmt.Printf("strtemp:%v\n",strtemp)
		return strtemp
	} else {
		//fmt.Printf("str:%v\n",str)
		return str
	}
}
func StringToBigInt(str string) *big.Int{
	//string转bigint
	big1,_ := new(big.Int).SetString(str,10)
	return big1
}

func BigIntToString(b *big.Int) string {
	//bigint转string
	return b.String()
}

func JoinString(strlist []string) string {
	return strings.Join(strlist,"")
}

func SplitString(str string) []string {
	return strings.Split(str,"|")
}

func AddBig(str1 string, str2 string) string {
	big1 := StringToBigInt(str1)
	big2 := StringToBigInt(str2)
	big2.Add(big1,big2)
	return big2.String()
}

func Updata(str1 string,str2 string) string {
	//接受的值位str1，数据库中的值位str2，返回str1+str2得出新值
	var str string
	if HitSplitYi(str2) {
		//数据库的字符串能被8整除
		strlist2 := SplitString(str2)
		str2 := JoinString(strlist2)
		//big2 := StringToBigInt(str2)
		str = AddBig(str1,str2)
	} else {
		//数据库的字符串不能被8整除
		str = AddBig(str1,str2)
	}
	return str
}