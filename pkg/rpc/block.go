package sdk_rpc

import (
	"context"

	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// EthBlockNumber return the number of the most recent block
// https://developers.tron.network/reference/eth_blocknumber
func (r *RpcClient) EthBlockNumber(ctx context.Context) (uint64, error) {
	var resp string
	if err := r.Call(ctx, EthBlockNumber, nil, &resp); err != nil {
		return 0, err
	}
	return common.HexToUint64(resp), nil
}

// EthGetBlockByHash return information a block by hash
// https://developers.tron.network/reference/eth_getblockbyhash
func (r *RpcClient) EthGetBlockByHash(ctx context.Context, blockHash string, full bool) (*EthGetBlockByHashResult, error) {
	var result = new(EthGetBlockByHashResult)
	err := r.Call(ctx, EthGetBlockByHash, []interface{}{blockHash, full}, result)
	return result, err
}

// EthGetBlockByNumber return information a block by block number
// https://developers.tron.network/reference/eth_getblockbynumber
func (r *RpcClient) EthGetBlockByNumber(ctx context.Context, blockNumber string, full bool) (*EthGetBlockByHashResult, error) {
	var result = new(EthGetBlockByHashResult)
	err := r.Call(ctx, EthGetBlockByNumber, []interface{}{blockNumber, full}, result)
	return result, err
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash
// https://developers.tron.network/reference/eth_getblocktransactioncountbyhash
func (r *RpcClient) EthGetBlockTransactionCountByHash(ctx context.Context, blockHash string) (string, error) {
	var result string
	err := r.Call(ctx, EthGetBlockTransactionCountByHash, []interface{}{blockHash}, &result)
	return result, err
}

// EthGetBlockTransactionCountByNumber return the number of transactions in a block matching the given block number.
// https://developers.tron.network/reference/eth_getblocktransactioncountbynumber
func (r *RpcClient) EthGetBlockTransactionCountByNumber(ctx context.Context, blockNumber string) (string, error) {
	var result string
	err := r.Call(ctx, EthGetBlockTransactionCountByNumber, []interface{}{blockNumber}, &result)
	return result, err
}

// EthGetWork return the hash of the current block
// https://developers.tron.network/reference/eth_getwork
func (r *RpcClient) EthGetWork(ctx context.Context) ([]string, error) {
	var result []string
	err := r.Call(ctx, EthGetWork, []interface{}{}, &result)
	return result, err
}
