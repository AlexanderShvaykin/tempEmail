package secmail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseUrl = "https://1secmail.com/api/v1/"

func GetMails(login string, domain string) []Mail {
	var mails []Mail

	params := map[string]string{
		"action": "getMessages",
		"login":  login,
		"domain": domain,
	}
	endpoint := buildURL(baseUrl, params)
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	if err := json.Unmarshal(body, &mails); err != nil {
		fmt.Println(err)
	}

	return mails
}

func buildURL(base string, p map[string]string) string {
	base += "?"
	for key, value := range p {
		base += key
		base += "="
		base += url.QueryEscape(value)
		base += "&"
	}
	return base[0:(len(base) - 1)]
}
