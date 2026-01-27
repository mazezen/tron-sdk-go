package sdk_rpc

import "context"

type EthEstimateGasParams struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

// EthEstiMateGas get the required energy through triggerConstantContract.
// https://developers.tron.network/reference/eth_estimategas
func (r *RpcClient) EthEstiMateGas(ctx context.Context, call EthEstimateGasParams) (string, error) {
	var result string
	err := r.Call(ctx, EthEstiMateGas, []interface{}{call}, &result)
	return result, err
}
