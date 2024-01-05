package http

import (
	"fmt"
	"net/url"
	"testing"
)

var httpClient = NewClient()

func TestGet(t *testing.T) {
	URL := "https://cn.bing.com"
	body, _, err := httpClient.Get(URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GET", string(body))
}

func TestPost(t *testing.T) {
	URL := "https://cn.bing.com"
	values := url.Values{}
	values.Set("name", "star")
	body, _, err := httpClient.Post(URL, values)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST", string(body))
}

func TestPostJSON(t *testing.T) {
	URL := "https://cn.bing.com"
	values := map[string]interface{}{
		"name": "star",
	}
	body, _, err := httpClient.PostJSON(URL, values)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST-JSON", string(body))
}
