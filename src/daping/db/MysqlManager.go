package db

import (
	"daping/src/daping/db/bean"
	"daping/src/daping/db/mysql"
)

func MYSQL_FindDeviceNum(addrId string) ([]*bean.Device, error) {
	return mysql.DB_FindAllDevice(addrId)
}
func MYSQL_FindWaring(addrId string) ([]*bean.Warning, error) {
	return mysql.DB_FindWaring(addrId)
}
func MYSQL_FindDeviceId(addrId string) ([]*bean.Device, error) {

	return mysql.DB_FindDeviceIdBytypeId(addrId, "")
}

func MYSQL_FindAttributeNameById(mgrId string, attributeId string, addrId string) string {
	return mysql.DB_FindAttributeNameById(mgrId, attributeId, addrId)
}

func MYSQL_FindIdByTypeId(addrId string) ([]*bean.Device, error) {
	return mysql.DB_FindDeviceIdBytypeId(addrId, "00000000-0PUE-0000-0000-000000000000")
}

func MYSQL_FindAddrName() ([]*bean.AddrName, error) {
	return mysql.DB_FindAddrName()
}

func MYSQL_FindAddrNameById(addrId string) (*bean.AddrName, error) {
	return mysql.DB_FindAddrNameById(addrId)
}
func MYSQL_FindPUE(addrId string) (*bean.PUE, error) {
	return mysql.DB_FindPUE(addrId)
}
