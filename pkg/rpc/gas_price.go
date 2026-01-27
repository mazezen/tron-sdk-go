package sdk_rpc

import "context"

// EthGasPrice return the current energy price in sun
// https://developers.tron.network/reference/eth_gasprice
func (r *RpcClient) EthGasPrice(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, EthGasPrice, nil, &result)
	return result, err
}
