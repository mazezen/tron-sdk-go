package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// CreateTransaction create a TRX transaction. If the recipient account(to_address)
// does not exist, it will be automatically activated,
// -------------------------------------------------------------
// call map[string]interface{}
// call['owner_address'] = "T..." sender address.(Format: Base58 or Hex)
// call['to_address'] = "T..."    Recipient  address.(Format: Base58 or Hex)
// call['amount'] = 100           amount of TRX to transfer. (Unit: sun)
// call['permission_id'] = 0      the permission ID of the account. Defaults to owner permission if unspecified
// call['visible'] = false        set to true to format addresses in Base58; set to false for hex format. (Default: false)
// call['extra_data'] = ”         remarks or notes for the transaction, in hex format.
// -------------------------------------------------------------
// -------------------------------------------------------------
// https://developers.tron.network/reference/createtransaction
// -------------------------------------------------------------
func (a *Client) CreateTransaction(ctx context.Context, call map[string]interface{}) (*internet.Transaction, error) {
	var result = new(internet.Transaction)
	rs, err := a.client.R().SetContext(ctx).SetBody(call).SetResult(result).Post(methods.CreateTransaction)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("create transaction status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("create transaction error: %s", result.Error),
		}
	}

	return result, nil
}

// BroadTransaction broadcasts a signed transaction to the network
// -------------------------------------------------------------
// https://developers.tron.network/reference/broadcasttransaction
// -------------------------------------------------------------
func (a *Client) BroadTransaction(ctx context.Context, tx *internet.Transaction) (string, error) {
	var result struct {
		Code    string `json:"code"`
		TxId    string `json:"txid"`
		Message string `json:"message"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(tx).SetResult(&result).Post(methods.Broadcasttransaction)
	if err != nil {
		return "", &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return "", &SendRequestError{
			Err: fmt.Errorf("broadcast transaction status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Code != "" || result.Message != "" {
		return "", &SendRequestError{
			Err: fmt.Errorf("error: %s message: %s", result.Code, result.Message),
		}
	}

	return result.TxId, nil
}

// BroadcastHex
// -------------------------------------------------------------
// https://developers.tron.network/reference/broadcasthex
// -------------------------------------------------------------
func (a *Client) BroadcastHex(ctx context.Context, transactionHex string) (string, error) {
	var result struct {
		Code        string               `json:"code"`
		TxId        string               `json:"txid"`
		Message     string               `json:"message"`
		Result      bool                 `json:"result"`
		Transaction internet.Transaction `json:"transaction"`
	}
	rs, err := a.client.R().SetContext(ctx).SetBody(transactionHex).SetResult(&result).Post(methods.BroadcastHex)
	if err != nil {
		return "", &SendRequestError{Err: err}
	}
	if rs.StatusCode() != 200 {
		return "", &SendRequestError{
			Err: fmt.Errorf("broadcast transaction status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Result != true {
		return "", &SendRequestError{
			Err: fmt.Errorf("broadcast transaction code: %s: error %s", result.Code, result.Message),
		}
	}

	return result.TxId, nil
}
