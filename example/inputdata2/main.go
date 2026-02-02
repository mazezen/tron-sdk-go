package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mazezen/tron-sdk-go/pkg/abi"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 无 0x 前缀
	inputHex := "23b872dd000000000000000000000000322e657a3f1401994efa0b612e5c1af0c40fb3e4000000000000000000000000521aefb920d5562ce96ecee0210867ca95d3c40700000000000000000000000000000000000000000000000000000000076bbff0"

	inputBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 取方法签名（前4字节）
	methodSelector := inputBytes[:4]
	fmt.Printf("method id: %s\n", hex.EncodeToString(methodSelector))

	// 获取contractABI
	c := client.NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = c.Start(dialOptions...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Stop()

	var value = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	smartContract, err := c.GetContract(context.Background(), value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("smart contract ABI: %s\n", smartContract.Abi)

	inputs, err := abi.GetInputsParser(smartContract.GetAbi(), "transferFrom")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解码参数（从第5字节开始）
	unpacked, err := inputs.Unpack(inputBytes[4:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(unpacked) != 3 {
		fmt.Println(fmt.Errorf("unexpected param count"))
		return
	}
	from := address.HexToBase58Address(unpacked[0].(common.Address).Hex())
	to := address.HexToBase58Address(unpacked[1].(common.Address).Hex())
	v2 := unpacked[2].(*big.Int)

	fmt.Printf("transferFrom(\n  from:  %s,\n  to:    %s,\n  value: %s\n)\n", from, to, v2.String())
	return
}
