package main

import (
	"chao/count"
	"chao/lun"
	"io/ioutil"
	"strconv"
	"strings"

	"fmt"
	"github.com/spf13/viper"

	"time"
)

type Config struct {
	Count int
}
type ConfigFile struct {
	Config
}

func main() {
	/////////////读取配置文件/////////////////
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	var conf ConfigFile
	viper.Unmarshal(&conf) //反序列化到结构体
	/////////////读取配置文件/////////////////

	///////////////////初始化轮子/////////////////
	lunfilename := "./config.json"
	ag := lun.Luninit(lunfilename)
	///////////////////初始化轮子/////////////////

	//////////////////////test///////////////////
	/*
		//Dispay := count.Base(ag)
		test := make([][]int,5,5)
		test[0]=[]int{20,20,2}
		test[1]=[]int{2,20,2}
		test[2]=[]int{20,20,20}
		test[3]=[]int{2,20,3}
		test[4]=[]int{20,4,2}
	//fmt.Println("轮：",Dispay)
	//count.Linechart(ag,Dispay)
	//fmt.Println(test,ag)
	//sum := count.Linechart(ag,test)
	//fmt.Println(ag.Lines,ag.ScoreCfg,ag.ScoreBase)
	 */
	//////////////////////////test/////////////////////
	starttime := time.Now().String()
	var sum int
	sum = 0
	for i := 0; i < conf.Count; i++ {
		Dispay := count.Base(ag)
		//fmt.Println("#################################################################################################")
		//fmt.Println(Dispay)
		linescore := count.Linechart(ag, Dispay)
		sum = sum + linescore
	}
	sumfloat := float64(sum)
	countflaot := float64(conf.Count)
	spinavgscore := sumfloat / countflaot
	lineavgscore := spinavgscore / float64(len(ag.Lines))
	outfile := "out.txt"
	endtime := time.Now().String()
	//fmt.Printf("endtime:%v\n",time.Now())
	///////
	/*
		fmt.Printf("scorebase:%v\n",ag.ScoreBase)
		fmt.Printf("linecount:%v\n",len(ag.Lines))
		fmt.Printf("score sum: %v\n",sum)
		fmt.Printf("every spin avg score:%v\n",spinavgscore)
		fmt.Printf("line avg score:%v\n",lineavgscore)
	*/
	/////////////////////字符串拼接/////////////////////////
	var b strings.Builder
	b.WriteString("开始时间：")
	b.WriteString(starttime)
	b.WriteString("\n")
	b.WriteString("结束时间：")
	b.WriteString(endtime)
	b.WriteString("\n")
	b.WriteString("scorebase：")
	b.WriteString(strconv.Itoa(ag.ScoreBase))
	b.WriteString("\n")
	b.WriteString("总线数：")
	b.WriteString(strconv.Itoa(len(ag.Lines)))
	b.WriteString("\n")
	b.WriteString("总得分：")
	b.WriteString(strconv.Itoa(sum))
	b.WriteString("\n")
	b.WriteString("每次平均得分：")
	b.WriteString(strconv.FormatFloat(spinavgscore, 'f', -1, 64))
	b.WriteString("\n")
	b.WriteString("每条线平均得分：")
	b.WriteString(strconv.FormatFloat(lineavgscore, 'f', -1, 64))
	b.WriteString("\n")
	/////////////////////字符串拼接/////////////////////////

	/////////////////////写入文件/////////////////////////
	data := []byte(b.String())
	err = ioutil.WriteFile(outfile, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
	/////////////////////写入文件/////////////////////////
	fmt.Printf("%v\n保存文件名为：%v\n", b.String(),outfile)
	time.Sleep(time.Second*3)
}
