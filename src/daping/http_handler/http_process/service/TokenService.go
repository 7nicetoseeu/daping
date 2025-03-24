package service

import (
	"daping/src/daping/db"
	"daping/src/daping/http_handler/bean"
)

func CheckToken(token string) (*bean.BaseResponse, error) {
	return db.MONGO_CheckToken(token)
}
