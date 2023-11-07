package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GetNumber 随机id生成
func GetNumber() string {
	orderDate := time.Now().Format("20060102150405")
	orderID := orderDate + fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return orderID
}
