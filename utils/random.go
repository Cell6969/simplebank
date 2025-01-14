package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
*
For Generate random int
@min int64
@max int64

return int64
*/
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

/*
*
For generate random string
@n int

return string
*/
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

/*
*
For generate name

return string
*/
func RandomName() string {
	return RandomString(6)
}

/*
*
For generate money or balance

return int64
*/
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
