package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetBurnTrx(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")

	burnTrx, err := client.GetBurnTrx()
	assert.NoError(t, err, "failed to get burn trx")
	assert.NotNil(t, burnTrx, "burn trx should not be nil")

	t.Logf("burn trx: %v", burnTrx.GetNum())
}
