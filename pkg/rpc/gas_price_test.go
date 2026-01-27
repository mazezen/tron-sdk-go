package sdk_rpc

import (
	"context"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthGasPrice(t *testing.T) {
	client := New("")

	gasPrice, err := client.EthGasPrice(context.Background())
	assert.NoError(t, err, "should not return error")
	t.Logf("gasPrice: %d", common.HexToUint64(gasPrice))
}
