package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// GetTransactionFromPending get transaction details from the pending pool.
func (c *GrpcClient) GetTransactionFromPending(value string) (*tronpb.Transaction, error) {
	var err error
	var req = new(tronpb.BytesMessage)

	if req.Value, err = common.Hex2Bytes(value); err != nil {
		return nil, fmt.Errorf("invalid transaction hash: %w", err)
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetTransactionFromPending(ctx, req)
}

// GetTransactionListFromPending get transaction list information from pending pool.
// https://developers.tron.network/reference/gettransactionlistfrompending
func (c *GrpcClient) GetTransactionListFromPending() (*tronpb.TransactionIdList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetTransactionListFromPending(ctx, &tronpb.EmptyMessage{})
}

// GetPendingSize return the size of pending pool queue
func (c *GrpcClient) GetPendingSize() (*tronpb.NumberMessage, error) {
	var req = new(tronpb.EmptyMessage)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetPendingSize(ctx, req)
}
