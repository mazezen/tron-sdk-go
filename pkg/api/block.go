package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetBlock query block header information or entire block information according to block height or block hash
// -------------------------------------------------------------
// https://developers.tron.network/reference/getblock-1
// -------------------------------------------------------------
func (a *Client) GetBlock(ctx context.Context, idOrNum string, detail bool) (*internet.BlockInternet, error) {
	var result = new(internet.BlockInternet)
	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"id_or_num": idOrNum,
			"detail":    detail,
		}).
		SetResult(result).
		Post(methods.GetBlock)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get block status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetBlockByNum returns the Block Object corresponding to the 'Block Height' specified (number of blocks preceding it).
// --------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getblockbynum
// --------------------------------------------------------------
func (a *Client) GetBlockByNum(ctx context.Context, num int64) (*internet.BlockInternet, error) {
	var result = new(internet.BlockInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int64{
		"num": num,
	}).SetResult(result).Post(methods.GetBlockByNum)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get block by num: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetBlockById query block by ID (Hash)
// --------------------------------------------------------------
// https://developers.tron.network/reference/getblockbyid
// --------------------------------------------------------------
func (a *Client) GetBlockById(ctx context.Context, blockHash string) (*internet.BlockInternet, error) {
	var result = new(internet.BlockInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]string{
		"value": blockHash,
	}).SetResult(result).Post(methods.GetBlockById)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get block by id: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetBlockByLatestNum returns a list of block objects
// --------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getblockbylatestnum
// --------------------------------------------------------------------
func (a *Client) GetBlockByLatestNum(ctx context.Context, num int64) (*internet.BlockListInternet, error) {
	var result = new(internet.BlockListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int64{
		"num": num,
	}).SetResult(result).Post(methods.GetBlockByLatestNum)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get block by latest num: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetBlockByLimitNext returns the list of Block Objects included in the 'Block Height' range specified.
// --------------------------------------------------------------------
// https://developers.tron.network/reference/getblockbylimitnext
// --------------------------------------------------------------------
func (a *Client) GetBlockByLimitNext(ctx context.Context, startNum, endNum int) (*internet.BlockListInternet, error) {
	var result = new(internet.BlockListInternet)

	rx, err := a.client.R().SetContext(ctx).SetBody(map[string]int{
		"startNum": startNum,
		"endNum":   endNum,
	}).SetResult(result).Post(methods.GetBlockByLimitNext)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rx.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get block by limit next: %d error: %s", rx.StatusCode(), rx.Status()),
		}
	}

	return result, nil
}

// GetNowBlock query the latest block information
// --------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getnowblock
// --------------------------------------------------------------------
func (a *Client) GetNowBlock(ctx context.Context) (*internet.BlockInternet, error) {
	var result = new(internet.BlockInternet)

	rs, err := a.client.R().SetContext(ctx).SetResult(result).Post(methods.GetNowBlock)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get now block: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetEnergyPrices query historical energy unit price
// --------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getnowblock
// --------------------------------------------------------------------
func (a *Client) GetEnergyPrices(ctx context.Context) (string, error) {
	var result struct {
		Prices string `json:"prices,omitempty"`
		Error  string `json:"error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetResult(&result).Post(methods.GetEnergyPrices)
	if err != nil {
		return "", &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return "", &SendRequestError{
			Err: fmt.Errorf("get energy prices: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("get energy prices error: %s", result.Error),
		}
	}

	return result.Prices, nil
}

// GetBandwidthPrices query historical bandwidth unit price.
// --------------------------------------------------------------------
// https://developers.tron.network/reference/getbandwidthprices
// --------------------------------------------------------------------
func (a *Client) GetBandwidthPrices(ctx context.Context) (string, error) {
	var result struct {
		Prices string `json:"prices,omitempty"`
		Error  string `json:"error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetResult(&result).Post(methods.GetBandwidthPrices)
	if err != nil {
		return "", &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return "", &SendRequestError{
			Err: fmt.Errorf("get energy prices: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("get energy prices error: %s", result.Error),
		}
	}

	return result.Prices, nil
}

// GetBurnTRX query the amount of TRX burned from on-chain transaction fees since No. 54 Committee Proposal took effect
// --------------------------------------------------------------------
// https://developers.tron.network/reference/getburntrx
// --------------------------------------------------------------------
func (a *Client) GetBurnTRX(ctx context.Context) (int64, error) {
	var result struct {
		BurnTrxAmount int64  `json:"burnTrxAmount,omitempty"`
		Error         string `json:"error,omitempty"`
	}
	rs, err := a.client.R().SetContext(ctx).SetResult(&result).Post(methods.GetBurnTRX)
	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get burn trx: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get burn trx error: %s", result.Error),
		}
	}

	return result.BurnTrxAmount, nil
}

// GetApprovedList query the account address list which signed the transaction based on the transaction content
// and signature information, which can be used for transaction verification.
// --------------------------------------------------------------------
// https://developers.tron.network/reference/http-getapprovedlist
// --------------------------------------------------------------------
func (a *Client) GetApprovedList(
	ctx context.Context,
	signature []string,
	rawData string,
	visible bool,
) (*internet.ApproveListInternet, error) {
	var result = new(internet.ApproveListInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"signature": signature,
		"rawData":   rawData,
		"visible":   visible,
	}).SetResult(result).Post(methods.GetApprovedList)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}
