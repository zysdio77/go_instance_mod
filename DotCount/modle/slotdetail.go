package modle

func (writedata *WriteSlotDetailData) DataExist(tablename string) (bool ,*ReadSlotDetailData){
	//数据是否存在，存在true，不存在false
	readrow := writedata.GormSelectRow(tablename)
	if GDb.Table(tablename).NewRecord(readrow) {
		return false ,nil
	} else {
		return true,readrow
	}
}
func (writedata *WriteSlotDetailData)GormInsertRow(tablename string)  {
	//插入数据
	GDb.Table(tablename).Create(&writedata)
}
func(writedata *WriteSlotDetailData) GormSelectRow(tablename string) *ReadSlotDetailData {
	//查询数据
	var readrow  ReadSlotDetailData
	GDb.Table(tablename).Where("uid = ? and date_str = ? and slot_id = ?",&writedata.Uid,&writedata.DateStr ,&writedata.SlotId).Find(&readrow)
	return &readrow
}
func (readdata *ReadSlotDetailData) GormUpdateRow(tablename string)  {
	//更新数据
	GDb.Table(tablename).Save(&readdata)
}