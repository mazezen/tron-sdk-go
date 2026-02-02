package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// ListProposals all proposals
// -------------------------------------------------------------
// https://developers.tron.network/reference/wallet-listproposals
// -------------------------------------------------------------
func (a *Client) ListProposals(ctx context.Context) (*internet.ProposalListInternet, error) {
	var result = new(internet.ProposalListInternet)

	rs, err := a.client.R().SetContext(ctx).SetResult(result).Get(methods.ListProposals)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{Err: fmt.Errorf("list proposals status code: %d error: %s", rs.StatusCode(), rs.Status())}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// GetProposalById queries proposal based on ID and returns proposal details
// -------------------------------------------------------------
// https://developers.tron.network/reference/getproposalbyid
// -------------------------------------------------------------
func (a *Client) GetProposalById(ctx context.Context, id int32, visible bool) (*internet.ProposalInternet, error) {
	var result = new(internet.ProposalInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"id":      id,
		"visible": visible,
	}).SetResult(result).Get(methods.GetProposalById)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{Err: fmt.Errorf("get proposal by id status code: %d error: %s", rs.StatusCode(), rs.Status())}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// ProposalCreate creates a proposal transaction
// -------------------------------------------------------------
// https://developers.tron.network/reference/proposalcreate
// -------------------------------------------------------------
func (a *Client) ProposalCreate(
	ctx context.Context,
	ownerAddress string,
	Parameters string,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"parameters":    Parameters,
		"visible":       visible,
	}).SetResult(result).Post(methods.ProposalCreate)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{Err: fmt.Errorf("create proposal status code: %d error: %s", rs.StatusCode(), rs.Status())}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// ProposalApprove approves proposed transaction
// -------------------------------------------------------------
// https://developers.tron.network/reference/proposalapprove
// -------------------------------------------------------------
func (a *Client) ProposalApprove(
	ctx context.Context,
	ownerAddress string,
	proposalId int32,
	isAddApprove bool,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":  ownerAddress,
		"proposal_id":    proposalId,
		"is_add_approve": isAddApprove,
		"visible":        visible,
	}).SetResult(result).Post(methods.ProposalApprove)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{Err: fmt.Errorf("proposal approve status code: %d error: %s", rs.StatusCode(), rs.Status())}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// ProposalDelete deletes Proposal Transaction
// -------------------------------------------------------------
// https://developers.tron.network/reference/proposaldelete
// -------------------------------------------------------------
func (a *Client) ProposalDelete(
	ctx context.Context,
	ownerAddress string,
	proposalId int32,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address": ownerAddress,
		"proposal_id":   proposalId,
		"visible":       visible,
	}).SetResult(result).Post(methods.ProposalDelete)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("proposal delete status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}
