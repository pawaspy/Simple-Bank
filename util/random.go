package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcedifghijklmnopqrstuvwxyz"

func RandomInt(max, min int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() int {
	return rand.Intn(1000)
}

func RandomCurrency() string {
	curr := []string{"EUR", "INR", "YEN"}
	n := len(curr)
	return curr[rand.Intn(n)]
}
