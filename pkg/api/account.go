package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// CreateAccount activate address
// ---------------------------------------------------------------
// https://developers.tron.network/reference/account-createaccount
// ---------------------------------------------------------------
func (a *Client) CreateAccount(ctx context.Context, ownerAddress, accountAddress string, visible bool) error {
	var result = new(internet.CreateAccountInternet)
	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"owner_address":   ownerAddress,
			"account_address": accountAddress,
			"visible":         visible,
		}).
		SetResult(result).
		Post(methods.CreateAccount)
	if err != nil {
		return &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return &SendRequestError{
			Err: fmt.Errorf("create account status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return &SendRequestError{
			Err: fmt.Errorf("create account error: %s", result.Error),
		}
	}

	return nil
}

// GetAccount an account’s on-chain information, including its TRX balance, TRC10 assets, staked information,
// voting status, and permission settings
// ---------------------------------------------------------------
// https://developers.tron.network/reference/account-getaccount
// ---------------------------------------------------------------
func (a *Client) GetAccount(ctx context.Context, address string, visible bool) (*internet.AccountInternet, error) {
	var result = new(internet.AccountInternet)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"address": address,
		"visible": visible,
	}).SetResult(result).Post(methods.GetAccount)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get account status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}
	return result, nil
}

// UpdateAccount modify the name of account
// ---------------------------------------------------------------
// https://developers.tron.network/reference/updateaccount
// ---------------------------------------------------------------
func (a *Client) UpdateAccount(
	ctx context.Context,
	ownerAddress string,
	accountName string,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"account_name":  accountName,
		"visible":       visible,
	}).SetResult(result).Post(methods.UpdateAccount)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update account status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}
	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update account error: %s", result.Error),
		}
	}
	return result, nil
}

// GetAccountResource the resource of an account.(bandwidth, energy, etc.)
// -------------------------------------------------------------
// https://developers.tron.network/reference/getaccountresource
// -------------------------------------------------------------
func (a *Client) GetAccountResource(
	ctx context.Context,
	address string,
	visible bool,
) (*internet.AccountResourceInternet, error) {
	var result = new(internet.AccountResourceInternet)
	var in = struct {
		Address string `json:"address"`
		Visible bool   `json:"visible"`
	}{
		Address: address,
		Visible: visible,
	}
	rs, err := a.client.R().SetContext(ctx).SetBody(in).SetResult(result).Post(methods.GetAccountResource)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get account resource status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}
	return result, nil
}

// GetAccountBalance query the historical balance of an account at a specific block.
// -------------------------------------------------------------
// https://developers.tron.network/reference/getaccountbalance
// -------------------------------------------------------------
func (a *Client) GetAccountBalance(
	ctx context.Context,
	address,
	hash string,
	visible bool,
) (*internet.AccountBalanceInternet, error) {
	var result = new(internet.AccountBalanceInternet)
	var in = struct {
		AccountIdentifier struct {
			Address string `json:"address"`
		} `json:"account_identifier"`
		BlockIdentifier struct {
			Hash string `json:"hash"`
		} `json:"block_hash"`
		Visible bool `json:"visible"`
	}{
		AccountIdentifier: struct {
			Address string `json:"address"`
		}{Address: address},
		BlockIdentifier: struct {
			Hash string `json:"hash"`
		}{Hash: hash},
		Visible: visible,
	}

	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(in).SetResult(result).Post(methods.GetAccountBalance)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get account balance status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetAccountNet query bandwidth information of an account.
// -------------------------------------------------------------
// https://developers.tron.network/reference/getaccountnet
// -------------------------------------------------------------
func (a *Client) GetAccountNet(ctx context.Context, address string, visible bool) (*internet.AccountResourceInternet, error) {
	var result = new(internet.AccountResourceInternet)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"address": address,
		"visible": visible,
	}).SetResult(result).Post(methods.GetAccountNet)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get account net status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetDelegatedResource returns all resources delegations during stake1.0 phase from an account to another account
// -------------------------------------------------------------
// https://developers.tron.network/reference/getdelegatedresource
// -------------------------------------------------------------
func (a *Client) GetDelegatedResource(
	ctx context.Context,
	from, to string,
	visible bool,
) (*internet.DelegatedResourceInternet, error) {
	var result = new(internet.DelegatedResourceInternet)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"fromAddress": from,
		"toAddress":   to,
		"visible":     visible,
	}).SetResult(result).Post(methods.GetDelegatedResource)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get delegated resource status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetDelegatedResourceAccountIndex query the resource delegation by an account during stake1.0 phase.
// i.e. list all addresses that have delegated resources to an account.
// --------------------------------------------------------------------------
// https://developers.tron.network/reference/getdelegatedresourceaccountindex
// --------------------------------------------------------------------------
func (a *Client) GetDelegatedResourceAccountIndex(
	ctx context.Context,
	value string,
	visible bool,
) (*internet.DelegatedResourceInternet, error) {
	var result = new(internet.DelegatedResourceInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   value,
		"visible": visible,
	}).SetResult(result).Post(methods.GetDelegatedResourceAccountIndex)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get delegated resource account index: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetDelegatedResourceV2 in Stake2.0, query the detail of resource share delegated from fromAddress to toAddress.
// -------------------------------------------------------------
// https://developers.tron.network/reference/getdelegatedresourcev2
// -------------------------------------------------------------
func (a *Client) GetDelegatedResourceV2(
	ctx context.Context,
	from, to string,
	visible bool,
) (*internet.DelegatedResourceInternet, error) {
	var result = new(internet.DelegatedResourceInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"fromAddress": from,
		"toAddress":   to,
		"visible":     visible,
	}).SetResult(result).Post(methods.GetDelegatedResourceV2)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get delegated resource status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetDelegatedResourceAccountIndexV2 In Stake2.0, query the resource delegation index of an account.
// Two lists will return, one is the list of addresses the account has delegated its resources(toAddress),
// and the other is the list of addresses that have delegated resources to the account(fromAddress).
// ------------------------------------------------------------------------------
// https://developers.tron.network/reference/getdelegatedresourceaccountindexv2-1
// ------------------------------------------------------------------------------
func (a *Client) GetDelegatedResourceAccountIndexV2(
	ctx context.Context,
	value string,
	visible bool,
) (*internet.DelegatedResourceInternet, error) {
	var result = new(internet.DelegatedResourceInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   value,
		"visible": visible,
	}).SetResult(result).Post(methods.GetDelegatedResourceAccountIndexV2)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get delegated resource account index v2: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// DelegateResource delegate bandwidth or energy resources to other accounts in Stake2.0.
// -------------------------------------------------------------
// https://developers.tron.network/reference/delegateresource-1
// -------------------------------------------------------------
func (a *Client) DelegateResource(
	ctx context.Context,
	ownerAddress, receiveAddress string,
	balance int64, // Amount of TRX to delegate for resources. (Unit: sun)
	resource string, // Type of resource. (Enum: "BANDWIDTH" or "ENERGY")
	lock bool, // Optional. Whether to lock the resource delegation. If true, the delegation cannot be canceled during the lock_period.
	lockPeriod int64, // Lock duration in blocks (1 block ≈ 3 seconds). Only valid if lock is true. (e.g., 1 day = 28800)
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":   ownerAddress,
		"receive_address": receiveAddress,
		"balance":         balance,
		"resource":        resource,
		"lock_period":     lockPeriod,
		"lock":            lock,
		"visible":         visible,
	}).SetResult(result).Post(methods.DelegateResource)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("delegate resource status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// UnDelegateResource cancel the delegation of bandwidth or energy resources to other accounts in Stake2.0.
// -------------------------------------------------------------
// https://developers.tron.network/reference/undelegateresource-1
// -------------------------------------------------------------
func (a *Client) UnDelegateResource(
	ctx context.Context,
	ownerAddress, receiveAddress string,
	balance int64, // Amount of TRX to delegate for resources. (Unit: sun)
	resource string, // Type of resource. (Enum: "BANDWIDTH" or "ENERGY")
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":   ownerAddress,
		"receive_address": receiveAddress,
		"balance":         balance,
		"resource":        resource,
		"visible":         visible,
	}).SetResult(result).Post(methods.UnDelegateResource)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("un delegate resource status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetAvailableUnfreezeCount remaining times of executing unstake operation in Stake2.0.
// ---------------------------------------------------------------------
// https://developers.tron.network/reference/getavailableunfreezecount-1
// ---------------------------------------------------------------------
func (a *Client) GetAvailableUnfreezeCount(ctx context.Context, ownerAddress string, visible bool) (int64, error) {
	var result struct {
		Count int64 `json:"count,omitempty"` // 剩余可用取消质押次数
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"visible":       visible,
	}).SetResult(&result).Post(methods.GetAvailableUnfreezeCount)

	if err != nil {
		return 0, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get available unfree count: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result.Count, nil
}

// GetCanWithdrawUnfreezeAmount query the withdrawable balance at the specified timestamp In Stake2.0.
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/getcanwithdrawunfreezeamount-1
// ------------------------------------------------------------------------
func (a *Client) GetCanWithdrawUnfreezeAmount(
	ctx context.Context,
	ownerAddress string,
	timestamp int64, // Timestamp to query. (Unit: millGisecond)
	visible bool,
) (int64, error) {
	var result struct {
		Amount int64 `json:"amount,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"timestamp":     timestamp,
		"visible":       visible,
	}).SetResult(&result).Post(methods.GetCanWithdrawUnfreezeAmount)

	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get can withdraw unfree amount: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result.Amount, nil
}

// GetCanDelegatedMaxSize in Stake2.0, query the amount of delegatable resources share of the specified resource
// type for an address, unit is sun
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/getcandelegatedmaxsize
// ------------------------------------------------------------------------
func (a *Client) GetCanDelegatedMaxSize(
	ctx context.Context,
	ownerAddress string,
	typ int64, // resource type, 0 is bandwidth, 1 is energy.
	visible bool,
) (int64, error) {
	var result struct {
		MaxSize int64 `json:"max_size,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"type":          typ,
		"visible":       visible,
	}).SetResult(&result).Post(methods.GetCanDelegatedMaxSize)

	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get can delegated max size: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result.MaxSize, nil
}
