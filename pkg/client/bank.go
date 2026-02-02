package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// FreezeBalance  stake TRX to obtain resources
// please use FreezeBalance2 instead of this function.
func (c *GrpcClient) FreezeBalance(from, to string, balance int64, resource tronpb.ResourceCode) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.FreezeBalanceContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if len(to) > 0 {
		if in.ReceiverAddress, err = c.convert(to); err != nil {
			return nil, err
		}
	}
	in.FrozenBalance = balance
	in.Resource = resource
	in.FrozenDuration = 3

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.FreezeBalance(ctx, in)
}

// FreezeBalance2 use this function instead of FreezeBalance.
func (c *GrpcClient) FreezeBalance2(from, to string, balance int64, resource tronpb.ResourceCode) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.FreezeBalanceContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if len(to) > 0 {
		if in.ReceiverAddress, err = c.convert(to); err != nil {
			return nil, err
		}
	}
	in.FrozenBalance = balance
	in.Resource = resource
	in.FrozenDuration = 3

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.FreezeBalance2(ctx, in)
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

// FreezeBalanceV2 use this function when FreezeBalanceV2.
func (c *GrpcClient) FreezeBalanceV2(from string, balance int64, resource tronpb.ResourceCode) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.FreezeBalanceV2Contract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.FrozenBalance = balance
	in.Resource = resource

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.FreezeBalanceV2(ctx, in)
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

// UnfreezeBalance unstake the TRX staked during Stake1.0, release the obtained bandwidth or energy and TP.
// This operation will cause automatically cancel all votes
// Please use UnfreezeBalance2 instead of this function.
func (c *GrpcClient) UnfreezeBalance(from, to string, resource tronpb.ResourceCode) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.UnfreezeBalanceContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if len(to) > 0 {
		if in.ReceiverAddress, err = c.convert(to); err != nil {
			return nil, err
		}
	}
	in.Resource = resource
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.UnfreezeBalance(ctx, in)
}

// UnfreezeBalance2  unstake some TRX staked in Stake2.0, release the corresponding amount of bandwidth or energy,and voting rights (TP).
// use this function instead of UnfreezeBalance.
func (c *GrpcClient) UnfreezeBalance2(from, to string, resource tronpb.ResourceCode) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.UnfreezeBalanceContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if len(to) > 0 {
		if in.ReceiverAddress, err = c.convert(to); err != nil {
			return nil, err
		}
	}
	in.Resource = resource

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UnfreezeBalance2(ctx, in)
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

// UnfreezeBalanceV2
// Use this function when UnfreezeBalanceV2.
func (c *GrpcClient) UnfreezeBalanceV2(from string, balance int64, resource tronpb.ResourceCode) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.UnfreezeBalanceV2Contract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.UnfreezeBalance = balance
	in.Resource = resource

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UnfreezeBalanceV2(ctx, in)
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

// GetAvailableUnfreezeCount remaining times of executing unstake operation in Stake2.0.
// https://developers.tron.network/reference/getavailableunfreezecount-1
func (c *GrpcClient) GetAvailableUnfreezeCount(from string) (*tronpb.GetAvailableUnfreezeCountResponseMessage, error) {
	var err error
	var in = new(tronpb.GetAvailableUnfreezeCountRequestMessage)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetAvailableUnfreezeCount(ctx, in)
}

// GetCanWithdrawUnfreezeAmount query the withdrawable balance at the specified timestamp In Stake2.0.
// https://developers.tron.network/reference/getcanwithdrawunfreezeamount-1
func (c *GrpcClient) GetCanWithdrawUnfreezeAmount(from string, timestamp int64) (*tronpb.CanWithdrawUnfreezeAmountResponseMessage, error) {
	var err error
	var in = new(tronpb.CanWithdrawUnfreezeAmountRequestMessage)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Timestamp = timestamp
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetCanWithdrawUnfreezeAmount(ctx, in)
}
