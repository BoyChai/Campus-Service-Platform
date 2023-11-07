package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GetNumber 随机id生成
func GetNumber() string {
	orderDate := time.Now().Format("20060102150405")
	orderID := orderDate + fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return orderID
}

func GetUint(value string) uint {
	parseUint, err := strconv.ParseUint(fmt.Sprint(value), 10, 64)
	if err != nil {
		fmt.Println("类型转换错误:", err.Error())
	}
	return uint(parseUint)
}
