package utils

import (
	"fmt"
	"strconv"
)

func stringToint() {
	var a = "10"
	var a1 = 10
	var a3 = "true"
	var a4 = "3.14"
	//字符串解析为int类型
	b, _ := strconv.Atoi(a)
	//int类型转换为字符串
	b1 := strconv.Itoa(a1)
	//字符串解析布尔类型
	b2, _ := strconv.ParseBool(a3)
	//字符串解析浮点类型64位的
	b3, _ := strconv.ParseFloat(a4, 64)
	fmt.Printf("%T,%v", b, b)   //int,10
	fmt.Printf("%T,%v", b1, b1) //string,10
	fmt.Printf("%T,%v", b2, b2) //bool,true
	fmt.Printf("%T,%v", b3, b3) //float64,3.14
}

func GetAddrNum(str string) (string, error) {
	str1 := str[6:9]
	// fmt.Printf("str1: %v\n", len(str))
	b, err := strconv.Atoi(str1)
	if err != nil {
		return "", err
	}
	b1 := strconv.Itoa(b)
	return b1, nil
}
