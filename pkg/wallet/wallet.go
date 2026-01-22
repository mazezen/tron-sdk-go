package wallet

import (
	"fmt"
	"path"

	bip39 "github.com/mazezen/go-bip39"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/mazezen/tron-sdk-go/pkg/keystore"
	mn "github.com/mazezen/tron-sdk-go/pkg/mnemonic"
)

// CreateWallet from mnemonic and save keystore dir
func CreateWallet(
	keystoreDir string, // your keystore path，example: ~/.tronctl/keystore
	password string, // user input password for encrypt keystore
	mnemonic string, // if mnemonic is empty, auto generate mnemonic of 12/24
	passphrase string, // Additional password for mnemonic words, usually empty
	wordCount int, // How many words (12 or 24) are generated if mnemonic is empty
	accountIndex int, // usually 0
) (string, string, string, error) {
	var finalMnemonic string
	if mnemonic != "" {
		// auto generate mnemonic, default 12
		if wordCount != 12 && wordCount != 24 {
			return "", "", "", fmt.Errorf("wordCount must be 12 or 24")
		}
		if wordCount == 12 {
			finalMnemonic = mn.Generate12()
		} else {
			finalMnemonic = mn.Generate24()
		}
	} else {
		if bip39.IsMnemonicValid(mnemonic) {
			return "", "", "", fmt.Errorf("invalid mnemonic")
		}
		finalMnemonic = mnemonic
	}

	// derive private key
	priv, pub := keys.FromMnemonicSeedAndPassphrase(finalMnemonic, passphrase, accountIndex)
	if priv == nil || pub == nil {
		return "", "", "", fmt.Errorf("failed to derive private/public key")
	}

	// derive TRON base58 address
	if keystoreDir == "" {
		keystoreDir = common.DefaultConfigDirName
	}
	addr := address.PubkeyToAddress(*pub.ToECDSA()).String()

	// private key store keystore
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.ImportECDSA(priv.ToECDSA(), password) // ks.NewAccount(password)
	if err != nil {
		return "", "", "", err
	}

	// account url example: file://path
	keystoreFilePath := path.Base(account.URL.Path)

	return finalMnemonic, addr, keystoreFilePath, nil
}
