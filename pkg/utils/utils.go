package utils

import (
	"encoding/hex"
	"math/rand"
)

func GenerateAccountId() string {

}

func GenerateSalt() string {

}

func GenerateRandom(n int) (string, error ) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}