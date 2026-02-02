package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetContract(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	var value = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	smartContract, err := client.GetContract(value)
	assert.Nil(t, err, fmt.Sprintf("get contract should not be err: %v", err))
	assert.NotNil(t, smartContract, "smartContract should not be nil")

	t.Logf("smartContract: %v", smartContract)

	j, _ := json.Marshal(smartContract.GetAbi())
	t.Logf("smartContract_ABI: %v", common.JSONPrettyFormat(string(j)))
}

func TestGrpcClient_GetContractsInfo(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	var value = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	smartContractDataWrapper, err := client.GetContractInfo(value)
	assert.Nil(t, err, fmt.Sprintf("get contract should not be err: %v", err))
	assert.NotNil(t, smartContractDataWrapper, "smartContract should not be nil")

	j, _ := json.Marshal(smartContractDataWrapper)
	t.Logf("smartContractDataWrapper: %v", common.JSONPrettyFormat(string(j)))
}

func TestGrpcClient_TriggerConstantContract_USDT_BalanceOf(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	usdtContract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // 主网 USDT 合约地址

	queryAddress := "T..." // 要查的地址
	functionSelector := "balanceOf(address)"

	jsonParams := fmt.Sprintf(`[{"address": "%s"}]`, queryAddress)
	result, err := client.TriggerConstantContract(
		queryAddress, // ownerAddress
		usdtContract, // contractAddress
		functionSelector,
		jsonParams, // jsonString
	)
	assert.NoError(t, err, "TriggerConstantContract should not return error")
	assert.NotNil(t, result, "TransactionExtention should not be nil")
	assert.NotNil(t, result.Result, "result.Result should not be nil")
	assert.Equal(t, tronpb.Return_SUCCESS, result.Result.Code,
		fmt.Sprintf("contract failed: %s", string(result.Result.Message)))

	if result.Result.Result {
		t.Log("Contract call succeeded (result.Result = true)")
	} else {
		t.Fatalf("Contract call indicated failure despite code=0: %s", string(result.Result.Message))
	}

	t.Logf("Energy used: %d", result.EnergyUsed)
	t.Logf("Energy penalty: %d", result.EnergyPenalty)
	t.Logf("Result message: %s", string(result.Result.Message))

	if len(result.ConstantResult) == 0 {
		t.Log("No constant result returned")
		return
	}

	balanceBytes := result.ConstantResult[0]
	t.Logf("Raw balance bytes (hex): %s", hex.EncodeToString(balanceBytes))

	if len(balanceBytes) > 0 {
		balanceHex := hex.EncodeToString(balanceBytes)
		t.Logf("USDT balance (raw hex): 0x%s", balanceHex)

		balance := new(big.Int).SetBytes(balanceBytes)
		readable := new(big.Float).Quo(
			new(big.Float).SetInt(balance),
			big.NewFloat(1e6),
		)
		t.Logf("Readable USDT balance: %s", readable.Text('f', 6))
	}

	j, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("TriggerConstantContract full response:\n%s", string(j))
}

func TestGrpcClient_TriggerConstantContract_USDC_BalanceOf(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	usdcContract := "TLZSucJRjnqBKwvQz6n5hd29gbS4P7u7w8" // 主网 USDT 合约地址

	queryAddress := "T..." // 要查的地址
	functionSelector := "balanceOf(address)"

	jsonParams := fmt.Sprintf(`[{"address": "%s"}]`, queryAddress)
	result, err := client.TriggerConstantContract(
		queryAddress, // ownerAddress
		usdcContract, // contractAddress
		functionSelector,
		jsonParams, // jsonString
	)
	assert.NoError(t, err, "TriggerConstantContract should not return error")
	assert.NotNil(t, result, "TransactionExtention should not be nil")
	assert.NotNil(t, result.Result, "result.Result should not be nil")
	assert.Equal(t, tronpb.Return_SUCCESS, result.Result.Code,
		fmt.Sprintf("contract failed: %s", string(result.Result.Message)))

	if result.Result.Result {
		t.Log("Contract call succeeded (result.Result = true)")
	} else {
		t.Fatalf("Contract call indicated failure despite code=0: %s", string(result.Result.Message))
	}

	t.Logf("Energy used: %d", result.EnergyUsed)
	t.Logf("Energy penalty: %d", result.EnergyPenalty)
	t.Logf("Result message: %s", string(result.Result.Message))

	if len(result.ConstantResult) == 0 {
		t.Log("No constant result returned")
		return
	}

	balanceBytes := result.ConstantResult[0]
	t.Logf("Raw balance bytes (hex): %s", hex.EncodeToString(balanceBytes))

	if len(balanceBytes) > 0 {
		balanceHex := hex.EncodeToString(balanceBytes)
		t.Logf("USDC balance (raw hex): 0x%s", balanceHex)

		balance := new(big.Int).SetBytes(balanceBytes)
		readable := new(big.Float).Quo(
			new(big.Float).SetInt(balance),
			big.NewFloat(1e6),
		)
		t.Logf("Readable USDC balance: %s", readable.Text('f', 6))
	}

	j, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("TriggerConstantContract full response:\n%s", string(j))
}

func TestGrpcClient_EstimateEnergy_USDT_BalanceOf(t *testing.T) {
	// vm.estimateEnergy = true
	// vm.supportConstant = true
	client := NewGrpcClient("grpc.trongrid.io:50051") // 当前节点不支持能量估算。存在失败。
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	usdtContract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // 主网 USDT 合约地址
	from := "T..."
	to := ""
	transferAmount := int64(1_000_000) // 1 USDT

	functionSelector := "transfer(address,uint256)"
	jsonParams := fmt.Sprintf(`[
        {"address": "%s"},
        {"uint256": "%d"}
    ]`, to, transferAmount)
	t.Logf("Using jsonParams: %s", jsonParams)

	result, err := client.EstimateEnergy(
		from,
		usdtContract,
		functionSelector,
		0, // callValue → TRC-20 转账不需要附带 TRX
		jsonParams,
		0, // tokenAmount → TRC-10 用，不需要
		0, // tokenId    → TRC-10 用，不需要
	)
	assert.NoError(t, err, "EstimateEnergy should not return error")
	assert.NotNil(t, result, "TransactionExtention should not be nil")

	if result.Result != nil {
		assert.Equal(t, tronpb.Return_SUCCESS, result.Result.Code,
			fmt.Sprintf("estimate failed: %s", string(result.Result.Message)))
	}

	// 输出能量估算结果
	energyRequired := result.GetEnergyRequired()
	t.Logf("Estimated Energy Required: %d", energyRequired)

	if energyRequired > 0 {
		t.Logf("预计本次 USDT 转账需要消耗约 %d Energy", energyRequired)
		t.Logf("（当前主网 1 Energy ≈ 280 SUN，根据网络拥堵程度会有浮动）")
	} else {
		t.Log("EnergyRequired 为 0，可能调用失败或返回异常，请检查日志")
	}

	// 完整响应（调试用）
	j, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("EstimateEnergy full response:\n%s", string(j))
}
