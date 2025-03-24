package rw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errCode "daping/src/daping/http_handler/error_code"

	logs "github.com/cihub/seelog"
)

func ReadBody(request *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Println(" ioutil.ReadAll is error........")
		return err
	}
	defer request.Body.Close()
	fmt.Println(string(data))
	if data == nil {
		fmt.Println("data is nil........................")
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		logs.Error(err)
		return errCode.New(errCode.JSON_UNMARSHAL_ERROR, "JSON 解析错误")
	}
	return nil
}
