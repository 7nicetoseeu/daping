package mongo

import (
	"context"
	bbean "daping/src/daping/db/bean"
	"daping/src/daping/http_handler/bean"
	"daping/src/db"
	"time"

	"github.com/cihub/seelog"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckToken(token string) (*bean.BaseResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
		}
	}()
	res := &bean.BaseResponse{
		Code:         200,
		Message:      "",
		ResponseTime: 0,
	}
	col := db.GetMongoDb().Collection("daping_login")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	data := bbean.User{}
	err := col.FindOne(ctx, bson.M{"accessToken": token}).Decode(&data)
	if err != nil {
		seelog.Error("token:", err.Error())
		res.Code = 404
	}
	return res, nil
}
