package main

import (
	"fmt"
	"math/rand"
	"nightclub/WriteToFile"
	"nightclub/handlelogic"
	"nightclub/readconf"
	"nightclub/statistics"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	starttime := time.Now().UnixNano()
	conf := readconf.Initconf()
	ag := handlelogic.Gearsinit("config.json")
	ScoreCfgMap := ag.InitScoreCifToMap()
	filename := conf.OutFile
	if WriteToFile.FileExist(filename){
		WriteToFile.DeleteFile(filename)
	}
	f :=WriteToFile.CreatFile(filename)
	defer f.Close()
	prizesummap := statistics.InitPrizeSumMap(ScoreCfgMap)
	scoresum := 0
	for i := 0; i < conf.Count; i++ {
		sum := ag.CalculateLine(ScoreCfgMap, conf, prizesummap)
		scoresum = scoresum + sum
	}

	scorebase := ag.ScoreBase
	linecount := conf.Assignline[1] - conf.Assignline[0] + 1
	linescorebase := float64(scorebase) /float64(linecount)
	linecountsum:=linecount*conf.Count
	//fmt.Println(conf)
	endtime := time.Now().UnixNano()
	spendtime := endtime - starttime
	statistics.SortPrizeSumMap(prizesummap, ScoreCfgMap,linecountsum,linescorebase,f)
	//fmt.Printf("start time: %v\n", starttime)
	//fmt.Printf("end time:%v\n", endtime)
	fmt.Printf("用时: %v 秒\n", spendtime/1e9)
	fmt.Printf("ScoreBase: %v\n",scorebase)
	fmt.Printf("每次Spin中奖线: %v\n",linecount)
	fmt.Printf("总中奖线: %v\n",linecountsum)
	fmt.Printf("每条中奖线的押注: %v\n",linescorebase)
	fmt.Printf("总得分：%v\n", scoresum*scorebase)
	fmt.Printf("%v,%T\n", prizesummap, prizesummap)

}
