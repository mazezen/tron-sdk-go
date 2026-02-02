package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetAssetIssueByAccount query the TRC10 token information issued by an account
// ----------------------------------------------------------------
// https://developers.tron.network/reference/getassetissuebyaccount
// ----------------------------------------------------------------
func (a *Client) GetAssetIssueByAccount(
	ctx context.Context,
	address string,
	visible bool,
) (*internet.AssetIssueListInternet, error) {
	var result = new(internet.AssetIssueListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"address": address,
		"visible": visible,
	}).SetResult(result).Post(methods.GetAssetIssueByAccount)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get asset issue by account status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetAssetIssueById query a token by token id. Returns the token object, which contains the token name.
// ----------------------------------------------------------------
// https://developers.tron.network/reference/getassetissuebyid
// ----------------------------------------------------------------
func (a *Client) GetAssetIssueById(ctx context.Context, id int32) (*internet.AssetIssueInternet, error) {
	var result = new(internet.AssetIssueInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int32{
		"value": id,
	}).SetResult(result).Post(methods.GetAssetIssueById)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get asset issue by id status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetAssetIssueByName query a token by name, returns token info
// ------------------------------------------------------------------
// https://developers.tron.network/reference/getassetissuebyname-copy
// ------------------------------------------------------------------
func (a *Client) GetAssetIssueByName(ctx context.Context, tokenName string, visible bool) (*internet.AssetIssueInternet, error) {
	var result = new(internet.AssetIssueInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   tokenName,
		"visible": visible,
	}).SetResult(result).Post(methods.GetAssetIssueByName)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get asset issue by name status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetAssetIssueList query the list of all the TRC10 tokens.
// ------------------------------------------------------------------
// https://developers.tron.network/reference/getassetissuelist
// ------------------------------------------------------------------
func (a *Client) GetAssetIssueList(ctx context.Context) (*internet.AssetIssueListInternet, error) {
	var result = new(internet.AssetIssueListInternet)

	rs, err := a.client.R().SetContext(ctx).SetResult(result).Get(methods.GetAssetIssueList)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get asset issue list status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetAssetIssueListByName query the list of all the TRC10 tokens by a name
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/getassetissuelistbyname-copy
// ----------------------------------------------------------------------
func (a *Client) GetAssetIssueListByName(ctx context.Context, tokenName string, visible bool) (*internet.AssetIssueListInternet, error) {
	var result = new(internet.AssetIssueListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   tokenName,
		"visible": visible,
	}).SetResult(result).Post(methods.GetAssetIssueListByName)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get asset issue list by name status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetPaginatedAssetIssueList query the list of all the tokens by pagination.Returns a list of Tokens
// that succeed the Token located at offset
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/getpaginatedassetissuelist
// ----------------------------------------------------------------------
func (a *Client) GetPaginatedAssetIssueList(ctx context.Context, offset, limit int32) (*internet.AssetIssueListInternet, error) {
	var result = new(internet.AssetIssueListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int32{
		"offset": offset,
		"limit":  limit,
	}).SetResult(result).Post(methods.GetPaginatedAssetIssueList)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get paginated asset issue list status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// TransferAsset transfer TRC-10 token.
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/transferasset
// ----------------------------------------------------------------------
func (a *Client) TransferAsset(
	ctx context.Context,
	from string,
	to string,
	assetName string,
	amount int64,
	visible bool,
	extraData string, // (Optional) Totes on the transaction. (Format: Hex)
) (string, error) {
	var result = new(internet.Transaction)

	var in = map[string]interface{}{
		"owner_address": from,
		"to_address":    to,
		"asset_name":    assetName,
		"amount":        amount,
		"visible":       visible,
	}

	if extraData != "" {
		in["extra_data"] = extraData
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(in).SetResult(result).Post(methods.TransferAsset)
	if err != nil {
		return "", &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return "", &SendRequestError{
			Err: fmt.Errorf("transfer asset status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result.TxID, nil
}

// CreateAssetIssue issues a TRC-10 token. An account can only issue a TRC-10 token once.
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/createassetissue
// ----------------------------------------------------------------------
func (a *Client) CreateAssetIssue(
	ctx context.Context,
	ownerAddress string,
	name string,
	abbr string,
	totalSupply int64,
	trxNum int64,
	num int64,
	startTime int64,
	endTime int64,
	description string,
	url string,
	freeAssetNetLimit int64,
	publicFreeAssetNetLimit int64,
	FrozenSupply int64,
	precision int32,
	visible bool,
) (string, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":               ownerAddress,
		"name":                        name,
		"abbr":                        abbr,
		"total_supply":                totalSupply,
		"trx_num":                     trxNum,
		"num":                         num,
		"start_time":                  startTime,
		"end_time":                    endTime,
		"description":                 description,
		"url":                         url,
		"free_asset_net_limit":        freeAssetNetLimit,
		"public_free_asset_net_limit": publicFreeAssetNetLimit,
		"Frozen_supply":               FrozenSupply,
		"precision":                   precision,
		"visible":                     visible,
	}).SetResult(result).Post(methods.CreateAssetIssue)

	if err != nil {
		return "", &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return "", &SendRequestError{
			Err: fmt.Errorf("create asset issue status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result.TxID, nil
}

// ParticipateAssetIssue participate in an asset issue.
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/participateassetissue
// ----------------------------------------------------------------------
func (a *Client) ParticipateAssetIssue(
	ctx context.Context,
	ownerAddress string,
	toAddress string,
	amount int64,
	assetName string,
	visible bool,
) (string, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"to_address":    toAddress,
		"asset_name":    assetName,
		"amount":        amount,
		"visible":       visible,
	}).SetResult(result).Post(methods.ParticipateAssetIssue)

	if err != nil {
		return "", &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return "", &SendRequestError{}
	}

	if result.Error != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result.TxID, nil
}

// UnfreezeAsset unstake a token that has passed the minimum freeze duration.
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/unfreezeasset
// ----------------------------------------------------------------------
func (a *Client) UnfreezeAsset(ctx context.Context, ownerAddress string, visible bool) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"visible":       visible,
	}).SetResult(result).Post(methods.UnfreezeAsset)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("unfreeze asset status code: %d", rs.StatusCode()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// UpdateAsset update basic TRC10 token information.
// ----------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-updateasset
// ----------------------------------------------------------------------
func (a *Client) UpdateAsset(
	ctx context.Context,
	ownerAddress string,
	description string,
	url string,
	newLimit int32,
	newPublicLimit int32,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":    ownerAddress,
		"description":      description,
		"url":              url,
		"new_limit":        newLimit,
		"new_public_limit": newPublicLimit,
		"visible":          visible,
	}).SetResult(result).Post(methods.UpdateAsset)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("status code: %d error %s", rs.StatusCode(), rs.Error()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}
