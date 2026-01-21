package mnemonic

import bip39 "github.com/mazezen/go-bip39"

// Generate12 with 12 words
func Generate12() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

// Generate24 with 24 words
func Generate24() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}
