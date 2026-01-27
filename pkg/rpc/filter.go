package sdk_rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/common"
)

type EthNewFilterParams struct {
	Address   []string      `json:"address,omitempty"`
	FromBlock string        `json:"fromBlock,omitempty"`
	ToBlock   string        `json:"toBlock,omitempty"`
	Topics    []interface{} `json:"topics,omitempty"`
}

// EthNewFilter creates a filter object, based on filter options, to notify when the state changes (logs).
// https://developers.tron.network/reference/eth_newfilter
func (r *RpcClient) EthNewFilter(ctx context.Context, call EthNewFilterParams) (string, error) {
	var result string
	bytes, _ := json.Marshal(call)
	fmt.Println(common.JSONPrettyFormat(string(bytes)))
	err := r.Call(ctx, EthNewFilter, []interface{}{call}, &result)
	return result, err
}

// EthNewBlockFilter creates a new block filter on the node, enabling clients to be notified when a new block arrives.
// https://developers.tron.network/reference/eth_newblockfilter
func (r *RpcClient) EthNewBlockFilter(ctx context.Context) (string, error) {
	var result string
	err := r.Call(ctx, EthNewBlockFilter, []interface{}{}, &result)
	return result, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
// https://developers.tron.network/reference/eth_getfilterchanges
func (r *RpcClient) EthGetFilterChanges(ctx context.Context, filterId string) ([]*Log, error) {
	var raw json.RawMessage
	if err := r.Call(ctx, EthGetFilterChanges, []interface{}{filterId}, &raw); err != nil {
		return nil, err
	}

	var err error
	var strLogs []string
	if err = json.Unmarshal(raw, &strLogs); err == nil {
		result := make([]*Log, len(strLogs))
		for i, s := range strLogs {
			result[i] = &Log{
				Data:    s,
				Removed: false,
			}
		}
		return result, nil
	}

	var logs []*Log
	if err = json.Unmarshal(raw, &logs); err == nil {
		return logs, nil
	}

	return nil, fmt.Errorf("failed to parse eth_getFilterChanges response: %w\nraw: %s", err, string(raw))
}

// EthGetFilterLogs returns all log objects that match the filter identified by the given ID.
func (r *RpcClient) EthGetFilterLogs(ctx context.Context, filterId string) (*Log, error) {
	var result = new(Log)
	err := r.Call(ctx, EthGetFilterLogs, []interface{}{filterId}, &result)
	return result, err
}

// UninstallFilter Uninstalls the filter by the given ID. It should always be called when the filter is no longer required.
// Filters automatically time out if they are not queried using eth_getFilterChanges within a certain period
// https://developers.tron.network/reference/eth_uninstallfilter
func (r *RpcClient) UninstallFilter(ctx context.Context, filterId string) (bool, error) {
	var result bool
	err := r.Call(ctx, EthUninstallFilter, []interface{}{filterId}, &result)
	return result, err
}
