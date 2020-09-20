package main

import (
	"fmt"
	"math/rand"
	"tempEmail/internal/secmail"
	"time"
)

const (
	nameLen = 10
)

var (
	tails = [3]string{"com", "net", "org"}
)

func main() {
	userData := GenerateEmail()
	fmt.Println(userData.Email())

	for {
		mails := secmail.GetMails(userData.Login, userData.Domain)
		fmt.Println(mails)
		time.Sleep(5 * time.Second)
	}
}

type Email struct {
	Login  string
	Domain string
}

func (e Email) Email() string {
	return fmt.Sprintf("%s@%s", e.Login, e.Domain)
}

func GenerateEmail() Email {
	userName := RandomString(nameLen)
	tld := tails[rand.Intn(len(tails))]
	domain := fmt.Sprintf("1secmail.%s", tld)
	return Email{Login: userName, Domain: domain}
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
