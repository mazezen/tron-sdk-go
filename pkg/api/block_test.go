package api

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	abi2 "github.com/mazezen/tron-sdk-go/pkg/abi"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/stretchr/testify/assert"
)

func TestApiClient_GetBlock(t *testing.T) {
	result, err := NewApiClient("").GetBlock(context.Background(), "0000000004bf3832b41b5982fa8102e09f3bfcdd9259a44a81e952440e0bc9ae", true)
	assert.NoError(t, err, err)
	assert.NotNil(t, result, "GetBlock should not return nil")
	assert.NotNil(t, result.BlockHeader.RawData, "GetBlock raw data should not be nil")
	assert.NotEqual(t, 0, len(result.Transactions), "GetBlock transaction should not be 0")

	t.Logf(" ---------------------------------------------------------------------------------------------------- ")
	t.Logf("block hash: %s", result.BlockID)
	t.Logf("block number: %d", result.BlockHeader.RawData.Number)
	t.Logf("block time: %v", time.Unix(result.BlockHeader.RawData.Timestamp/1000, (result.BlockHeader.RawData.Timestamp%1000)*1000000).Format(time.DateTime))
	t.Logf(" ---------------------------------------------------------------------------------------------------- ")

	for _, transaction := range result.Transactions {
		t.Logf(" ************************************************************************************************ ")
		t.Logf("ret status :%s", transaction.Ret[0].ContractRet)
		t.Logf("tx id: %s", transaction.TxID)
		t.Logf("type: %s", transaction.RawData.Contract[0].Type)
		if transaction.RawData.Contract[0].Type == "UnDelegateResourceContract" {
			t.Logf("from: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.OwnerAddress).String())
			t.Logf("to: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.ReceiverAddress).String())
			t.Logf("balance: %d", transaction.RawData.Contract[0].Parameter.Value.Balance)
			t.Logf("resource: %s", transaction.RawData.Contract[0].Parameter.Value.Resource)
		}

		if transaction.RawData.Contract[0].Type == "DelegateResourceContract" {
			t.Logf("from: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.OwnerAddress).String())
			t.Logf("to: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.ReceiverAddress).String())
			t.Logf("balance: %d", transaction.RawData.Contract[0].Parameter.Value.Balance)
			t.Logf("resource: %s", transaction.RawData.Contract[0].Parameter.Value.Resource)
		}

		if transaction.RawData.Contract[0].Type == "TriggerSmartContract" {
			t.Logf("合约地址: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.ContractAddress).String())
			t.Logf("data: %s", transaction.RawData.Contract[0].Parameter.Value.Data)
			inputBytes, err := hex.DecodeString(transaction.RawData.Contract[0].Parameter.Value.Data)
			if err != nil {
				t.Logf("hex decode failed: %v", err)
				os.Exit(1)
			}
			// TRC-20 / ERC-20 标准 transfe	rFrom 方法的 ABI 定义（只需这一段即可）

			// 解析ABI
			contractABI, err := abi.JSON(strings.NewReader(abi2.Trc20AbiFragment))
			if err != nil {
				t.Logf("abi json parse failed: %v", err)
				os.Exit(1)
			}

			// 根据前 4 字节找到方法
			method, err := contractABI.MethodById(inputBytes[:4])
			if err != nil {
				t.Logf("Method not found: %v", err)
				os.Exit(1)
			}

			// 解码参数（从第 5 个字节开始）
			args, err := method.Inputs.Unpack(inputBytes[4:])
			if err != nil {
				t.Logf("Unpack failed: %v", err)
				os.Exit(1)
			}
			switch method.Name {
			case "transfer":
				if len(args) != 2 {
					continue
				}
				to := args[0].(common.Address).Hex()
				value := args[1].(*big.Int)
				t.Logf("Transfer → To: %s", address.HexToBase58Address(to))
				t.Logf("Amount:   %s", value)
			case "transferFrom":
				if len(args) != 3 {
					continue
				}
				from := args[0].(common.Address).Hex()
				to := args[1].(common.Address).Hex()
				value := args[2].(*big.Int)
				t.Logf("TransferFrom → From: %s", address.HexToBase58Address(from))
				t.Logf("To:               %s", address.HexToBase58Address(to))
				t.Logf("Amount:           %s", value)
			case "approve":
				if len(args) != 2 {
					continue
				}
				spender := args[0].(common.Address).Hex()
				value := args[1].(*big.Int)
				t.Logf("Approve → Spender: %s", address.HexToBase58Address(spender))
				t.Logf("Amount:           %s", value)

			case "increaseAllowance", "decreaseAllowance":
				// 类似 approve 处理

			case "balanceOf", "allowance", "totalSupply", "name", "symbol", "decimals":
				t.Logf("View function call: %s (通常不带 value，但偶尔出现)", method.Name)

			default:
				t.Logf("Unhandled method: %s", method.Name)
			}
			//// args 是一个 []interface{}，按顺序对应 inputs
			//if len(args) != 3 {
			//	t.Fatal("参数数量不匹配")
			//}
			//from := args[0].(common.Address).Hex()
			//to := args[1].(common.Address).Hex()
			//value := args[2].(*big.Int)
			//
			//t.Logf("From:  %s\n", address.HexToBase58Address(from))
			//t.Logf("To:    %s\n", address.HexToBase58Address(to))
			//t.Logf("Value: %s\n", value.String())  // 原始数值
			//t.Logf("Value (decimal): %d\n", value) // 十进制显示

		}

		if transaction.RawData.Contract[0].Type == "TransferContract" {
			t.Logf("from: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.OwnerAddress).String())
			t.Logf("to: %s", address.HexToAddress(transaction.RawData.Contract[0].Parameter.Value.ToAddress).String())
			t.Logf("amount: %d", transaction.RawData.Contract[0].Parameter.Value.Amount)
		}

		t.Logf(" ************************************************************************************************ ")
	}

}
