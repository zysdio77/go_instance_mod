package statistics

import (
	"fmt"
	"nightclub/WriteToFile"
	"os"
	"sort"
	"strconv"
)

type PrizSumMap map[int]map[int]int

func InitPrizeSumMap(ScoreCfgMap map[int][]int) map[int]map[int]int{
	//var prizesummap map[int]map[int]int
	prizesummap := make(map[int]map[int]int)
	fmt.Println(ScoreCfgMap)
	for i:=1; i<=len(ScoreCfgMap[0]);i++{
		prizesummap[i] = make(map[int]int)
	}
	for k,v := range ScoreCfgMap{
		for i:=1;i<=len(v);i++{
			//prizesummap[i] = make(map[int]int)
			prizesummap[i][k]=0
			//fmt.Println(prizesummap)
		}

	}
	//prizesummap.SortPrizeSumMap()
	fmt.Printf("----------------%v\n",prizesummap)
	return prizesummap
}

func StatisticsPrizeSumMap(prizesummap map[int]map[int]int ,prizemap map[int]int)  {
	for v,k := range prizemap{
		prizesummap[k][v]++
	}
}

func SortPrizeSumMap(prizesummap map[int]map[int]int,ScoreCfgMap  map[int][]int,linescountsum int,linescorebase float64,f *os.File)  {
	sortkeys := make([]int, 0)
	Reversesortkeys := make([]int, 0)
	sortkeys1 := make([]int, 0)
	for k, _ := range prizesummap {
		sortkeys = append(sortkeys, k)
	}

	for k, _ := range prizesummap[1] {
		sortkeys1 = append(sortkeys1, k)
	}

	sort.Ints(sortkeys)
	sort.Ints(sortkeys1)
	for k, _ := range sortkeys { //把排序好的切片反过来实现倒序
		Reversesortkeys = append(Reversesortkeys, sortkeys[len(sortkeys)-1-k])
	}

	for _, v := range Reversesortkeys {
		if v == 1 {
			continue
		}
		fmt.Printf("Base 中%v:\n", v)
		data := fmt.Sprintf("Base中%v:\n", v)
		WriteToFile.WriteTOFile(f, data)
		for _, vv := range sortkeys1 {
			//avg := float64(prizesummap[v][vv])/float64(conf.Count)*100
			avg := float64(prizesummap[v][vv]) / float64(linescountsum) * 100
			avgstr := strconv.FormatFloat(avg, 'f', 6, 64)
			hitscoreCfg := ScoreCfgMap[vv][v-1]
			scoreout := float64(hitscoreCfg) * float64(prizesummap[v][vv]) * linescorebase
			scoreoutstr := strconv.FormatFloat(scoreout,'f',-1,64)
			avgout := float64(scoreout) / float64(linescountsum) / linescorebase * 100
			avgoutstr := strconv.FormatFloat(avgout,'f',6,64)
			fmt.Printf("        图标%v，次数%v，平均出现几率: %v%v 总分(OUT): %v , 平均OUT: %v%v\n", vv, prizesummap[v][vv],avgstr,"%",scoreoutstr,avgoutstr,"%")
			data := fmt.Sprintf("        图标%v，次数%v，平均出现几率: %v%v 总分(OUT): %v , 平均OUT: %v%v\n", vv, prizesummap[v][vv], avgstr, "%", scoreoutstr, avgoutstr, "%")
			WriteToFile.WriteTOFile(f, data)
		}

	}
	fmt.Printf("prizenumsum:%v,sortkeys:%v,sortkeys1:%v,Reversesortkeys:%v\n", prizesummap, sortkeys, sortkeys1, Reversesortkeys)
}