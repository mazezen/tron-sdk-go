package client

import (
	"fmt"
	"math/big"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/address"
)

func makePermission(name string, pType tronpb.Permission_PermissionType, id int32,
	threshold int64, operations map[string]bool, keys map[string]int64) (*tronpb.Permission, error) {

	pKey := make([]*tronpb.Key, 0)

	if len(keys) > 5 {
		return nil, fmt.Errorf("cant have more than 5 keys")
	}
	totalWeight := int64(0)
	for k, w := range keys {
		totalWeight += w
		addr, err := address.Base58ToAddress(k)
		if err != nil {
			return nil, fmt.Errorf("invalid address: %s", k)
		}
		pKey = append(pKey, &tronpb.Key{
			Address: addr,
			Weight:  w,
		})
	}
	var bigOP *big.Int
	if len(operations) > 0 {
		bigOP = big.NewInt(0)
		for k, o := range operations {
			if o {
				value, b := tronpb.Transaction_Contract_ContractType_value[k]
				if !b {
					return nil, fmt.Errorf("permission not found: %s", k)
				}
				bigOP.SetBit(bigOP, int(value), 1)
			}
		}
	} else {
		bigOP = nil
	}

	if threshold > totalWeight {
		return nil, fmt.Errorf("invalid key/threshold size (%d/%d)", threshold, totalWeight)
	}
	var bOP []byte
	if bigOP != nil {
		bOP = make([]byte, 32)
		l := len(bigOP.Bytes()) - 1
		for i, b := range bigOP.Bytes() {
			bOP[l-i] = b
		}
	}

	return &tronpb.Permission{
		Type:           pType,
		Id:             id,
		PermissionName: name,
		Threshold:      threshold,
		Operations:     bOP,
		Keys:           pKey,
	}, nil
}
