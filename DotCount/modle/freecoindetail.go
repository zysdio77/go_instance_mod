package modle

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)
func (writedata *WriteFreeData) DataExist(tablename string) (bool ,*ReadFreeData){
	//数据是否存在，存在true，不存在false
	readrow := writedata.GormSelectRow(tablename)
	if GDb.Table(tablename).NewRecord(readrow) {
		return false ,nil
	} else {
		return true,readrow
	}
}
func (writedata *WriteFreeData)GormInsertRow(tablename string)  {
	//插入数据
	GDb.Table(tablename).Create(&writedata)
}
func(writedata *WriteFreeData) GormSelectRow(tablename string) *ReadFreeData {
	//查询数据
	var readrow  ReadFreeData
	GDb.Table(tablename).Where("uid = ? and date_str = ?",&writedata.Uid,&writedata.DateStr).Find(&readrow)
	return &readrow
}
func (readdata *ReadFreeData) GormUpdateRow(tablename string,freedata *ReadFreeData)  {
	//更新数据
	GDb.Table(tablename).Save(&readdata)
}

type FreeDetail struct {
	TimeStamp int64    `json:"timeStamp"`
	CoinNum   int    `json:"coinNum"`
	CoinName  string `json:"coinName"`
}
func FreeCoinDetails(readdata string,writedata string) string {
	var readdetails []FreeDetail
	var tempdetails []FreeDetail
	var writedetails []FreeDetail
	var err error
	err =json.Unmarshal([]byte(readdata),&readdetails)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"FreeCoinDetails下freedetail":"json解析错误",}).Error(err)
	}
	err = json.Unmarshal([]byte(writedata),&tempdetails)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"FreeCoinDetails下writedetail":"json解析错误",}).Error(err)
	}
	for _,r_v := range readdetails {
		for _, w_v := range tempdetails {
			if w_v.CoinName == r_v.CoinName {
				w_v.CoinNum = w_v.CoinNum + r_v.CoinNum
				writedetails = append(writedetails,w_v)
				//fmt.Printf("CoinName:%v ,w_v.CoinNum:%v\n",w_v.CoinName,w_v.CoinNum)
			}
		}
	}
	bytedata ,err := json.Marshal(&writedetails)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"FreeCoinDetails下writedetails":"json解析错误",}).Error(err)
	}
	//fmt.Printf("writedetails fanaily:%v,%v\n",writedetails,string(bytedata))
	return string(bytedata)
}