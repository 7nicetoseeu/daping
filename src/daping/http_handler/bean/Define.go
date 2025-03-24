package bean

//控制报文类型
type PacketType uint8

const (
	RESERVED_0         = PacketType(iota) //保留
	CONNECT                               //客户端请求连接服务端(客户端到服务端)
	CONNACK                               //连接报文确认(服务端到客户端)
	PING                                  //心跳请求(客户端到服务端)
	PONG                                  //心跳响应(服务端到客户端)
	PUSH                                  //实时数据推送(服务端到客户端)
	PUSP                                  //实时数据推送响应(客户端到服务端)
	SUB                                   //订阅消息
	EVENT                                 //8 报警事件推送(服务端到客户端)
	OTHER_CLIENT_ClOSE                    //关闭连接
	LOGIN
	LOGINBACK
	TOKEN
)

const (
	CONNECT_ACCEPTED               = iota //连接已接受
	UN_ACCEPTABLE_PROTOCOL_VERSION        //服务端不支持客户端请求的协议级别
	IDENTIFIER_REJECTED                   //不合格的客户端标识符,客户端标识符是正确的UTF-8编码，但服务端不允许使用
	SERVER_UNAVAILABLE                    //服务端不可用，网络连接已建立，但服务不可用
	BAD_ACCESS_TOKEN                      //无效的授权令牌
	NOT_AUTHORIZED                        //客户端未被授权连接到此服务器
	CONNECT_FAILED                        //登录失败（自定义）
	LOG_BY_OTHER_CLIENT                   //其他设备登录
	UN_ACCEPTABLE_PROTOCOL_TYPE           //服务端不支持客户端请求的协议类型
	NOT_LOGGED_IN
)

const (
	WS_LEVEL_1       = 1
	WS_LEVEL_CURRENT = 1 //websocket 当前版本号
)

const (
	USER_STATE_NORMAL  = 1 //正常
	USER_STATE_DISABLE = 2 //不可用
)
