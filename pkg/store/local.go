package store

import (
	"fmt"
	"os"
	"path"

	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/keystore"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

func InitConfigDir() {
	uDir, _ := homedir.Dir()
	tronCTLDir := path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		err = os.MkdirAll(tronCTLDir, 0700)
		if err != nil {
			fmt.Printf("create keystore dir error: %v\n", err)
		}
	}
}

// LocalAccounts returns a slice of local account alias names
func LocalAccounts() []string {
	uDir, _ := homedir.Dir()
	files, _ := os.ReadDir(path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName))
	var accounts []string
	for _, f := range files {
		if f.IsDir() {
			accounts = append(accounts, path.Base(f.Name()))
		}
	}
	return accounts
}

var (
	describe = fmt.Sprintf("%-24s\t\t%23\n", "NAME", "ADDRESS")
	// ErrNoUnlockBadPassphrase for bad password
	ErrNoUnlockBadPassphrase = fmt.Errorf("could not unlock account with passphrase, perhaps need different phrase")
)

// DescribeLocalAccounts will display all the account alias name and their corresponding one address
func DescribeLocalAccounts() {
	for _, name := range LocalAccounts() {
		ks := FromAccountName(name)
		allAccounts := ks.Accounts()
		for _, account := range allAccounts {
			fmt.Printf("%-48s\t%s\n", name, account.Address)
		}
	}
}

// DoesNamedAccountExist return true if the given string name is an alias account already define,
// and return false otherwise
func DoesNamedAccountExist(name string) bool {
	for _, account := range LocalAccounts() {
		if account == name {
			return true
		}
	}
	return false
}

// AddressFromAccountName Returns address for account name if exists
func AddressFromAccountName(name string) (string, error) {
	ks := FromAccountName(name)
	// FIXME: Assume 1 account per keystore for now
	for _, account := range ks.Accounts() {
		return account.Address.String(), nil
	}
	return "", fmt.Errorf("keystore not found")
}

func FromAddress(addr string) *keystore.KeyStore {
	for _, account := range LocalAccounts() {
		ks := FromAccountName(account)
		allAccounts := ks.Accounts()
		for _, a := range allAccounts {
			if addr == a.Address.String() {
				return ks
			}
		}
	}
	return nil
}

// FromAccountName get account from name
func FromAccountName(name string) *keystore.KeyStore {
	uDir, _ := homedir.Dir()
	p := path.Join(uDir, common.DefaultConfigDirName, name)
	return keystore.ForPath(p)
}

// DefaultLocation get default location
func DefaultLocation() string {
	uDir, _ := homedir.Dir()
	return path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
}

// SetDefaultLocation set default location
func SetDefaultLocation(directory string) error {
	common.DefaultConfigDirName = directory
	uDir, err := homedir.Dir()
	if err != nil {
		return err
	}
	tronCTLDir := path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
	if _, err = os.Stat(tronCTLDir); os.IsNotExist(err) {
		err = os.MkdirAll(tronCTLDir, 0700)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnlockedKeystore return keystore unlocked
func UnlockedKeystore(from, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	sender, err := address.Base58ToAddress(from)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid address: %s", from)
	}
	ks := FromAddress(from)
	if ks == nil {
		return nil, nil, fmt.Errorf("could not open local keystore for %s", from)
	}
	account, lookupErr := ks.Find(keystore.Account{Address: sender})
	if lookupErr != nil {
		return nil, nil, fmt.Errorf("could not find %s in keystore", from)
	}
	if unlockError := ks.Unlock(account, passphrase); unlockError != nil {
		return nil, nil, errors.Wrap(ErrNoUnlockBadPassphrase, unlockError.Error())
	}
	return ks, &account, nil
}
