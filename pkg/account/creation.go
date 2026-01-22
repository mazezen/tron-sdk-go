package account

import (
	"fmt"
	"regexp"
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
func CreateNewLocalAccount() error {

	return nil
}
