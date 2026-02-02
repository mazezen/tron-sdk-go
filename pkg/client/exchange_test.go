package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_ListExchanges(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	list, err := client.ListExchanges()
	assert.NoError(t, err, "failed to list exchanges")

	for _, exchange := range list.GetExchanges() {
		t.Logf("%+v", exchange)
	}
}

func TestGrpcClient_GetPaginatedExchangeList(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	list, err := client.GetPaginatedExchangeList(1, 5)
	assert.NoError(t, err, "failed to list exchange list")
	for _, exchange := range list.GetExchanges() {
		t.Logf("%+v", exchange)
	}
}

func TestGrpcClient_GetExchangeById(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	exchange, err := client.GetExchangeById(2)
	assert.NoError(t, err, "failed to get exchange by id")
	t.Logf("%+v", exchange)
}
