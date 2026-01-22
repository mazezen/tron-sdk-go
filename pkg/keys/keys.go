package keys

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/keystore"
	"github.com/mitchellh/go-homedir"
)

func CheckAndMakeKeyDirIfNeeded() string {
	userDir, _ := homedir.Dir()
	tronCTLDir := path.Join(userDir, ".tronctl", "keystore")
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		// Double check with Leo what is right file persmission
		err = os.Mkdir(tronCTLDir, 0700)
		if err != nil {
			return fmt.Sprintf("create keystore dir error: %v\n", err)
		}
	}
	return tronCTLDir
}

func ListKeys() {
	tronCTLDir := CheckAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	// keystore.Keystore
	allAccounts := ks.Accounts()
	fmt.Printf("Tron Address: %s File URL:\n", strings.Repeat(" ", address.AddressLengthBase58))
	for _, account := range allAccounts {
		fmt.Printf("%s\t\t %s\n", account.Address, account.URL)
	}
}

func AddNewKey(password string) {
	tronCTLDir := CheckAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("NewAccount error: %v\n", err)
	}
	fmt.Printf("New Account created account: %s\n", account.Address)
	fmt.Printf("New Account created URL: %s\n", account.URL)
}

func GetPrivateKeyFromBytes(privateKeyBytes []byte) (*btcec.PrivateKey, error) {
	if len(privateKeyBytes) != 32 {
		return nil, fmt.Errorf("private key must be 32 bytes long. length: %d", len(privateKeyBytes))
	}

	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, common.ErrBadKeyLength
	}

	// btcec.PrivKeyFromBytes only returns a secret key and public key
	private, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	if private == nil {
		return nil, fmt.Errorf("failed to create private key from bytes: %v", privateKeyBytes)
	}
	return private, nil
}
