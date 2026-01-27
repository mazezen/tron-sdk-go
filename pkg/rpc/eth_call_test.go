package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthCall(t *testing.T) {
	client := New("")

	usdtContract := "0xa614f803b6fd780986a42c78ec9c7f77e6ded13c"

	// userAddr := "0x8f8f1af0e596777f332be2a0e52701b010005d22"
	// 0x70a08231 (balanceOf) + 8f8f1af0e596777f332be2a0e52701b010005d22
	data := "0x70a082318f8f1af0e596777f332be2a0e52701b010005d22"

	req := EthCallParams{
		To:   usdtContract,
		Data: data,
	}
	resultHex, err := client.EthCall(context.Background(), req, "latest")
	assert.NoError(t, err, "should not error")
	assert.NotEmpty(t, resultHex)

	t.Logf("resultHex: %s", resultHex)
}
