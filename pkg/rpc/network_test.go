package sdk_rpc

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestRpcClient_EthGetCode(t *testing.T) {
	client := New("")

	code, err := client.EthGetCode(context.Background(), "0x4170082243784DCDF3042034E7B044D6D342A91360", "latest")

	assert.NoError(t, err, "EthGetCode err")

	t.Logf("code: %s", code)
}

func TestRpcClient_EthProtocolVersion(t *testing.T) {
	client := New("")

	version, err := client.EthProtocolVersion(context.Background())
	assert.NoError(t, err, "EthProtocolVersion err")
	t.Logf("version: %s", version)
}

func TestRpcClient_EthSyncing(t *testing.T) {
	client := New("")

	result, err := client.EthSyncing(context.Background())
	assert.NoError(t, err, "EthSyncing err")
	t.Logf("result: %s", result)
}

func TestRpcClient_NetListening(t *testing.T) {
	client := New("")

	result, err := client.NetListening(context.Background())
	assert.NoError(t, err, "NetListening err")
	t.Logf("result: %v", result)
}

func TestRpcClient_NetPeerCount(t *testing.T) {
	client := New("")
	result, err := client.NetPeerCount(context.Background())
	assert.NoError(t, err, "NetPeerCount err")
	t.Logf("result: %d", result)
}

func TestRpcClient_NetVersion(t *testing.T) {
	client := New("")
	result, err := client.NetVersion(context.Background())
	assert.NoError(t, err, "NetVersion err")
	b, err := common.HexToByte(result[2:])
	assert.NoError(t, err, "HexToByte err")
	t.Logf("result: %s", hex.EncodeToString(b))
}

func TestRpcClient_Web3ClientVersion(t *testing.T) {
	client := New("")
	result, err := client.Web3ClientVersion(context.Background())
	assert.NoError(t, err, "Web3ClientVersion err")
	t.Logf("result: %s", result)
}

func TestRpcClient_Web3Sha3(t *testing.T) {
	client := New("")
	result, err := client.Web3Sha3(context.Background(), "0x68656c6c6f20776f726c64")
	assert.NoError(t, err, "Web3Sha3 err")
	t.Logf("result: %s", result)
}
