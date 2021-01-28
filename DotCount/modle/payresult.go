package modle

func (writedata *SelectPayResoltData) DataExist(tablename string) (bool ,*ReadPayResoltData){
	//数据是否存在，存在true，不存在false
	readrow := writedata.GormSelectRow(tablename)
	if GDb.Table(tablename).NewRecord(readrow) {
		return false ,nil
	} else {
		return true,readrow
	}
}
/*
func (writedata *WritePayResoltData)GormInsertRow(tablename string)  {
	//插入数据
	GDb.Table(tablename).Create(&writedata)
}

 */
func (writedata *SelectPayResoltData)GormSelectRow(tablename string) *ReadPayResoltData {
	//查询数据
	var readrow  ReadPayResoltData
	GDb.Table(tablename).Where("transaction_id = ?",&writedata.TransactionId).Find(&readrow)
	return &readrow
}
func (readdata *ReadPayResoltData) GormUpdateRow(tablename string) int64 {
	//更新数据
	db:=GDb.Table(tablename).Save(&readdata)
	//fmt.Printf("RowsAffected:%v\n",db.RowsAffected)
	//fmt.Printf("dberr:%v\n",db.Error)
	return db.RowsAffected
}