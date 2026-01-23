package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// GetAccount from base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// Query information about an account, including TRX balance,
// TRC-10 balances, stake information and vote information and permissions etc. (Confirmed state)
// https://developers.tron.network/reference/walletsolidity-getaccount
func (c *GrpcClient) GetAccount(addr string) (*tronpb.Account, error) {
	var account = &tronpb.Account{}
	var err error

	account.Address, err = c.convert(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetAccount(ctx, account)
}

// SetAccountId set account id if the account has no id. Account id is unique and case insensitive.
func (c *GrpcClient) SetAccountId(id, addr string) (*tronpb.Transaction, error) {
	var setAccountIdContract = &tronpb.SetAccountIdContract{}
	var err error

	setAccountIdContract.AccountId = []byte(id)
	setAccountIdContract.OwnerAddress, err = c.convert(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.SetAccountId(ctx, setAccountIdContract)
}

// GetAccountById from id
func (c *GrpcClient) GetAccountById(id string) (*tronpb.Account, error) {
	var account = &tronpb.Account{}

	account.AccountId = []byte(id)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetAccountById(ctx, account)
}

// GetAccountBalance Query the historical balance of an account at a specific block
// addr: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// hash: block hash
// number: block number int64
// https://developers.tron.network/reference/getaccountbalance
func (c *GrpcClient) GetAccountBalance(addr, hash string, number int64) (*tronpb.AccountBalanceResponse, error) {
	req := &tronpb.AccountBalanceRequest{
		AccountIdentifier: &tronpb.AccountIdentifier{},
		BlockIdentifier:   &tronpb.BlockBalanceTrace_BlockIdentifier{},
	}
	var err error

	a, err := c.convert(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid addrress: %w", err)
	}

	hexToByte, err := common.TronHexToByte(hash)
	if err != nil {
		return nil, fmt.Errorf("invalid hash: %w", err)
	}

	req.AccountIdentifier.Address = a
	req.BlockIdentifier.Hash = hexToByte
	req.BlockIdentifier.Number = number

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetAccountBalance(ctx, req)
}

// UpdateAccount change account name. Account name is not unique now
// addr: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// name: string
// Please use UpdateAccount2 instead of this function.
func (c *GrpcClient) UpdateAccount(addr, name string) (*tronpb.Transaction, error) {
	req := new(tronpb.AccountUpdateContract)
	var err error
	req.OwnerAddress, err = c.convert(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid addrress: %w", err)
	}
	req.AccountName = []byte(name)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	transaction, err := c.WalletClient.UpdateAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccount: %w", err)
	}

	if proto.Size(transaction) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	return transaction, nil
}

// UpdateAccount2 change account name. Account name is not unique now
// addr: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// name: string
// Use this function instead of UpdateAccount.
// https://developers.tron.network/reference/updateaccount
func (c *GrpcClient) UpdateAccount2(addr, name string) (*tronpb.TransactionExtention, error) {
	req := new(tronpb.AccountUpdateContract)
	var err error
	req.OwnerAddress, err = c.convert(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid addrress: %w", err)
	}
	req.AccountName = []byte(name)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	transaction, err := c.WalletClient.UpdateAccount2(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccount2: %w", err)
	}

	if proto.Size(transaction) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if transaction.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", transaction.GetResult().GetMessage())
	}

	return transaction, nil
}

func (c *GrpcClient) convert(addr string) (address.Address, error) {
	if len(addr) == 0 {
		return address.Address{}, fmt.Errorf("address is empty")
	}

	if common.HasTPrefix(addr) {
		return common.DecodeCheck(addr)
	} else {
		if common.Has0xPrefix(addr) {
			return common.EthHexToByte(addr)
		} else if common.Has41Prefix(addr) {
			return common.TronHexToByte(addr)
		}
	}

	return address.Address{}, fmt.Errorf("address is invalid")
}
