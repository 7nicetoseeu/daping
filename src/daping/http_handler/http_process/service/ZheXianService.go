package service

import (
	"daping/src/daping/db"
	"daping/src/daping/db/bean"
	"fmt"
	"time"
)

var Arrow *bean.ArrowDatas
var ArrowOldTime int64

func AllConsume(addrId string) ([]*bean.AllConsum, error) {
	return db.INFLUX_FindPUEById(addrId)
}
func AllConsumeFour(addrId string) (*bean.AllConsumFour, error) {
	ac, _ := db.INFLUX_FindPUEById(addrId)
	ac2, _ := db.INFLUX_FindPUEById(addrId)
	ac3, _ := db.INFLUX_FindPUEById(addrId)
	ac4, _ := db.INFLUX_FindPUEById(addrId)
	allConsumFour := &bean.AllConsumFour{
		AllConsum1: ac,
		AllConsum2: ac2,
		AllConsum3: ac3,
		AllConsum4: ac4,
	}
	return allConsumFour, nil
}

func Water(addrId string) (*bean.AllWater, error) {
	return db.INFLUX_FindWaterPresAndTem(addrId)
}
func ArrowData(addrId string) (*bean.ArrowDatas, error) {
	timeNow := time.Now().Unix()
	if (timeNow-ArrowOldTime) > 86400 || ArrowOldTime == 0 {
		//产生新数据
		fmt.Println("产生新数据")
		ArrowOldTime = timeNow
		arrow, _ := db.INFLUX_GetArrowData(addrId)
		Arrow = arrow
	}
	return Arrow, nil
}
