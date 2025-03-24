package controller

import (
	errCode "daping/src/daping/http_handler/error_code"
	"daping/src/daping/http_handler/http_process/service"
	rw "daping/src/daping/http_handler/http_read_write"
	"net/http"

	"github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
)

func GET_AllConsume(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------总耗能")
	res, err := service.AllConsumeFour(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	// if len(res) == 0 {
	// 	rw.WriteErrResp(errCode.DATA_NULL, "总耗能数据为空", response)
	// 	return
	// }
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}
func GET_Water(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------水压和水温")
	res, err := service.Water(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if res == nil {
		rw.WriteErrResp(errCode.DATA_NULL, "供回水温，供回水压数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}

func GET_ArrowData(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------同比环比")
	res, err := service.ArrowData(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if res == nil {
		rw.WriteErrResp(errCode.DATA_NULL, "同比环比数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}
