package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/common/decimal"
	"github.com/stretchr/testify/assert"
)

func TestApiClient_TransferTrx(t *testing.T) {
	pk := "" // owner_address private key

	amount, _ := decimal.Mul(big.NewFloat(1), big.NewFloat(1e6)).Int64()
	transactionPayload := map[string]interface{}{
		"owner_address": "T...",
		"to_address":    "T...",
		"amount":        amount, // 0.0001 TRX
		"visible":       true,
	}

	// 交易
	transaction, err := NewApiClient("").CreateTransaction(context.Background(), transactionPayload)
	assert.NoError(t, err, fmt.Sprintf("create transaction err: %v", err))
	j, _ := json.Marshal(transaction)
	t.Logf("transaction: %s", common.JSONPrettyFormat(string(j)))

	// 签名
	ecdsaPrivateKey, _ := crypto.HexToECDSA(pk)
	addressB58 := address.PubkeyToAddress(ecdsaPrivateKey.PublicKey).String()
	if addressB58 != transactionPayload["owner_address"].(string) {
		t.Fatalf("transaction owner address not match")
	}
	transaction, err = SingTransactionECDSA(transaction, ecdsaPrivateKey)
	j2, _ := json.Marshal(transaction)
	t.Logf("签名之后的transaction: %v", common.JSONPrettyFormat(string(j2)))

	// 广播
	txId, err := NewApiClient("").BroadTransaction(context.Background(), transaction)
	if err != nil {
		t.Errorf("broadcast transaction err: %v", err)
	} else {
		t.Logf("broadcast transaction id: %v", txId)
	}
}
