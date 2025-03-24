package httphandler

import (
	"daping/src/daping/http_handler/common"
	errCode "daping/src/daping/http_handler/error_code"
	"daping/src/daping/http_handler/http_process/service"
	rw "daping/src/daping/http_handler/http_read_write"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get(common.HeaderToken)
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		res, err := service.CheckToken(token)
		if err != nil {
			rw.WriteErrResp(errCode.SERVICE_BUSSY, "服务繁忙", w)
			return
		}
		if res.Code != 200 {
			rw.WriteErrResp(errCode.LOGIN_USER_UN_EXTIST, "token验证失败", w)
			return
		}
		h(w, r, ps)
		return
	}

}
