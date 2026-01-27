package sdk_rpc

import "context"

// EthChainId return the chainId of the TRON network which is the last four bytes of the genesis bloc hash
// https://developers.tron.network/reference/eth_chainid
func (r *RpcClient) EthChainId(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, EthChainId, nil, &result)
	return result, err
}
