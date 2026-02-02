package client

import (
	"fmt"
	"time"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// UpdateAsset update basic TRC-10 token information
// https://developers.tron.network/reference/wallet-updateasset
// please use UpdateAsset2 instead of this function.
func (c *GrpcClient) UpdateAsset(from, url, description string, newLimit, newPublicLimit int64) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.UpdateAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Url = []byte(url)
	in.Description = []byte(description)
	in.NewLimit = newLimit
	in.NewPublicLimit = newPublicLimit

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.UpdateAsset(ctx, in)
}

// UpdateAsset2 update basic TRC-10 token information
// https://developers.tron.network/reference/wallet-updateasset
// use this function instead of UpdateAsset.
func (c *GrpcClient) UpdateAsset2(from, url, description string, newLimit, newPublicLimit int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.UpdateAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Url = []byte(url)
	in.Description = []byte(description)
	in.NewLimit = newLimit
	in.NewPublicLimit = newPublicLimit

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UpdateAsset2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}

	return tx, nil
}

// TransferAsset transfer TRC-10 token
// https://developers.tron.network/reference/transferasset
// Please use TransferAsset2 instead of this function.
func (c *GrpcClient) TransferAsset(from, to, assetName string, amount int64) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.TransferAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ToAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.AssetName = []byte(assetName)
	in.Amount = amount

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.TransferAsset(ctx, in)
}

// TransferAsset2 transfer TRC-10 token
// https://developers.tron.network/reference/transferasset
// Use this function instead of TransferAsset.
func (c *GrpcClient) TransferAsset2(from, to, assetName string, amount int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.TransferAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ToAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.AssetName = []byte(assetName)
	in.Amount = amount

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.TransferAsset2(ctx, in)
	if err != nil {
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// ParticipateAssetIssue participate in an asset issue.
// https://developers.tron.network/reference/participateassetissue
// Please use ParticipateAssetIssue2 instead of this function.
func (c *GrpcClient) ParticipateAssetIssue(from, to, assetName string, amount int64) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.ParticipateAssetIssueContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ToAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.AssetName = []byte(assetName)
	in.Amount = amount

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.ParticipateAssetIssue(ctx, in)
}

// ParticipateAssetIssue2 participate in an asset issue.
// https://developers.tron.network/reference/participateassetissue
// Use this function instead of ParticipateAssetIssue.
func (c *GrpcClient) ParticipateAssetIssue2(from, to, assetName string, amount int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ParticipateAssetIssueContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ToAddress, err = c.convert(to); err != nil {
		return nil, err
	}
	in.AssetName = []byte(assetName)
	in.Amount = amount
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	tx, err := c.WalletClient.ParticipateAssetIssue2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// UnfreezeAsset unstake a token that has passed the minimum freeze duration.
// https://developers.tron.network/reference/unfreezeasset
// Please use UnfreezeAsset2 instead of this function.
func (c *GrpcClient) UnfreezeAsset(from string) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.UnfreezeAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.UnfreezeAsset(ctx, in)
}

// UnfreezeAsset2
// Use this function instead of UnfreezeAsset.
func (c *GrpcClient) UnfreezeAsset2(from string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.UnfreezeAssetContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UnfreezeAsset2(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// AssetIssue issues a TRC-10 token. An account can only issue a TRC-10 token once
// https://developers.tron.network/reference/createassetissue
func (c *GrpcClient) AssetIssue(
	from string, // Transaction initiator address (Default: Hex). Example: TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g .
	name string, // Token name (Default: Hex). Example: 0x6173736574497373756531353330383934333132313538.
	abbr string, // Token's abbreviation or symbol. Example: 0x6162627231353330383934333132313538.
	totalSupply int64, // Total supply of the TRC-10 token.
	trxNum int32, // Defines the price ratio as trx_num / num (Unit of 'trx_num': sun).
	num int32, // Defines the price ratio as trx_num / num (Unit of 'trx_num': sun).
	startTime int64, // ICO start time (Unit: millisecond).
	endTime int64, // ICO End time (Unit: millisecond).
	description string, // Description of the TRC-10 token (Default: Hex). Example: 0x4578616d706c654465736372697074696f6e.
	url string, // URL of the official website (Default: Hex). Example: 0x7777772e6578616d706c652e636f6d.
	freeAssetNetLimit int64, // Free Bandwidth limit for TRC-10 token transfers.
	publicFreeAssetNetLimit int64, // Total public free Bandwidth limit for a TRC-10 token.
	frozenSupply []*tronpb.AssetIssueContract_FrozenSupply, // List of frozen supply objects for the TRC-10 token. Example: {"frozen_amount":1,"frozen_days":2}.
	precision int32, // Token precision (number of decimal places).
	voteScore int32,
) (*tronpb.TransactionExtention, error) {
	var err error

	var assetIssue = new(tronpb.AssetIssueContract)

	if assetIssue.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	assetIssue.Name = []byte(name)
	assetIssue.Abbr = []byte(abbr)

	if totalSupply <= 0 {
		return nil, fmt.Errorf("totalSupply must be greater than zero")
	}
	assetIssue.TotalSupply = totalSupply
	assetIssue.TrxNum = trxNum

	if num <= 0 {
		return nil, fmt.Errorf("num must be greater than zero")
	}
	assetIssue.Num = num
	now := time.Now().UnixNano() / 1e6
	if startTime <= now {
		return nil, fmt.Errorf("startTime must be greater than current time")
	}
	assetIssue.StartTime = startTime
	if endTime <= startTime {
		return nil, fmt.Errorf("endTime must be greater than start time")
	}
	assetIssue.EndTime = endTime
	assetIssue.Description = []byte(description)
	assetIssue.Url = []byte(url)

	if assetIssue.FreeAssetNetLimit < 0 {
		return nil, fmt.Errorf("freeAssetNetLimit must be greater than zero")
	}
	assetIssue.FreeAssetNetLimit = freeAssetNetLimit

	if assetIssue.PublicFreeAssetNetUsage < 0 {
		return nil, fmt.Errorf("publicFreeAssetNetUsage must be greater than zero")
	}
	assetIssue.PublicFreeAssetNetUsage = publicFreeAssetNetLimit
	assetIssue.FrozenSupply = frozenSupply
	assetIssue.Precision = precision
	assetIssue.VoteScore = voteScore

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	tx, err := c.WalletClient.CreateAssetIssue2(ctx, assetIssue)
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

// GetAssetIssueByAccount query the TRC10 token information issued by an account.
// https://developers.tron.network/reference/getassetissuebyaccount
func (c *GrpcClient) GetAssetIssueByAccount(addr string) (*tronpb.AssetIssueList, error) {
	var err error
	var in = new(tronpb.Account)
	if in.Address, err = c.convert(addr); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetAssetIssueByAccount(ctx, in)
}

// GetAssetIssueByName query token by the name, return token info
// https://developers.tron.network/reference/getassetissuebyname-copy
func (c *GrpcClient) GetAssetIssueByName(tokenName string) (*tronpb.AssetIssueContract, error) {
	var in = &tronpb.BytesMessage{Value: []byte(tokenName)}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetAssetIssueByName(ctx, in)
}

// GetAssetIssueListByName query the list of all the TRC10 tokens by a name.
// https://developers.tron.network/reference/getassetissuelistbyname-copy
func (c *GrpcClient) GetAssetIssueListByName(tokenName string) (*tronpb.AssetIssueList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetAssetIssueListByName(ctx, &tronpb.BytesMessage{Value: []byte(tokenName)})
}

// GetAssetIssueById query a token by token id. return the token object, which contains the token name.
func (c *GrpcClient) GetAssetIssueById(tokenID string) (*tronpb.AssetIssueContract, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetAssetIssueById(ctx, &tronpb.BytesMessage{Value: []byte(tokenID)})
}

// GetAssetIssueList  query the list of all the TRC10 tokens.
// https://developers.tron.network/reference/getassetissuelist
func (c *GrpcClient) GetAssetIssueList() (*tronpb.AssetIssueList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetAssetIssueList(ctx, &tronpb.EmptyMessage{})
}

// GetPaginatedAssetIssueList query the list of all the tokens by pagination.Returns a list of Tokens that succeed the
// Token located at offset. (confirmed state)
// https://developers.tron.network/reference/getpaginatedassetissuelist-1
func (c *GrpcClient) GetPaginatedAssetIssueList(offset, limit int64) (*tronpb.AssetIssueList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetPaginatedAssetIssueList(ctx, &tronpb.PaginatedMessage{
		Offset: offset,
		Limit:  limit,
	})
}
