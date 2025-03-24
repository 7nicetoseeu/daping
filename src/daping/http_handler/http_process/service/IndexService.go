package service

import (
	"daping/src/daping/db"
	"daping/src/daping/db/bean"
)

func AddrName() ([]*bean.AddrName, error) {
	return db.MYSQL_FindAddrName()
}

func OneAddrName(addrId string) (*bean.AddrName, error) {
	return db.MYSQL_FindAddrNameById(addrId)
}
