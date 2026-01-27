package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_GetBlockNumber(t *testing.T) {
	client := New("")

	numberHeight, err := client.EthBlockNumber(context.Background())
	assert.NoError(t, err, "GetBlockNumber err")

	t.Logf("now block height %d", numberHeight)
}

func TestRpcClient_EthGetBlockByHash(t *testing.T) {
	client := New("")

	res, err := client.EthGetBlockByHash(context.Background(), "0x0000000004beb16317d497b4bd5a8a3e94e83fd31f0dd49aa1d0e68ccb3a3b9f", true)
	assert.NoError(t, err, "GetBlockByHash err")
	t.Logf("res %+v", res)
}

func TestRpcClient_EthGetBlockByNumber(t *testing.T) {
	client := New("")

	res, err := client.EthGetBlockByNumber(context.Background(), "0xF9CC56", true)
	assert.NoError(t, err, "GetBlockByNumber err")

	t.Logf("res %+v", res)
}

func TestRpcClient_EthGetBlockTransactionCountByHash(t *testing.T) {
	client := New("")

	count, err := client.EthGetBlockTransactionCountByHash(context.Background(), "0x0000000004beb16317d497b4bd5a8a3e94e83fd31f0dd49aa1d0e68ccb3a3b9f")
	assert.NoError(t, err, "GetBlockTransactionCountByHash err")

	t.Logf("count %s", count)
}

func TestRpcClient_EthGetBlockTransactionCountByNumber(t *testing.T) {
	client := New("")

	count, err := client.EthGetBlockTransactionCountByNumber(context.Background(), "0xF96B0F")
	assert.NoError(t, err, "GetBlockTransactionCountByNumber err")

	t.Logf("count %s", count)
}

func TestRpcClient_EthGetWork(t *testing.T) {
	client := New("")

	res, err := client.EthGetWork(context.Background())
	assert.NoError(t, err, "GetWork err")
	t.Logf("res %+v", res)
}
