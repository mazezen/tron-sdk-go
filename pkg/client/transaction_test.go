package client

import (
	"math/big"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/common/decimal"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_CreateTransaction2(t *testing.T) {
	from, err := address.Base58ToAddress("TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4")
	assert.NoError(t, err, "base58 address convert to Address should not error")
	assert.NotNil(t, from, "address should not be nil")
	t.Logf("from tron hex: %s", from.Hex())
	t.Logf("from eth hex: %s", from.EthHex())

	to, err := address.Base58ToAddress("TNngs5j7HG54C8DjzgGRuLFtHnooCjGukw")
	assert.NoError(t, err, "base58 address convert to Address should not error")
	assert.NotNil(t, to, "address should not be nil")
	t.Logf("to tron hex: %s", to.Hex())
	t.Logf("to eth hex: %s", to.EthHex())

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "client start should not error")
	defer client.Stop()

	tests := []struct {
		name   string
		from   string
		to     string
		amount *big.Float
	}{
		{
			name:   "create transaction - base58",
			from:   "TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4",
			to:     "TNngs5j7HG54C8DjzgGRuLFtHnooCjGukw",
			amount: big.NewFloat(1.0), // 1 TRX
		},
		{
			name:   "create transaction - tron hex",
			from:   "41edbbe86be140fd81327ddec7eb9d16f615fbee66",
			to:     "418c9c516a53c44a659047914954fa25bd288de34e",
			amount: big.NewFloat(0.1),
		},
		{
			name:   "create transaction - eth hex",
			from:   "0xedbbe86be140fd81327ddec7eb9d16f615fbee66",
			to:     "0x8c9c516a53c44a659047914954fa25bd288de34e",
			amount: big.NewFloat(0.2),
		},
		{
			name:   "create transaction - eth hex",
			from:   "0xedbbe86be140fd81327ddec7eb9d16f615fbee66",
			to:     "0x8c9c516a53c44a659047914954fa25bd288de34e",
			amount: big.NewFloat(0.0000001),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, accuracy := decimal.Mul(tt.amount, big.NewFloat(1000000)).Int64()
			if accuracy == big.Exact {
				t.Fatalf("lost accuracy")
			}
			t.Logf("amount: %d", amount)
			t.Logf("accuracy: %d", accuracy)
			tx, err := client.CreateTransaction2(tt.from, tt.to, amount)
			assert.NoError(t, err, "client create transaction should not error")
			assert.NotNil(t, tx, "client create transaction should not be nil")

			t.Logf("tx: %s", tx.GetResult())
			t.Logf("tx: %s", common.BytesToHexString(tx.GetTxid()))
		})
	}
}

func TestGrpcClient_GetTransactionInfoByBlockNum(t *testing.T) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "client start should not error")
	defer client.Stop()

	nowBlock2, err := client.GetNowBlock2()
	assert.NoError(t, err, "client GetNowBlock2 should not error")
	assert.NotNil(t, nowBlock2, "client GetNowBlock2 should not be nil")
	t.Logf("now block height: %d", nowBlock2.GetBlockHeader().GetRawData().GetNumber())

	transactionInfoList, err := client.GetTransactionInfoByBlockNum(nowBlock2.GetBlockHeader().GetRawData().GetNumber())
	assert.NoError(t, err, "client GetTransactionInfoByBlockNum should not error")
	assert.NotNil(t, transactionInfoList, "client GetTransactionInfoByBlockNum should not be nil")

	for _, txInfo := range transactionInfoList.GetTransactionInfo() {
		t.Logf("txInfo: %v", txInfo.BlockNumber)
		t.Logf("tx log:%v", txInfo.GetLog())
		t.Logf(" ------------------------------------------------------------------- ")
	}
}
