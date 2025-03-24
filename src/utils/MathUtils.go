package utils

import (
	"fmt"
	"strconv"
)

//取两位小数
func Decimal2(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}
func Decimal1(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", num), 64)
	return num
}
