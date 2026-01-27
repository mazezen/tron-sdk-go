package sdk_rpc

import (
	"context"
)

// EthAccounts return a list address owned by the client.
// https://developers.tron.network/reference/eth_accounts
func (r *RpcClient) EthAccounts(ctx context.Context) error {
	var resp []string

	if err := r.Call(ctx, EthAccounts, nil, &resp); err != nil {
		return err
	}

	return nil
}
