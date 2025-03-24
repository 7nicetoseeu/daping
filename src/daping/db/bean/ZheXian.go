package bean

import (
	"time"
)

type AllConsum struct {
	Time time.Time
	Num  float64
}
type AllConsumFour struct {
	AllConsum1 []*AllConsum `json:"allConsum1"`
	AllConsum2 []*AllConsum `json:"allConsum2"`
	AllConsum3 []*AllConsum `json:"allConsum3"`
	AllConsum4 []*AllConsum `json:"allConsum4"`
}

type AllWaterPressure struct {
	GiveWater1 []*WaterPressure
	BackWater1 []*WaterPressure
	GiveWater2 []*WaterPressure
	BackWater2 []*WaterPressure
}
type WaterPressure struct {
	Name string
	Time time.Time
	Num  float64
}

type AllWaterTem struct {
	GiveWater1 []*WaterTem
	BackWater1 []*WaterTem
	GiveWater2 []*WaterTem
	BackWater2 []*WaterTem
}

type WaterTem struct {
	Name string
	Time time.Time
	Num  float64
}

type PerAndTemMgr struct {
	Id          string
	AttributeId string
	Name        string
}

type Water struct {
	Name string
	Time time.Time
	Num  float64
}
type AllWater struct {
	GiveWaterPre1 []*Water
	BackWaterPre1 []*Water
	GiveWaterPre2 []*Water
	BackWaterPre2 []*Water
	GiveWaterTem1 []*Water
	BackWaterTem1 []*Water
	GiveWaterTem2 []*Water
	BackWaterTem2 []*Water
}

type ArrowData struct {
	Time int
	Num  float64
}

type ArrowDatas struct {
	ElecHuan    int     `json:"elechuan"`
	ColdHuan    int     `json:"coldhuan"`
	ElecTong    int     `json:"electong"`
	ColdTong    int     `json:"coldtong"`
	ElecHuanNum float64 `json:"elechuannum"`
	ColdHuanNum float64 `json:"coldhuannum"`
	ElecTongNum float64 `json:"electongnum"`
	ColdTongNum float64 `json:"coldtongnum"`
	IsEmpty     int16   `json:"isEmpty"`
}
