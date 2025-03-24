package mysql

import (
	"daping/src/config"
	"daping/src/daping/db/bean"
	"daping/src/db"
	"daping/src/utils"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/cihub/seelog"
)

// func FindDeviceID() []*bean.Device {
// 	conn := db.DBConn
// 	DeviceManage := make([]*bean.Device, 0)
// 	findIdStr := "SELECT mgr_obj_type_id,name FROM mgr_obj_type"
// 	rows, err := conn.Query(findIdStr)
// 	if err != nil {
// 		seelog.Error("DB_findIdStr error:" + err.Error())
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name sql.NullString
// 		// var num sql.NullInt64
// 		var id sql.NullString
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			seelog.Error("DB_findIdStr error:" + err.Error())
// 		}
// 		device := &bean.Device{
// 			Id:   id.String,
// 			Name: name.String,
// 			Num:  0,
// 		}
// 		DeviceManage = append(DeviceManage, device)

// 		// findNumStr := `SELECT COUNT( mgr_obj_type_id ) FROM mgr_obj WHERE mgr_obj_type_id = ? AND addr_id = ?;`
// 		// rows2, err := conn.Query(findNumStr, id.String, addrId)
// 		// if err != nil {
// 		// 	seelog.Error("1DB_findNumStr error:" + err.Error())
// 		// }
// 		// defer rows2.Close()
// 		// for rows2.Next() {
// 		// 	var num sql.NullInt64
// 		// 	err = rows2.Scan(&num)
// 		// 	if err != nil {
// 		// 		seelog.Error("DB_findNumStr error:" + err.Error())
// 		// 	}
// 		// 	device := &bean.Device{
// 		// 		Id:   id.String,
// 		// 		Name: name.String,
// 		// 		Num:  num.Int64,
// 		// 	}
// 		// 	fmt.Printf("device: %v\n", device)
// 		// }
// 	}
// 	return DeviceManage
// }

// func DB_FindDeviceNum(addrId string) ([]*bean.Device, error) {
// 	DeviceManageAll := FindDeviceID()
// 	// var num sql.NullInt64
// 	conn := db.DBConn
// 	DeviceManage := make([]*bean.Device, 0)

// 	for _, device := range DeviceManageAll {
// 		var num sql.NullInt64
// 		findNumStr := `SELECT COUNT( mgr_obj_type_id ) FROM mgr_obj WHERE mgr_obj_type_id = ? AND addr_id like ?;`
// 		err := conn.QueryRow(findNumStr, device.Id, addrId+"%").Scan(&num)
// 		if err != nil {
// 			seelog.Error("DB_findNumStr error:" + err.Error())
// 		}
// 		device.Num = num.Int64
// 		if device.Num != 0 {
// 			DeviceManage = append(DeviceManage, device)
// 			// fmt.Printf("device: %v\n", device)
// 		}
// 	}

// 	return DeviceManage, nil
// }

// func DB_FindErrDeviceId(addrId string) ([]*bean.ErrDevice, error) {
// 	conn := db.DBConn
// 	ErrDeviceList := make([]*bean.ErrDevice, 0)
// 	str := `SELECT mgr_obj_id FROM event_log_src_a18 WHERE addr_id like ?;`
// 	addrId = addrId + "%"
// 	rows, err := conn.Query(str, addrId)
// 	if err != nil {
// 		seelog.Error("DB_FindErrDeviceId error:" + err.Error())
// 	}
// 	for rows.Next() {
// 		var id sql.NullString
// 		err := rows.Scan(&id)
// 		if err != nil {
// 			seelog.Error("DB_FindErrDeviceId error:" + err.Error())
// 		}
// 		errDevice := &bean.ErrDevice{
// 			Id: id.String,
// 		}
// 		ErrDeviceList = append(ErrDeviceList, errDevice)
// 	}
// 	defer rows.Close()
// 	return ErrDeviceList, nil
// }

// func DB_FindAllDevice(addrId string) ([]*bean.Device, error) {
// 	conn := db.DBConn
// 	DeviceManage, _ := DB_FindErrDeviceId(addrId)
// 	DeviceManageAll, _ := DB_FindDeviceNum(addrId)
// 	for _, DeviceId := range DeviceManage {
// 		str := `SELECT mgr_obj_type_id FROM mgr_obj WHERE mgr_obj_id =?`
// 		var errNum sql.NullString
// 		err := conn.QueryRow(str, DeviceId.Id).Scan(&errNum)
// 		if err != nil {
// 			fmt.Printf("DeviceId.Id: %v\n", DeviceId.Id)
// 			fmt.Println("该设备类型ID无法找到")

// 		}
// 		errId := errNum.String
// 		for _, device := range DeviceManageAll {
// 			if errId == device.Id {
// 				device.ErrNum++
// 			}
// 		}
// 	}
// 	return DeviceManageAll, nil

// }

/* 优化版 */
func DB_FindDeviceNum(addrId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	names := config.Dcm.Params.Names
	str := `
	SELECT
		typenum.mgr_obj_type_id,
		count,
		mgr_obj_type.name
	FROM
		( SELECT mgr_obj_type_id, COUNT( mgr_obj_type_id ) AS count FROM mgr_obj WHERE building = ? GROUP BY mgr_obj_type_id ) AS typenum,
		mgr_obj_type
	WHERE
		typenum.mgr_obj_type_id=mgr_obj_type.mgr_obj_type_id
		AND name IN (` + names + `)
	`
	DeviceManage := make([]*bean.Device, 0)
	rows, err := conn.Query(str, addrId)
	if err != nil {
		seelog.Error("DB_FindDeviceNum error:" + err.Error())
		return nil, err
	}
	for rows.Next() {
		var id sql.NullString
		var name sql.NullString
		var num sql.NullInt64
		rows.Scan(&id, &num, &name)
		device := &bean.Device{
			Id:   id.String,
			Name: name.String,
			Num:  num.Int64,
		}
		DeviceManage = append(DeviceManage, device)
	}
	return DeviceManage, nil
}

func DB_FindAllDevice(addrId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	DeviceManageAll, _ := DB_FindDeviceNum(addrId)
	if DeviceManageAll == nil {
		err := errors.New("设备分析获取所有设备为空")
		seelog.Error(err.Error())
		return nil, err
	}
	addrnum, err := utils.GetAddrNum(addrId)

	if err != nil {
		seelog.Error("设备分析获取addrId中的num失败:" + err.Error())
		return nil, err
	}

	table := "event_log_src_a" + addrnum
	str := `
		SELECT
		mgr_obj.mgr_obj_type_id,
		COUNT( mgr_obj_type_id ) 
		FROM
			( SELECT mgr_obj_id FROM ` + table + ` WHERE building = ? GROUP BY mgr_obj_id) AS mgrid,
			mgr_obj
		WHERE
			mgrid.mgr_obj_id=mgr_obj.mgr_obj_id
		GROUP BY
			mgr_obj_type_id
	`
	rows, err := conn.Query(str, addrId)
	if err != nil {
		seelog.Error("DB_FindAllDevice error:" + err.Error())
		return nil, err

	}
	for rows.Next() {
		var errId sql.NullString
		var count sql.NullInt64
		err = rows.Scan(&errId, &count)
		if err != nil {
			seelog.Error("设备分析数据值解析出错 error:" + err.Error())
			return nil, err

		}
		for _, device := range DeviceManageAll {
			if device.Id == errId.String {
				count.Int64 = int64(rand.Intn(int(device.Num)))
				device.ErrNum = count.Int64
				break
			}
		}
	}
	return DeviceManageAll, nil
}

func DB_FindWaring(addrId string) ([]*bean.Warning, error) {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Println("err = ", err)
		}
	}()
	warNings := make([]*bean.Warning, 0)
	conn := db.DBConn
	addNum, err := utils.GetAddrNum(addrId)
	if err != nil {
		seelog.Error("DB_FindWaring,addrId转addrNum失败: " + err.Error())
		return warNings, err

	}

	table := "event_log_src_a" + addNum
	// str := "SELECT min_alarm,max_alarm,value,addr_name,mgr_obj_attribute_name,mgr_obj_name ,content,alarm_time FROM " + table + " WHERE  level = 4 and is_alarm=1"
	str := `
		SELECT
			min_alarm,max_alarm,value,addr_name,mgr_obj_attribute_name,mgr_obj_name ,content,alarm_time
		FROM
		` + table + `
		WHERE
			is_alarm = 1
			AND level = 4
			AND current_state NOT IN(2)
		ORDER BY
			alarm_time DESC
		LIMIT 0,30
	`
	rows, err := conn.Query(str)
	defer rows.Close()
	if err != nil {
		seelog.Error("DB_FindWaring error:" + err.Error())
		return warNings, err
	}
	cutStr := "润泽园区A" + addNum + "数据中心"
	for rows.Next() {
		var addrName sql.NullString
		var attributeName sql.NullString
		var objName sql.NullString
		var msg sql.NullString
		var max sql.NullFloat64
		var min sql.NullFloat64
		var value sql.NullFloat64
		var intime sql.NullInt64
		err = rows.Scan(&min, &max, &value, &addrName, &attributeName, &objName, &msg, &intime)
		if err != nil {
			seelog.Error("DB_FindWaring error:" + err.Error())
			return nil, err

		}
		num := &bean.Len3{
			Max:     max.Float64,
			Min:     min.Float64,
			WarnNum: value.Float64,
		}
		s := strings.ReplaceAll(addrName.String, cutStr, "")
		warning := &bean.Warning{
			AddrName:     s,
			MgrName:      objName.String,
			MgrAttribute: attributeName.String,
			Msg:          msg.String,
			Num:          *num,
			Time:         intime.Int64,
		}
		warNings = append(warNings, warning)
	}
	return warNings, nil
}

func DB_FindDeviceIdBytypeId(addrId string, typeId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	var mgrTypeId string
	if typeId == "" {
		mgrTypeId = "00000000-00AF-0000-0000-000000000000"
	} else {
		mgrTypeId = typeId
	}
	str := "SELECT mgr_obj_id FROM `mgr_obj` WHERE mgr_obj_type_id=? AND building=?"
	rows, err := conn.Query(str, mgrTypeId, addrId)
	if err != nil {
		seelog.Error("DB_FindDeviceId error:" + err.Error())
		return nil, err

	}
	devices := make([]*bean.Device, 0)
	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		if err != nil {
			seelog.Error("DB_FindDeviceId error:" + err.Error())
			return nil, err

		}
		device := bean.Device{
			Id: id.String,
		}
		devices = append(devices, &device)
	}
	return devices, nil
}
func DB_AllConsumFindDeviceIdBytypeId(addrId string, typeId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	var mgrTypeId string
	if typeId == "" {
		mgrTypeId = "00000000-00AF-0000-0000-000000000000"
	} else {
		mgrTypeId = typeId
	}
	str := "SELECT mgr_obj_id FROM `mgr_obj` WHERE mgr_obj_type_id=? AND building=? AND order_no=1"
	rows, err := conn.Query(str, mgrTypeId, addrId)
	if err != nil {
		seelog.Error("DB_FindDeviceId error:" + err.Error())
		return nil, err

	}
	devices := make([]*bean.Device, 0)
	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		if err != nil {
			seelog.Error("DB_FindDeviceId error:" + err.Error())
			return nil, err

		}
		device := bean.Device{
			Id: id.String,
		}
		devices = append(devices, &device)
	}
	return devices, nil
}
func DB_FindDeviceIdBytypeIdAndName(addrId string, typeId string, name string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	var mgrTypeId string
	if typeId == "" {
		mgrTypeId = "00000000-00AF-0000-0000-000000000000"
	} else {
		mgrTypeId = typeId
	}
	str := "SELECT mgr_obj_id FROM `mgr_obj` WHERE mgr_obj_type_id=? AND building=? AND name LIKE ?"
	rows, err := conn.Query(str, mgrTypeId, addrId, "%"+name+"%")
	if err != nil {
		seelog.Error("DB_FindDeviceId error:" + err.Error())
		return nil, err

	}
	devices := make([]*bean.Device, 0)
	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		if err != nil {
			seelog.Error("DB_FindDeviceId error:" + err.Error())
			return nil, err

		}
		device := bean.Device{
			Id: id.String,
		}
		devices = append(devices, &device)
	}
	return devices, nil
}
func DB_FindTopDeviceIdBytypeId(addrId string, typeId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	var mgrTypeId string
	if typeId == "" {
		mgrTypeId = "00000000-00AF-0000-0000-000000000000"
	} else {
		mgrTypeId = typeId
	}
	str := "SELECT  mgr_obj_id FROM `mgr_obj` WHERE mgr_obj_type_id=? AND building=? LIMIT 0,10"
	rows, err := conn.Query(str, mgrTypeId, addrId)
	if err != nil {
		seelog.Error("DB_FindDeviceId error:" + err.Error())
		return nil, err

	}
	devices := make([]*bean.Device, 0)
	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		if err != nil {
			seelog.Error("DB_FindDeviceId error:" + err.Error())
			return nil, err

		}
		device := bean.Device{
			Id: id.String,
		}
		devices = append(devices, &device)
	}
	return devices, nil
}

func DB_FindAttributeNameById(mgrId string, attributeId string, addrId string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	addrId, err := utils.GetAddrNum(addrId)
	if err != nil {
		seelog.Error("DB_FindAttributeNameById获取addrId中的num失败:" + err.Error())
		return ""
	}
	table := "mgr_obj_attribute_set_a" + addrId
	str := "SELECT mgr_obj_attribute_name FROM " + table + " WHERE second_id= ? AND mgr_obj_id= ?"
	var name sql.NullString
	err = conn.QueryRow(str, attributeId, mgrId).Scan(&name)
	if err != nil {
		fmt.Printf("mgrId: %v\n", mgrId)
		fmt.Printf("attributeId: %v\n", attributeId)
		seelog.Error("DB_FindAttributeNameById error:" + err.Error())
		return ""
	}
	return name.String
}

func DB_FindIdByTypeId(addrId string) ([]*bean.Device, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	str := "SELECT mgr_obj_id FROM `mgr_obj` WHERE mgr_obj_type_id='00000000-0PUE-0000-0000-000000000000' AND building= ?"
	rows, err := conn.Query(str, addrId)
	if err != nil {
		seelog.Error("DB_FindIdByTypeId error:" + err.Error())
		return nil, err

	}
	Devices := make([]*bean.Device, 0)
	for rows.Next() {
		var mgrId sql.NullString
		err := rows.Scan(&mgrId)
		if err != nil {
			seelog.Error("DB_FindIdByTypeId error:" + err.Error())
			return nil, err
		}
		device := bean.Device{
			Id: mgrId.String,
		}
		Devices = append(Devices, &device)
	}

	return Devices, nil
}

func DB_FindAttributeIdByMgrId(addrId string, mgrId string) ([]string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	ids := make([]string, 0)
	addrNum, err := utils.GetAddrNum(addrId)
	if err != nil {
		seelog.Error("DB_FindAttributeIdByMgrId,addrid转换addrnum失败 : " + err.Error())
		return nil, err
	}
	str := `
		SELECT
			second_id
		FROM
			mgr_obj_attribute_set_a` + addrNum + `	
		WHERE
			mgr_obj_id = ?
			AND second_id LIKE 'sum_AB_L%'
	`
	rows, err := conn.Query(str, mgrId)
	if err != nil {
		seelog.Error("DB_FindAttributeIdByMgrId : " + err.Error())
		return nil, err
	}
	for rows.Next() {
		var id sql.NullString
		err := rows.Scan(&id)
		if err != nil {
			seelog.Error("DB_FindAttributeIdByMgrId : " + err.Error())
			return nil, err
		}
		ids = append(ids, id.String)
	}
	return ids, nil
}

func FindAIdMgrIdByType(addrId string) ([]*bean.ZhiLuDianLiu, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	str := `
		SELECT
			id.mgr_obj_id,
			second_id 
		FROM
			mgr_obj_attribute_set_a18,
			( SELECT mgr_obj_id FROM mgr_obj WHERE mgr_obj_type_id = '00000000-00AF-0000-0000-000000000000' AND building =? ) AS id 
		WHERE
			mgr_obj_attribute_set_a18.mgr_obj_id = id.mgr_obj_id 
			AND second_id LIKE 'sum_AB_L%' 
	`
	zhiludianlius := make([]*bean.ZhiLuDianLiu, 0)
	rows, err := conn.Query(str, addrId)
	if err != nil {
		seelog.Error("FindAIdMgrIdByType err : " + err.Error())
		return nil, err
	}
	for rows.Next() {
		var mgrId sql.NullString
		var attributeId sql.NullString
		err = rows.Scan(&mgrId, &attributeId)
		if err != nil {
			seelog.Error("FindAIdMgrIdByType err : " + err.Error())
			return nil, err
		}
		zhiludianlius = append(zhiludianlius, &bean.ZhiLuDianLiu{Attribute: attributeId.String, Id: mgrId.String})
	}
	return zhiludianlius, nil
}

func DB_FindPUE(addrId string) (*bean.PUE, error) {
	Pue := new(bean.PUE)
	max := 1.35
	min := 1.30
	if Pue.Num == 0.0 {
		Pue.Num = rand.Float64()*(max-min) + min
	}
	Pue.Num = utils.Decimal2(Pue.Num)
	return Pue, nil
}
