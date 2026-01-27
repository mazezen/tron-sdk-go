package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthCoinbase(t *testing.T) {
	client := New("")

	result, err := client.EthCoinbase(context.Background())
	assert.NoError(t, err, "eth coinbase should not error")

	t.Logf("result: %v", result)
}
