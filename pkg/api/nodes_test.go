package api

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListNodes(t *testing.T) {
	nodes, err := NewApiClient("").ListNodes()
	assert.NoError(t, err)
	for _, node := range nodes.Nodes {
		t.Logf("[]host: %s, []port: %d", node.Address.Host, node.Address.Port)
	}
}

func TestClient_GetNodeInfo(t *testing.T) {
	nodeInfo, err := NewApiClient("").GetNodeInfo(context.Background())
	assert.NoError(t, err)

	j, _ := json.Marshal(nodeInfo)

	t.Logf("%s", common.JSONPrettyFormat(string(j)))
}
