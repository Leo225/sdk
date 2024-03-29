package http

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Post(url, contentType string, args map[string]string) (result string, err error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}
	resp, err := http.Post(url, contentType, strings.NewReader(encode(args)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

func Get(url string, args map[string]string) (result string, err error) {
	if args != nil {
		url += "?" + encode(args)
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

func encode(params map[string]string) string {
	args := url.Values{}
	for k, v := range params {
		args.Set(k, v)
	}
	return args.Encode()
}
