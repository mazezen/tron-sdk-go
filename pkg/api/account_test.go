package api

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestApiClient_CreateAccount(t *testing.T) {
	err := NewApiClient("").
		CreateAccount(
			context.Background(),
			"TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g",
			"TFgY1uN8buRxAtV2r6Zy5sG3ACko6pJT1y",
			true,
		)

	if err != nil {
		t.Logf("CreateAccount error: %v\n", err)
	} else {
		t.Logf("CreateAccount OK\n")
	}
}

func TestClient_GetAccountResource(t *testing.T) {
	res, err := NewApiClient("").GetAccountResource(context.Background(), "TK5LJ8JcAqFPsyH7BgHLvY1py3PyBGPCZk", true)
	assert.NoError(t, err, "GetAccountResource error")
	assert.NotNil(t, res)

	j, _ := json.Marshal(res)
	t.Logf("GetAccountResource OK\n, res: %v", common.JSONPrettyFormat(string(j)))
	for _, v := range res.AssetNetLimit {
		t.Logf("key: [%s]: value: %d", v.Key, v.Value)
	}

	for _, v := range res.AssetNetUsed {
		t.Logf("key: [%s]: value: %d", v.Key, v.Value)
	}
}

func TestApiClient_GetAccountBalance(t *testing.T) {
	res, err := NewApiClient("").GetAccountBalance(context.Background(),
		"TLLM21wteSPs4hKjbxgmH1L6poyMjeTbHm", "0000000000010c4a732d1e215e87466271e425c86945783c3d3f122bfa5affd9", true)

	assert.NoError(t, err, "get account balance should not error")

	t.Logf("get account balance res: %+v", res)
	t.Logf("get account balance : %+d", res.Balance)
	t.Logf("get account balance block hash : %s", res.BlockIdentifier.Hash)
	t.Logf("get account balance block number : %d", res.BlockIdentifier.Number)
}

func TestClient_GetAccountNet(t *testing.T) {
	res, err := NewApiClient("").GetAccountNet(context.Background(),
		"TLLM21wteSPs4hKjbxgmH1L6poyMjeTbHm", true)
	assert.NoError(t, err, "GetAccountNet error")
	assert.NotNil(t, res)

	j, _ := json.Marshal(res)
	t.Logf("GetAccountNet OK\n, res: %v", common.JSONPrettyFormat(string(j)))
}
