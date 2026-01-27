package sdk_rpc

import (
	"net/http"
	"strings"
	"time"
)

// RpcClient Tron JSON-RPC client
type RpcClient struct {
	client *http.Client
	host   string
}

func New(host string, opts ...Option) *RpcClient {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	return &RpcClient{
		host: strings.TrimRight(host, "/"),
		client: &http.Client{
			Timeout: cfg.timeout,
			Transport: &http.Transport{
				MaxIdleConns:        64,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
	}
}

type Option func(*config)

type config struct {
	timeout time.Duration
}

func defaultConfig() *config {
	return &config{
		timeout: 15 * time.Second,
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *config) {
		c.timeout = timeout
	}
}
