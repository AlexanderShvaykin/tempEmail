package until

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	rand.Seed(time.Now().UnixNano())

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
