package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_TransferContract(t *testing.T) {
	client := New("")

	call := TransferContractParams{
		From:  "0xC4DB2C9DFBCB6AA344793F1DDA7BD656598A06D8",
		To:    "0x95FD23D3D2221CFEF64167938DE5E62074719E54",
		Value: "0x1f4",
	}
	result, err := client.TransferContract(context.Background(), call)
	assert.NoError(t, err)

	t.Logf("result: %v", result)
}

func TestRpcClient_TransferAssetContract(t *testing.T) {
	client := New("")

	call := TransferAssetContractParams{
		From:       "0x5a4bd47536a94b64b0c44652cabda32a993d7f8b",
		To:         "0x95FD23D3D2221CFEF64167938DE5E62074719E54",
		TokenId:    1004536,
		TokenValue: 10,
	}
	result, err := client.TransferAssetContract(context.Background(), call)
	assert.NoError(t, err)

	t.Logf("result: %v", result)
}

func TestRpcClient_CreateSmartContract(t *testing.T) {
	client := New("")

	call := CreateSmartContractParams{
		From:                       "0xC4DB2C9DFBCB6AA344793F1DDA7BD656598A06D8",
		Name:                       "transferTokenContract",
		Gas:                        "0x245498",
		Abi:                        "[{\"constant\":false,\"inputs\":[],\"name\":\"getResultInCon\",\"outputs\":[{\"name\":\"\",\"type\":\"trcToken\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"trcToken\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokenTo\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"msgTokenValueAndTokenIdTest\",\"outputs\":[{\"name\":\"\",\"type\":\"trcToken\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]\n",
		Data:                       "6080604052d3600055d2600155346002556101418061001f6000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305c24200811461005b5780633be9ece71461008157806371dc08ce146100aa575b600080fd5b6100636100b2565b60408051938452602084019290925282820152519081900360600190f35b6100a873ffffffffffffffffffffffffffffffffffffffff600435166024356044356100c0565b005b61006361010d565b600054600154600254909192565b60405173ffffffffffffffffffffffffffffffffffffffff84169082156108fc029083908590600081818185878a8ad0945050505050158015610107573d6000803e3d6000fd5b50505050565bd3d2349091925600a165627a7a72305820a2fb39541e90eda9a2f5f9e7905ef98e66e60dd4b38e00b05de418da3154e7570029",
		ConsumeUserResourcePercent: 100,
		OriginEnergyLimit:          11111111111111,
		Value:                      "0x1f4",
		TokenId:                    1000033,
		TokenValue:                 100000,
	}
	result, err := client.CreateSmartContract(context.Background(), call)
	assert.NoError(t, err)
	t.Logf("result: %v", result)
}

func TestRpcClient_TriggerSmartContract(t *testing.T) {
	client := New("")

	call := TriggerSmartContractParams{
		From:       "0xC4DB2C9DFBCB6AA344793F1DDA7BD656598A06D8",
		To:         "0xf859b5c93f789f4bcffbe7cc95a71e28e5e6a5bd",
		Data:       "0x3be9ece7000000000000000000000000ba8e28bdb6e49fbb3f5cd82a9f5ce8363587f1f600000000000000000000000000000000000000000000000000000000000f42630000000000000000000000000000000000000000000000000000000000000001",
		Gas:        "0x245498",
		Value:      "0xA",
		TokenId:    1000035,
		TokenValue: 20,
	}
	result, err := client.TriggerSmartContract(context.Background(), call)
	assert.NoError(t, err)
	t.Logf("result: %v", result)
}
