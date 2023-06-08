package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算给定字节切片的MD5哈希值，并将结果以十六进制字符串的形式返回
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
