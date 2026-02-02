package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetTransactionListFromPending get transaction list information from pending pool.
// -----------------------------------------------------------------------
// https://developers.tron.network/reference/gettransactionlistfrompending
// -----------------------------------------------------------------------
func (a *Client) GetTransactionListFromPending(ctx context.Context) ([]string, error) {
	var result struct {
		TxId  []string `json:"txId,omitempty"`
		Error string   `json:"Error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetResult(result).Get(methods.GetTransactionListFromPending)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get transaction list from pending status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result.TxId, nil
}

// GetTransactionFromPending get transaction details from the pending pool.
// -----------------------------------------------------------------------
// https://developers.tron.network/reference/gettransactionfrompending
// -----------------------------------------------------------------------
func (a *Client) GetTransactionFromPending(ctx context.Context, hash string) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]string{
		"value": hash,
	}).SetResult(result).Get(methods.GetTransactionFromPending)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get transaction from pending status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// GetPendingSize get the size of the pending pool queue.
// -----------------------------------------------------------------------
// https://developers.tron.network/reference/getpendingsize
// -----------------------------------------------------------------------
func (a *Client) GetPendingSize(ctx context.Context) (int64, error) {
	var result struct {
		PendingSize int64  `json:"pendingSize,omitempty"`
		Error       string `json:"Error,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetResult(&result).Get(methods.GetPendingSize)
	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return 0, &SendRequestError{
			Err: fmt.Errorf("get pending size status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return 0, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result.PendingSize, nil
}
