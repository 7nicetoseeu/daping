package service

import (
	"daping/src/daping/db"
	"daping/src/daping/db/bean"
	"time"
)

var QuanOldTime int64
var QuanData *bean.Quan

func DeviceNum(addrId string) ([]*bean.Device, error) {
	return db.MYSQL_FindDeviceNum(addrId)
}

func Quan(addrId string) (*bean.Quan, error) {
	timeNow := time.Now().Unix()
	if (timeNow-QuanOldTime) > 86400 || QuanOldTime == 0 {
		//产生新数据
		QuanOldTime = timeNow
		quan, _ := db.MONGO_FindQuan(addrId)
		QuanData = quan
	}
	return QuanData, nil
}

func Dianliu(addrId string) ([]*bean.ZhiLuDianLiu, error) {
	return db.MONGO_FindDianliu(addrId)
}

func Warning(addrId string) ([]*bean.Warning, error) {
	return db.MYSQL_FindWaring(addrId)
}
func FindPUE(addrId string) (*bean.PUE, error) {
	return db.MYSQL_FindPUE(addrId)
}
