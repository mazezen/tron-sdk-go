package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// UnfreezeBalance unstake unstake the TRX staked during Stake1.0, release the obtained bandwidth or energy and TP. This operation will
// cause automatically cancel all votes.
// -------------------------------------------------------------
// https://developers.tron.network/reference/account-resources-unfreezebalance
// -------------------------------------------------------------
func (a *Client) UnfreezeBalance(
	ctx context.Context,
	ownerAddress string,
	resource string, // Type of resource. (Enum: "BANDWIDTH" or "ENERGY")
	receiveAddress string, // Optional, in Base58Check format or HEX format.
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	var in = map[string]interface{}{
		"owner_address": ownerAddress,
		"resource":      resource,
		"visible":       visible,
	}
	if len(receiveAddress) > 0 {
		in["receive_address"] = receiveAddress
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(in).SetResult(result).Post(methods.UnfreezeBalance)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("unfreeze balance status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, err
}

// FreezeBalanceV2 in Stake2.0, stake an amount of TRX to obtain bandwidth or energy, and obtain equivalent
// TRON Power(TP) according to the staked amount.
// -------------------------------------------------------------
// https://developers.tron.network/reference/freezebalancev2-1
// -------------------------------------------------------------
func (a *Client) FreezeBalanceV2(
	ctx context.Context,
	ownerAddress string,
	frozenBalance int64, // TRX stake amount, the unit is sun
	resource string, // TRX stake type, 'BANDWIDTH' or 'ENERGY'
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":  ownerAddress,
		"frozen_balance": frozenBalance,
		"resource":       resource,
		"visible":        visible,
	}).SetResult(result).Post(methods.FreezeBalanceV2)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("freeze balance status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// UnfreezeBalanceV2 unstake some TRX staked in Stake2.0, release the corresponding amount of bandwidth or energy, and voting rights (TP).
// -------------------------------------------------------------
// https://developers.tron.network/reference/unfreezebalancev2-1
// -------------------------------------------------------------
func (a *Client) UnfreezeBalanceV2(
	ctx context.Context,
	ownerAddress string,
	frozenBalance int64, // TRX stake amount, the unit is sun
	resource string, // TRX stake type, 'BANDWIDTH' or 'ENERGY'
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":  ownerAddress,
		"frozen_balance": frozenBalance,
		"resource":       resource,
		"visible":        visible,
	}).SetResult(result).Post(methods.UnfreezeBalanceV2)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("unfreeze balance status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}
	return result, nil
}

// CancelAllUnfreezeV2 Cancel unstakings, all unstaked funds still in the waiting period will be re-staked,
// all unstaked funds that exceeded the 14-day waiting period will be automatically withdrawn to the owner’s account.
// -------------------------------------------------------------
// https://developers.tron.network/reference/cancelallunfreezev2
// -------------------------------------------------------------
func (a *Client) CancelAllUnfreezeV2(ctx context.Context, ownerAddress string, visible bool) (*internet.Transaction, error) {
	var result = new(internet.Transaction)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"visible":       visible,
	}).SetResult(result).Post(methods.CancelAllUnfreezeV2)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("cancel all balance status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}
	return result, nil
}

// WithdrawExpireUnfreeze withdraw unfrozen balance in Stake2.0, the user can call this API to get back their funds
// after executing /wallet/unfreezebalancev2 transaction and waiting N days, N is a network parameter.
// ------------------------------------------------------------------------------
// https://developers.tron.network/reference/withdrawexpireunfreeze
// ------------------------------------------------------------------------------
func (a *Client) WithdrawExpireUnfreeze(ctx context.Context, ownerAddress string, visible bool) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"visible":       visible,
	}).SetResult(result).Post(methods.WithdrawExpireUnfreeze)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("withdraw expire balance status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}
