package utils

import (
	"math/rand"
	"strings"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"
const COLOR_FFFFFF = 16777215

func RandomString(n int) string {
	var builder strings.Builder
	k := len(ALPHABET)
	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomColorInt() int32 {
	return int32(RandomInt(0, COLOR_FFFFFF))
}
