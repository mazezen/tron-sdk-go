package client

import (
	"context"
	"fmt"
	"time"

	"github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GrpcClient grpc client structure
type GrpcClient struct {
	Host         string
	Conn         *grpc.ClientConn
	WalletClient tronpb.WalletClient
	timeout      time.Duration
	opts         []grpc.DialOption
	apiKey       string
}

// NewGrpcClient create grpc client
func NewGrpcClient(host string) *GrpcClient {
	client := &GrpcClient{
		Host:    host,
		timeout: time.Second * 5,
	}
	return client
}

// NewGrpcClientWithTimeout create grpc client with timeout
func (c *GrpcClient) NewGrpcClientWithTimeout(host string, timeout time.Duration) *GrpcClient {
	client := &GrpcClient{
		Host:    host,
		timeout: timeout,
	}
	return client
}

// SetTimeout set timeout on connection
func (c *GrpcClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// SetApiKey set api key on connection
func (c *GrpcClient) SetApiKey(apiKey string) {
	c.apiKey = apiKey
}

func (c *GrpcClient) getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	if len(c.apiKey) > 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", c.apiKey)
	}
	return ctx, cancel
}

func (c *GrpcClient) Stop() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *GrpcClient) Reconnect(host string) error {
	c.Stop()
	if len(c.Host) > 0 {
		c.Host = host
	}
	return c.Start(c.opts...)
}

// Start initiate grpc connection
func (c *GrpcClient) Start(opts ...grpc.DialOption) error {
	var err error
	if len(c.Host) == 0 {
		c.Host = "grpc.trongrid.io:50051"
	}
	c.opts = opts
	c.Conn, err = grpc.NewClient(c.Host, opts...)
	if err != nil {
		return fmt.Errorf("create grpc client: %w", err)
	}
	c.WalletClient = tronpb.NewWalletClient(c.Conn)
	return nil
}
