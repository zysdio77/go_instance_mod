package count

import (
	"chao/WriteToFile"
	"chao/lun"
	"chao/readconf"
	"fmt"
	"os"
)

func Linechart(ag *lun.AutoGenerated, Dispay [][]int, conf *readconf.ConfigFile, file *os.File, prizenumsum map[int]map[int]int, scorecfg map[int][]int) int {
	//计算base上线的得分
	wild := ag.Wild.Type[0]
	/*v := reflect.ValueOf(ag.ScoreCfg)
	l := v.NumField()
	scoremap := make(map[int]string)
	namelist := []string{"Num0", "Num1", "Num2", "Num3", "Num4", "Num5", "Num6", "Num7", "Num8", "Num9", "Num20"}
	for i := 0; i < l; i++ {
		scoremap[i] = namelist[i]
	}

	*/
	var sum int
	//var winnumber map[int]int
	var count int
	var wildcount int
	var firsticon int
	var icon int
	//scorecfg:=ScorecfgToMap(ag)
	assignline := conf.Assignline //从配置文件读取中奖线信息指定遍历的中奖线
	//fmt.Println(assignline)
	//////////////////////////指定中奖线///////////////////////
	for i := assignline[0]; i <= assignline[1]; i++ {
		//fmt.Println(ag.Lines[i])
		count = 0
		wildcount = 0
		winnumber := make(map[int]int, 5) //中奖号码key是中奖图标，value是中几（例如：连续5个1，就是1:5）
		var maxscoreline int
		for i, j := range ag.Lines[i] {
			if i == 0 {
				firsticon = Dispay[0][j] //每行的第一个图标
				count++
				if firsticon == wild {
					wildcount++
					winnumber[firsticon] = wildcount
				} else {
					winnumber[firsticon] = count
				}

			}

			if i > 0 {
				icon = Dispay[i][j] //取图标
				if firsticon == wild {
					count++
					if icon == wild {
						wildcount++
						winnumber[firsticon] = wildcount
					} else {
						winnumber[icon] = count
						firsticon = icon
					}
				} else {
					if icon == wild {
						count++
						winnumber[firsticon] = count
					} else {
						if firsticon == icon {
							count++
							winnumber[icon] = count
							firsticon = icon
						} else {
							break
						}
					}
				}
			}
		}
		maxscoreline = Score(scorecfg, winnumber, prizenumsum) //一条线上的最大分数
		sum = sum + maxscoreline                               //总分
		if conf.Details {
			//fmt.Sprintf("中奖线:%v,展示的图标:%v,中奖图标次数:%v,firsticon:%v,icon:%v\n", ag.Lines[i], Dispay, winnumber,firsticon,icon)
			out := fmt.Sprintf("中奖线:%v,展示的图标:%v,中奖图标次数:%v,中奖线得分：%v\n", ag.Lines[i], Dispay, winnumber, maxscoreline)
			WriteToFile.WriteTOFile(file, out)
		}
	}
	//////////////////////////指定中奖线///////////////////////

	/////////////////遍历所有中奖线//////////////////
	/*
		for _, line := range ag.Lines {
			count = 0
			wildcount = 0
			winnumber = make(map[int]int, 5)
			for i, j := range line {
				if i == 0 {
					firsticon = Dispay[0][j] //每行的第一个图标
					count++
					if firsticon == wild {
						wildcount++
						winnumber[firsticon] = wildcount
					} else {
						winnumber[firsticon] = count
					}

				}

				if i > 0 {
					icon = Dispay[i][j] //取图标
					if firsticon == wild {
						count++
						if icon == wild {
							wildcount++
							winnumber[firsticon] = wildcount
						} else {
							winnumber[icon] = count
							firsticon = icon
						}
					} else {
						if icon == wild {
							count++
							winnumber[firsticon] = count
						} else {
							if firsticon == icon {
								count++
								winnumber[icon] = count
								firsticon = icon
							} else {
								break
							}
						}
					}
				}
			}
			//fmt.Printf("line:%v,count:%v,winnumber:%v,icon:%v,firsticon:%v\n", line, count, winnumber, icon, firsticon)
			maxscoreline := Score(ag, winnumber)	//一条线上的最大分数
			sum = sum + maxscoreline	//总分
		}
	*/
	/////////////////遍历所有中奖线//////////////////
	return sum
	//W.Done()
}

func FreeLinechart(ag *lun.AutoGenerated, FreeDisplay [][]int, conf *readconf.ConfigFile, file *os.File, freeprizenumsum map[int]map[int]int, scorecfg map[int][]int) int {
	//计算free一条线上的分数，（每条线的scorebase为1的情况，）
	wild := ag.Wild.Type[0]
	/*v := reflect.ValueOf(ag.ScoreCfg)
	l := v.NumField()
	scoremap := make(map[int]string)
	namelist := []string{"Num0", "Num1", "Num2", "Num3", "Num4", "Num5", "Num6", "Num7", "Num8", "Num9", "Num20"}
	for i := 0; i < l; i++ {
		scoremap[i] = namelist[i]
	}

	*/
	var sum int
	//var winnumber map[int]int
	var count int
	var wildcount int
	var firsticon int
	var icon int
	//scorecfg:=ScorecfgToMap(ag)
	freeassignline := conf.FreeAssignline //从配置文件读取中奖线信息指定遍历的中奖线
	//fmt.Println(freeassignline)
	//////////////////////////指定中奖线///////////////////////
	for i := freeassignline[0]; i <= freeassignline[1]; i++ {
		//fmt.Println(ag.Lines[i])
		count = 0
		wildcount = 0
		winnumber := make(map[int]int, 10) //中奖号码key是中奖图标，value是中几（例如：连续5个1，就是1:5）
		var maxscoreline int
		for i, j := range ag.Lines[i] {
			if i == 0 {
				firsticon = FreeDisplay[0][j] //每行的第一个图标
				count++
				if firsticon == wild {
					wildcount++
					winnumber[firsticon] = wildcount
				} else {
					winnumber[firsticon] = count
				}

			}

			if i > 0 {
				icon = FreeDisplay[i][j] //取图标
				if firsticon == wild {
					count++
					if icon == wild {
						wildcount++
						winnumber[firsticon] = wildcount
					} else {
						winnumber[icon] = count
						firsticon = icon
					}
				} else {
					if icon == wild {
						count++
						winnumber[firsticon] = count
					} else {
						if firsticon == icon {
							count++
							winnumber[icon] = count
							firsticon = icon
						} else {
							break
						}
					}
				}
			}
		}
		maxscoreline = FreeScore(scorecfg, winnumber, freeprizenumsum) //一条线上的最大分数
		sum = sum + maxscoreline                                       //总分
		//fmt.Printf("FreeDisplay: %v , maxscoreline: %v ,winnumber: %v ,sum: %v \n",FreeDisplay,maxscoreline,winnumber,sum)
		if conf.Details {
			//fmt.Sprintf("中奖线:%v,展示的图标:%v,中奖图标次数:%v,firsticon:%v,icon:%v\n", ag.Lines[i], Dispay, winnumber,firsticon,icon)
			out := fmt.Sprintf("Free:中奖线:%v,展示的图标:%v,中奖图标次数:%v,中奖线得分：%v\n", ag.Lines[i], FreeDisplay, winnumber, maxscoreline)
			WriteToFile.WriteTOFile(file, out)
		}
	}
	//fmt.Printf("sum111111111:%v\n",sum)
	return sum
	//////////////////////////指定中奖线///////////////////////

	/////////////////遍历所有中奖线//////////////////
	/*
		for _, line := range ag.Lines {
			count = 0
			wildcount = 0
			winnumber = make(map[int]int, 5)
			for i, j := range line {
				if i == 0 {
					firsticon = Dispay[0][j] //每行的第一个图标
					count++
					if firsticon == wild {
						wildcount++
						winnumber[firsticon] = wildcount
					} else {
						winnumber[firsticon] = count
					}

				}

				if i > 0 {
					icon = Dispay[i][j] //取图标
					if firsticon == wild {
						count++
						if icon == wild {
							wildcount++
							winnumber[firsticon] = wildcount
						} else {
							winnumber[icon] = count
							firsticon = icon
						}
					} else {
						if icon == wild {
							count++
							winnumber[firsticon] = count
						} else {
							if firsticon == icon {
								count++
								winnumber[icon] = count
								firsticon = icon
							} else {
								break
							}
						}
					}
				}
			}
			//fmt.Printf("line:%v,count:%v,winnumber:%v,icon:%v,firsticon:%v\n", line, count, winnumber, icon, firsticon)
			maxscoreline := Score(ag, winnumber)	//一条线上的最大分数
			sum = sum + maxscoreline	//总分
		}
	*/
	/////////////////遍历所有中奖线//////////////////
}
