package ecrypto

import "crypto/md5"

func MD5(key string) string {
	data := md5.Sum([]byte(key))
	return string(data[:])
}

func Sha256(key string) string {
	data := md5.Sum([]byte(key))
	return string(data[:])
}