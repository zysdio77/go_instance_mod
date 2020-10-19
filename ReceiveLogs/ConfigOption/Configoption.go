package ConfigOption

//package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConif() *viper.Viper {
	v := viper.New()
	//v.SetDefault("port", "9999")       //设置默认值
	//v.SetDefault("logfile", "out.log") //设置默认值
	v.AddConfigPath(".")      //设置路径，.表示当前目录
	v.SetConfigName("config") //  设置配置文件名 (不带后缀)
	v.SetConfigType("yaml")   //设置配置文件类型

	return v
}

func ReadConfig(v *viper.Viper) { //读取配置文件
	err := v.ReadInConfig() //搜索路径并读取配置
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": "Configoption.ReadInConfig()",
		}).Error(err)
		//fmt.Errorf("Fatal error config file: %s \n", err)
	}
	//fmt.Println(v.Get("port"), v.Get("logfile"))
}

func MonitorConfig() { //监视配置文件，重新读取配置数据
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})
}

type ConfigInfo struct {
	Port    int
	LogFile string
}

//把配置文件内容解析到结构体
func Unmarshalconfig(v *viper.Viper) (ConfigInfo, error) {
	var c ConfigInfo       //声明一个对应配置文件的结构体
	ReadConfig(v)          //读取配置文件
	err := v.Unmarshal(&c) //解析配置并写入结构体
	if err != nil {
		return c, err
	}

	return c, err
}

/*
func main() {
	v := InitConif()
	//ReadConfig(v)
	c, err := Unmarshalconfig(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
*/
