package util

import (
	"fmt"
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

func RandomPassword() string {
	return RandomString(10)
}

func RandomEmail(username string) string {
	q := []string{"gmail.com", "yahoo.com", "hotmail.com"}
	return username + "@" + q[rand.Intn(len(q))]
}

func RandomPhoneNumber() string {
	// Generate random 3-digit area code
	areaCode := rand.Intn(900) + 100

	// Generate random 3-digit exchange code
	exchangeCode := rand.Intn(900) + 100

	// Generate random 4-digit line number
	lineNumber := rand.Intn(9000) + 1000

	phoneNumber := fmt.Sprintf("(%d) %d-%d", areaCode, exchangeCode, lineNumber)

	return phoneNumber
}
