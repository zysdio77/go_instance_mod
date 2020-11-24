package main

import (
	"chao/WriteToFile"
	"chao/count"
	"chao/lun"
	"chao/readconf"
	"fmt"
	"strconv"
	"strings"

	"time"
)

func main() {
	conf := readconf.ReadConfFile()	//读取配置文件
	///////////////////初始化轮子/////////////////
	lunfilename := "./config.json"
	ag := lun.Luninit(lunfilename)
	///////////////////初始化轮子/////////////////

	if WriteToFile.FileExist(conf.OutFile) {
		WriteToFile.DeleteFile(conf.OutFile)
	}
	file := WriteToFile.CreatFile(conf) //打开输出的文件
	defer file.Close()
	starttime := time.Now().String()
	var sum int
	//////////////////////初始化中奖线总数/////////////////////
	prizenumsum := make(map[int]map[int]int)	//中奖总数，例如，1中5共多少次
	scorecfg :=count.ScorecfgToMap(ag)
	for k,v :=range scorecfg{	//初始化map
		prizenumsum[k]= make(map[int]int)
		for i:=1;i<=len(v);i++{
			prizenumsum[k][i]=0
		}
	}

	//fmt.Println(prizenumsum)
	sum = 0
	for i := 0; i < conf.Count; i++ {
		Dispay := count.Base(ag)
		//fmt.Println("#################################################################################################")
		//fmt.Println(Dispay)
		linescore := count.Linechart(ag, Dispay,conf,file,prizenumsum)
		sum = sum + linescore
	}
	sumfloat := float64(sum)
	countflaot := float64(conf.Count)
	spinavgscore := sumfloat / countflaot
	lineavgscore := spinavgscore / float64(len(ag.Lines))
	endtime := time.Now().String()
	linescount:=conf.Assignline[1]-conf.Assignline[0]+1
	linescountsum := linescount*conf.Count
	//fmt.Println(float64(sum)/float64(linescountsum))
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
	b.WriteString("中奖线：")
	b.WriteString(strconv.Itoa(linescount))
	b.WriteString("\n")
	b.WriteString("总中奖线：")
	b.WriteString(strconv.Itoa(linescountsum))
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
	/*
	data := []byte(b.String())
	err := ioutil.WriteFile(outfile, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
	 */
	for v,k := range prizenumsum{
		//fmt.Println(v,k)
		for vv,kk := range k{
			//kk,_ := strconv.ParseFloat(fmt.Sprintf("%.50f",float64(kk)),64)
			//linescountsum,_ :=strconv.ParseFloat(fmt.Sprintf("%.50f",float64(linescountsum)),64)
			//avg := kk/linescountsum
			//avg ,_= strconv.ParseFloat(fmt.Sprintf("%.50f",float64(avg)),64)
			avg := float64(kk)/float64(linescountsum)
			avgstr :=strconv.FormatFloat(avg, 'f', -1, 64)
			data :=fmt.Sprintf("图标%v中%v的次数为%v,平均出现几率：%v\n",v,vv,kk,avgstr)
			WriteToFile.WriteTOFile(file,data)

		}
	}
	WriteToFile.WriteTOFile(file,(b.String()))
	/////////////////////写入文件/////////////////////////
	fmt.Println(prizenumsum)
	fmt.Printf("%v\n详情以保存到文件，名为：%v\n", b.String(),conf.OutFile)
	time.Sleep(time.Second*3)
}
