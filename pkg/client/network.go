package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
)

// BroadcastTransaction broad transaction
func (c *GrpcClient) BroadcastTransaction(tx *tronpb.Transaction) (*tronpb.Return, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	res, err := c.WalletClient.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	if !res.GetResult() {
		return res, fmt.Errorf("BroadcastTransaction error: %s", res.GetMessage())
	}
	if res.GetCode() != tronpb.Return_SUCCESS {
		return res, fmt.Errorf("BroadcastTransaction error(%s): %s", res.GetCode(), res.GetMessage())
	}
	return res, nil
}
