package config

import (
	"daping/src/utils"
	"flag"
)

var Dcm *DBConfig

var filePath string

func GetDBconfig() *DBConfig {
	if Dcm == nil {
		Dcm = &DBConfig{
			Mysql:  new(MysqlConfig),
			Mongo:  new(MongoConfig),
			Influx: new(InfluxConfig),
			Params: new(Params),
		}
	}
	utils.ReadJsonFile(filePath, Dcm)
	return Dcm
}
func init() {
	path := flag.String("path", "pdev.json", "setting file path")
	flag.Parse()
	filePath = *path
}
