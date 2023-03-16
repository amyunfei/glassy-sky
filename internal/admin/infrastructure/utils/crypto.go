package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(str string, salt string, iteration int) string {
	strByte := []byte(str)
	saltByte := []byte(salt)
	hash := md5.New()
	hash.Write(saltByte)
	hash.Write(strByte)
	result := hash.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		hash.Reset()
		hash.Write(result)
		result = hash.Sum(nil)
	}
	return hex.EncodeToString(result)
}
