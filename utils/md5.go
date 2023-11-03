package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func CalculateMD5Hash(input string) string {
	// 创建MD5哈希对象
	h := md5.New()
	// 将字符串转换为字节数组并计算哈希值
	h.Write([]byte(input))
	// 获取MD5哈希的字节数组
	hashBytes := h.Sum(nil)
	// 将字节数组转换为16进制字符串
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
