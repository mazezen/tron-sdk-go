package api

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetTransactionById(t *testing.T) {
	var value = "16492dfb8543b649d8e7ab3e5be1e0e2280f0d006004ca0f2f26ceede0e1a284"
	response, err := NewApiClient("").GetTransactionById(context.Background(), value)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.Marshal(response)
	assert.NoError(t, err, fmt.Sprintf("Error marshaling response err: %v", err))
	t.Logf("%v", common.JSONPrettyFormat(string(j)))
}

func TestClient_GetTransactionInfoById(t *testing.T) {
	//var value = "16492dfb8543b649d8e7ab3e5be1e0e2280f0d006004ca0f2f26ceede0e1a284"
	var value = "b54803b9e0ddfee48a24434393537c6ef153f3351ea31618fb6d5b3468068888"
	response, err := NewApiClient("").GetTransactionInfoById(context.Background(), value)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.Marshal(response)
	assert.NoError(t, err, fmt.Sprintf("Error marshaling response err: %v", err))
	t.Logf("%v", common.JSONPrettyFormat(string(j)))
}

func TestClient_GetTransactionInfoByBlockNum(t *testing.T) {
	response, err := NewApiClient("").GetTransactionInfoByBlockNum(context.Background(), 79640626)
	if err != nil {
		t.Fatal(err)
	}

	//j, err := json.Marshal(response)
	//assert.NoError(t, err, fmt.Sprintf("Error marshaling response err: %v", err))

	count, err := NewApiClient("").GetTransactionCountByBlockNum(context.Background(), 79640626)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Transaction Count: %v", count)

	t.Logf("block contains {%+d} transatins", len(response))
	for _, idResponse := range response {
		t.Logf(" ----------------------------------------------------------------- ")
		t.Logf("transaction id (hash): %s", idResponse.ID)
		t.Logf("合约地址: %s", address.HexToAddress(idResponse.ContractAddress).String())
		t.Logf("能量消耗: %d", idResponse.Receipt.EnergyUsage)
		t.Logf("带宽消耗: %d", idResponse.Receipt.NetUsage)
		t.Logf(" ----------------------------------------------------------------- ")
	}

}
