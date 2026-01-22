package mnemonic

import (
	"fmt"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/keys"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name         string
		mnemonic     string
		passphrase   string
		accountIndex int
	}{
		{
			name:         "standard 12 mnemonic - index 0 - no passphrase",
			mnemonic:     Generate12(),
			passphrase:   "",
			accountIndex: 0,
		},
		{
			name:         "standard 24 mnemonic - index 0 - no passphrase",
			mnemonic:     Generate24(),
			passphrase:   "",
			accountIndex: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mnemonic := tt.mnemonic
			t.Logf("Generated mnemonic: %s", mnemonic)

			priv, pub := keys.FromMnemonicSeedAndPassphrase(mnemonic, tt.passphrase, tt.accountIndex)
			require.NotNil(t, priv, "private key should not be nil")
			require.NotNil(t, pub, "public key should not be nil")

			privBytes := priv.Serialize()
			privHex := common.BytesToHexString(privBytes)

			addr := address.PubkeyToAddress(*pub.ToECDSA()).String()

			fmt.Println("私钥 (hex):", privHex)
			fmt.Println("TRON 地址 (Base58):", addr)
		})
	}
}
