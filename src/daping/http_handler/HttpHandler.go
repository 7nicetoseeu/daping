package httphandler

import (
	"fmt"
	"net/http"

	logs "github.com/cihub/seelog"
)

func Initservlet(port int) {

	defer func() {
		if err := recover(); err != nil {
			logs.Error("initServlet error:", err)
		}
	}()
	router := newRouter()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	logs.Infof("HTTP SERVER STARTING AND LISTENING AT: %d", port)
	err := server.ListenAndServe()
	if err != nil {
		logs.Error("Listen Server error:" + err.Error())
		panic(err.Error())
	}

}

// func Gin_StaticRun(port int) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			logs.Error("initServlet error:", err)
// 		}
// 	}()
// 	r := gin.Default()
// 	r.Static("/static", "dist/static")
// 	r.LoadHTMLGlob("./dist/index.html")
// 	r.GET("/", func(ctx *gin.Context) {
// 		ctx.HTML(http.StatusOK, "index.html", nil)
// 	})
// 	r.Run(fmt.Sprintf(":%d", port))
// }
