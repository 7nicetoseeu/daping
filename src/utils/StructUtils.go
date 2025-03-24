package utils

import (
	"daping/src/daping/db/bean"
	"reflect"

	"github.com/cihub/seelog"
)

const (
	ORDER = 50
)

func Max50Dianliu(pdus []*bean.Pdu) [ORDER]*bean.Pdu {
	AllPdus := pdus
	// OrderPdu := make([]*bean.Pdu, 10)
	var OrderPdu [ORDER]*bean.Pdu

	// fmt.Printf("len(OrderPdu): %v\n", len(OrderPdu))
	index := 0
	if len(AllPdus) <= 50 {
		for _, pdu := range AllPdus {
			OrderPdu[index] = pdu
			index++
		}
		OrderPdu = bubbleSort(OrderPdu)
		return OrderPdu
	}
	for _, pdu := range AllPdus {

		if index < ORDER {
			OrderPdu[index] = pdu
			index++
			continue
		} else if index == ORDER {
			OrderPdu = bubbleSort(OrderPdu)
		}
		if len(OrderPdu) == 0 {
			panic("前50条电流为空")
		}
		if pdu == nil {
			seelog.Error("pdu=nil")
			continue
		}
		tempdu := pdu
		for index, opdu := range OrderPdu {
			// fmt.Printf("opdu.Value: %v\n", opdu.Value)
			if tempdu.Value >= opdu.Value {
				//交换数据
				// fmt.Println("交换数据")
				OrderPdu[index] = tempdu
				break
			}
		}
	}
	return bubbleSort(OrderPdu)
}

func ForeachStruct(obj interface{}, name string, id string) []*bean.Pdu {
	t := reflect.TypeOf(obj) // 注意，obj不能为指针类型，否则会报：panic recovered: reflect: NumField of non-struct type
	v := reflect.ValueOf(obj)
	pdus := make([]*bean.Pdu, 0)
	for k := 0; k < t.NumField(); k++ {
		// fmt.Printf("%s -- %T \n", t.Field(k).Tag, v.Field(k).Interface())
		tempdu := v.Field(k).Interface()
		pdu := tempdu.(bean.Pdu)
		pdu.MgrObjName = name
		pdu.MgrObjId = id
		//11231
		if pdu.Value != 0 && pdu.Value > 10 {
			pdus = append(pdus, &pdu)
		}
	}
	return pdus
}

func IsIdIn(id string) bool {
	var ids = []string{"PDU_1I01", "PDU_1I01", "PDU_1I02", "PDU_1I03", "PDU_1I04", "PDU_1I05", "PDU_1I06", "PDU_1I07", "PDU_1I08", "PDU_1I09", "PDU_1I10", "PDU_1I11", "PDU_1I12", "PDU_1I13", "PDU_1I14", "PDU_1I15", "PDU_1I16", "PDU_1I17", "PDU_1I18", "PDU_1I19", "PDU_1I20", "PDU_1I21", "PDU_1I22", "PDU_1I23", "PDU_1I24", "PDU_1I25", "PDU_1I26", "PDU_1I27", "PDU_1I28", "PDU_1I29", "PDU_1I30", "PDU_1I31", "PDU_1I32", "PDU_1I33", "PDU_1I34", "PDU_1I35", "PDU_1I36", "PDU_1I37", "PDU_1I38", "PDU_1I39", "PDU_2I01", "PDU_2I02", "PDU_2I03", "PDU_2I04", "PDU_2I05", "PDU_2I06", "PDU_2I07", "PDU_2I08", "PDU_2I09", "PDU_2I10", "PDU_2I11", "PDU_2I12", "PDU_2I13", "PDU_2I14", "PDU_2I15", "PDU_2I16", "PDU_2I17", "PDU_2I18", "PDU_2I19", "PDU_2I20", "PDU_2I21", "PDU_2I22", "PDU_2I23", "PDU_2I24", "PDU_2I25", "PDU_2I26", "PDU_2I27", "PDU_2I28", "PDU_2I29", "PDU_2I30", "PDU_2I31", "PDU_2I32", "PDU_2I33", "PDU_2I34", "PDU_2I35", "PDU_2I36", "PDU_2I37", "PDU_2I38", "PDU_2I39"}
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}

// type Monitor struct {
// 	Pdu1_1  Pdu `bson:"PDU_1I01"`
// 	Pdu1_2  Pdu `bson:"PDU_1I02"`
// 	Pdu1_3  Pdu `bson:"PDU_1I03"`
// 	Pdu1_4  Pdu `bson:"PDU_1I04"`
// 	Pdu1_5  Pdu `bson:"PDU_1I05"`
// 	Pdu1_6  Pdu `bson:"PDU_1I06"`
// 	Pdu1_7  Pdu `bson:"PDU_1I07"`
// 	Pdu1_8  Pdu `bson:"PDU_1I08"`
// 	Pdu1_9  Pdu `bson:"PDU_1I09"`
// 	Pdu1_10 Pdu `bson:"PDU_1I10"`
// 	Pdu1_11 Pdu `bson:"PDU_1I11"`
// 	Pdu1_12 Pdu `bson:"PDU_1I12"`
// 	Pdu1_13 Pdu `bson:"PDU_1I13"`
// 	Pdu1_14 Pdu `bson:"PDU_1I14"`
// 	Pdu1_15 Pdu `bson:"PDU_1I15"`
// 	Pdu1_16 Pdu `bson:"PDU_1I16"`
// 	Pdu1_17 Pdu `bson:"PDU_1I17"`
// 	Pdu1_18 Pdu `bson:"PDU_1I18"`
// 	Pdu1_19 Pdu `bson:"PDU_1I19"`
// 	Pdu1_20 Pdu `bson:"PDU_1I20"`
// 	Pdu1_21 Pdu `bson:"PDU_1I21"`
// 	Pdu1_22 Pdu `bson:"PDU_1I22"`
// 	Pdu1_23 Pdu `bson:"PDU_1I23"`
// 	Pdu1_24 Pdu `bson:"PDU_1I24"`
// 	Pdu1_25 Pdu `bson:"PDU_1I25"`
// 	Pdu1_26 Pdu `bson:"PDU_1I26"`
// 	Pdu1_27 Pdu `bson:"PDU_1I27"`
// 	Pdu1_28 Pdu `bson:"PDU_1I28"`
// 	Pdu1_29 Pdu `bson:"PDU_1I29"`
// 	Pdu1_30 Pdu `bson:"PDU_1I30"`
// 	Pdu1_31 Pdu `bson:"PDU_1I31"`
// 	Pdu1_32 Pdu `bson:"PDU_1I32"`
// 	Pdu1_33 Pdu `bson:"PDU_1I33"`
// 	Pdu1_34 Pdu `bson:"PDU_1I34"`
// 	Pdu1_35 Pdu `bson:"PDU_1I35"`
// 	Pdu1_36 Pdu `bson:"PDU_1I36"`
// 	Pdu1_37 Pdu `bson:"PDU_1I37"`
// 	Pdu1_38 Pdu `bson:"PDU_1I38"`
// 	Pdu1_39 Pdu `bson:"PDU_1I39"`
// 	Pdu2_1  Pdu `bson:"PDU_2I01"`
// 	Pdu2_2  Pdu `bson:"PDU_2I02"`
// 	Pdu2_3  Pdu `bson:"PDU_2I03"`
// 	Pdu2_4  Pdu `bson:"PDU_2I04"`
// 	Pdu2_5  Pdu `bson:"PDU_2I05"`
// 	Pdu2_6  Pdu `bson:"PDU_2I06"`
// 	Pdu2_7  Pdu `bson:"PDU_2I07"`
// 	Pdu2_8  Pdu `bson:"PDU_2I08"`
// 	Pdu2_9  Pdu `bson:"PDU_2I09"`
// 	Pdu2_10 Pdu `bson:"PDU_2I10"`
// 	Pdu2_11 Pdu `bson:"PDU_2I11"`
// 	Pdu2_12 Pdu `bson:"PDU_2I12"`
// 	Pdu2_13 Pdu `bson:"PDU_2I13"`
// 	Pdu2_14 Pdu `bson:"PDU_2I14"`
// 	Pdu2_15 Pdu `bson:"PDU_2I15"`
// 	Pdu2_16 Pdu `bson:"PDU_2I16"`
// 	Pdu2_17 Pdu `bson:"PDU_2I17"`
// 	Pdu2_18 Pdu `bson:"PDU_2I18"`
// 	Pdu2_19 Pdu `bson:"PDU_2I19"`
// 	Pdu2_20 Pdu `bson:"PDU_2I20"`
// 	Pdu2_21 Pdu `bson:"PDU_2I21"`
// 	Pdu2_22 Pdu `bson:"PDU_2I22"`
// 	Pdu2_23 Pdu `bson:"PDU_2I23"`
// 	Pdu2_24 Pdu `bson:"PDU_2I24"`
// 	Pdu2_25 Pdu `bson:"PDU_2I25"`
// 	Pdu2_26 Pdu `bson:"PDU_2I26"`
// 	Pdu2_27 Pdu `bson:"PDU_2I27"`
// 	Pdu2_28 Pdu `bson:"PDU_2I28"`
// 	Pdu2_29 Pdu `bson:"PDU_2I29"`
// 	Pdu2_30 Pdu `bson:"PDU_2I30"`
// 	Pdu2_31 Pdu `bson:"PDU_2I31"`
// 	Pdu2_32 Pdu `bson:"PDU_2I32"`
// 	Pdu2_33 Pdu `bson:"PDU_2I33"`
// 	Pdu2_34 Pdu `bson:"PDU_2I34"`
// 	Pdu2_35 Pdu `bson:"PDU_2I35"`
// 	Pdu2_36 Pdu `bson:"PDU_2I36"`
// 	Pdu2_37 Pdu `bson:"PDU_2I37"`
// 	Pdu2_38 Pdu `bson:"PDU_2I38"`
// 	Pdu2_39 Pdu `bson:"PDU_2I39"`
// }
