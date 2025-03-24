package mysql

import (
	"daping/src/daping/db/bean"
	"daping/src/db"
	"database/sql"
	"fmt"

	"github.com/cihub/seelog"
)

func DB_FindAddrName() ([]*bean.AddrName, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	str := `
	SELECT name,addr_id FROM address WHERE addr_type=1
	`
	addrNames := make([]*bean.AddrName, 0)
	rows, err := conn.Query(str)
	if err != nil {
		seelog.Error("DB_FindrAllAddName err:" + err.Error())
		return nil, err

	}
	// 	k8s:
	// http://10.5.1.200:30258/
	// account:	
	// password:Admin@1233
	for rows.Next() {
		var name sql.NullString
		var addrId sql.NullString
		err = rows.Scan(&name, &addrId)
		if err != nil {
			seelog.Error("DB_FindrAllAddName err:" + err.Error())
			return nil, err
		}
		addrName := &bean.AddrName{
			Name:   name.String,
			AddrId: addrId.String,
		}
		addrNames = append(addrNames, addrName)
	}
	return addrNames, nil
}

func DB_FindAddrNameById(addrId string) (*bean.AddrName, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	conn := db.DBConn
	str := `
		SELECT
			name
		FROM
			address 
		WHERE
			addr_type = 1 
			AND addr_id =?
	`
	var name sql.NullString
	err := conn.QueryRow(str, addrId).Scan(&name)
	if err != nil {
		seelog.Error("DB_FindAddrNameById err:" + err.Error())
		return nil, err
	}
	addrName := &bean.AddrName{
		Name:   name.String,
		AddrId: addrId,
	}
	return addrName, nil
}
