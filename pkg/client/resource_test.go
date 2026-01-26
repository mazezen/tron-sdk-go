package client

import (
	"math/big"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/mazezen/tron-sdk-go/pkg/common/decimal"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetDelegatedResourceAccountIndex(t *testing.T) {
	valueAddress, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse address")
	//t.Logf("hex address: %s", valueAddress.Hex())
	//t.Logf("eth address: %s", valueAddress.EthHex())

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")

	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "get delegated resource account index - b58",
			value: "T...",
		},
		{
			name:  "get delegated resource account index - hex address",
			value: valueAddress.Hex(),
		},
		{
			name:  "get delegated resource account index - eth address",
			value: valueAddress.EthHex(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountIndex, err := client.GetDelegatedResourceAccountIndex(tt.value)
			assert.NoError(t, err, "failed to get delegated resource account index")
			assert.NotNil(t, accountIndex, "failed to get delegated resource account index")

			for _, fromA := range accountIndex.GetFromAccounts() {
				t.Logf("from account: %s", common.BytesToHexString(fromA))
			}
			t.Log(" ------------------------------------------------------------------------- ")
			for _, toA := range accountIndex.GetToAccounts() {
				t.Logf("to account: %s", common.BytesToHexString(toA))
			}
		})
	}
}

func TestGrpcClient_GetDelegatedResourceAccountIndex2(t *testing.T) {
	valueAddress, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse address")
	//t.Logf("hex address: %s", valueAddress.Hex())
	//t.Logf("eth address: %s", valueAddress.EthHex())

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")

	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "get delegated resource account index - b58",
			value: "T...",
		},
		{
			name:  "get delegated resource account index - hex address",
			value: valueAddress.Hex(),
		},
		{
			name:  "get delegated resource account index - eth address",
			value: valueAddress.EthHex(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountIndex, err := client.GetDelegatedResourceAccountIndexV2(tt.value)
			assert.NoError(t, err, "failed to get delegated resource account index")
			assert.NotNil(t, accountIndex, "failed to get delegated resource account index")

			for _, fromA := range accountIndex.GetFromAccounts() {
				t.Logf("from account: %s", address.HexToBase58Address(common.BytesToHexString(fromA)))
			}
			t.Log(" ------------------------------------------------------------------------- ")
			for _, toA := range accountIndex.GetToAccounts() {
				t.Logf("to account: %s", address.HexToAddress(common.BytesToHexString(toA)))
			}
		})
	}
}

func TestGrpcClient_GetDelegatedResourceByGivenTo(t *testing.T) {
	addressFrom, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse from b58 address")
	//t.Logf("from hex is: %s", addressFrom.Hex())
	//t.Logf("from eth hex: %s", addressFrom.EthHex())

	addressTo, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse to b58 address")
	//t.Logf("to hex is: %s", addressTo.Hex())
	//t.Logf("to eth hex: %s", addressTo.EthHex())

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	tests := []struct {
		name string
		from string
		to   string
	}{
		{
			name: "get delegated resource - base58",
			from: "T...",
			to:   "T...",
		},
		{
			name: "get delegated resource - hex",
			from: addressFrom.Hex(),
			to:   addressTo.Hex(),
		},
		{
			name: "get delegated resource - eth address",
			from: addressTo.EthHex(),
			to:   addressFrom.EthHex(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resourceList, err := client.GetDelegatedResourceByGivenTo(test.from, test.to)
			assert.NoError(t, err, "failed get delegated resource")
			assert.NotNil(t, resourceList, "failed get delegated resource")

			for _, resource := range resourceList.DelegatedResource {
				t.Logf("resource from: %s", common.BytesToHexString(resource.From))
				t.Logf("resource to: %s", common.BytesToHexString(resource.To))
				t.Logf("resource expire time for bandwidth %d", resource.ExpireTimeForBandwidth)
				t.Logf("resource forzen balance for bandwidth: %d", resource.FrozenBalanceForBandwidth)

				t.Logf("resource expire time for energy %d", resource.ExpireTimeForEnergy)
				t.Logf("resource forzen balance for energy %d", resource.ExpireTimeForEnergy)

			}
		})
	}
}

func TestGrpcClient_GetDelegatedResourceV2ByGivenTo(t *testing.T) {
	addressFrom, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse from b58 address")

	addressTo, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse to b58 address")

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	tests := []struct {
		name string
		from string
		to   string
	}{
		{
			name: "get delegated resource - base58",
			from: "T...",
			to:   "T...",
		},
		{
			name: "get delegated resource - hex",
			from: addressFrom.Hex(),
			to:   addressTo.Hex(),
		},
		{
			name: "get delegated resource - eth address",
			from: addressFrom.EthHex(),
			to:   addressTo.EthHex(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resourceList, err := client.GetDelegatedResourceV2ByGivenTo(test.from, test.to)
			assert.NoError(t, err, "failed get delegated resource")
			assert.NotNil(t, resourceList, "failed get delegated resource")

			for _, resource := range resourceList.DelegatedResource {
				t.Logf("resource from: %s", address.HexToBase58Address(common.BytesToHexString(resource.From)))
				t.Logf("resource to: %s", address.HexToBase58Address(common.BytesToHexString(resource.To)))
				t.Logf("resource expire time for bandwidth %d", resource.ExpireTimeForBandwidth)
				t.Logf("resource forzen balance for bandwidth: %d", resource.FrozenBalanceForBandwidth)

				t.Logf("resource expire time for energy %d", resource.ExpireTimeForEnergy)
				t.Logf("resource forzen balance for energy %d", resource.ExpireTimeForEnergy)

			}
		})
	}
}

func TestGrpcClient_GetDelegatedResource(t *testing.T) {
	addressFrom, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse from b58 address")
	t.Logf("from hex is: %s", addressFrom.Hex())
	t.Logf("from eth hex: %s", addressFrom.EthHex())

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	tests := []struct {
		name string
		from string
		to   string
	}{
		{
			name: "get delegated resource - base58",
			from: "T...",
		},
		{
			name: "get delegated resource - hex",
			from: addressFrom.Hex(),
		},
		{
			name: "get delegated resource - eth address",
			from: addressFrom.EthHex(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resourceList, err := client.GetDelegatedResource(test.from)
			assert.NoError(t, err, "failed get delegated resource")
			assert.NotNil(t, resourceList, "failed get delegated resource")

			for _, resource := range resourceList {
				for _, delegatedResource := range resource.GetDelegatedResource() {
					t.Logf("resource from: %s", common.BytesToHexString(delegatedResource.From))
					t.Logf("resource to: %s", common.BytesToHexString(delegatedResource.To))
					t.Logf("resource expire time for bandwidth %d", delegatedResource.ExpireTimeForBandwidth)
					t.Logf("resource forzen balance for bandwidth: %d", delegatedResource.FrozenBalanceForBandwidth)

					t.Logf("resource expire time for energy %d", delegatedResource.ExpireTimeForEnergy)
					t.Logf("resource forzen balance for energy %d", delegatedResource.ExpireTimeForEnergy)
				}
			}
		})
	}
}

func TestGrpcClient_GetDelegatedResourceV2(t *testing.T) {
	addressFrom, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse from b58 address")
	//t.Logf("from hex is: %s", addressFrom.Hex())
	//t.Logf("from eth hex: %s", addressFrom.EthHex())

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	tests := []struct {
		name string
		from string
		to   string
	}{
		{
			name: "get delegated resource - base58",
			from: "T...",
		},
		{
			name: "get delegated resource - hex",
			from: addressFrom.Hex(),
		},
		{
			name: "get delegated resource - eth address",
			from: addressFrom.EthHex(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resourceList, err := client.GetDelegatedResourceV2(test.from)
			assert.NoError(t, err, "failed get delegated resource")
			assert.NotNil(t, resourceList, "failed get delegated resource")

			for _, resource := range resourceList {
				for _, delegatedResource := range resource.GetDelegatedResource() {
					t.Logf("resource from: %s", address.HexToBase58Address(common.BytesToHexString(delegatedResource.From)))
					t.Logf("resource to: %s", address.HexToBase58Address(common.BytesToHexString(delegatedResource.To)))
					t.Logf("resource expire time for bandwidth %d", delegatedResource.ExpireTimeForBandwidth)
					t.Logf("resource forzen balance for bandwidth: %d", delegatedResource.FrozenBalanceForBandwidth)

					t.Logf("resource expire time for energy %d", delegatedResource.ExpireTimeForEnergy)
					t.Logf("resource forzen balance for energy %d", delegatedResource.ExpireTimeForEnergy)
				}
			}
		})
	}
}

func TestGrpcClient_GetCanDelegatedMaxSize(t *testing.T) {
	addr, err := address.Base58ToAddress("T...")
	assert.NoError(t, err, "failed to parse from b58 address")

	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err = client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	tests := []struct {
		name         string
		addr         string
		resourceType int32
	}{
		{
			name:         "get delegated resource - base58",
			addr:         "T...",
			resourceType: 0,
		},
		{
			name:         "get delegated resource - hex",
			addr:         addr.Hex(),
			resourceType: 0,
		},
		{
			name:         "get delegated resource - eth address",
			addr:         addr.EthHex(),
			resourceType: 0,
		},
		{
			name:         "get delegated resource - base58",
			addr:         "T...",
			resourceType: 1,
		},
		{
			name:         "get delegated resource - hex",
			addr:         addr.Hex(),
			resourceType: 1,
		},
		{
			name:         "get delegated resource - eth address",
			addr:         addr.EthHex(),
			resourceType: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resouce, err := client.GetCanDelegatedMaxSize(test.addr, test.resourceType)
			assert.NoError(t, err, "failed get delegated resource")
			assert.NotNil(t, resouce, "failed get delegated resource")

			if test.resourceType == 0 {
				t.Logf("can delegated max size: %d : type: %s", resouce.MaxSize, "band width")
			}
			if test.resourceType == 1 {
				t.Logf("can delegated max size: %d : type: %s", resouce.MaxSize, "energy")
			}

		})
	}
}

func TestGrpcClient_DelegateResource(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	fromAddress, _ := address.Base58ToAddress("T...")
	fromPk := "" // for signature
	toAddress, _ := address.Base58ToAddress("T...")
	tests := []struct {
		name       string
		from       string
		to         string
		core       tronpb.ResourceCode
		balance    *big.Float
		lock       bool
		lockPeriod int64
	}{
		{
			name:       "delegate bandwidth resource - base58",
			from:       fromAddress.String(),
			to:         toAddress.String(),
			core:       tronpb.ResourceCode_BANDWIDTH,
			balance:    big.NewFloat(300.00),
			lock:       false,
			lockPeriod: 0,
		},
		{
			name:       "delegate energy resource - base58",
			from:       fromAddress.String(),
			to:         toAddress.String(),
			core:       tronpb.ResourceCode_ENERGY,
			balance:    big.NewFloat(10000.00),
			lock:       false,
			lockPeriod: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var balance int64
			if test.core == tronpb.ResourceCode_BANDWIDTH {
				resource, _ := client.GetAccountResource(test.to)
				exchangeRate := decimal.Div(big.NewFloat(float64(resource.TotalNetLimit)), big.NewFloat(float64(resource.TotalNetWeight)))
				exchangeRate = decimal.Sub(exchangeRate, big.NewFloat(0.06))
				// convert to sun
				trxToBandWidth := decimal.Div(test.balance, exchangeRate)
				balance, _ = decimal.Mul(trxToBandWidth, big.NewFloat(1_000_000)).Int64()
				t.Logf("balance: %d", balance)
			}

			if test.core == tronpb.ResourceCode_ENERGY {
				resource, err := client.GetAccountResource(test.to)
				assert.NoErrorf(t, err, "failed get [%s] account resource", test.to)
				assert.NotNil(t, resource, "failed get [%s] account resource", test.to)
				t.Logf("resource: %s", resource)
				exchangeRate := decimal.Div(big.NewFloat(float64(resource.TotalEnergyLimit)), big.NewFloat(float64(resource.TotalEnergyWeight)))
				exchangeRate = decimal.Sub(exchangeRate, big.NewFloat(0.01))
				t.Logf("exchangeRate: %0.6f", exchangeRate)

				balanceF := decimal.Div(test.balance, exchangeRate)
				t.Logf("balanceF: %0.6f", balanceF)

				// convert to sun
				balance, _ = decimal.Mul(balanceF, big.NewFloat(1_000_000)).Int64()
				t.Logf("balance: %d", balance)
			}

			txEx, err := client.DelegateResource(test.from, test.to, test.core, balance, test.lock, test.lockPeriod)
			assert.NoError(t, err, "failed delegate resource")
			assert.NotNil(t, txEx, "failed delegate resource")

			toBytes, _ := common.HexToByte(fromPk)
			btcecPk, _ := btcec.PrivKeyFromBytes(toBytes)
			addr := address.BTCECPrivkeyToAddress(btcecPk)
			//t.Logf("addr: %s", addr.String())
			if addr.String() != test.from {
				txEx.Transaction.RawData.Contract[0].PermissionId = 3
			}

			// signature
			ecdsa, err := crypto.HexToECDSA(fromPk)
			assert.NoError(t, err, "failed hex to ECDSA")
			tx, err := SignTransactionECDSA(txEx.Transaction, ecdsa)
			assert.NoError(t, err, "failed sign transaction")
			assert.NotNil(t, tx, "failed sign transaction")

			// broadcast
			res, err := client.BroadcastTransaction(tx)
			assert.NoError(t, err, "failed broadcast transaction")
			assert.NotNil(t, res, "failed broadcast transaction")

			t.Logf("delegate resource from %s to %s txid:%s", fromAddress.String(), toAddress.String(), common.BytesToHexString(txEx.Txid))
		})
	}
}

func TestGrpcClient_UnDelegateResource(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client.SetTimeout(10 * time.Second)
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed start grpc client")

	fromAddress, _ := address.Base58ToAddress("T...")
	fromPk := "" // for signature
	toAddress, _ := address.Base58ToAddress("T...")
	tests := []struct {
		name    string
		from    string
		to      string
		core    tronpb.ResourceCode
		balance *big.Float
	}{
		{
			name:    "unDelegate bandwidth resource - base58",
			from:    fromAddress.String(),
			to:      toAddress.String(),
			core:    tronpb.ResourceCode_BANDWIDTH,
			balance: big.NewFloat(300.00),
		},
		{
			name:    "unDelegate energy resource - base58",
			from:    fromAddress.String(),
			to:      toAddress.String(),
			core:    tronpb.ResourceCode_ENERGY,
			balance: big.NewFloat(10000.00),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var balance int64
			if test.core == tronpb.ResourceCode_BANDWIDTH {
				// convert to sun
				resource, _ := client.GetAccountResource(test.to)
				exchangeRate := decimal.Div(big.NewFloat(float64(resource.TotalNetLimit)), big.NewFloat(float64(resource.TotalNetWeight)))
				exchangeRate = decimal.Sub(exchangeRate, big.NewFloat(0.06))
				trxToBandWidth := decimal.Div(test.balance, exchangeRate)
				balance, _ = decimal.Mul(trxToBandWidth, big.NewFloat(1_000_000)).Int64()
				t.Logf("balance: %d", balance)
			}
			if test.core == tronpb.ResourceCode_ENERGY {
				resource, err := client.GetAccountResource(test.to)
				assert.NoErrorf(t, err, "failed get [%s] account resource", test.to)
				assert.NotNil(t, resource, "failed get [%s] account resource", test.to)
				t.Logf("resource: %s", resource)
				exchangeRate := decimal.Div(big.NewFloat(float64(resource.TotalEnergyLimit)), big.NewFloat(float64(resource.TotalEnergyWeight)))
				exchangeRate = decimal.Sub(exchangeRate, big.NewFloat(0.01))
				t.Logf("exchangeRate: %0.6f", exchangeRate)

				balanceF := decimal.Div(test.balance, exchangeRate)
				t.Logf("balanceF: %0.6f", balanceF)

				// convert to sun
				balance, _ = decimal.Mul(balanceF, big.NewFloat(1_000_000)).Int64()
				t.Logf("balance: %d", balance)
			}

			txEx, err := client.UnDelegateResource(test.from, test.to, test.core, balance)
			assert.NoError(t, err, "failed delegate resource")
			assert.NotNil(t, txEx, "failed delegate resource")

			toBytes, _ := common.HexToByte(fromPk)
			btcecPk, _ := btcec.PrivKeyFromBytes(toBytes)
			addr := address.BTCECPrivkeyToAddress(btcecPk)
			//t.Logf("addr: %s", addr.String())
			if addr.String() != test.from {
				txEx.Transaction.RawData.Contract[0].PermissionId = 3
			}

			// signature
			ecdsa, err := crypto.HexToECDSA(fromPk)
			assert.NoError(t, err, "failed hex to ECDSA")
			tx, err := SignTransactionECDSA(txEx.Transaction, ecdsa)
			assert.NoError(t, err, "failed sign transaction")
			assert.NotNil(t, tx, "failed sign transaction")

			// broadcast
			res, err := client.BroadcastTransaction(tx)
			assert.NoError(t, err, "failed broadcast transaction")
			assert.NotNil(t, res, "failed broadcast transaction")

			t.Logf("delegate resource from %s to %s txid:%s", fromAddress.String(), toAddress.String(), common.BytesToHexString(txEx.Txid))
		})
	}
}
