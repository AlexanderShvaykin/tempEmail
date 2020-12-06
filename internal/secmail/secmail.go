package secmail

import (
	"encoding/json"
	"fmt"
	"tempEmail/pkg/http"
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

type Mail struct {
	ID      int64
	From    string
	Subject string
	Date    string
}
