package api

import (
	"time"

	"resty.dev/v3"
)

type Client struct {
	client *resty.Client
}

const (
	defaultRetryCount = 3
)

func NewApiClient(host string) *Client {
	ac := &Client{
		client: resty.New(),
	}

	ac.client.SetHeader("Accept", "application/json")
	ac.client.SetHeader("Content-Type", "application/json")
	ac.client.SetRetryCount(defaultRetryCount)

	if host == "" {
		ac.client.SetBaseURL("https://api.trongrid.io")
	} else {
		ac.client.SetBaseURL(host)
	}

	return ac
}

func (a *Client) WithTimeout(time time.Duration) *Client {
	a.client.SetTimeout(time)
	return a
}

func (a *Client) WithKey(key, val string) *Client {
	a.client.SetHeader(key, val)
	return a
}
