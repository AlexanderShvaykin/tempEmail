package main

import (
	"fmt"
	"math/rand"
	"os"
	"tempEmail/pkg/cmd/root"
	"tempEmail/pkg/until"
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
	cmd := root.NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
	userName := until.RandomString(nameLen)
	tld := tails[rand.Intn(len(tails))]
	domain := fmt.Sprintf("1secmail.%s", tld)
	return Email{Login: userName, Domain: domain}
}
