package readconf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Count          int
	Assignline     []int
	OutFile        string
	Debug        bool
	FreeAssignline []int
	HitScatter     int
	Hitbonus int
	RespinCount int
}
type ConfigFile struct {
	Config
}

func Initconf() *ConfigFile{
	/////////////读取配置文件/////////////////
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Initconf viper.ReadInConfig Error:%v\n",err)
	}
	var conf ConfigFile
	viper.Unmarshal(&conf) //反序列化到结构体
	//fmt.Printf("%v,%T\n",conf.Details,conf.Details)
	/////////////读取配置文件/////////////////
	return &conf
}
