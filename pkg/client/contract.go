package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/abi"
	"github.com/mazezen/tron-sdk-go/pkg/common"
	"google.golang.org/protobuf/proto"
)

// DeployContract deploys a contract. Returns TransactionExtention, which contains an unsigned transaction.
// https://developers.tron.network/reference/wallet-deploycontract
func (c *GrpcClient) DeployContract(from, contractName string,
	abi *tronpb.SmartContract_ABI,
	codeStr string,
	feeLimit, curPercent, oeLimit int64,
) (*tronpb.TransactionExtention, error) {
	var err error

	fromAddr, err := c.convert(from)
	if err != nil {
		return nil, err
	}

	if curPercent > 100 || curPercent < 0 {
		return nil, fmt.Errorf("curPercent must be between 0 and 100")
	}
	if oeLimit <= 0 {
		return nil, fmt.Errorf("oeLimit must be greater than zero")
	}

	cs, err := common.FromHex(codeStr)
	if err != nil {
		return nil, err
	}

	ct := &tronpb.CreateSmartContract{
		OwnerAddress: fromAddr.Bytes(),
		NewContract: &tronpb.SmartContract{
			OriginAddress:              fromAddr.Bytes(),
			Abi:                        abi,
			Bytecode:                   cs,
			ConsumeUserResourcePercent: curPercent,
			OriginEnergyLimit:          oeLimit,
		},
		CallTokenValue: 0,
		TokenId:        0,
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.DeployContract(ctx, ct)
	if err != nil {
		return nil, err
	}

	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
	}

	return tx, nil
}

// GetContract Fetches comprehensive information for a specified smart contract deployed on the blockchain.
// The returned details include the contract's bytecode, Application Binary Interface (ABI), and configuration parameters.
// https://developers.tron.network/reference/wallet-getcontract
func (c *GrpcClient) GetContract(value string) (*tronpb.SmartContract, error) {
	var err error
	var in = new(tronpb.BytesMessage)

	in.Value, err = c.convert(value)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetContract(ctx, in)
}

// GetContractInfo queries and returns the complete information of a contract from the blockchain.
// Unlike the GetContract interface, this endpoint returns both the contract's bytecode and its runtime bytecode.
// https://developers.tron.network/reference/getcontractinfo
func (c *GrpcClient) GetContractInfo(value string) (*tronpb.SmartContractDataWrapper, error) {
	var err error
	var in = new(tronpb.BytesMessage)

	in.Value, err = c.convert(value)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetContractInfo(ctx, in)
}

func (c *GrpcClient) TriggerContract(
	from string,
	contractAddress string,
	method string,
	feeLimit, callValue int64,
	jsonData string,
	tokenAmount int64,
	tokenId int64,
) (*tronpb.TransactionExtention, error) {
	if jsonData == "" && method == "" {
		return nil, fmt.Errorf("jsonData or method must be set")
	}

	var err error
	var in = new(tronpb.TriggerSmartContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	if in.ContractAddress, err = c.convert(contractAddress); err != nil {
		return nil, err
	}

	param, err := abi.LoadFromJson(jsonData)
	if err != nil {
		return nil, err
	}

	dataBytes, err := abi.Pack(method, param)
	if err != nil {
		return nil, err
	}
	in.Data = dataBytes

	if callValue > 0 {
		in.CallValue = callValue
	}
	if tokenId > 0 && tokenAmount > 0 {
		in.TokenId = tokenId
		in.CallTokenValue = tokenAmount
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.TriggerContract(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}

	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
	}

	return tx, nil
}

// TriggerConstantContract this interface is used for pre-execution (or simulation) of smart contracts on the blockchain:
// it can call a contract's read-only functions for data queries, call non-read-only functions to predict transaction
// success or estimate required Energy consumption
// https://developers.tron.network/reference/triggerconstantcontract
func (c *GrpcClient) TriggerConstantContract(
	ownerAddress string,
	contractAddress string,
	functionSelector string,
	jsonString string,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.TriggerSmartContract)

	if in.OwnerAddress, err = c.convert(ownerAddress); err != nil {
		return nil, err
	}

	if in.ContractAddress, err = c.convert(contractAddress); err != nil {
		return nil, err
	}

	param, err := abi.LoadFromJson(jsonString)
	if err != nil {
		return nil, err
	}

	dataBytes, err := abi.Pack(functionSelector, param)
	if err != nil {
		return nil, err
	}

	in.Data = dataBytes

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.TriggerConstantContract(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// EstimateEnergy estimate the energy required for the successful execution of smart contract transactions or deploying a contract
// https://developers.tron.network/reference/estimateenergy
func (c *GrpcClient) EstimateEnergy(
	from string,
	contractAddress string,
	functionSelector string,
	callValue int64,
	jsonData string,
	tokenAmount int64,
	tokenId int64,
) (*tronpb.EstimateEnergyMessage, error) {
	var err error
	var in = new(tronpb.TriggerSmartContract)
	if jsonData == "" && functionSelector == "" {
		return nil, fmt.Errorf("jsonData or method must be set")
	}

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ContractAddress, err = c.convert(contractAddress); err != nil {
		return nil, err
	}
	param, err := abi.LoadFromJson(jsonData)
	if err != nil {
		return nil, err
	}
	dataBytes, err := abi.Pack(functionSelector, param)
	if err != nil {
		return nil, err
	}
	in.Data = dataBytes
	in.CallValue = callValue
	in.CallTokenValue = tokenAmount
	in.TokenId = tokenId

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.EstimateEnergy(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// ClearContractABI clear from address contract abi
func (c *GrpcClient) ClearContractABI(from, contractAddress string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ClearABIContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	if in.ContractAddress, err = c.convert(contractAddress); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	tx, err := c.WalletClient.ClearContractABI(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}
