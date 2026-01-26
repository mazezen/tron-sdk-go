package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetBandwidthPrices(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	bandwidthPrices, err := client.GetBandwidthPrices()
	assert.NoError(t, err, "failed to get bandwidth prices")
	assert.NotNil(t, bandwidthPrices, "bandwidth prices should not be nil")

	t.Logf("bandwidth price: %s", bandwidthPrices.GetPrices())
}

func TestGrpcClient_GetEnergyPrices(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	energyPrices, err := client.GetEnergyPrices()
	assert.NoError(t, err, "failed to get energy prices")
	assert.NotNil(t, energyPrices, "energy prices should not be nil")

	t.Logf("energy price: %s", energyPrices.GetPrices())
}

func TestGrpcClient_GetMemoPrice(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	memoPrice, err := client.GetMemoPrice()
	assert.NoError(t, err, "failed to get memo price")
	assert.NotNil(t, memoPrice, "no memo price")

	t.Logf("memo price: %v", memoPrice.GetPrices())
}
