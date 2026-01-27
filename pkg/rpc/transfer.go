package sdk_rpc

import "context"

type TransferContractParams struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

// TransferContract transfer TRX
// https://developers.tron.network/reference/buildtransaction
func (r *RpcClient) TransferContract(ctx context.Context, call TransferContractParams) (*TransferResult, error) {
	var result = new(TransferResult)
	err := r.Call(ctx, BuildTransaction, []interface{}{call}, result)
	return result, err
}

type TransferAssetContractParams struct {
	From       string `json:"from"`
	To         string `json:"to"`
	TokenId    int64  `json:"tokenId"`
	TokenValue int64  `json:"tokenValue"`
}

// TransferAssetContract transfer trc10
func (r *RpcClient) TransferAssetContract(ctx context.Context, call TransferAssetContractParams) (*TransferResult, error) {
	var res = new(TransferResult)
	err := r.Call(ctx, BuildTransaction, []interface{}{call}, res)
	return res, err
}

type CreateSmartContractParams struct {
	From                       string `json:"from"`
	Name                       string `json:"name"`
	Gas                        string `json:"gas"`
	Abi                        string `json:"abi"`
	Data                       string `json:"data"`
	ConsumeUserResourcePercent int64  `json:"consumeUserResourcePercent"`
	OriginEnergyLimit          int64  `json:"originEnergyLimit"`
	Value                      string `json:"value"`
	TokenId                    int64  `json:"tokenId"`
	TokenValue                 int64  `json:"tokenValue"`
}

// CreateSmartContract create smart contract
func (r *RpcClient) CreateSmartContract(ctx context.Context, call CreateSmartContractParams) (*CreateSmartContractResult, error) {
	var result = new(CreateSmartContractResult)
	err := r.Call(ctx, BuildTransaction, []interface{}{call}, result)
	return result, err
}

type TriggerSmartContractParams struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Data       string `json:"data"`
	Gas        string `json:"gas"`
	Value      string `json:"value"`
	TokenId    int64  `json:"tokenId"`
	TokenValue int64  `json:"tokenValue"`
}

// TriggerSmartContract transfer trc20
func (r *RpcClient) TriggerSmartContract(ctx context.Context, call TriggerSmartContractParams) (*TriggerSmartContractResult, error) {
	var result = new(TriggerSmartContractResult)
	err := r.Call(ctx, BuildTransaction, []interface{}{call}, result)
	return result, err
}
