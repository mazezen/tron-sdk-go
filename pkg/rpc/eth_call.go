package sdk_rpc

import "context"

type EthCallParams struct {
	From     string `json:"from,omitempty"`     // option. Called address(0x...). Hex format address, remove the prefix "41"
	To       string `json:"to"`                 // required. Contract address(0x...). Hex format address, remove the prefix "41"
	Gas      string `json:",omitempty"`         // option. The value is 0x0
	GasPrice string `json:"gasPrice,omitempty"` // option. The value is 0x0
	Value    string `json:"value,omitempty"`    // option. The value is 0x0
	Data     string `json:"data,omitempty"`     // option. Hash of the method signature and encoded parameters.
}

// EthCall executes a message call immediately without creating a transaction on the block chain
// https://developers.tron.network/reference/eth_call
func (r *RpcClient) EthCall(ctx context.Context, callParams EthCallParams, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	var result string
	if err := r.Call(ctx, EthCall, []interface{}{callParams, block}, &result); err != nil {
		return "", err
	}
	return result, nil
}
