package api

import (
	"context"
	"fmt"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetChainParameters retrieve all parameters that can be configured by the blockchain committee along with their values
// -------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getchainparameters
// -------------------------------------------------------------------
func (a *Client) GetChainParameters(ctx context.Context) (*internet.ChainParametersInternet, error) {
	var result = new(internet.ChainParametersInternet)

	rs, err := a.client.R().SetContext(ctx).SetResult(result).Get(methods.GetChainParameters)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != 200 {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get chain parameters: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}
