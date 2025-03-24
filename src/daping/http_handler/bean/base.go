package bean

type BaseRequest struct {
	Signature   string `json:"signature"`   //用户签名标记
	RequestTime int64  `json:"requestTime"` //请求时间
}

type BaseResponse struct {
	Code         int    `json:"code"`               //返回码
	Message      string `json:"message,omitempty"`  //错误信息
	ResponseTime int64  `json:"respTime,omitempty"` //响应时间
}

//响应参数
type DataResult struct {
	*BaseResponse
	Result interface{} `json:"result"`
}

type WSHeader struct {
	WSPacketType PacketType `json:"packetType"` //报文类型
	Level        uint8      `json:"level"`      //协议级别
}

type WSBase struct {
	*WSHeader
	Protocol interface{} `json:"protocol"` //协议内容
}
type Protocol interface {
}

type WSLoginBack struct {
	Code    int32  `json:"code"`
	CodeMes string `json:"codeMes"`
}
type WSToken struct {
	Token string `json:"token"`
}

func NewWSMessage(packetType PacketType, level uint8, value interface{}) *WSBase {
	return &WSBase{
		WSHeader: &WSHeader{
			WSPacketType: packetType,
			Level:        level,
		},
		Protocol: value,
	}
}
