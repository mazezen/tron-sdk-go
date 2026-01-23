package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
)

// CreateTransaction
// from: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// to: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// amount: Amount of TRX to transfer. (Unit: sun, 1 TRX = 1,000,000 sun).
// Please use CreateTransaction2 instead of this function.
func (c *GrpcClient) CreateTransaction(from, to string, amount int64) (*tronpb.Transaction, error) {
	var err error
	var req = new(tronpb.TransferContract)

	if req.OwnerAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid from: %w", err)
	}

	if req.ToAddress, err = c.convert(to); err != nil {
		return nil, fmt.Errorf("invalid to: %w", err)
	}
	req.Amount = amount

	ctx, cancelFun := c.getContext()
	defer cancelFun()

	return c.WalletClient.CreateTransaction(ctx, req)
}

// CreateTransaction2
// from: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// to: account address base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// amount: Amount of TRX to transfer. (Unit: sun, 1 TRX = 1,000,000 sun).
// Use this function instead of CreateTransaction.
// https://developers.tron.network/reference/createtransaction
func (c *GrpcClient) CreateTransaction2(from, to string, amount int64) (*tronpb.TransactionExtention, error) {
	var err error
	var req = new(tronpb.TransferContract)
	if req.OwnerAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid from: %w", err)
	}
	if req.ToAddress, err = c.convert(to); err != nil {
		return nil, fmt.Errorf("invalid to: %w", err)
	}
	req.Amount = amount

	ctx, cancelFun := c.getContext()
	defer cancelFun()
	tx, err := c.WalletClient.CreateTransaction2(ctx, req)
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
