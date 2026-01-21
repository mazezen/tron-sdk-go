package client

import (
	"fmt"

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

	acc, err := c.WalletClient.GetAccount(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}

	return acc, nil
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

	transaction, err := c.WalletClient.SetAccountId(ctx, setAccountIdContract)
	if err != nil {
		return nil, fmt.Errorf("SetAccountId: %w", err)
	}

	return transaction, nil
}

// GetAccountById from base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
func (c *GrpcClient) GetAccountById(id string) (*tronpb.Account, error) {
	var account = &tronpb.Account{}
	var err error

	account.AccountId = []byte(id)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	acc, err := c.WalletClient.GetAccountById(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("GetAccountById: %w", err)
	}
	return acc, nil
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
