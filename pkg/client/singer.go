package client

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

func SignTransaction(tx *tronpb.Transaction, signer *btcec.PrivateKey) (*tronpb.Transaction, error) {
	return SignTransactionECDSA(tx, signer.ToECDSA())
}

func SignTransactionECDSA(tx *tronpb.Transaction, signer *ecdsa.PrivateKey) (*tronpb.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}

	s256 := sha256.New()
	s256.Write(rawData)
	hash := s256.Sum(nil)

	signature, err := crypto.Sign(hash, signer)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}
