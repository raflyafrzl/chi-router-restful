package utils

import (
	"math/rand"
	"time"
)

var alphanumeric string = "123abcefghijklPQRSTUVWXYmnopqrstu456789vwxyzABCDEFGHIJKLMNOZ"

func GetRandomID() string {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	var randomstring string = ""
	var temp int
	for i := 1; i <= 8; i++ {
		temp = r.Intn(len(alphanumeric)-1) + 1
		randomstring = randomstring + string(alphanumeric[temp])
	}

	return randomstring
}
