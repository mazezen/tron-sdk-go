package api

import (
	"context"
	"fmt"
	"net/http"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/api/internet"
	"github.com/mazezen/tron-sdk-go/pkg/api/methods"
)

// GetContract fetches comprehensive information for a specified smart contract deployed on the blockchain.
// The returned details include the contract's bytecode, Application Binary Interface (ABI), and configuration parameters.
// -------------------------------------------------------------
// https://developers.tron.network/reference/wallet-getcontract
// -------------------------------------------------------------
func (a *Client) GetContract(ctx context.Context, value string, visible bool) (*internet.ContractInternet, error) {
	var result = new(internet.ContractInternet)
	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   value,
		"visible": visible,
	}).SetResult(result).Post(methods.GetContract)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get contract status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result, nil
}

// GetContractInfo
// -------------------------------------------------------------
// https://developers.tron.network/reference/getcontractinfo
// -------------------------------------------------------------
func (a *Client) GetContractInfo(ctx context.Context, value string, visible bool) (*internet.ContractInfoInternet, error) {
	var result = new(internet.ContractInfoInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"value":   value,
		"visible": visible,
	}).SetResult(result).Post(methods.GetContractInfo)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("get contract info status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// TriggerSmartContract returns a Transaction object, which encapsulates the unsigned transaction data
// --------------------------------------------------------------
// https://developers.tron.network/reference/triggersmartcontract
// --------------------------------------------------------------
func (a *Client) TriggerSmartContract(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	functionSelector string,
	parameter string,
	data string,
	feeLimit string,
	callValue int64,
	callTokenValue int64,
	tokenId int64,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":     ownerAddress,
		"contract_address":  contractAddress,
		"function_selector": functionSelector,
		"parameter":         parameter,
		"data":              data,
		"fee_limit":         feeLimit,
		"call_value":        callValue,
		"call_token_value":  callTokenValue,
		"token_id":          tokenId,
		"visible":           visible,
	}).SetResult(result).Post(methods.TriggerSmartContract)

	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("trigger smartcontract status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// TriggerConstantContract
// --------------------------------------------------------------
// https://developers.tron.network/reference/triggersmartcontract
// --------------------------------------------------------------
func (a *Client) TriggerConstantContract(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	functionSelector string,
	parameter string,
	data string,
	callValue int64,
	callTokenValue int64,
	tokenId int64,
	visible bool,
) (*internet.TriggerConstantContractInternet, error) {
	var result = new(internet.TriggerConstantContractInternet)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":     ownerAddress,
		"contract_address":  contractAddress,
		"function_selector": functionSelector,
		"parameter":         parameter,
		"data":              data,
		"call_value":        callValue,
		"call_token_value":  callTokenValue,
		"token_id":          tokenId,
		"visible":           visible,
	}).SetResult(result).Post(methods.TriggerConstantContract)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("trigger constantcontract status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Result.Code != tronpb.ReturnResponseCode_name[0] {
		return nil, &SendRequestError{Err: fmt.Errorf("code: %s error: %s", result.Result.Code, result.Result.Message)}
	}

	return result, nil
}

// DeployContract deploys a contract. Returns Transaction, which contains an unsigned transaction.
// ---------------------------------------------------------------
// https://developers.tron.network/reference/wallet-deploycontract
// ---------------------------------------------------------------
func (a *Client) DeployContract(
	ctx context.Context,
	ownerAddress string,
	abiJson string,
	bytecode string,
	feeLimit int64,
	parameter string,
	originEnergyLimit int64,
	name string,
	callValue int64,
	consumerUserResourcePercent int64,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":                  ownerAddress,
		"abi":                            abiJson,
		"bytecode":                       bytecode,
		"fee_limit":                      feeLimit,
		"parameter":                      parameter,
		"origin_energy_limit":            originEnergyLimit,
		"name":                           name,
		"call_value":                     callValue,
		"consumer_user_resource_percent": consumerUserResourcePercent,
		"visible":                        visible,
	}).SetResult(result).Post(methods.DeployContract)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("deploy contract status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// UpdateSetting
// ---------------------------------------------------------------
// https://developers.tron.network/reference/wallet-updatesetting
// ---------------------------------------------------------------
func (a *Client) UpdateSetting(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	consumeUserResourcePercent int64,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":                 ownerAddress,
		"contract_address":              contractAddress,
		"consume_user_resource_percent": consumeUserResourcePercent,
		"visible":                       visible,
	}).SetResult(result).Post(methods.UpdateSetting)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}
	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update setting status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{
			Err: fmt.Errorf("%s", result.Error),
		}
	}

	return result, nil
}

// UpdateEnergyLimit update the smart contract's origin_energy_limit parameter
// ------------------------------------------------------------------
// https://developers.tron.network/reference/wallet-updateenergylimit
// ------------------------------------------------------------------
func (a *Client) UpdateEnergyLimit(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	originEnergyLimit int64,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":       ownerAddress,
		"contract_address":    contractAddress,
		"origin_energy_limit": originEnergyLimit,
		"visible":             visible,
	}).SetResult(result).Post(methods.UpdateEnergyLimit)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("update energylimit status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// ClearAbi remove or delete the ABI information associated with the smart contract.
// ------------------------------------------------------------------
// https://developers.tron.network/reference/clearabi
// ------------------------------------------------------------------
func (a *Client) ClearAbi(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	visible bool,
) (*internet.Transaction, error) {
	var result = new(internet.Transaction)

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":    ownerAddress,
		"contract_address": contractAddress,
		"visible":          visible,
	}).SetResult(result).Post(methods.ClearAbi)
	if err != nil {
		return nil, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return nil, &SendRequestError{
			Err: fmt.Errorf("clear abi status code: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	if result.Error != "" {
		return nil, &SendRequestError{Err: fmt.Errorf("%s", result.Error)}
	}

	return result, nil
}

// EstimateEnergy estimate the energy required for the successful execution of smart contract transactions or deploying a contract
// ------------------------------------------------------------------
// https://developers.tron.network/reference/estimateenergy
// ------------------------------------------------------------------
func (a *Client) EstimateEnergy(
	ctx context.Context,
	ownerAddress string,
	contractAddress string,
	functionSelector string,
	parameter string,
	data string,
	callValue int64,
	callTokenValue int64,
	tokenId int64,
	visible bool,
) (int64, error) {
	var result struct {
		Result         internet.Result `json:"result,omitempty"`
		EnergyRequired int64           `json:"energy_required,omitempty"`
	}

	rs, err := a.client.R().SetContext(ctx).SetBody(map[string]interface{}{
		"owner_address":     ownerAddress,
		"contract_address":  contractAddress,
		"function_selector": functionSelector,
		"parameter":         parameter,
		"data":              data,
		"call_value":        callValue,
		"call_token_value":  callTokenValue,
		"token_id":          tokenId,
		"visible":           visible,
	}).SetResult(&result).Post(methods.EstimateEnergy)
	if err != nil {
		return 0, &SendRequestError{Err: err}
	}

	if rs.StatusCode() != http.StatusOK {
		return 0, &SendRequestError{
			Err: fmt.Errorf("estimate energy: %d error: %s", rs.StatusCode(), rs.Status()),
		}
	}

	return result.EnergyRequired, nil
}
