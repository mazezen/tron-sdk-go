package account

import (
	"fmt"
	"regexp"

	"github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/mazezen/tron-sdk-go/pkg/mnemonic"
	"github.com/mazezen/tron-sdk-go/pkg/store"
)

// Creation struct for account
type Creation struct {
	Name               string
	Passphrase         string
	Mnemonic           string
	MnemonicPassphrase string
	HdAccountNumber    *uint32
	HdIndexNumber      *uint32
}

// New create new name
func New() string {
	return "New Account"
}

// CheckPassphraseStrong check passphrase if wrong
func CheckPassphraseStrong(passphrase string) error {
	if len(passphrase) < 8 {
		return fmt.Errorf("passphrase too short")
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(passphrase)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(passphrase)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\./?]`).MatchString(passphrase)

	if !hasNumber || !hasUpper || !hasSpecial {
		return fmt.Errorf("passphrase must contain numbers and uppercase letters, and special characters")
	}
	return nil
}

// CreateNewLocalAccount create a new local account
func CreateNewLocalAccount(candidate *Creation) error {
	if err := CheckPassphraseStrong(candidate.Passphrase); err != nil {
		return err
	}
	ks := store.FromAccountName(candidate.Name)

	if candidate.Mnemonic == "" {
		candidate.Mnemonic = mnemonic.Generate24()
	}
	// Hardcoded index of 0 for brandnew account.
	private, _ := keys.FromMnemonicSeedAndPassphrase(candidate.Mnemonic, candidate.MnemonicPassphrase, 0)
	_, err := ks.ImportECDSA(private.ToECDSA(), candidate.Passphrase)
	if err != nil {
		return err
	}
	return nil
}
