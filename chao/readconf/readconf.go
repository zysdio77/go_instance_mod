package readconf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Count int
	Assignline []int
	OutFile string
	Details bool
	FreeAssignline []int
	HitScatter int
}
type ConfigFile struct {
	Config
}

func ReadConfFile() *ConfigFile {
	/////////////读取配置文件/////////////////
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	var conf ConfigFile
	viper.Unmarshal(&conf) //反序列化到结构体
	//fmt.Printf("%v,%T\n",conf.Details,conf.Details)
	/////////////读取配置文件/////////////////
	return &conf
}

