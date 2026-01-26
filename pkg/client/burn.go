package client

import tronpb "github.com/mazezen/tron-sdk-go/pb/tron"

// GetBurnTrx query the amount of TRX burned due to on-chain transaction fees
// since No. 54 Committee Proposal took effect. (Confirmed state)
// https://developers.tron.network/reference/getburntrx-1
func (c *GrpcClient) GetBurnTrx() (*tronpb.NumberMessage, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetBurnTrx(ctx, &tronpb.EmptyMessage{})
}
