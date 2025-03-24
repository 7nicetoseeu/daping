package utils

import "daping/src/daping/db/bean"

func bubbleSort(list [ORDER]*bean.Pdu) [ORDER]*bean.Pdu {
	lenth := len(list)
	for i := 0; i <= lenth; i++ { //循环对比的轮数
		exchange := false
		for j := 1; j < lenth-i; j++ { //当前轮相邻元素循环对比
			if list[j-1] == nil || list[j] == nil {
				continue
			}
			if list[j-1].Value < list[j].Value { //如果前边的大于后边的
				list[j-1], list[j] = list[j], list[j-1] //交换数据
				exchange = true
			}
		}
		if !exchange {
			break
		}
	}
	return list
}

//切片排序
func BubbleSort1(list []*bean.Pdu) []*bean.Pdu {
	lenth := len(list)
	for i := 0; i <= lenth; i++ { //循环对比的轮数
		exchange := false
		for j := 1; j < lenth-i; j++ { //当前轮相邻元素循环对比
			if list[j-1].Value < list[j].Value { //如果前边的大于后边的
				list[j-1], list[j] = list[j], list[j-1] //交换数据
				exchange = true
			}
		}
		if !exchange {
			break
		}
	}
	return list
}
