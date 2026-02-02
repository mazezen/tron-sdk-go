package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mazezen/tron-sdk-go/pkg/address"
)

func main() {

	// 无 0x 前缀
	inputHex := "23b872dd000000000000000000000000322e657a3f1401994efa0b612e5c1af0c40fb3e4000000000000000000000000521aefb920d5562ce96ecee0210867ca95d3c40700000000000000000000000000000000000000000000000000000000076bbff0"

	inputBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		log.Fatalf("hex decode failed: %v", err)
		return
	}

	// TRC-20 / ERC-20 标准 transfe	rFrom 方法的 ABI 定义（只需这一段即可）
	abiJSON := `[
		{
			"constant": false,
			"inputs": [
				{"name": "from", "type": "address"},
				{"name": "to", "type": "address"},
				{"name": "value", "type": "uint256"}
			],
			"name": "transferFrom",
			"outputs": [{"name": "", "type": "bool"}],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`

	// 解析ABI
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatalf("abi json parse failed: %v", err)
		return
	}

	// 根据前 4 字节找到方法
	method, err := contractABI.MethodById(inputBytes[:4])
	if err != nil {
		log.Fatalf("Method not found: %v", err)
	}

	fmt.Printf("方法名称: %s\n", method.Name) // transferFrom

	// 解码参数（从第 5 个字节开始）
	args, err := method.Inputs.Unpack(inputBytes[4:])
	if err != nil {
		log.Fatalf("Unpack failed: %v", err)
	}

	// args 是一个 []interface{}，按顺序对应 inputs
	if len(args) != 3 {
		log.Fatal("参数数量不匹配")
	}

	from := args[0].(common.Address).Hex()
	to := args[1].(common.Address).Hex()
	value := args[2].(*big.Int)

	fmt.Printf("From:  %s\n", address.HexToBase58Address(from))
	fmt.Printf("To:    %s\n", address.HexToBase58Address(to))
	fmt.Printf("Value: %s\n", value.String())  // 原始数值
	fmt.Printf("Value (decimal): %d\n", value) // 十进制显示

}
