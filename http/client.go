package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultTimeout = 3 * time.Second
)

type Client struct {
	Header *Header
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		Header: new(Header),
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func NewHeaderClient(header *Header) *Client {
	return &Client{
		Header: header,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (c *Client) SetHeader(header *Header) {
	c.Header = header
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

func (c *Client) SetTransport(transport http.RoundTripper) {
	c.client.Transport = transport
}

func (c *Client) Get(url string, values url.Values) (body []byte, statusCode int, err error) {
	if values != nil {
		url += "?" + values.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	return c.do(req)
}

func (c *Client) Post(url string, values url.Values) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	return c.do(req)
}

func (c *Client) PostData(url string, values []byte) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(values))
	if err != nil {
		return
	}
	return c.do(req)
}

func (c *Client) PostJSON(url string, values interface{}) (body []byte, statusCode int, err error) {
	js, err := json.Marshal(values)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(js))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

func (c *Client) Do(r *http.Request) (response *http.Response, err error) {
	for k, v := range *c.Header {
		r.Header.Set(k, v)
	}
	response, err = c.client.Do(r)
	return
}

func (c *Client) do(r *http.Request) (body []byte, statusCode int, err error) {
	resp, err := c.Do(r)
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	return
}
