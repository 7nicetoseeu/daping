package controller

import (
	errCode "daping/src/daping/http_handler/error_code"
	"daping/src/daping/http_handler/http_process/service"
	rw "daping/src/daping/http_handler/http_read_write"
	"net/http"

	"github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
)

/* 直接返回一个切片包括所有楼的信息 */
func GET_AddrName(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	seelog.Infof("---------------------------首页1")
	res, err := service.AddrName()
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if len(res) == 0 {
		rw.WriteErrResp(errCode.DATA_NULL, "楼信息数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}

/* 返回一个楼的信息 */
func GET_OneAddrName(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	seelog.Infof("---------------------------首页2")
	addrId := ps.ByName("addrId")
	res, err := service.OneAddrName(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}
