package handlelogic

import (
	"encoding/json"
	"fmt"
	"nightclub/readconf"
)

func (ag *AutoGenerated)InitScoreCifToMap() map[int][]int {
	//struct 转 map
	ScoreCfgMap := make(map[int][]int)
	ScoreCfgJson ,err := json.Marshal(ag.ScoreCfg)
	if err != nil {
		fmt.Printf("InitScoreCifToMap Error:%v\n",err)
	}
	json.Unmarshal(ScoreCfgJson,&ScoreCfgMap)
	//fmt.Println(ScoreCfgMap)
	return ScoreCfgMap
}

func (ag *AutoGenerated	)Score(ScoreCfgMap map[int][]int,prizemap  map[int]int,conf *readconf.ConfigFile,prizesummap map[int]map[int]int )  int{
	maxscore := 0
	for k,v := range prizemap{
		if k ==21 || k == 22 {
			continue
		}
		prizesummap[v][k]++
		//prizenummap.StatisticsPrizeSumMap(prizemap)
		score :=ScoreCfgMap[k][v-1]
		if score >= maxscore{
			maxscore = score
		}
		if conf.Debug{
			fmt.Printf("key : %v value : %v ,scorecfg : %v ,score : %v\n",k,v,ScoreCfgMap[k],score)
		}
	}
	if conf.Debug{
		fmt.Printf("maxscore:%v\n",maxscore)
	}
	return maxscore

}
