package db

import (
	"daping/src/daping/db/bean"
	"daping/src/daping/db/influx"
	"daping/src/daping/db/mysql"
	"daping/src/utils"
	"math/rand"
	"time"

	"github.com/cihub/seelog"
)

func INFLUX_FindPUEById(addrId string) ([]*bean.AllConsum, error) {
	devices, err := mysql.DB_AllConsumFindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
	// for _, device := range devices {
	// 	fmt.Printf("device: %v\n", device)
	// }
	if err != nil {
		seelog.Error("INFLUX_FindPUEById通过mgrid找typeId失败")
	}
	AllConsums, err := influx.INFLUX_FindPUEById(devices, addrId)
	//40000-50000
	max := 12000.0
	min := 11000.0
	if len(AllConsums) == 0 {
		ts := time.Now().AddDate(0, 0, -1)
		for i := 0; i < 13; i++ {
			AllConsum := &bean.AllConsum{
				Time: time.Date(ts.Year(), ts.Month(), ts.Day(), 0, i*2, 0, 0, ts.Location()),
				Num:  utils.Decimal2(rand.Float64()*(max-min) + min),
			}
			AllConsums = append(AllConsums, AllConsum)
		}
	}
	return AllConsums, nil
}
func INFLUX_FindWaterPresAndTem(addrId string) (*bean.AllWater, error) {
	aw, err := influx.INFLUX_FindWaterPresAndTem(addrId)
	if len(aw.BackWaterPre1) == 0 {
		//水压
		//供水3-4，回水3-3.2
		max := 3.9
		min := 3.7
		max1 := 3.2
		min1 := 3.0
		ts := time.Now().AddDate(0, 0, -1)
		var timeLine time.Time
		for i := 0; i < 13; i++ {
			timeLine = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, i*2, 0, 0, ts.Location())
		}
		for i := 0; i < 13; i++ {
			GiveWaterPre1 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max-min) + min),
				Name: "供水压力1",
			}
			aw.GiveWaterPre1 = append(aw.GiveWaterPre1, GiveWaterPre1)
		}

		for i := 0; i < 13; i++ {
			BackWaterPre1 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max1-min1) + min1),
				Name: "回水压力1",
			}
			aw.BackWaterPre1 = append(aw.BackWaterPre1, BackWaterPre1)
		}
		for i := 0; i < 13; i++ {
			GiveWaterPre2 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max-min) + min),
				Name: "供水压力2",
			}
			aw.GiveWaterPre2 = append(aw.GiveWaterPre2, GiveWaterPre2)
		}
		for i := 0; i < 13; i++ {
			BackWaterPre2 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max1-min1) + min1),
				Name: "回水压力2",
			}
			aw.BackWaterPre2 = append(aw.BackWaterPre2, BackWaterPre2)
		}
		//水温
		//供水11-13，回水18-19tem
		max = 13
		min = 11
		max1 = 19
		min1 = 18
		for i := 0; i < 13; i++ {
			GiveWaterTem1 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max-min) + min),
				Name: "供水温度1",
			}
			aw.GiveWaterTem1 = append(aw.GiveWaterTem1, GiveWaterTem1)
		}
		for i := 0; i < 13; i++ {
			BackWaterTem1 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max1-min1) + min1),
				Name: "回水温度1",
			}
			aw.BackWaterTem1 = append(aw.BackWaterTem1, BackWaterTem1)
		}
		for i := 0; i < 13; i++ {
			GiveWaterTem2 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max-min) + min),
				Name: "供水温度2",
			}
			aw.GiveWaterTem2 = append(aw.GiveWaterTem2, GiveWaterTem2)
		}
		for i := 0; i < 13; i++ {
			BackWaterTem2 := &bean.Water{
				Time: timeLine,
				Num:  utils.Decimal2(rand.Float64()*(max1-min1) + min1),
				Name: "回水温度2",
			}
			aw.BackWaterTem2 = append(aw.BackWaterTem2, BackWaterTem2)
		}
		return aw, nil
	}
	return aw, err
}
func INFLUX_GetArrowData(addrId string) (*bean.ArrowDatas, error) {
	return influx.GetArrowData(addrId)
}
