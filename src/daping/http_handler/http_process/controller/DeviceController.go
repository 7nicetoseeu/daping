package controller

import (
	errCode "daping/src/daping/http_handler/error_code"
	"daping/src/daping/http_handler/http_process/service"
	rw "daping/src/daping/http_handler/http_read_write"
	"net/http"
	"time"

	"github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
)

func GET_DeviceNum(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	// fmt.Printf("addrId: %v\n", addrId)
	seelog.Infof("---------------------------设备分析")
	res, err := service.DeviceNum(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if len(res) == 0 {
		rw.WriteErrResp(errCode.DATA_NULL, "设备分析数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)
}

func GET_DianLiu(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------所有设备支路电流")
	res, err := service.Dianliu(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if len(res) == 0 {
		rw.WriteErrResp(errCode.DATA_NULL, "支路电流数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)

}

// func GET_Dianliu10(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
// 	addrId := ps.ByName("addrId")
// 	seelog.Infof("---------------------------10个设备支路电流")
// 	res, err := service.Dianliu10(addrId)
// 	if err != nil {
// 		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
// 		return
// 	}
// 	if len(res) == 0 {
// 		rw.WriteErrResp(errCode.DATA_NULL, "数据为空", response)
// 		return
// 	}
// 	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)

// }

func GET_Quan(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	time := time.Now().Format(" 2006年1月2日15:04:05 ")
	seelog.Infof("---------------------------" + time)
	seelog.Infof("---------------------------左右圈数据")
	res, err := service.Quan(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}

	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)

}
func GET_Waring(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------设备告警")
	res, err := service.Warning(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if len(res) == 0 {
		rw.WriteErrResp(errCode.DATA_NULL, "设备告警数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)

}
func GET_FindPUE(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	addrId := ps.ByName("addrId")
	seelog.Infof("---------------------------PUE查询")
	res, err := service.FindPUE(addrId)
	if err != nil {
		rw.WriteErrResp(errCode.DATA_ERR, "失败", response)
		return
	}
	if res.Num == 0.0 {
		rw.WriteErrResp(errCode.DATA_NULL, "PUE查询数据为空", response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS, "成功", res, response)

}
