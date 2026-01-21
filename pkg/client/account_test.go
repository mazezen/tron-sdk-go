package client

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetAccount(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	client.SetTimeout(20 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	const testBase58 = "TTiWGY3myYPyjTYrBHeACfvRaGppNbLVr9"

	tests := []struct {
		name     string
		addr     string
		wantAddr string
	}{
		{
			name:     "GetAccount - base58",
			addr:     testBase58,
			wantAddr: "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5",
		},
		{
			name:     "GetAccount - tron hex address",
			addr:     "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5",
			wantAddr: "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5",
		},
		{
			name:     "GetAccount - eth hex address",
			addr:     "0xc2aa7035af217ee90e227b99f60cbf2ca45c84d5",
			wantAddr: "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := client.GetAccount(tt.addr)
			assert.NoError(t, err, "GetAccount should not return error")
			assert.NotNil(t, acc, "Account should not be nil")

			if acc != nil {
				assert.Equal(t, tt.wantAddr, hex.EncodeToString(acc.Address), "Account address mismatch")
			}

			t.Logf("[]%s Account is: %v", tt.name, acc)
		})
	}
}

func TestGrpcClient_SetAccountId(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	client.SetTimeout(20 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	const testBase58 = "TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4"
	const privateKey = "" // write private key
	tests := []struct {
		name      string
		accountId string
		addr      string
	}{
		{
			name:      "SetAccountId - base58",
			accountId: "testbhoi", // ad890b90d337842b2f573c5830c878fcefd5b3fac1cc5facffcd3ea67487bdf4
			addr:      testBase58,
		},
		{
			name:      "SetAccountId - tron hex address",
			accountId: "testbhob",
			addr:      "41edbbe86be140fd81327ddec7eb9d16f615fbee66",
		},
		{
			name:      "SetAccountId - eth hex address",
			accountId: "testbhoi3",
			addr:      "0xedbbe86be140fd81327ddec7eb9d16f615fbee66",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := client.SetAccountId(tt.accountId, tt.addr)
			assert.NoError(t, err, "SetAccount should not return error")
			assert.NotNil(t, tx, "SetAccount should not be nil")

			// signature
			ecdsa, err := crypto.HexToECDSA(privateKey)
			assert.NoError(t, err, "HexToECDSA should not return error")
			tx, err = SignTransactionECDSA(tx, ecdsa)
			assert.NoError(t, err, "SignTransactionECDSA should not return error")
			assert.NotNil(t, tx, "SignTransactionECDSA should not be nil")

			// Broadcast
			res, err := client.BroadcastTransaction(tx)
			assert.NoError(t, err, "BroadcastTransaction should not return error")
			assert.NotNil(t, res, "BroadcastTransaction should not be nil")

			t.Logf("[]%s BroadcastTransaction result is: %v", tt.name, res.GetResult())
			t.Logf("[]%s BroadcastTransaction result code is: %v", tt.name, res.GetCode())
			t.Logf("[]%s BroadcastTransaction result message is: %v", tt.name, res.GetMessage())
		})
	}
}

func TestGrpcClient_GetAccountById(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	client.SetTimeout(20 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	tests := []struct {
		name string
		id   string
	}{
		{
			name: "get account by id",
			id:   "testbhoi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := client.GetAccountById(tt.id)
			assert.NoError(t, err, "GetAccountById should not return error")
			assert.NotNil(t, acc, "Account should not be nil")

			t.Logf("[]%s Account id is: %v", tt.name, acc)
		})
	}
}
