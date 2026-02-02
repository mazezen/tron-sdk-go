package api

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
)

func SingTransactionECDSA(tx *internet.Transaction, signer *ecdsa.PrivateKey) (*internet.Transaction, error) {
	if tx.RawDataHex == "" {
		return nil, fmt.Errorf("raw_data_hex is empty, must use the value from /wallet/createtransaction")
	}

	rawBytes, err := hex.DecodeString(tx.RawDataHex)
	if err != nil {
		return nil, fmt.Errorf("decode raw_data_hex failed: %w", err)
	}

	hash := sha256.Sum256(rawBytes)

	signature, err := crypto.Sign(hash[:], signer)
	if err != nil {
		return nil, fmt.Errorf("sign failed: %w", err)
	}

	sigHex := hex.EncodeToString(signature)
	tx.Signature = append(tx.Signature, sigHex)

	return tx, nil
}
