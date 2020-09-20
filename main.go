package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nameLen = 10
)

var (
	tails = [3]string{"com", "net", "org"}
)

func main() {
	fmt.Println(GenerateEmail())
	fmt.Println(GenerateEmail())
	fmt.Println(GenerateEmail())
}

func GenerateEmail() string {
	userName := RandomString(nameLen)
	tld := tails[rand.Intn(len(tails))]
	return fmt.Sprintf("%s@1secmail.%s", userName, tld)
}

func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	rand.Seed(time.Now().UnixNano())

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
