package sdk_rpc

import (
	"context"

	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// EthGetCode  return runtime code of a given smart contract address
// https://developers.tron.network/reference/eth_getcode
func (r *RpcClient) EthGetCode(ctx context.Context, contractAddress string, tag string) (string, error) {
	var result string
	err := r.Call(ctx, EthGetCode, []interface{}{contractAddress, tag}, &result)
	return result, err
}

// EthProtocolVersion return the java-tron block version
// https://developers.tron.network/reference/eth_protocolversion
func (r *RpcClient) EthProtocolVersion(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, EthProtocolVersion, []interface{}{}, &result)
	return result, err
}

// EthSyncing return an object with data about the sync status of the node
// https://developers.tron.network/reference/eth_syncing
func (r *RpcClient) EthSyncing(ctx context.Context) (*SyncingResult, error) {
	var result = new(SyncingResult)
	err := r.Call(ctx, EthSyncing, []interface{}{}, result)
	return result, err
}

// NetListening returns true if the client is actively listening for network connections.
// https://developers.tron.network/reference/net_listening
func (r *RpcClient) NetListening(ctx context.Context) (bool, error) {
	var result bool
	err := r.Call(ctx, NetListening, []interface{}{}, &result)
	return result, err
}

// NetPeerCount returns number of peers currently connected to the client.
// https://developers.tron.network/reference/net_peercount
func (r *RpcClient) NetPeerCount(ctx context.Context) (uint64, error) {
	var result string
	err := r.Call(ctx, NetPeerCount, []interface{}{}, &result)
	c := common.HexToUint64(result)
	return c, err
}

// NetVersion returns the hash of the genesis block.
// https://developers.tron.network/reference/net_version
func (r *RpcClient) NetVersion(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, NetVersion, []interface{}{}, &result)
	return result, err
}

// Web3ClientVersion return the current client version
// https://developers.tron.network/reference/web3_clientversion
func (r *RpcClient) Web3ClientVersion(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, Web3ClientVersion, []interface{}{}, &result)
	return result, err
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data
// https://developers.tron.network/reference/web3_sha3
func (r *RpcClient) Web3Sha3(ctx context.Context, data string) (string, error) {
	var result string
	err := r.Call(ctx, Web3Sha3, []interface{}{data}, &result)
	return result, err
}
