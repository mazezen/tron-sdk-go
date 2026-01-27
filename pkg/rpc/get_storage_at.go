package sdk_rpc

import "context"

type EthGetStorageAtParams struct {
	Address string `json:"address"` //  contract address
	Index   string `json:"index"`   // integer of the position in the storage
	Tag     string `json:"tag"`
}

// EthGetStorageAt returns the value from a storage position at a given address.
// It can be used to get the value of a variable in a contract
// https://developers.tron.network/reference/eth_getstorageat
func (r *RpcClient) EthGetStorageAt(ctx context.Context, call EthGetStorageAtParams) (string, error) {
	var result string
	err := r.Call(ctx, EthGetStorageAt, []interface{}{call.Address, call.Index, call.Tag}, &result)
	return result, err
}
