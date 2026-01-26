package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetTransactionListFromPending(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")

	idList, err := client.GetTransactionListFromPending()
	assert.NoError(t, err, "failed to get transaction list from grpc client")

	for _, txId := range idList.GetTxId() {
		//t.Logf("transaction id: %s", txId)
		pendingTx, err := client.GetTransactionFromPending(txId)
		assert.NoError(t, err, "failed to get transaction from grpc client")
		if txId != "" {
			t.Logf("pending transaction: %v", pendingTx.GetRawData().GetRefBlockNum())
		}

	}
}

func TestGrpcClient_GetPendingSize(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")

	numberMessage, err := client.GetPendingSize()
	assert.NoError(t, err, "failed to get pending size")
	assert.NotEqual(t, numberMessage, int64(0))

	t.Logf("successfully got pending size: %d", numberMessage.GetNum())
}
