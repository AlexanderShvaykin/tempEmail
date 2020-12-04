package http

import (
	"io/ioutil"
	"net/http"
	uri "net/url"
)

type Methods interface {
	Get(url string, params map[string]string) ([]byte, error)
}

type Client struct {
}

func (c Client) Get(url string, params map[string]string) ([]byte, error) {
	if len(params) != 0 {
		for key, value := range params {
			url += key
			url += "="
			url += uri.QueryEscape(value)
			url += "&"
		}
		url = url[0:(len(url) - 1)]
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
