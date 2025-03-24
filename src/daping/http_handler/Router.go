package httphandler

import (
	c "daping/src/daping/http_handler/http_process/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func router_options(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	origin := request.Header.Get("origin")
	response.Header().Add("Access-Control-Allow-Origin", origin)
	response.Header().Add("Access-Control-Allow-Credentials", "true")
	response.Header().Add("Access-Control-Allow-Methods", "*")
	response.Header().Add("Access-Control-Allow-Headers", "User-Agent,X-Requested-With,Cache-Control,Content-Type,"+
		"Access-Token,Authorization,daping-token")
	response.Header().Add("Access-Control-Expose-Headers", "*")
	response.Header().Add("Content-Type", "application/json")
	response.Write(nil)
}

func newRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = new(notFound)
	router.MethodNotAllowed = new(methodNotAllow)

	router.OPTIONS("/*path", router_options)
	router_DeviceManager(router)//设备分析
	router_ZheXian(router)//折线
	router_index(router)//主页
	return router
}

func router_DeviceManager(router *httprouter.Router) {
	router.GET("/daping/device/:addrId", auth(c.GET_DeviceNum))//设备数量
	router.GET("/daping/quan/:addrId", auth(c.GET_Quan))//圈百分比
	router.GET("/daping/dianliuall/:addrId", auth(c.GET_DianLiu))//电流
	router.GET("/daping/warning/:addrId", auth(c.GET_Waring))//设备告警
	router.GET("/daping/arraydata/:addrId", auth(c.GET_ArrowData))//箭头
	router.GET("/daping/puedata/:addrId", auth(c.GET_FindPUE))//pue数据
}

func router_ZheXian(router *httprouter.Router) {
	router.GET("/daping/allconsum/:addrId", auth(c.GET_AllConsume))
	router.GET("/daping/water/:addrId", auth(c.GET_Water))
}

func router_index(router *httprouter.Router) {
	router.GET("/daping/index", auth(c.GET_AddrName))
}
