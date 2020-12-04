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
	prizesummap := handlelogic.InitPrizeSumMap(ScoreCfgMap)
	freeprizesummap := handlelogic.InitPrizeSumMap(ScoreCfgMap)
	scattermap := ag.InitScattermap()
	scoresum := 0
	freecountsum:= 0
	freescoresum := 0
	for i := 0; i < conf.Count; i++ {
		gears := ag.Base()
		//gears := ag.TestBase()
		freecount , ok :=ag.FreeTrigger(gears,conf,scattermap)
		if ok {
			freecount2,freesum :=ag.FreeGames(freecount,conf,scattermap,ScoreCfgMap,freeprizesummap)
			freecountsum = freecountsum +freecount2
			freescoresum = freescoresum + freesum
		}
		//fmt.Println(scattermutil)
		sum := ag.CalculateLine(gears,ScoreCfgMap, conf, prizesummap)
		scoresum = scoresum + sum
	}

	scorebase := ag.ScoreBase
	linecount := conf.Assignline[1] - conf.Assignline[0] + 1
	linescorebase := float64(scorebase) /float64(linecount)
	linecountsum:=linecount*conf.Count
	//fmt.Println(conf)
	endtime := time.Now().UnixNano()
	spendtime := endtime - starttime
	statistics.SortPrizeSumMap(prizesummap, ScoreCfgMap,linecountsum,linescorebase,f,"Base")

	scatterscoresum :=  statistics.SortScatterSumMap(ag,scattermap,conf,f)



	fmt.Printf("用时: %v 秒\n", spendtime/1e9)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("用时: %v 秒\n", spendtime/1e9))
	fmt.Printf("ScoreBase: %v\n",scorebase)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("ScoreBase: %v\n",scorebase))
	fmt.Printf("每次Spin中奖线: %v\n",linecount)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("每次Spin中奖线: %v\n",linecount))
	fmt.Printf("总中奖线: %v\n",linecountsum)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("总中奖线: %v\n",linecountsum))
	fmt.Printf("每条中奖线的押注: %v\n",linescorebase)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("每条中奖线的押注: %v\n",linescorebase))

	fmt.Printf("Free总次数:%v, free总分: %v \n",freecountsum,freescoresum)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("Free总次数:%v, free总分: %v \n",freecountsum,freescoresum))
	fmt.Printf("Base得分：%v , scatter得分: %v ,总得分：%v \n", scoresum*scorebase,scatterscoresum,scoresum*scorebase + freescoresum + scatterscoresum)
	WriteToFile.WriteTOFile(f,fmt.Sprintf("Base得分：%v , scatter得分: %v ,总得分：%v \n", scoresum*scorebase,scatterscoresum,scoresum*scorebase + freescoresum + scatterscoresum))
	//fmt.Printf("总得分：%v\n",scoresum*scorebase + freescoresum + scatterscoresum)




	fmt.Printf("%v,%T\n", prizesummap, prizesummap)
	fmt.Printf("scattermap: %v\n",scattermap)

	fmt.Printf("freeprizesummap: %v \n",freeprizesummap)
	freelinecount := conf.FreeAssignline[1] - conf.FreeAssignline[0] + 1
	freelinecountsum := freelinecount * freecountsum
	freelinescorebase :=float64(scorebase) /float64(freelinecount)
	fmt.Printf("freelinecount:%v,freelinecountsum:%v ,freelinescorebase:%v\n ",freelinecount,freelinecountsum,freelinescorebase)
	statistics.SortPrizeSumMap(freeprizesummap,ScoreCfgMap,freelinecountsum,freelinescorebase,f,"Free")

}
