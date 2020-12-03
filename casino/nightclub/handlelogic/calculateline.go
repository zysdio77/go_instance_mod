package handlelogic

import (
	"fmt"
	"nightclub/readconf"
)

func(ag *AutoGenerated ) CalculateLine (ScoreCfgMap map[int][]int,conf *readconf.ConfigFile,prizesummap map[int]map[int]int ) int{
	gears := ag.Base()
	//gears := TestBase() //测试用的轮子，方便测试
	beforeicon := 0
	icon := 0
	wild := ag.Wild.Type[0]
	sum := 0
	//var prizemap map[int]int
	for i:=conf.Assignline[0];i<=conf.Assignline[1];i++{
		count := 0
		prizemap := make(map[int]int)
		for ii,jj := range ag.Lines[i]{
			icon = gears[ii][jj]
			if ii == 0 {
				beforeicon =icon
				count++
				prizemap[icon]=count
			}
			if ii >0 {
				if beforeicon == wild{
					count++
					prizemap[icon]=count
					beforeicon = icon
				} else {
					if icon == wild{
						count++
						prizemap[beforeicon]=count

					} else if icon == beforeicon  {
						count++
						prizemap[icon]=count

					} else {
						break
					}
				}
			}
		}
		if conf.Debug{
			fmt.Printf("line:%v,gears%v,prizemap:%v\n",ag.Lines[i],gears,prizemap)
		}
		//prizenummap.StatisticsPrizeSumMap(prizemap)
		//statistics.StatisticsPrizeSumMap(prizenummap,prizemap)
		//fmt.Printf("prizemap:%v,prizesummap%v\n",prizemap,prizenummap)
		maxscore :=ag.Score(ScoreCfgMap,prizemap,conf,prizesummap)
		sum = sum + maxscore
	}
	return sum
}