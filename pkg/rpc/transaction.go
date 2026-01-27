package sdk_rpc

import "context"

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position
// https://developers.tron.network/reference/eth_gettransactionbyblockhashandindex
func (r *RpcClient) EthGetTransactionByBlockHashAndIndex(ctx context.Context, blockHash string, index string) (*TransactionResult, error) {
	var info = new(TransactionResult)
	err := r.Call(ctx, EthGetTransactionByBlockHashAndIndex, []interface{}{blockHash, index}, &info)
	return info, err
}

// EthGetTransactionByBlockNumberAndIndex
// https://developers.tron.network/reference/eth_gettransactionbyblocknumberandindex
func (r *RpcClient) EthGetTransactionByBlockNumberAndIndex(ctx context.Context, blockNumber string, index string) (*TransactionResult, error) {
	var info = new(TransactionResult)
	err := r.Call(ctx, EthGetTransactionByBlockNumberAndIndex, []interface{}{blockNumber, index}, &info)
	return info, err
}

// EthGetTransactionByHash return the information about a transaction requested by transaction hash
// https://developers.tron.network/reference/eth_gettransactionbyhash
func (r *RpcClient) EthGetTransactionByHash(ctx context.Context, hash string) (*TransactionResult, error) {
	var info = new(TransactionResult)
	err := r.Call(ctx, EthGetTransactionByHash, []interface{}{hash}, &info)
	return info, err
}

// EthGetTransactionReceipt returns the transaction info: receipt, transaction fee, block height ... by transaction hash.
// https://developers.tron.network/reference/eth_gettransactionreceipt
func (r *RpcClient) EthGetTransactionReceipt(ctx context.Context, hash string) (*TransactionReceiptResult, error) {
	var info = new(TransactionReceiptResult)
	err := r.Call(ctx, EthGetTransactionReceipt, []interface{}{hash}, &info)
	return info, err
}
