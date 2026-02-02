package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// ListNodes list of peers connected to the current node.
// -------------------------------------------------------------
// https://developers.tron.network/reference/wallet-listnodes
// -------------------------------------------------------------
func (a *Client) ListNodes() (*internet.ListNodesInternet, error) {
	var result = new(internet.ListNodesInternet)
	rs, err := a.client.R().SetResult(result).Post(methods.ListNodes)
	if err != nil {
		return nil, &SendRequestError{
			Err: err,
		}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("list nodes status code %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	for _, node := range result.Nodes {
		bytes, _ := common.HexToByte(node.Address.Host)
		node.Address.Host = string(bytes)
	}

	return result, nil
}

// GetNodeInfo return node information.
// -------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getnodeinfo
// -------------------------------------------------------------
func (a *Client) GetNodeInfo(ctx context.Context) (*internet.NodeInfoInternet, error) {
	var result = new(internet.NodeInfoInternet)

	rs, err := a.client.R().SetResult(result).Get(methods.GetNodeInfo)
	
	if err != nil {
		return nil, &SendRequestError{
			Err: err,
		}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get node info status code %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}
