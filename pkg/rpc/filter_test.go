package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthNewFilter(t *testing.T) {
	client := New("")

	call := EthNewFilterParams{
		Address: []string{
			"cc2e32f2388f0096fae9b055acffd76d4b3e5532",
			"E518C608A37E2A262050E10BE0C9D03C7A0877F3",
		},
		FromBlock: "0x989680",
		ToBlock:   "0x9959d0",
		Topics: []interface{}{
			"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			nil,
			[]interface{}{
				"0x0000000000000000000000001806c11be0f9b9af9e626a58904f3e5827b67be7",
				"0x0000000000000000000000003c8fb6d064ceffc0f045f7b4aee6b3a4cefb4758",
			},
		},
	}
	filter, err := client.EthNewFilter(context.Background(), call)
	assert.Nil(t, err)

	t.Logf("filter: %v", filter)
}

func TestRpcClient_EthNewBlockFilter(t *testing.T) {
	client := New("")

	fid, err := client.EthNewBlockFilter(context.Background())
	assert.Nil(t, err)
	t.Logf("fid: %v", fid)
}

func TestRpcClient_EthGetFilterChanges(t *testing.T) {
	client := New("")

	changes, err := client.EthGetFilterChanges(context.Background(), "")
	assert.Nil(t, err)
	t.Logf("changes: %v", changes)
}

func TestRpcClient_EthGetFilterLogs(t *testing.T) {
	client := New("")

	logs, err := client.EthGetFilterLogs(context.Background(), "")
	assert.Nil(t, err)
	t.Logf("logs: %v", logs)
}

func TestRpcClient_UninstallFilter(t *testing.T) {
	client := New("")
	uninstallFilter, err := client.UninstallFilter(context.Background(), "")
	assert.Nil(t, err)
	t.Logf("uninstall filter: %v", uninstallFilter)
}
