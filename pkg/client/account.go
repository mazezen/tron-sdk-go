package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/address"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"google.golang.org/protobuf/proto"
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

	hexToByte, err := common.HexToByte(hash)
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

// CreateAccount activate an account
// https://developers.tron.network/reference/account-createaccount
// Please use CreateAccount2 instead of this function.
func (c *GrpcClient) CreateAccount(from, to string, typ int) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.AccountCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.AccountAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.Type = tronpb.AccountType(typ)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.CreateAccount(ctx, in)
}

// CreateAccount2 activate an account
// https://developers.tron.network/reference/account-createaccount
// Use this function instead of CreateAccount.
func (c *GrpcClient) CreateAccount2(from, to string, typ int) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.AccountCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.AccountAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.Type = tronpb.AccountType(typ)
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.CreateAccount2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// WithdrawBalance super Representative or user withdraw rewards, usable every 24 hours.
// https://developers.tron.network/reference/withdrawbalance
// Please use WithdrawBalance2 instead of this function.
func (c *GrpcClient) WithdrawBalance(from string) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.WithdrawBalanceContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.WithdrawBalance(ctx, in)
}

// WithdrawBalance2 use this function instead of WithdrawBalance.
func (c *GrpcClient) WithdrawBalance2(from string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.WithdrawBalanceContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.WithdrawBalance2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// WithdrawExpireUnfreeze withdraw unfrozen balance in Stake2.0, the user can call this API to get back their funds
// after executing /wallet/unfreezebalancev2 transaction and waiting N days, N is a network parameter
// https://developers.tron.network/reference/withdrawexpireunfreeze
func (c *GrpcClient) WithdrawExpireUnfreeze(from string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.WithdrawExpireUnfreezeContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.WithdrawExpireUnfreeze(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// CancelAllUnfreezeV2 cancel unstake all unfreeze
// https://developers.tron.network/reference/cancelallunfreezev2
func (c *GrpcClient) CancelAllUnfreezeV2(from string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.CancelAllUnfreezeV2Contract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.CancelAllUnfreezeV2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// GetAccountNet query bandwidth information of an account
// https://developers.tron.network/reference/getaccountnet
func (c *GrpcClient) GetAccountNet(addr string) (*tronpb.AccountNetMessage, error) {
	var err error
	var in = new(tronpb.Account)
	if in.Address, err = c.convert(addr); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetAccountNet(ctx, in)
}

// AccountPermissionUpdate update the account's permission.
// https://developers.tron.network/reference/accountpermissionupdate
func (c *GrpcClient) AccountPermissionUpdate(
	from string,
	owner,
	witness map[string]interface{},
	actives []map[string]interface{},
) (*tronpb.TransactionExtention, error) {
	if len(actives) > 8 {
		return nil, fmt.Errorf("cant have more than 8 active operations")
	}

	if owner == nil {
		return nil, fmt.Errorf("owner is manadory")
	}
	ownerPermission, err := makePermission(
		"owner",
		tronpb.Permission_Owner,
		0,
		owner["threshold"].(int64),
		nil,
		owner["keys"].(map[string]int64),
	)
	if err != nil {
		return nil, err
	}
	contract := &tronpb.AccountPermissionUpdateContract{
		Owner: ownerPermission,
	}

	if contract.OwnerAddress, err = common.DecodeCheck(from); err != nil {
		return nil, err
	}

	if actives != nil {
		activesPermission := make([]*tronpb.Permission, 0)
		for i, active := range actives {
			activeP, err := makePermission(
				active["name"].(string),
				tronpb.Permission_Active,
				int32(2+i),
				active["threshold"].(int64),
				active["operations"].(map[string]bool),
				active["keys"].(map[string]int64),
			)
			if err != nil {
				return nil, err
			}
			activesPermission = append(activesPermission, activeP)
		}
		contract.Actives = activesPermission
	}

	if witness != nil {
		witnessPermission, err := makePermission(
			"witness",
			tronpb.Permission_Witness,
			1,
			witness["threshold"].(int64),
			nil,
			witness["keys"].(map[string]int64),
		)
		if err != nil {
			return nil, err
		}
		contract.Witness = witnessPermission
	}

	ctx, cancel := c.getContext()
	defer cancel()

	tx, err := c.WalletClient.AccountPermissionUpdate(ctx, contract)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// GetTransactionSignWeight queries transaction sign weight
func (c *GrpcClient) GetTransactionSignWeight(tx *tronpb.Transaction) (*tronpb.TransactionSignWeight, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetTransactionSignWeight(ctx, tx)
}

// GetTransactionApprovedList query approve transaction list
func (c *GrpcClient) GetTransactionApprovedList(tx *tronpb.Transaction) (*tronpb.TransactionApprovedList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetTransactionApprovedList(ctx, tx)
}

// GetRewardInfo get reward info
func (c *GrpcClient) GetRewardInfo(addr string) (*tronpb.NumberMessage, error) {
	var err error
	var in = new(tronpb.BytesMessage)
	if in.Value, err = c.convert(addr); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetRewardInfo(ctx, in)
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
			return common.HexToByte(addr)
		}
	}

	return address.Address{}, fmt.Errorf("address is invalid")
}
