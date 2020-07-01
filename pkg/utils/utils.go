package utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
)

func GenerateAccountId() string {
	return ""
}

func GenerateSalt() string {
	return ""
}

func GenerateRandom(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "6b0d767330ea1b312e65a4965db85ae1"
	}
	return hex.EncodeToString(bytes)
}

func MD5(key string) string {
	data := md5.Sum([]byte(key))
	return string(data[:])
}

func Sha512(key string) string {
	data := sha512.Sum512([]byte(key))
	return string(data[:])
}
