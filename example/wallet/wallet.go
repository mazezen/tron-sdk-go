package main

import (
	"fmt"
	"path"

	"github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/mazezen/tron-sdk-go/pkg/wallet"
)

// 示例：main 函数 或 CLI 命令实现
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
