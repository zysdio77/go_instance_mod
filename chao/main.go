package main

import (
	"chao/WriteToFile"
	"chao/count"
	"chao/lun"
	"chao/readconf"
	"fmt"
	"math/rand"
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

	//prizenumsum,scorecfg := count.ScoreInit(ag)//初始化中奖线总数
	prizenumsum,scorecfg := count.ScoreInit2(ag)//初始化中奖线总数
	//fmt.Println(prizenumsum)
	sum = 0
	rand.Seed(time.Now().UnixNano()) //初始化种子一次即可，重复调用可能会产生重复的随机数
	for i := 0; i < conf.Count; i++ {
		Dispay := count.Base(ag)
		linescore := count.Linechart(ag, Dispay,conf,file,prizenumsum,scorecfg)
		sum = sum + linescore
	}
	sumfloat := float64(sum)
	countflaot := float64(conf.Count)
	spinavgscore := sumfloat / countflaot
	endtime := time.Now().String()
	linescount:=conf.Assignline[1]-conf.Assignline[0]+1
	linescountsum := linescount*conf.Count
	lineavgscore := sumfloat / float64(linescountsum)
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
	for v,k := range prizenumsum{
		sortv := make([]int,0)
		sortv	= append(sortv, v)
		//fmt.Println(v,k)
		for vv,kk := range k{
			avg := float64(kk)/float64(linescountsum)*100
			avgstr :=strconv.FormatFloat(avg, 'f', -1, 64)
			data :=fmt.Sprintf("图标%v中%v的次数为%v,平均出现几率：%v%v \n",v,vv,kk,avgstr,"%")
			WriteToFile.WriteTOFile(file,data)
		}
		sort.Ints(sortv)
		fmt.Println(sortv)
	}

	 */
	linescorebase := float64(ag.ScoreBase)/float64(linescount)
	//count.PrizeNumSumSort(prizenumsum,file,linescountsum,scorecfg,linescorebase)	//排序后写入文件
	count.PrizeNumSumSort2(prizenumsum,file,linescountsum,scorecfg,linescorebase)//降序后后写入文件，中5，中4，中3...
	WriteToFile.WriteTOFile(file,(b.String()))
	/////////////////////写入文件/////////////////////////
	//fmt.Println(prizenumsum)
	fmt.Printf("%v\n详情以保存到文件，名为：%v\n", b.String(),conf.OutFile)
	time.Sleep(time.Second*3)
}
