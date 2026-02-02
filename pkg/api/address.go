package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// ValidateAddress validate address. return either true or false
// -------------------------------------------------------------
// https://developers.tron.network/reference/validateaddress
// -------------------------------------------------------------
func (a *Client) ValidateAddress(ctx context.Context, address string, visible bool) (*internet.ValidateAddressInternet, error) {
	var result = new(internet.ValidateAddressInternet)
	rs, err := a.client.R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"address": address,
			"visible": visible,
		}).SetResult(result).
		Post(methods.Validateaddress)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, fmt.Errorf("validate address status code: %d error: %s", rs.StatusCode(), rs.Status())
	}

	return result, err
}
