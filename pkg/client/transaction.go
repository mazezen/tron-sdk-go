package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"google.golang.org/protobuf/proto"
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

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

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

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
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

// GetTransactionInfoByBlockNum return all transactions in the specified block height.
// https://developers.tron.network/reference/gettransactioninfobyblocknum
func (c *GrpcClient) GetTransactionInfoByBlockNum(num int64) (*tronpb.TransactionInfoList, error) {
	var req = new(tronpb.NumberMessage)
	req.Num = num

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetTransactionInfoByBlockNum(ctx, req)
}

// GetTransactionById query transaction information by transaction id.(Confirmed)
// https://developers.tron.network/reference/gettransactionbyid
func (c *GrpcClient) GetTransactionById(hash string) (*tronpb.Transaction, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	hex, err := common.FromHex(hash)
	if err != nil {
		return nil, err
	}

	return c.WalletClient.GetTransactionById(ctx, &tronpb.BytesMessage{Value: hex})
}

// TotalTransaction total transaction count
func (c *GrpcClient) TotalTransaction() (*tronpb.NumberMessage, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.TotalTransaction(ctx, &tronpb.EmptyMessage{})
}

// GetTransactionInfoById query information by transaction hash
func (c *GrpcClient) GetTransactionInfoById(hash string) (*tronpb.TransactionInfo, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	hex, err := common.FromHex(hash)
	if err != nil {
		return nil, err
	}
	return c.WalletClient.GetTransactionInfoById(ctx, &tronpb.BytesMessage{Value: hex})
}

func (c *GrpcClient) CreateCommonTransaction(tx *tronpb.Transaction) (*tronpb.TransactionExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.CreateCommonTransaction(ctx, tx)
}
