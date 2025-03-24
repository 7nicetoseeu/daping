package db

import (
	"context"
	"daping/src/config"
	"database/sql"
	"fmt"
	"os"
	"time"

	logs "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DBType int = 0
var DBConn *sql.DB
var Mongodb *mongo.Client
var InfluxMap map[string]influxdb2.Client
var InfluxBucket string

const (
	_ = iota
	TIDB
)
const (
	TIDBSTR = "TIDB"
)

func init() {
	config.GetDBconfig()
	mysqlConfig()
	Mongodb = mongoConfig()
	InfluxConfig()
}

var AddrName = config.GetDBconfig().Params.AddrName

func mysqlConfig() {
	dbc := config.Dcm
	var url string = ""
	var drivername string = ""
	drivername = dbc.Mysql.Driver
	if dbc.DBType == TIDBSTR {
		DBType = TIDB
		url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbc.Mysql.Account, dbc.Mysql.Pwd, dbc.Mysql.Host, dbc.Mysql.Port, dbc.Mysql.DBName)
	}

	dbconn, err := sql.Open(drivername, url)
	if err != nil {
		fmt.Println(AddrName, "mysql数据库连接失败", err.Error())
		os.Exit(-1)
	}
	if err := dbconn.Ping(); err != nil {
		fmt.Println(AddrName, "dbconn ping is err: ", err.Error())
		os.Exit(-1)
	} else {
		logs.Debug(AddrName, "mysql conn success")
		DBConn = dbconn
		DBConn.SetMaxIdleConns(50)
		DBConn.SetMaxOpenConns(512)
	}
}
func mongoConfig() *mongo.Client {
	mgo := config.Dcm.Mongo
	mongoConfig := &options.ClientOptions{}
	mongoConfig.SetMaxConnIdleTime(10 * time.Second)
	mongoConfig.SetMaxPoolSize(512)
	for _, v := range mgo.Host {
		mongoConfig.Hosts = append(mongoConfig.Hosts, fmt.Sprintf("%s:%d", v, mgo.Port))
	}
	if mgo.Auth {
		mongoConfig.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    mgo.Database,
			Username:      mgo.Account,
			Password:      mgo.Pwd,
			PasswordSet:   true,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, mongoConfig)
	if err != nil {
		fmt.Println(AddrName, "mongo db init err : "+err.Error())
		os.Exit(-1)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(AddrName, "db mongo ping err: "+err.Error())
		os.Exit(-1)
	}

	logs.Debug(AddrName, "mongo conn success")

	return client
}
func GetMongoDb() *mongo.Database {
	return Mongodb.Database(config.Dcm.Mongo.Database)
}

func InfluxConfig() {
	InfluxMap = make(map[string]influxdb2.Client)
	dbc := config.Dcm.Influx

	for addrId, host := range dbc.Host {
		url := fmt.Sprintf("http://%s:%d", host, dbc.Port)
		if dbc.Port == 0 {
			url = fmt.Sprintf("http://%s", host)
		}
		token := dbc.Token[addrId]
		Influx := influxdb2.NewClient(url, token)
		_, err := Influx.Ping(context.Background())
		if err != nil {
			fmt.Println(AddrName, "Influxdb conn ping is err: ", err.Error())
			os.Exit(-1)
		}
		InfluxMap[addrId] = Influx
		logs.Debug(addrId + " : influx conn success")
	}
	InfluxBucket = dbc.Bucket
}
