package sdk_rpc

import (
	"context"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthGetBalance(t *testing.T) {
	client := New("")

	balance, err := client.EthGetBalance(context.Background(), "0x41f0cc5a2a84cd0f68ed1667070934542d673acbd8", "latest")
	assert.NoError(t, err, "should not error")

	b := common.HexToUint64(balance)
	t.Logf("balance: %d", b)
}
