package sdk_rpc

import "context"

// EthGetBalance return the balance of the account of the given address
// https://developers.tron.network/reference/eth_getbalance
func (r *RpcClient) EthGetBalance(ctx context.Context, address string, block string) (string, error) {
	var result string
	if block == "" {
		block = "latest"
	}
	err := r.Call(ctx, EthGetBalance, []interface{}{address, block}, &result)
	return result, err
}
