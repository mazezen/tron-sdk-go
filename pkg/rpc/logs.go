package sdk_rpc

import "context"

type EthGetLogsParams struct {
	FromBlock string   `json:"fromBlock,omitempty"`
	ToBlock   string   `json:"toBlock,omitempty"`
	Address   string   `json:"address,omitempty"`
	Topics    []string `json:"topics,omitempty"`
	BlockHash string   `json:"blockHash,omitempty"`
}

// EthGetLogs return an array of all logs that match the given filter object.
// https://developers.tron.network/reference/eth_getlogs
func (r *RpcClient) EthGetLogs(ctx context.Context) (*Log, error) {
	var result = new(Log)
	err := r.Call(ctx, EthGetLogs, []interface{}{}, result)
	return result, err
}
