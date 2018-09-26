package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	var charset string
	charset = Words.RuneCharInKey
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
