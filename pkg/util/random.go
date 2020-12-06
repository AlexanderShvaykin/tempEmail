package util

import (
	"math/rand"
)

var (
	tails   = [3]string{"com", "net", "org"}
	letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

func RandomString(n int, seed int64) string {
	rand.Seed(seed)
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomTail(seed int64) string {
	rand.Seed(seed)
	return tails[rand.Intn(len(tails))]
}
