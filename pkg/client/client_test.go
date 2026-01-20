package client

import (
	"testing"

	"github.com/go-playground/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClient(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)
	block, err := client.GetNowBlock()
	t.Logf("block: %v\n", block)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
}

func TestNewGrpcClient(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)
	block, err := client.GetNowBlock2()
	t.Logf("block: %v\n", block)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
}
