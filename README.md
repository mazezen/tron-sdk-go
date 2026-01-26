# tron-sdk-go
[![Go](https://img.shields.io/badge/Go-1.24.1-00ADD8.svg?logo=go&logoColor=white)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/mazezen/tron-sdk-go)](https://goreportcard.com/report/github.com/mazezen/tron-sdk-go)
[![License: LGPL v3](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0)
> tron-sdk-go 是一个的Go SDK用于与TRON区块链进行交互.

## 使用
```shell
go get github.com/mazezen/tron-sdk-go
```

## 生成助记词 > 私钥 > 地址
```go
import (
    "fmt"
    
    "github.com/mazezen/tron-sdk-go/pkg/address"
    "github.com/mazezen/tron-sdk-go/pkg/common"
    "github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/mazezen/tron-sdk-go/pkg/mnemonic"
)

func main() {
    // mnemonic.Generate12() // 生成12位助记词 
	mn := mnemonic.Generate24() // 生成24位助记词
    fmt.Printf("Generated mnemonic: %s\n", mn)
    
    priv, pub := keys.FromMnemonicSeedAndPassphrase(mn, "", 0)
    
    privBytes := priv.Serialize()
    privHex := common.BytesToHexString(privBytes)
    
    addr := address.PubkeyToAddress(*pub.ToECDSA()).String()
    
    fmt.Println("私钥 (hex):", privHex)
    fmt.Println("TRON 地址 (Base58):", addr)
}
```

## 根据助记词创建钱包
```go
package main

import (
	"fmt"
	"path"

	"github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/mazezen/tron-sdk-go/pkg/wallet"
)

// 示例：main 函数
func main() {
	keystoreDir := keys.CheckAndMakeKeyDirIfNeeded()

	mnemonic, addr, filePath, err := wallet.CreateWallet(
		keystoreDir,
		"123456",
		"", // 留空 = 自动生成
		"", // passphrase 通常空
		12, // 或 24
		0,  // 第一个账户
	)
	if err != nil {
		fmt.Printf("创建钱包失败: %v\n", err)
		return
	}

	fmt.Println("钱包创建成功！")
	fmt.Println("助记词（请务必备份，丢失无法找回）：")
	fmt.Println(mnemonic)
	fmt.Println("\nTRON 地址：")
	fmt.Println(addr)
	fmt.Println("\nKeystore 文件已保存到：")
	fmt.Println(path.Join(keystoreDir, filePath))
	fmt.Println("\n请使用密码解锁此文件进行签名交易。")
}

```

## 安装protoc依赖 protoc-gen-go && protoc-gen-go-grpc
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## protobuf 依赖
* https://github.com/tronprotocol/protocol
* https://github.com/googleapis/googleapis
