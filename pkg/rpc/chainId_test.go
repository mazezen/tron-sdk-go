package sdk_rpc

import (
	"context"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthChainId(t *testing.T) {
	client := New("")

	chainId, err := client.EthChainId(context.Background())
	assert.NoError(t, err, "get chain id should not error")

	t.Logf("chain id is: %d", common.HexToUint64(chainId))
}
