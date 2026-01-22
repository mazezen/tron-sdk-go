package keys

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	bip39 "github.com/mazezen/go-bip39"
	"github.com/mazezen/tron-sdk-go/pkg/keys/hd"
)

func FromMnemonicSeedAndPassphrase(mnemonic, passphrase string, index int) (*btcec.PrivateKey, *btcec.PublicKey) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	//fmt.Printf("[DEBUG] mnemonic   : %s\n", mnemonic)
	//fmt.Printf("[DEBUG] passphrase : '%s'\n", passphrase)
	//fmt.Printf("[DEBUG] seed (64 bytes hex): %x\n", seed)

	//if len(seed) != 64 {
	//	fmt.Printf("[ERROR] seed length wrong! got %d bytes\n", len(seed))
	//}

	master, ch := hd.ComputeMastersFromSeed(seed, []byte("Bitcoin seed"))
	//fmt.Printf("[DEBUG] master priv (32 bytes hex): %x\n", master)
	//fmt.Printf("[DEBUG] chain code  (32 bytes hex): %x\n", ch)

	path := fmt.Sprintf("m/44'/195'/0'/0/%d", index)
	//fmt.Printf("[DEBUG] deriving path: %s\n", path)
	private, err := hd.DerivePrivateKeyForPath(
		btcec.S256(),
		master,
		ch,
		path,
	)
	if err != nil {
		//fmt.Printf("[ERROR] DerivePrivateKeyForPath failed: %v\n", err)
		return nil, nil
	}
	//fmt.Printf("[DEBUG] derived priv (32 bytes hex): %x\n", private)
	return btcec.PrivKeyFromBytes(private[:])
}
