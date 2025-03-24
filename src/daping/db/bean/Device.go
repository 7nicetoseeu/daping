package bean

type DeviceManage struct {
	Device []Device
}

type Device struct {
	Id     string
	Name   string
	Num    int64
	ErrNum int64
}

type Quan struct {
	ColdOpenNum  int64
	ColdCloseNum int64
	ElecPer      int64
	ColdPer      int64
	IsEmpty      int16
}

type ZhiLuDianLiuOrder struct {
	num1  ZhiLuDianLiu
	num2  ZhiLuDianLiu
	num3  ZhiLuDianLiu
	num4  ZhiLuDianLiu
	num5  ZhiLuDianLiu
	num6  ZhiLuDianLiu
	num7  ZhiLuDianLiu
	num8  ZhiLuDianLiu
	num9  ZhiLuDianLiu
	num10 ZhiLuDianLiu
}

type ZhiLuDianLiu struct {
	Name      string
	Num       float64
	Attribute string
	Id        string
}

type Warning struct {
	AddrName     string
	MgrName      string
	MgrAttribute string
	Msg          string
	Time         int64
	Num          Len3
}

type Len3 struct {
	Max     float64
	Min     float64
	WarnNum float64
}

type ErrDevice struct {
	Id string
}

type Mgr struct {
	Monitor    map[string]Pdu `bson:"monitor"`
	MgrObjName string         `bson:"mgrObjName"`
	MgrObjId   string         `bson:"mgrObjId"`
}

type Monitor struct {
	Pdu1_1  Pdu `bson:"PDU_1I01"`
	Pdu1_2  Pdu `bson:"PDU_1I02"`
	Pdu1_3  Pdu `bson:"PDU_1I03"`
	Pdu1_4  Pdu `bson:"PDU_1I04"`
	Pdu1_5  Pdu `bson:"PDU_1I05"`
	Pdu1_6  Pdu `bson:"PDU_1I06"`
	Pdu1_7  Pdu `bson:"PDU_1I07"`
	Pdu1_8  Pdu `bson:"PDU_1I08"`
	Pdu1_9  Pdu `bson:"PDU_1I09"`
	Pdu1_10 Pdu `bson:"PDU_1I10"`
	Pdu1_11 Pdu `bson:"PDU_1I11"`
	Pdu1_12 Pdu `bson:"PDU_1I12"`
	Pdu1_13 Pdu `bson:"PDU_1I13"`
	Pdu1_14 Pdu `bson:"PDU_1I14"`
	Pdu1_15 Pdu `bson:"PDU_1I15"`
	Pdu1_16 Pdu `bson:"PDU_1I16"`
	Pdu1_17 Pdu `bson:"PDU_1I17"`
	Pdu1_18 Pdu `bson:"PDU_1I18"`
	Pdu1_19 Pdu `bson:"PDU_1I19"`
	Pdu1_20 Pdu `bson:"PDU_1I20"`
	Pdu1_21 Pdu `bson:"PDU_1I21"`
	Pdu1_22 Pdu `bson:"PDU_1I22"`
	Pdu1_23 Pdu `bson:"PDU_1I23"`
	Pdu1_24 Pdu `bson:"PDU_1I24"`
	Pdu1_25 Pdu `bson:"PDU_1I25"`
	Pdu1_26 Pdu `bson:"PDU_1I26"`
	Pdu1_27 Pdu `bson:"PDU_1I27"`
	Pdu1_28 Pdu `bson:"PDU_1I28"`
	Pdu1_29 Pdu `bson:"PDU_1I29"`
	Pdu1_30 Pdu `bson:"PDU_1I30"`
	Pdu1_31 Pdu `bson:"PDU_1I31"`
	Pdu1_32 Pdu `bson:"PDU_1I32"`
	Pdu1_33 Pdu `bson:"PDU_1I33"`
	Pdu1_34 Pdu `bson:"PDU_1I34"`
	Pdu1_35 Pdu `bson:"PDU_1I35"`
	Pdu1_36 Pdu `bson:"PDU_1I36"`
	Pdu1_37 Pdu `bson:"PDU_1I37"`
	Pdu1_38 Pdu `bson:"PDU_1I38"`
	Pdu1_39 Pdu `bson:"PDU_1I39"`
	Pdu2_1  Pdu `bson:"PDU_2I01"`
	Pdu2_2  Pdu `bson:"PDU_2I02"`
	Pdu2_3  Pdu `bson:"PDU_2I03"`
	Pdu2_4  Pdu `bson:"PDU_2I04"`
	Pdu2_5  Pdu `bson:"PDU_2I05"`
	Pdu2_6  Pdu `bson:"PDU_2I06"`
	Pdu2_7  Pdu `bson:"PDU_2I07"`
	Pdu2_8  Pdu `bson:"PDU_2I08"`
	Pdu2_9  Pdu `bson:"PDU_2I09"`
	Pdu2_10 Pdu `bson:"PDU_2I10"`
	Pdu2_11 Pdu `bson:"PDU_2I11"`
	Pdu2_12 Pdu `bson:"PDU_2I12"`
	Pdu2_13 Pdu `bson:"PDU_2I13"`
	Pdu2_14 Pdu `bson:"PDU_2I14"`
	Pdu2_15 Pdu `bson:"PDU_2I15"`
	Pdu2_16 Pdu `bson:"PDU_2I16"`
	Pdu2_17 Pdu `bson:"PDU_2I17"`
	Pdu2_18 Pdu `bson:"PDU_2I18"`
	Pdu2_19 Pdu `bson:"PDU_2I19"`
	Pdu2_20 Pdu `bson:"PDU_2I20"`
	Pdu2_21 Pdu `bson:"PDU_2I21"`
	Pdu2_22 Pdu `bson:"PDU_2I22"`
	Pdu2_23 Pdu `bson:"PDU_2I23"`
	Pdu2_24 Pdu `bson:"PDU_2I24"`
	Pdu2_25 Pdu `bson:"PDU_2I25"`
	Pdu2_26 Pdu `bson:"PDU_2I26"`
	Pdu2_27 Pdu `bson:"PDU_2I27"`
	Pdu2_28 Pdu `bson:"PDU_2I28"`
	Pdu2_29 Pdu `bson:"PDU_2I29"`
	Pdu2_30 Pdu `bson:"PDU_2I30"`
	Pdu2_31 Pdu `bson:"PDU_2I31"`
	Pdu2_32 Pdu `bson:"PDU_2I32"`
	Pdu2_33 Pdu `bson:"PDU_2I33"`
	Pdu2_34 Pdu `bson:"PDU_2I34"`
	Pdu2_35 Pdu `bson:"PDU_2I35"`
	Pdu2_36 Pdu `bson:"PDU_2I36"`
	Pdu2_37 Pdu `bson:"PDU_2I37"`
	Pdu2_38 Pdu `bson:"PDU_2I38"`
	Pdu2_39 Pdu `bson:"PDU_2I39"`
}
type Pdu struct {
	MgrObjName string
	MgrObjId   string

	Value    float64 `bson:"value"`
	SecondId string  `bson:"secondId"`
}

type MgrC struct {
	Monitor MonitorC `bson:"monitor"`
}
type MonitorC struct {
	Attribute     Attribute `bson:"BA_Status_DDCLJ_Run"`
	ElecAttribute Attribute `bson:"ZB_Active_P"`
	ColdAttribute Attribute `bson:"BA_Param_LJ_DLB"`
}

type Attribute struct {
	Value float64 `bson:"value"`
}

type Devicem struct {
	MgrObjName string         `bson:"mgrObjName"`
	MgrObjId   string         `bson:"mgrObjId"`
	Monitor    map[string]Pdu `bson:"monitor"`
}

// type Monitorm struct {
// 	Pdu1_1 Pdu `bson:"PDU_1I01"`
// }
type PUE struct {
	Num float64 `json:"num"`
}
