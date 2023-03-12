package util

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return strings.Title(RandomString(10))
}

// RandomMoney returns a random amount of money
func RandomMoney() int64 {
	return int64(rand.Intn(1000000))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP"}
	return currencies[rand.Intn(len(currencies))]
}
