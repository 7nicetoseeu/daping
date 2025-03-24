package db

import (
	"daping/src/daping/db/bean"
	"daping/src/daping/db/mongo"
	"daping/src/daping/db/mysql"
	hbean "daping/src/daping/http_handler/bean"
	"daping/src/utils"
	"errors"
	"strings"

	"github.com/cihub/seelog"
)

const (
	//查询电流条数
	ORDER int = 50
)

func MONGO_FindDianliu(addrId string) ([]*bean.ZhiLuDianLiu, error) {
	deviceIds, err := mysql.DB_FindDeviceIdBytypeId(addrId, "")
	if deviceIds == nil {
		err1 := errors.New("支路电流通过类型id获取所有的设备id为空")
		seelog.Error(err1.Error() + err.Error())
		return nil, err1
	}
	AllDianLiu := mongo.FindAllDianliu(deviceIds, addrId)
	if len(AllDianLiu) == 0 {
		return nil, err
	}
	Dianliu50 := utils.Max50Dianliu(AllDianLiu)

	zhiLus := make([]*bean.ZhiLuDianLiu, 0)
	for _, dianliu := range Dianliu50 {
		if dianliu == nil {
			continue
		}
		name := MYSQL_FindAttributeNameById(dianliu.MgrObjId, dianliu.SecondId, addrId)
		zhilu := &bean.ZhiLuDianLiu{
			Id:        dianliu.MgrObjId,
			Name:      strings.ReplaceAll(name, "电流", ""),
			Num:       dianliu.Value,
			Attribute: dianliu.MgrObjName,
		}
		zhiLus = append(zhiLus, zhilu)
	}

	return zhiLus, nil

}

// func MONGO_Find10Dianliu(addrId string) ([]*bean.ZhiLuDianLiu, error) {
// 	deviceIds, err := mysql.DB_FindTopDeviceIdBytypeId(addrId, "")
// 	if err != nil {
// 		seelog.Error("MONGO_FindAllDianliu,获取设备ID失败")
// 	}

// 	//测试单个设备
// 	// device := bean.Device{
// 	// 	Id: "A18M10000-00AF-002F-R001-000000000002",
// 	// }
// 	// deviceIds := make([]*bean.Device, 0)
// 	// deviceIds = append(deviceIds, &device)

// 	AllDianLiu := mongo.FindAllDianliu(deviceIds, addrId)

// 	// for num, dianliu := range AllDianLiu {
// 	// 	fmt.Printf("%v----dianliu: %v\n", num, dianliu)
// 	// }

// 	Dianliu10 := utils.StructMax10(AllDianLiu)
// 	zhiLus := make([]*bean.ZhiLuDianLiu, 0)
// 	for _, dianliu := range Dianliu10 {

// 		name := MYSQL_FindAttributeNameById(dianliu.MgrObjId, dianliu.SecondId, addrId)
// 		zhilu := &bean.ZhiLuDianLiu{
// 			Name:      name,
// 			Num:       dianliu.Value,
// 			Attribute: dianliu.MgrObjName,
// 		}
// 		zhiLus = append(zhiLus, zhilu)
// 	}

// 	return zhiLus, nil
// }

func MONGO_FindQuan(addrId string) (*bean.Quan, error) {
	device, err := mysql.DB_FindDeviceIdBytypeId(addrId, "00000000-00LJ-0000-0000-000000000000")
	if err != nil {
		seelog.Error("MONGO_FindQuan通过设备类型Id找设备Id失败")
	}
	return mongo.FindQuan(device, addrId)
}

func MONGO_CheckToken(token string) (*hbean.BaseResponse, error) {
	return mongo.CheckToken(token)
}
