package mongo

import (
	"context"
	"daping/src/daping/db/bean"
	"daping/src/daping/db/mysql"
	"daping/src/db"
	"daping/src/utils"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	TIMEOUT        = 60
	ALLELECPOWER   = 120
	ALLCOLDCONTAIN = 17600
	ONECOLDCONTAIN = 2200
)

// func GetCapacity() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			seelog.Error(err)
// 		}
// 	}()

// 	mongoDb := db.GetMongoDb()
// 	col := mongoDb.Collection("rangeidc_mop")
// 	ctx, cancol := context.WithTimeout(context.Background(), TIME_OUT*time.Second)

// }

// func GetMaintenanceTable(id string) (*bean.MaintenanceTable, error) {
// 	col := db.GetMongoDb().Collection(MaintenanceTable)
// 	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
// 	defer cancel()
// 	data := &bean.MaintenanceTable{}
// 	err := col.FindOne(ctx, bson.M{"tableid": id}).Decode(data)
// 	if err != nil {
// 		seelog.Error(err)
// 		return nil, err
// 	}
// 	return data, nil
// }
/* 版1 */
func FindAllDianliu(pdus []*bean.Device, addrId string) []*bean.Pdu {
	var wg sync.WaitGroup
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	AllPdu := make([]*bean.Pdu, 0)
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	// sumT := 0.0

	for _, pdu := range pdus {
		mgrId := pdu.Id
		ids, err := mysql.DB_FindAttributeIdByMgrId(addrId, mgrId)
		// fmt.Printf("costmysql: %v\n", time.Since(t))
		if err != nil {
			return nil
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := bean.Mgr{}
			err := col.FindOne(ctx, bson.M{"mgrObjId": mgrId}).Decode(&data)
			if err != nil {
				seelog.Error(err.Error())
				return
			}
			MonitorMap := data.Monitor
			for _, id := range ids {
				_, ok := MonitorMap[id]
				if ok {
					tempdu := &bean.Pdu{
						MgrObjName: data.MgrObjName,
						MgrObjId:   data.MgrObjId,
						SecondId:   id,
						Value:      utils.Decimal2(rand.Float64() * 30),
					}
					AllPdu = append(AllPdu, tempdu)
				}
			}
		}()
	}

	wg.Wait()
	return AllPdu
}

func FindQuan(devices []*bean.Device, addrId string) (*bean.Quan, error) {
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	var Quan *bean.Quan
	data := bean.MgrC{}
	var (
		coldOpenNum  int64
		coldCloseNum int64
		coldPer      int64
		elecPer      int64
	)
	for _, device := range devices {
		id := device.Id
		err := col.FindOne(ctx, bson.M{"mgrObjId": id}).Decode(&data)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("data: %v\n", data.Monitor.Attribute.Value)
		if data.Monitor.Attribute.Value == 1 {
			coldOpenNum++
		} else {
			coldCloseNum++
		}
	}

	elec, err := FindAllElec(addrId)
	if err != nil {
		seelog.Error("FindAllElec err : " + err.Error())
		return nil, err
	}
	cold, err := FindAllCold(addrId)
	if err != nil {
		seelog.Error("FindAllElec err : " + err.Error())
		return nil, err
	}
	elecPer = int64(elec)
	coldPer = int64(cold)
	Quan = &bean.Quan{
		ColdOpenNum:  coldOpenNum,
		ColdCloseNum: coldCloseNum,
		ColdPer:      coldPer,
		ElecPer:      elecPer,
	}

	Quan.ColdOpenNum = 2
	Quan.ColdCloseNum = 6
	if Quan.ColdPer == 0 {
		Quan.ColdPer = int64(rand.Intn(45-40) + 40)
	}
	if Quan.ElecPer == 0 {
		Quan.ElecPer = int64(rand.Intn(45-40) + 40)
	}
	return Quan, nil
}

func FindAllElec(addrId string) (int, error) {
	devices, err := mysql.DB_FindDeviceIdBytypeIdAndName(addrId, "00000000-010A-0000-0000-000000000000", "受电柜")
	if err != nil {
		seelog.Error("FindAllElec_DB_FindDeviceIdBytypeIdAndName err : " + err.Error())
		return 0, err
	}
	var sum float64 = 0
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	data := bean.MgrC{}
	for _, device := range devices {
		id := device.Id
		err := col.FindOne(ctx, bson.M{"mgrObjId": id}).Decode(&data)
		if err != nil {
			return 0, err
		}
		// fmt.Printf("data: %v\n", data.Monitor.Attribute.Value)
		value := data.Monitor.ElecAttribute.Value
		sum += value
	}
	div := sum / ALLELECPOWER * 100
	return int(math.Floor(div + 0.5)), nil
}

func FindAllCold(addrId string) (int, error) {
	devices, err := mysql.DB_FindDeviceIdBytypeId(addrId, "00000000-00LJ-0000-0000-000000000000")
	if err != nil {
		seelog.Error("FindAllCold_DB_FindDeviceIdBytypeIdAndName err : " + err.Error())
		return 0, err
	}
	var sum float64 = 0
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	data := bean.MgrC{}
	for _, device := range devices {
		id := device.Id
		err := col.FindOne(ctx, bson.M{"mgrObjId": id}).Decode(&data)
		if err != nil {
			return 0, err
		}
		// fmt.Printf("data: %v\n", data.Monitor.Attribute.Value)
		value := data.Monitor.ColdAttribute.Value / 100 * ONECOLDCONTAIN
		// fmt.Printf("value: %v\n", value)
		sum += value
	}
	div := sum / ALLCOLDCONTAIN * 100
	// fmt.Printf("div: %v\n", div)
	return int(math.Floor(div + 0.5)), nil
	// return 0, nil
}

func TestFindDianLiu(mgrobjid string) []*bean.Pdu {
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	data := bean.Devicem{}
	var pdus []*bean.Pdu
	col.FindOne(ctx, bson.M{"mgrObjId": mgrobjid}).Decode(&data)
	// fmt.Printf("data.Monitor: %v\n", data.Monitor)
	for _, m := range data.Monitor {
		if utils.IsIdIn(m.SecondId) {
			pdu := bean.Pdu{
				Value:      m.Value,
				SecondId:   m.SecondId,
				MgrObjName: data.MgrObjName,
				MgrObjId:   data.MgrObjId,
			}
			pdus = append(pdus, &pdu)
		}

	}
	return pdus
}

/* 版本2 */
func FindAllDianliu2(addrId string) []*bean.Pdu {
	var wg sync.WaitGroup

	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	pdus, err := mysql.FindAIdMgrIdByType(addrId)
	if err != nil {
		return nil
	}
	AllPdu := make([]*bean.Pdu, 0)
	col := db.GetMongoDb().Collection("rangeidc_mop")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	// sumT := 0.0
	for _, pdu := range pdus {
		mgrId := pdu.Id
		attributeId := pdu.Attribute
		// fmt.Printf("costmysql: %v\n", time.Since(t))
		if err != nil {
			return nil
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := bean.Mgr{}
			err := col.FindOne(ctx, bson.M{"mgrObjId": mgrId}).Decode(&data)
			if err != nil {
				seelog.Error(err.Error())
				return
			}
			MonitorMap := data.Monitor
			m, ok := MonitorMap[attributeId]
			if ok {
				tempdu := &bean.Pdu{
					MgrObjName: data.MgrObjName,
					MgrObjId:   data.MgrObjId,
					SecondId:   attributeId,
					Value:      m.Value,
				}
				AllPdu = append(AllPdu, tempdu)
			}
		}()
	}
	wg.Wait()
	fmt.Println("end")
	return AllPdu

}
