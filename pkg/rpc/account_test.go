package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthAccounts(t *testing.T) {
	client := New("")

	err := client.EthAccounts(context.Background())
	assert.NoError(t, err, "eth_accounts should not error")

}
