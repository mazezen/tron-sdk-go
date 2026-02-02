package api

import (
	"context"
	"fmt"
	"math/big"
	"net/http"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// ListWitnesses all super Representatives
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/listwitnesses
// ------------------------------------------------------------------------
func (a *Client) ListWitnesses(ctx context.Context, visible bool) (*internet.WitnessListInternet, error) {
	var result = new(internet.WitnessListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]bool{"visible": visible}).Get(methods.ListWitnesses)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("list witnesses status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// CreateWitness apply to become a super representative candidate
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/createwitness
// ------------------------------------------------------------------------
func (a *Client) CreateWitness(
	ctx context.Context,
	ownerAddress string,
	url string,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"url":           url,
		"visible":       visible,
	}).SetResult(result).Post(methods.CreateWitness)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("create witness status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// UpdateWitness edit the URL of the SR's official website
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/updatewitness
// ------------------------------------------------------------------------
func (a *Client) UpdateWitness(
	ctx context.Context,
	ownerAddress string,
	url string,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"url":           url,
		"visible":       visible,
	}).SetResult(result).Post(methods.UpdateWitness)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update witness status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// GetBrokerage SR brokerage ratio
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getbrokerage
// ------------------------------------------------------------------------
func (a *Client) GetBrokerage(ctx context.Context, address string, visible bool) (int64, error) {
	var result struct {
		Brokerage int64  `json:"brokerage,omitempty"`
		Error     string `json:"error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"address": address,
		"visible": visible,
	}).SetResult(&result).Get(methods.GetBrokerage)

	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get brokerage status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return 0, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result.Brokerage, nil
}

// UpdateBrokerage update the SR's brokerage setting
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-updatebrokerage
// ------------------------------------------------------------------------
func (a *Client) UpdateBrokerage(ctx context.Context, ownerAddress string, brokerage int32, visible bool) (*internet.Transaction, error) {
	result := new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"brokerage":     brokerage,
		"visible":       visible,
	}).SetResult(result).Post(methods.UpdateBrokerage)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update brokerage status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// VoteWitnessAccount
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/votewitnessaccount
// ------------------------------------------------------------------------
func (a *Client) VoteWitnessAccount(
	ctx context.Context,
	ownerAddress string,
	votes []map[string]interface{},
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	res, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"votes":         votes,
		"visible":       visible,
	}).SetResult(result).Post(methods.VoteWitnessAccount)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if res.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("vote witness account status code: %d error %s", res.StatusCode(), res.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// GetReward get the rewards that a SR or a user has not yet withdrawn
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getreward
// ------------------------------------------------------------------------
func (a *Client) GetReward(ctx context.Context, address string, visible bool) (*big.Int, error) {
	var result struct {
		Reward *big.Int `json:"reward,omitempty"`
		Error  string   `json:"Error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"address": address,
		"visible": visible,
	}).SetResult(result).Post(methods.GetReward)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get reward status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result.Reward, nil
}

// WithdrawBalance
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/withdrawbalance
// ------------------------------------------------------------------------
func (a *Client) WithdrawBalance(ctx context.Context, ownerAddress string, visible bool) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"visible":       visible,
	}).SetResult(result).Post(methods.WithdrawBalance)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("withdraw balance status code: %d error %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// GetNextMaintenanceTime returns the timestamp of the next voting time in milliseconds
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/getnextmaintenancetime
// ------------------------------------------------------------------------
func (a *Client) GetNextMaintenanceTime(ctx context.Context) (int64, error) {
	var result struct {
		Num   int64  `json:"num,omitempty"`
		Error string `json:"error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetResult(&result).Post(methods.GetNextMaintenanceTime)
	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get next maintenance time status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return 0, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result.Num, nil
}
