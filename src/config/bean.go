package config

type DBConfig struct {
	DBType    string        `json:"dbType"` //数据库类型
	Mysql     *MysqlConfig  `json:"mysql"`  //mysql数据库
	Mongo     *MongoConfig  `json:"mongo"`  //mongoDB数据库
	Influx    *InfluxConfig `json:"influx"`
	Params    *Params       `json:"parmas"`
	LimitData *LimitData    `json:"limitdata"`
}
type Params struct {
	Names    string `json:"names"`    //设备分析需要查询的设备名
	AddrName string `json:"addrname"` //大屏运行地点
	HttpPort int    `json:httpport`   //http端口号
}
type MysqlConfig struct {
	Driver  string `json:"driver"`  //数据库驱动
	Host    string `json:"host"`    //数据库地址
	Port    int    `json:"port"`    //端口号
	Account string `json:"account"` //账号
	Pwd     string `json:"pwd"`     //密码
	DBName  string `json:"dbname"`  //数据库名
}

type MongoConfig struct {
	Host     []string `json:"host"`     //数据库地址
	Port     int      `json:"port"`     //端口号
	Account  string   `json:"account"`  //账号
	Pwd      string   `json:"pwd"`      //密码
	Database string   `json:"database"` //数据库名
	Auth     bool     `json:"auth"`     //是否需要认证
}

type InfluxConfig1 struct {
	Host  string `json:host`
	Port  int    `json:port`
	Token string `json:token`
}

type InfluxConfig struct {
	Host   map[string]string `json:"host"`
	Port   int64             `json:"port"`
	Token  map[string]string `json:"token"`
	Bucket string            `json:"bucket"`
}
type LimitData struct {
	Coldnum     int `json:"coldnum"`
	Percent100  int `json:"percent100"`
	Maxdianliu  int `json:"maxdianliu"`
	Zeroorone   int `json:"zeroorone"`
	Percent200  int `json:"percent200"`
	Maxpue      int `json:"maxpue"`
	Maxconsume  int `json:"maxconsume"`
	Minconsume  int `json:"minconsume"`
	Maxwaterpre int `json:"maxwaterpre"`
	Minwaterpre int `json:"minwaterpre"`
	Maxwatertem int `json:"maxwatertem"`
	Minwatertem int `json:"minwatertem"`
}
