package client

import (
	"encoding/json"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_ListNodes(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	nodes, err := client.ListNodes()
	assert.NoError(t, err, "failed to list nodes")

	for _, node := range nodes.Nodes {
		t.Logf("node: %v", node)
	}
}

func TestGrpcClient_GetChainParameters(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()
	p, err := client.GetChainParameters()
	assert.NoError(t, err, "failed to list nodes")

	j, _ := json.Marshal(p)
	t.Logf("chain parameters: %s", common.JSONPrettyFormat(string(j)))
}
