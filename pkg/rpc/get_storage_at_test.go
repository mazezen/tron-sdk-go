package sdk_rpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthGetStorageAt(t *testing.T) {
	client := New("")

	call := EthGetStorageAtParams{
		Address: "0xE94EAD5F4CA072A25B2E5500934709F1AEE3C64B",
		Index:   "0x29313b34b1b4beab0d3bad2b8824e9e6317c8625dd4d9e9e0f8f61d7b69d1f26",
		Tag:     "latest",
	}
	res, err := client.EthGetStorageAt(context.Background(), call)
	assert.NoError(t, err, "GetStorageAt err")
	assert.NotNil(t, res, "GetStorageAt result is nil")

	t.Logf("GetStorageAt result: %s", res)
}
