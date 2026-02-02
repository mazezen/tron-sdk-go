package client

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
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

func TestGrpcClient_GetAccountBalance(t *testing.T) {
	toAddress, _ := address.Base58ToAddress("TCypds3XB6zjo6dpyLLKd6rkz4btrEfkob")
	t.Logf("tron address is: %v", toAddress.Hex())
	t.Logf("eth address is: %v", toAddress.EthHex())

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	client.SetTimeout(20 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "should not return error")
	defer client.Stop()

	tests := []struct {
		name   string
		addr   string
		hash   string
		number int64
	}{
		{
			name:   "get account balance - base58",
			addr:   "TCypds3XB6zjo6dpyLLKd6rkz4btrEfkob",
			hash:   "0000000004bd01cdce66bc05738cb77155afaaa6846502c621c4e52cb5d2e84e",
			number: 79495629,
		},
		{
			name:   "get account balance - tron hex",
			addr:   "4121061fa3592c73c4c1692afc3e2e8ae81ae94911",
			hash:   "0000000004bd01cdce66bc05738cb77155afaaa6846502c621c4e52cb5d2e84e",
			number: 79495629,
		},
		{
			name:   "get account balance - eth hex",
			addr:   "0x21061fa3592c73c4c1692afc3e2e8ae81ae94911",
			hash:   "0000000004bd01cdce66bc05738cb77155afaaa6846502c621c4e52cb5d2e84e",
			number: 79495629,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, err := client.GetAccountBalance(tt.addr, tt.hash, tt.number)
			assert.NoError(t, err, "GetAccountBalance should not return error")
			assert.NotNil(t, balance, "GetAccountBalance should not be nil")

			t.Logf("[]%s GetAccountBalance result is: %d", tt.name, balance.GetBalance())
		})
	}
}

func TestGrpcClient_UpdateAccount2(t *testing.T) {
	toAddress, _ := address.Base58ToAddress("TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4")
	t.Logf("tron address is: %v", toAddress.Hex())
	t.Logf("eth address is: %v", toAddress.EthHex())

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "should not return error")
	defer client.Stop()

	tests := []struct {
		name        string
		addr        string
		accountName string
	}{
		{
			name:        "update account name - base58",
			addr:        "TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4",
			accountName: "TXeEFbJpGM6zgWFgcUD1",
		},
		{
			name:        "update account name - tron hex",
			addr:        "41edbbe86be140fd81327ddec7eb9d16f615fbee66",
			accountName: "TXeEFbJpGM6zgWFgcUD2",
		},
		{
			name:        "update account name - eth hex",
			addr:        "0xedbbe86be140fd81327ddec7eb9d16f615fbee66",
			accountName: "TXeEFbJpGM6zgWFgcUD3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction, err := client.UpdateAccount2(tt.addr, tt.accountName)
			assert.NoError(t, err, "UpdateAccount2 should not return error")
			assert.NotNil(t, transaction, "UpdateAccount2 should not be nil")

			t.Logf("[]%s UpdateAccount2 result is: %v", tt.name, transaction.GetResult())
			t.Logf("[]%s UpdateAccount2 tx is :%s", tt.name, common.BytesToHexString(transaction.GetTxid()))
		})
	}
}

func TestGrpcClient_CreateAccount(t *testing.T) {
	var from = "T..."
	var to = "T..."
	var typ = 1

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "should not return error")
	defer client.Stop()

	tests := []struct {
		name string
		from string
		to   string
		typ  int
	}{
		{
			name: "create account",
			from: from,
			to:   to,
			typ:  typ,
		},
		{
			name: "create account2",
			from: from,
			to:   to,
			typ:  typ,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "create account" {
				tx, err := client.CreateAccount(tt.from, tt.to, tt.typ)
				assert.NoError(t, err, "CreateAccount should not return error")
				assert.NotNil(t, tx, "CreateAccount should not be nil")
				t.Logf("[]%s CreateAccount result is: %v", tt.name, tx)
			}
			if tt.name == "create account2" {
				tx, err := client.CreateAccount2(tt.from, tt.to, tt.typ)
				assert.NoError(t, err, "CreateAccount should not return error")
				assert.NotNil(t, tx, "CreateAccount should not be nil")
				t.Logf("[]%s CreateAccount result is: %v", tt.name, tx)
			}
		})
	}
}

func TestGrpcClient_GetAccountNet(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "should not return error")
	defer client.Stop()

	net, err := client.GetAccountNet("")
	assert.NoError(t, err, "should not return error")
	assert.NotNil(t, net, "GetAccountNet should not be nil")

	t.Logf("net is: %v", net)
}

func TestGrpcClient_AccountPermissionUpdate(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.shasta.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "should not return error")
	defer client.Stop()

	// === 测试账户 ===
	// 已激活账户，有 owner 权限
	fromAddress := "TYourTestAddressHereLikeTS..."

	// === Owner Permission（必须提供） ===
	// 示例：保持原 owner，但 threshold=1，1/1 key（最简单）
	owner := map[string]interface{}{
		"threshold": int64(1),
		"keys": map[string]int64{
			fromAddress: 1, // 自己给自己权重1
		},
	}

	// === Active Permission 示例 ===
	// 示例：2个 active 权限组
	// Active0: threshold=2，需要 2/3 签名才能转账等操作
	actives := []map[string]interface{}{
		{
			"name":      "active0",
			"threshold": int64(2),
			"operations": map[string]bool{
				"TransferContract":        true,
				"TransferAssetContract":   true,
				"TriggerSmartContract":    true,
				"FreezeBalanceContract":   true,
				"UnfreezeBalanceContract": true,
				// 可添加更多：AccountUpdateContract, AccountPermissionUpdateContract 等
				// 但 AccountPermissionUpdateContract 通常只留给 owner
			},
			"keys": map[string]int64{
				"TSignerAddress1...": 1,
				"TSignerAddress2...": 1,
				"TSignerAddress3...": 1,
			},
		},
		// 可选：第二个 active 组（例如只允许特定操作）
		{
			"name":      "active1",
			"threshold": 1,
			"operations": map[string]bool{
				"UpdateAccountContract": true,
			},
			"keys": map[string]int64{
				fromAddress: 1,
			},
		},
	}
	// === Witness Permission（可选，如果是 SR 账户才需要，通常 nil） ===
	var witness map[string]interface{} = nil // 大多数账户不需要

	// === 调用 ===
	txExt, err := client.AccountPermissionUpdate(
		fromAddress,
		owner,
		witness,
		actives,
	)
	assert.NoError(t, err, "AccountPermissionUpdate 调用不应报错")
	assert.NotNil(t, txExt, "TransactionExtention 不应为空")

	if txExt.Result != nil {
		assert.Equal(t, tronpb.Return_SUCCESS, txExt.Result.Code,
			fmt.Sprintf("权限更新模拟失败: %s", string(txExt.Result.Message)))
		t.Log("模拟更新成功（未签名），code=0")
	} else {
		t.Fatal("txExt.Result 为 nil，请检查 protobuf 定义")
	}

	// 输出调试信息
	t.Logf("Transaction ID (txid): %x", txExt.GetTxid())
	t.Logf("Energy used (估算): %d", txExt.EnergyUsed)
	t.Logf("完整响应: %+v", txExt)

	// 可选：打印 raw transaction hex（用于后续签名测试）
	// rawDataHex := hex.EncodeToString(txExt.Transaction.GetRawData().GetRawData())
	// t.Logf("Raw Data Hex: %s", rawDataHex)
}
