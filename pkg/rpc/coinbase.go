package sdk_rpc

import "context"

// EthCoinbase return the super representative address of the current node
// 只有部署自己的 FullNode（且该节点被配置为 SR）才会返回一个地址
func (r *RpcClient) EthCoinbase(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, EthCoinbase, nil, &result)
	return result, err
}
