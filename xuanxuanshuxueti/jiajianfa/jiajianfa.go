package jiajianfa

import (
	"strconv"
	"strings"
)

func JiaFa ()[]string{	//10以内加法
	var tiku []string
	tiku = make([]string, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			var b strings.Builder
			if i+j > 10 {
				continue
			} else {
				//fmt.Printf("%v + %v,=%v\n",i,j,i+j)
				ii := strconv.Itoa(i)
				jj := strconv.Itoa(j)
				b.WriteString(ii)
				b.WriteString("+")
				b.WriteString(jj)
				b.WriteString("=")
				tiku = append(tiku, b.String())
			}
		}
	}
	return tiku
}

func JianFa () []string {	//10以内加法
	var tiku []string
	tiku = make([]string, 0)
	for i := 10; i >= 1; i-- {
		for j := 10; j >= 1; j-- {
			var b strings.Builder
			if i-j >= 0 {
				//fmt.Printf("%v - %v,=%v\n", i, j, i-j)
				ii := strconv.Itoa(i)
				jj := strconv.Itoa(j)
				b.WriteString(ii)
				b.WriteString("-")
				b.WriteString(jj)
				b.WriteString("=")
				tiku = append(tiku, b.String())
			} else {
				continue
			}
		}
	}
	return tiku
}

func JiaAndJianfa()[]string{	//10以内加减法混合
	jia :=JiaFa()
	jian := JianFa()
	for _ ,j := range jian{
		jia = append(jia, j)
	}
	return jia
}