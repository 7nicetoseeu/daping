package influx

import (
	"context"
	"daping/src/config"
	"daping/src/daping/db/bean"
	"daping/src/daping/db/mysql"
	"daping/src/db"
	"daping/src/utils"
	"errors"
	"fmt"
	"math/rand"

	// "fmt"
	"math"
	"strconv"

	"github.com/cihub/seelog"
)

//48MWh--48000
func INFLUX_FindPUEById(devices []*bean.Device, addrId string) ([]*bean.AllConsum, error) {
	// mgrTypeId := "00000000-0PUE-0000-0000-000000000000"
	allConsums := make([]*bean.AllConsum, 0)
	client, ok := db.InfluxMap[addrId]
	if !ok {
		return nil, nil
	}
	defer client.Close()
	queryAPI := client.QueryAPI(db.InfluxBucket)
	for _, device := range devices {
		id := "\"" + device.Id + "\""
		start := strconv.Itoa(int(utils.GetUnixNowTime(-1, -2)))
		end := strconv.Itoa(int(utils.GetUnixNowTime(0, 0)))
		mid := utils.MonthZero()
		measurement := "\"" + mid + addrId + "\""
		// seelog.Debug("measurement:" + measurement) //"202207P001_A018"
		str := `
			from(bucket: "` + db.InfluxBucket + `")
				|> range(start: ` + start + `, stop: ` + end + `)
				|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
				|> filter(fn: (r) => r["mgr_obj_id"] == ` + id + `)
				|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "PUE_Real_Total_Value")
				|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
				|> yield(name: "mean")
		`
		result, err := queryAPI.Query(context.Background(), str)
		if err != nil {
			seelog.Error("INFLUX_FindPUEById error:" + err.Error())
			return nil, err
		}
		for result.Next() {
			// fmt.Printf("value: %v\n", result.Record().Value())
			// fmt.Printf("time: %v\n", result.Record().Time().Local())
			allConsum := bean.AllConsum{
				Num:  result.Record().Value().(float64),
				Time: result.Record().Time().Local(),
			}
			allConsums = append(allConsums, &allConsum)
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
			return nil, err
		}
	}

	return allConsums, nil
}

func INFLUX_FindWaterPresAndTem(addrId string) (*bean.AllWater, error) {
	PerAndTemMgr := mysql.DB_ZheXianMgrId(addrId)
	if PerAndTemMgr == nil {
		err := errors.New("water_pre_tem表出错")
		seelog.Error("err: %v\n", err.Error())
		return nil, err
	}
	var allWater *bean.AllWater
	allWater = &bean.AllWater{
		GiveWaterPre1: []*bean.Water{},
		BackWaterPre1: []*bean.Water{},
		GiveWaterPre2: []*bean.Water{},
		BackWaterPre2: []*bean.Water{},
		GiveWaterTem1: []*bean.Water{},
		BackWaterTem1: []*bean.Water{},
		GiveWaterTem2: []*bean.Water{},
		BackWaterTem2: []*bean.Water{},
	}
	for _, mgr := range PerAndTemMgr {
		client, ok := db.InfluxMap[addrId]
		if !ok {
			seelog.Error("配置文件中未导入", addrId, "influxdb相关数据")
			return allWater, nil
		}
		defer client.Close()

		bucket := config.Dcm.Influx.Bucket
		queryAPI := client.QueryAPI(bucket)
		start := strconv.Itoa(int(utils.GetUnixNowTime(-1, 0)))
		end := strconv.Itoa(int(utils.GetUnixNowTime(0, 0)))
		//返回202208
		mid := utils.MonthZero()
		measurement := "\"" + mid + addrId + "\""
		str := `
			from(bucket: "` + bucket + `")
				|> range(start: ` + start + `, stop: ` + end + `)
				|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
				|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr.Id + `")
				|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr.AttributeId + `")
				|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
				|> yield(name: "mean")
			`
		result, err := queryAPI.Query(context.Background(), str)
		if err != nil {
			seelog.Error("INFLUX_FindGiveWaterPres error:" + err.Error())
			return nil, err
		}
		for result.Next() {
			water := bean.Water{
				Name: mgr.Name,
				Time: result.Record().Time().Local(),
				Num:  result.Record().Value().(float64),
			}
			switch water.Name {
			case "回水压力1":
				allWater.BackWaterPre1 = append(allWater.BackWaterPre1, &water)
			case "回水压力2":
				allWater.BackWaterPre2 = append(allWater.BackWaterPre2, &water)
			case "供水压力1":
				allWater.GiveWaterPre1 = append(allWater.GiveWaterPre1, &water)
			case "供水压力2":
				allWater.GiveWaterPre2 = append(allWater.GiveWaterPre2, &water)
			case "回水温度1":
				allWater.BackWaterTem1 = append(allWater.BackWaterTem1, &water)
			case "回水温度2":
				allWater.BackWaterTem2 = append(allWater.BackWaterTem2, &water)
			case "供水温度1":
				allWater.GiveWaterTem1 = append(allWater.GiveWaterTem1, &water)
			case "供水温度2":
				allWater.GiveWaterTem2 = append(allWater.GiveWaterTem2, &water)
			}
		}

		// seelog.Debug(len(waters))
	}
	return allWater, nil
}

func GetArrowData(addrId string) (*bean.ArrowDatas, error) {
	arrowDatas := &bean.ArrowDatas{}
	// arrowDatas.ColdHuan, arrowDatas.ColdHuanNum = FindMonthColdArrow(addrId)
	// arrowDatas.ColdTong, arrowDatas.ColdTongNum = FindDayColdArrow(addrId)
	// arrowDatas.ElecHuan, arrowDatas.ElecHuanNum = FindMonthElecArrow(addrId)
	// arrowDatas.ElecTong, arrowDatas.ElecTongNum = FindMonthColdArrow(addrId)
	arrowDatas = &bean.ArrowDatas{
		ElecHuan:    rand.Intn(2),
		ColdHuan:    rand.Intn(2),
		ElecTong:    rand.Intn(2),
		ColdTong:    rand.Intn(2),
		ElecHuanNum: utils.Decimal2(rand.Float64()*(50-45) + 45),
		ColdHuanNum: utils.Decimal2(rand.Float64()*(50-45) + 45),
		ElecTongNum: utils.Decimal2(rand.Float64()*(50-45) + 45),
		ColdTongNum: utils.Decimal2(rand.Float64()*(50-45) + 45),
		IsEmpty:     0,
	}
	if arrowDatas.ElecHuan == 0 {
		arrowDatas.ElecHuan = -1
	}
	if arrowDatas.ColdHuan == 0 {
		arrowDatas.ColdHuan = -1
	}
	if arrowDatas.ElecTong == 0 {
		arrowDatas.ElecTong = -1
	}
	if arrowDatas.ColdTong == 0 {
		arrowDatas.ColdTong = -1
	}
	if arrowDatas.ElecHuan == 0 && arrowDatas.ElecTong == 0 && arrowDatas.ColdHuan == 0 && arrowDatas.ColdTong == 0 {
		err := errors.New("数据全为空")
		seelog.Error("err: ", err.Error())
		return arrowDatas, nil
	}
	return arrowDatas, nil
}

//00000000-0PUE-0000-0000-000000000018
//同比
func FindDayColdArrow(addrId string) (int, float64) {
	devices, _ := mysql.DB_AllConsumFindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
	device := &bean.Device{}
	for _, d := range devices {
		device = d
	}

	client, ok := db.InfluxMap[addrId]
	if !ok {
		seelog.Error("没有当前地址对应的influxdb")
		return 0, 0.0
	}
	defer client.Close()
	bucket := config.Dcm.Influx.Bucket
	queryAPI := client.QueryAPI(bucket)
	//返回202208
	mid := utils.MonthZero()
	measurement := "\"" + mid + addrId + "\""
	str := `
			from(bucket: "` + bucket + `")
			|> range(start: -2d)
			|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
			|> filter(fn: (r) => r["mgr_obj_id"] == "` + device.Id + `")
			|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "PUE_Day_Cool_Value")
			|> aggregateWindow(every: 1d, fn: mean, createEmpty: false)
			|> yield(name: "mean")
		`
	result, err := queryAPI.Query(context.Background(), str)
	if err != nil {
		seelog.Error("INFLUX_FindDayColdArrow error:" + err.Error())
		return 0, 0.0
	}
	TodayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}
	YesterdayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}

	for result.Next() {
		if YesterdayData.Time == 0 {
			YesterdayData.Time = int(result.Record().Time().Local().Unix())
			YesterdayData.Num = result.Record().Value().(float64)
		} else {
			TodayData.Time = int(result.Record().Time().Local().Unix())
			TodayData.Num = result.Record().Value().(float64)
		}
	}
	percent := Percent(YesterdayData.Num, TodayData.Num)
	if YesterdayData.Num > TodayData.Num {
		return -1, percent
	}
	if YesterdayData.Num < TodayData.Num {
		return 1, percent
	}
	return 0, 0.0
}

//环比
func FindMonthColdArrow(addrId string) (int, float64) {
	devices, _ := mysql.DB_AllConsumFindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
	device := &bean.Device{}
	for _, d := range devices {
		device = d
	}
	client, ok := db.InfluxMap[addrId]
	if !ok {
		return 0, 0.0
	}
	defer client.Close()
	bucket := config.Dcm.Influx.Bucket
	queryAPI := client.QueryAPI(bucket)
	// start := strconv.Itoa(int(utils.GetUnixNowTime(-2, 23)))
	// end := strconv.Itoa(int(utils.GetUnixNowTime(0, 0)))
	//返回202208
	mid := utils.MonthSubOne()
	measurement := "\"" + mid + addrId + "\""
	str := `
			from(bucket: "` + bucket + `")
			|> range(start: -32d)
			|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
			|> filter(fn: (r) => r["mgr_obj_id"] == "` + device.Id + `")
			|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "PUE_Day_Cool_Value")
			|> aggregateWindow(every: 1d, fn: mean, createEmpty: false)
			|> yield(name: "mean")
		`
	result, err := queryAPI.Query(context.Background(), str)
	if err != nil {
		seelog.Error("INFLUX_FindMonthColdArrow error:" + err.Error())
		return 0, 0.0
	}
	TodayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}
	YesterdayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}

	for result.Next() {
		if YesterdayData.Time == 0 {
			YesterdayData.Time = int(result.Record().Time().Local().Unix())
			YesterdayData.Num = result.Record().Value().(float64)
		} else {
			TodayData.Time = int(result.Record().Time().Local().Unix())
			TodayData.Num = result.Record().Value().(float64)
		}
	}
	percent := Percent(YesterdayData.Num, TodayData.Num)
	if YesterdayData.Num > TodayData.Num {
		return -1, percent
	}
	if YesterdayData.Num < TodayData.Num {
		return 1, percent
	}
	return 0, 0.0
}

//同比
func FindDayElecArrow(addrId string) (int, float64) {
	devices, _ := mysql.DB_AllConsumFindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
	device := &bean.Device{}
	for _, d := range devices {
		device = d
	}
	client, ok := db.InfluxMap[addrId]
	if !ok {
		return 0, 0.0
	}
	defer client.Close()
	bucket := config.Dcm.Influx.Bucket
	queryAPI := client.QueryAPI(bucket)
	// start := strconv.Itoa(int(utils.GetUnixNowTime(-2, 23)))
	// end := strconv.Itoa(int(utils.GetUnixNowTime(0, 0)))
	//返回202208
	mid := utils.MonthZero()
	measurement := "\"" + mid + addrId + "\""
	str := `
			from(bucket: "` + bucket + `")
			|> range(start: -2d)
			|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
			|> filter(fn: (r) => r["mgr_obj_id"] == "` + device.Id + `")
			|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "PUE_Day_Total_Value")
			|> aggregateWindow(every: 1d, fn: mean, createEmpty: false)
			|> yield(name: "mean")
		`
	result, err := queryAPI.Query(context.Background(), str)
	if err != nil {
		seelog.Error("INFLUX_FindDayColdArrow error:" + err.Error())
		return 0, 0.0
	}
	TodayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}
	YesterdayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}

	for result.Next() {
		if YesterdayData.Time == 0 {
			YesterdayData.Time = int(result.Record().Time().Local().Unix())
			YesterdayData.Num = result.Record().Value().(float64)
		} else {
			TodayData.Time = int(result.Record().Time().Local().Unix())
			TodayData.Num = result.Record().Value().(float64)
		}
	}
	percent := Percent(YesterdayData.Num, TodayData.Num)
	if YesterdayData.Num > TodayData.Num {
		return -1, percent
	}
	if YesterdayData.Num < TodayData.Num {
		return 1, percent
	}
	return 0, 0.0
}

//环比
func FindMonthElecArrow(addrId string) (int, float64) {
	devices, _ := mysql.DB_AllConsumFindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
	device := &bean.Device{}
	for _, d := range devices {
		device = d
	}
	client, ok := db.InfluxMap[addrId]
	if !ok {
		return 0, 0.0
	}
	defer client.Close()
	bucket := config.Dcm.Influx.Bucket
	queryAPI := client.QueryAPI(bucket)
	// start := strconv.Itoa(int(utils.GetUnixNowTime(-2, 23)))
	// end := strconv.Itoa(int(utils.GetUnixNowTime(0, 0)))
	//
	//返回202208
	mid := utils.MonthZero()
	measurement := "\"" + mid + addrId + "\""
	str := `
			from(bucket: "` + bucket + `")
			|> range(start: -31d)
			|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
			|> filter(fn: (r) => r["mgr_obj_id"] == "` + device.Id + `")
			|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "PUE_Day_Total_Value")
			|> aggregateWindow(every: 1d, fn: mean, createEmpty: false)
			|> yield(name: "mean")
		`
	result, err := queryAPI.Query(context.Background(), str)
	if err != nil {
		seelog.Error("INFLUX_FindMonthColdArrow error:" + err.Error())
		return 0, 0.0
	}
	TodayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}
	YesterdayData := bean.ArrowData{
		Time: 0,
		Num:  0.0,
	}

	for result.Next() {
		if YesterdayData.Time == 0 {
			YesterdayData.Time = int(result.Record().Time().Local().Unix())
			YesterdayData.Num = result.Record().Value().(float64)
		} else {
			TodayData.Time = int(result.Record().Time().Local().Unix())
			TodayData.Num = result.Record().Value().(float64)
		}
	}
	percent := Percent(YesterdayData.Num, TodayData.Num)
	if YesterdayData.Num > TodayData.Num {
		return -1, percent
	}
	if YesterdayData.Num < TodayData.Num {
		return 1, percent
	}
	return 0, 0.0
}
func Percent(old float64, new float64) float64 {
	if old == new {
		return 0
	}
	if old == 0.0 {
		return 100
	}
	data := math.Abs((new/old - 1) * 100)
	return utils.Decimal2(data)
}

// //供水压力1
// func INFLUX_FindGiveWaterPres1(addrId string) ([]*bean.WaterPressure, error) {
// 	waterPressures := make([]*bean.WaterPressure, 0)
// 	client := db.InfluxMap[addrId]
// 	defer client.Close()
// 	bucket := config.GetDBconfig().Influx.Bucket
// 	queryAPI := client.QueryAPI(bucket)
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	//返回202208
// 	mid := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + mid + addrId + "\""
// 	mgr_obj_id, mgr_obj_attribute_id := mysql.DB_ZheXianMgrId(addrId, "供水压力1")
// 	str := `
// 	from(bucket: "` + bucket + `")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr_obj_id + `")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr_obj_attribute_id + `")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterPres error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterPressure := bean.WaterPressure{
// 			Name: "供水压力1",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterPressures = append(waterPressures, &waterPressure)
// 	}
// 	return waterPressures, nil
// }
// //回水压力1
// func INFLUX_FindBackWaterPres1(addrId string) ([]*bean.WaterPressure, error) {
// 	waterPressures := make([]*bean.WaterPressure, 0)
// 	client := db.InfluxMap[addrId]
// 	defer client.Close()
// 	bucket := config.GetDBconfig().Influx.Bucket
// 	queryAPI := client.QueryAPI(bucket)
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	mid := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + mid + addrId + "\""
// 	mgr_obj_id, mgr_obj_attribute_id := mysql.DB_ZheXianMgrId(addrId, "回水压力1")
// 	str := `
// 	from(bucket: "` + bucket + `")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr_obj_id + `")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr_obj_attribute_id + `")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterPres error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterPressure := bean.WaterPressure{
// 			Name: "回水压力1",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterPressures = append(waterPressures, &waterPressure)
// 	}
// 	return waterPressures, nil
// }
// func INFLUX_FindGiveWaterPres2(addrId string) ([]*bean.WaterPressure, error) {
// 	waterPressures := make([]*bean.WaterPressure, 0)
// 	client := db.InfluxMap[addrId]
// 	defer client.Close()
// 	bucket := config.GetDBconfig().Influx.Bucket
// 	queryAPI := client.QueryAPI(bucket)
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	//返回202208
// 	mid := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + mid + addrId + "\""
// 	mgr_obj_id, mgr_obj_attribute_id := mysql.DB_ZheXianMgrId(addrId, "供水压力2")
// 	str := `
// 	from(bucket: "` + bucket + `")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr_obj_id + `")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr_obj_attribute_id + `")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterPres error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterPressure := bean.WaterPressure{
// 			Name: "供水压力2",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterPressures = append(waterPressures, &waterPressure)
// 	}
// 	return waterPressures, nil
// }
// //回水压力2
// func INFLUX_FindBackWaterPres2(addrId string) ([]*bean.WaterPressure, error) {
// 	waterPressures := make([]*bean.WaterPressure, 0)
// 	client := db.InfluxMap[addrId]
// 	bucket := config.GetDBconfig().Influx.Bucket
// 	defer client.Close()
// 	queryAPI := client.QueryAPI(bucket)
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	mid := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + mid + addrId + "\""
// 	mgr_obj_id, mgr_obj_attribute_id := mysql.DB_ZheXianMgrId(addrId, "回水压力2")
// 	str := `
// 	from(bucket: "` + bucket + `")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr_obj_id + `")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr_obj_attribute_id + `")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterPres error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterPressure := bean.WaterPressure{
// 			Name: "回水压力2",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterPressures = append(waterPressures, &waterPressure)
// 	}
// 	return waterPressures, nil
// }
// //供水温度1
// func INFLUX_FindGiveWaterTem1(addrId string) ([]*bean.WaterTem, error) {
// 	waterTems := make([]*bean.WaterTem, 0)
// 	client := db.InfluxMap[addrId]
// 	bucket := config.GetDBconfig().Influx.Bucket
// 	defer client.Close()
// 	queryAPI := client.QueryAPI(bucket)
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	year := strconv.Itoa(int(time.Now().Year()))
// 	month := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + year + month + addrId + "\""
// 	mgr_obj_id, mgr_obj_attribute_id := mysql.DB_ZheXianMgrId(addrId, "供水温度1")
// 	str := `
// 	from(bucket: "` + bucket + `")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "` + mgr_obj_id + `")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "` + mgr_obj_attribute_id + `")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterTem error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterTem := bean.WaterTem{
// 			Name: "供水温度1",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterTems = append(waterTems, &waterTem)
// 	}
// 	return waterTems, nil
// }
// func INFLUX_FindBackWaterTem1(addrId string) ([]*bean.WaterTem, error) {
// 	waterTems := make([]*bean.WaterTem, 0)
// 	client := db.InfluxMap[addrId]
// 	defer client.Close()
// 	queryAPI := client.QueryAPI("sspaas_test")
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	year := strconv.Itoa(int(time.Now().Year()))
// 	month := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + year + month + addrId + "\""

// 	str := `
// 	from(bucket: "sspaas_test")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "A18-10000-00ZG-0000-0000-100020006")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "BA_Param_Temp")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindBackWaterTem error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterTem := bean.WaterTem{
// 			Name: "回水温度1",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterTems = append(waterTems, &waterTem)
// 	}
// 	return waterTems, nil
// }
// func INFLUX_FindGiveWaterTem2(addrId string) ([]*bean.WaterTem, error) {
// 	waterTems := make([]*bean.WaterTem, 0)
// 	// client := influxdb2.NewClient("http://10.3.0.109:8086", "hmc72eicR3wRt1VJI4gYvMA3LmrocLR6A9Qu3XGFOexXYy7j5ujthEhslU4ehVg5sDTe6jCUdF5IdHe2OWOH5w==")
// 	client := db.InfluxMap[addrId]

// 	defer client.Close()
// 	queryAPI := client.QueryAPI("sspaas_test")
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	year := strconv.Itoa(int(time.Now().Year()))
// 	month := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + year + month + addrId + "\""

// 	str := `
// 	from(bucket: "sspaas_test")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "A18-20000-00ZG-0000-0000-100010001")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "BA_Param_Temp")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindGiveWaterTem error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterTem := bean.WaterTem{
// 			Name: "供水温度2",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterTems = append(waterTems, &waterTem)
// 	}

// 	return waterTems, nil
// }
// func INFLUX_FindBackWaterTem2(addrId string) ([]*bean.WaterTem, error) {
// 	waterTems := make([]*bean.WaterTem, 0)
// 	client := db.InfluxMap[addrId]

// 	defer client.Close()
// 	queryAPI := client.QueryAPI("sspaas_test")
// 	start := strconv.Itoa(int(utils.GetUnixNowTime(-1)))
// 	end := strconv.Itoa(int(utils.GetUnixNowTime(0)))
// 	year := strconv.Itoa(int(time.Now().Year()))
// 	month := utils.MonthZero(int(time.Now().Month()))
// 	measurement := "\"" + year + month + addrId + "\""

// 	str := `
// 	from(bucket: "sspaas_test")
// 		|> range(start: ` + start + `, stop: ` + end + `)
// 		|> filter(fn: (r) => r["_measurement"] == ` + measurement + `)
// 		|> filter(fn: (r) => r["mgr_obj_id"] == "A18-20000-00ZG-0000-0000-100020005")
// 		|> filter(fn: (r) => r["mgr_obj_attribute_id"] == "BA_Param_Temp")
// 		|> aggregateWindow(every: 2h, fn: mean, createEmpty: false)
// 		|> yield(name: "mean")
// 	`
// 	result, err := queryAPI.Query(context.Background(), str)
// 	if err != nil {
// 		seelog.Error("INFLUX_FindBackWaterTem error:" + err.Error())
// 	}
// 	for result.Next() {
// 		waterTem := bean.WaterTem{
// 			Name: "回水温度2",
// 			Time: result.Record().Time().Local(),
// 			Num:  result.Record().Value().(float64),
// 		}
// 		waterTems = append(waterTems, &waterTem)
// 	}
// 	return waterTems, nil
// }
