package client

import (
	"slices"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_VoteWitnessAccount2(t *testing.T) {
	from, err := address.Base58ToAddress("TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4")
	assert.NoError(t, err, "base58 address convert to Address should not error")
	assert.NotNil(t, from, "address should not be nil")
	t.Logf("from tron hex: %s", from.TronHex())
	t.Logf("from eth hex: %s", from.EthHex())

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client := NewGrpcClient("grpc.trongrid.io:50051")
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start client")
	defer client.Stop()

	tests := []struct {
		name string
		from string
		vote map[string]int64
	}{
		// 必须先质押 TRX 获得投票权
		// 投票是可变的：每次调用 VoteWitnessAccount2 会覆盖你之前的投票(不是增量)
		// 撤票：把对应 SR 的票数设为 0
		// 资源消耗：只耗 Bandwidth（带宽），有免费 1500 点/天，通常免费；不足时烧很少 TRX
		{
			name: "Vote for super representatives",     // 投票者地址
			from: "TXeEFbJpGM6zgWFgcUD1Prar2hK3iAuvN4", // 投票的超级代表列表 (map: SR地址 → 投票数量)
			vote: map[string]int64{
				"TN2W4cc7a4dsYyTLiLMWa9m7jVpdLjGvYs": int64(100000000), // 例如 100 TRX 的票（单位 sun）
				"TVFKwzE8qeETLaZEHMx2tjEsdnujAgAWaA": int64(50000000),  // 50 TRX 的票
			},
		},
		{
			name: "Vote for super representatives",
			from: "41edbbe86be140fd81327ddec7eb9d16f615fbee66",
			vote: map[string]int64{
				"TC6qGw3d6h25gjcM64KLuZn1cznNi5NR6t": int64(50000000), // 50 TRX 的票
				"TQopP5GM68QoqLzpz8YReDfSoCMkvwcZYd": int64(50000000), // 50 TRX 的票
			},
		},
		{
			name: "Vote for super representatives",
			from: "0xedbbe86be140fd81327ddec7eb9d16f615fbee66",
			vote: map[string]int64{
				"TQ4bh4nQknQp33vuf1mUAKu5M5TWW8cTAD": int64(50000000), // 50 TRX 的票
				"TMafrJCuNoYq3mg9dDThfg7c9VP6enZN6j": int64(50000000), // 50 TRX 的票
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 验证 SR 地址 是否是 Super Representatives
			witnessList, err := client.ListWitnesses()
			assert.NoError(t, err, "failed to list witnesses")
			assert.NotNil(t, witnessList, "list witnesses should not be nil")
			var superRepresentative []string
			for _, witness := range witnessList.Witnesses {
				//t.Logf("witness address: %s", address.HexToBase58Address(common.BytesToHexString(witness.Address)))
				superRepresentative = append(superRepresentative, address.HexToBase58Address(common.BytesToHexString(witness.Address)))
			}

			for voteAddress := range tt.vote {
				if !slices.Contains(superRepresentative, voteAddress) {
					t.Errorf("[]%s is not super representatives", voteAddress)
				}
			}

			tx, err := client.VoteWitnessAccount2(tt.from, tt.vote)
			assert.NoError(t, err, "VoteWitnessAccount2 Should not return error")
			assert.NotNil(t, tx, "VoteWitnessAccount2 should not be nil")

			t.Logf("[]%s VoteWitnessAccount2 result is: %v", tt.name, tx.GetResult())
			t.Logf("[]%s VoteWitnessAccount2 result is: %v", tt.name, common.BytesToHexString(tx.Txid))
		})
	}

}
