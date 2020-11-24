package WriteToFile

import (
	"chao/readconf"
	"fmt"
	"os"
)

func CreatFile(conf *readconf.ConfigFile) *os.File {	//创建文件
	file ,err := os.OpenFile(conf.OutFile,os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("写入文件失败：%v\n",err)
	}
	return file
}

func WriteTOFile(file *os.File,data string)  {	//写入文件
	_,err :=file.WriteString(data)
	if err != nil {
		fmt.Printf("写入文件失败：%v\n",err)
	}
}

func DeleteFile (filename string){	//删除文件
	os.Remove(filename)
}

func FileExist(filename string) bool {		//判断文件是否存在
	_ ,err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}