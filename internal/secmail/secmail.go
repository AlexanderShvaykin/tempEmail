package secmail

import (
	"encoding/json"
	"fmt"
	"github.com/AlexanderShvaykin/tempemail/pkg/http"
)

const baseUrl = "https://1secmail.com/api/v1/"

func GetMails(login string, domain string, client http.Methods) []Mail {
	var mails []Mail

	params := map[string]string{
		"action": "getMessages",
		"login":  login,
		"domain": domain,
	}
	body, err := client.Get(baseUrl, params)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &mails); err != nil {
		fmt.Println(err)
	}

	return mails
}

func GetMail(login string, domain string, id string, client http.Methods) Mail {
	var mail Mail
	params := map[string]string{
		"action": "readMessage",
		"login":  login,
		"domain": domain,
		"id":     id,
	}
	body, err := client.Get(baseUrl, params)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &mail); err != nil {
		fmt.Println(err)
	}

	return mail
}

type Mail struct {
	ID      int64
	From    string
	Subject string
	Date    string
	Body    string
}
