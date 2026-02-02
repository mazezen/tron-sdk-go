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

// GetBandwidthPrices return bandwidth price
// https://developers.tron.network/reference/getbandwidthprices
func (c *GrpcClient) GetBandwidthPrices() (*tronpb.PricesResponseMessage, error) {
	var req = new(tronpb.EmptyMessage)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetBandwidthPrices(ctx, req)
}

// GetEnergyPrices return energy price
// https://developers.tron.network/reference/getenergyprices
func (c *GrpcClient) GetEnergyPrices() (*tronpb.PricesResponseMessage, error) {
	var req = new(tronpb.EmptyMessage)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetEnergyPrices(ctx, req)
}

// GetMemoPrice return memo price
func (c *GrpcClient) GetMemoPrice() (*tronpb.PricesResponseMessage, error) {
	var req = new(tronpb.EmptyMessage)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetMemoFee(ctx, req)
}

// GetChainParameters retrieve all parameters that can be configured by the blockchain committee along with their values.
// https://developers.tron.network/reference/wallet-getchainparameters
func (c *GrpcClient) GetChainParameters() (*tronpb.ChainParameters, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetChainParameters(ctx, &tronpb.EmptyMessage{})
}

// GetNextMaintenanceTime returns the timestamp of the next voting time in milliseconds
// https://developers.tron.network/reference/getnextmaintenancetime
func (c *GrpcClient) GetNextMaintenanceTime() (*tronpb.NumberMessage, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetNextMaintenanceTime(ctx, &tronpb.EmptyMessage{})
}
