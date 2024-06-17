package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// CalculateMD5 计算给定文件的 MD5 值
func CalculateMD5(file *os.File) (string, error) {
	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hashInBytes := hasher.Sum(nil)[:16]
	hashInString := hex.EncodeToString(hashInBytes)
	return hashInString, nil
}
