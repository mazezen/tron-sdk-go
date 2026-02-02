package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetTransactionById query transaction information by transaction hash.(Confirmed state)
// -------------------------------------------------------------
// https://developers.tron.network/reference/gettransactionbyid
// -------------------------------------------------------------
func (a *Client) GetTransactionById(ctx context.Context, hash string) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(map[string]string{
			"value": hash,
		}).SetResult(result).
		Post(methods.GetTransactionById)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get transaction by id status code %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetTransactionInfoById query transaction info by transaction hash
// -------------------------------------------------------------
// https://developers.tron.network/reference/gettransactioninfobyid-1
// -------------------------------------------------------------
func (a *Client) GetTransactionInfoById(ctx context.Context, hash string) (*internet.TransactionInfoByIdInternet, error) {
	var result = new(internet.TransactionInfoByIdInternet)

	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(map[string]string{
			"value": hash,
		}).SetResult(result).
		Post(methods.GetTransactionInfoById)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get transaction info by id status code %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetTransactionInfoByBlockNum returns the TransactionInfo data of all transactions contained in the block of the specified height
// ------------------------------------------------------------------------
// https://developers.tron.network/reference/gettransactioninfobyblocknum-1
// ------------------------------------------------------------------------
func (a *Client) GetTransactionInfoByBlockNum(ctx context.Context, num int32) ([]*internet.TransactionInfoByIdInternet, error) {
	var result = make([]*internet.TransactionInfoByIdInternet, 0, 200)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int32{
		"num": num,
	}).SetResult(&result).Post(methods.GetTransactionInfoByBlockNum)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get transaction info by block num %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetTransactionCountByBlockNum get transaction count in a block by block number.(Confirmed state)
// -----------------------------------------------------------------------
// https://developers.tron.network/reference/gettransactioncountbyblocknum
// -----------------------------------------------------------------------
func (a *Client) GetTransactionCountByBlockNum(ctx context.Context, num int32) (int64, error) {
	var result struct {
		Count int64 `json:"count"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]int32{
		"num": num,
	}).SetResult(&result).Post(methods.GetTransactionCountByBlockNum)
	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get transaction count by block num %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result.Count, nil
}
