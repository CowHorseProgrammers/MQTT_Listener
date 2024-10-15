package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetHmacSha256(sessionKey string) string {
	key := []byte(sessionKey)
	message := []byte("")

	// 创建一个新的 HMAC 使用 SHA256 哈希算法
	h := hmac.New(sha256.New, key)

	// 写入消息数据
	h.Write(message)

	// 计算 HMAC 值
	mac := h.Sum(nil)

	// 将结果以十六进制编码并打印
	fmt.Println(hex.EncodeToString(mac))
	return hex.EncodeToString(mac)
}
