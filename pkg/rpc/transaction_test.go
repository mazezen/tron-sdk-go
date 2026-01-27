package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthGetTransactionByBlockHashAndIndex(t *testing.T) {
	client := New("")

	tx, err := client.EthGetTransactionByBlockHashAndIndex(context.Background(), "00000000020ef11c87517739090601aa0a7be1de6faebf35ddb14e7ab7d1cc5b", "0x0")
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	t.Log(tx)
}

func TestRpcClient_EthGetTransactionByBlockNumberAndIndex(t *testing.T) {
	client := New("")

	tx, err := client.EthGetTransactionByBlockNumberAndIndex(context.Background(), "0xfb82f0", "0x0")
	assert.NoError(t, err)
	assert.NotNil(t, tx)
	t.Log(tx)
}

func TestRpcClient_EthGetTransactionByHash(t *testing.T) {
	client := New("")

	tx, err := client.EthGetTransactionByHash(context.Background(), "c9af231ad59bcd7e8dcf827afd45020a02112704dce74ec5f72cb090aa07eef0")
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	t.Log(tx)
}

func TestRpcClient_EthGetTransactionReceipt(t *testing.T) {
	client := New("")

	receipt, err := client.EthGetTransactionReceipt(context.Background(), "c9af231ad59bcd7e8dcf827afd45020a02112704dce74ec5f72cb090aa07eef0")
	assert.NoError(t, err)
	assert.NotNil(t, receipt)

	t.Log(receipt)
}
