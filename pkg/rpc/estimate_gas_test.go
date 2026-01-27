package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthEstiMateGas(t *testing.T) {
	client := New("")

	call := EthEstimateGasParams{
		From: "0x41F0CC5A2A84CD0F68ED1667070934542D673ACBD8",
		To:   "0x4170082243784DCDF3042034E7B044D6D342A91360",
		//Gas:      "0x01",
		GasPrice: "0x0",
		Value:    "0x10000000000",
		Data:     "0x70a08231000000000000000000000000f0cc5a2a84cd0f68ed1667070934542d673acbd8",
	}
	estiMateGas, err := client.EthEstiMateGas(context.Background(), call)
	assert.NoError(t, err, "should not return error")

	t.Log(estiMateGas)
}
