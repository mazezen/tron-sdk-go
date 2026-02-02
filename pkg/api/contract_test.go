package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetContract(t *testing.T) {

	value := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	response, err := NewApiClient("").GetContract(context.Background(), value, true)
	assert.NoError(t, err, fmt.Sprintf("get contract should not err: %v", err))
	assert.NotNil(t, value, response, fmt.Sprintf("response should not be nil"))

	t.Logf("response: %v", response)
	for s, i := range response.ABI {
		t.Logf("k: %s, v: %v\n", s, i)
	}
}
