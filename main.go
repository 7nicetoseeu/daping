package main

import (
	"daping/src/config"
	httphandler "daping/src/daping/http_handler"
	"math/rand"
	"time"

	"fmt"

	logs "github.com/cihub/seelog"
)

//	httphandler "daping/src/daping/http_handler"
func main() {
	rand.Seed(time.Now().Unix())
	defer logs.Flush()
	logger, err := logs.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		fmt.Println("parse config.xml error: " + err.Error())
	}
	logs.ReplaceLogger(logger)
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()
	httphandler.Initservlet(config.Dcm.Params.HttpPort)

}
