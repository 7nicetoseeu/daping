package mysql

import (
	"daping/src/daping/db/bean"
	"daping/src/db"
	"database/sql"
	"fmt"

	"github.com/cihub/seelog"
)

func DB_ZheXianMgrId(addrId string) []*bean.PerAndTemMgr {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	PerAndTemMgr := make([]*bean.PerAndTemMgr, 0)
	str := "SELECT mgr_obj_id,mgr_obj_attribute_id,name FROM water_pre_tem WHERE building =? "
	var mgrid sql.NullString
	var attributeId sql.NullString
	var name sql.NullString
	rows, err := conn.Query(str, addrId)
	if err != nil {
		seelog.Error("DB_ZheXianMgrId err : " + err.Error())
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&mgrid, &attributeId, &name)
		if err != nil {
			return nil
		}
		mgr := &bean.PerAndTemMgr{
			Name:        name.String,
			AttributeId: attributeId.String,
			Id:          mgrid.String,
		}
		PerAndTemMgr = append(PerAndTemMgr, mgr)
	}
	return PerAndTemMgr
}
