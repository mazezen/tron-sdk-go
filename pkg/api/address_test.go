package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiClient_ValidateAddress(t *testing.T) {
	response, err := NewApiClient("").ValidateAddress(context.Background(), "TTvDaNWWGRWUa4nEwnaM88bPvPiF4RuR4T", true)
	assert.NoError(t, err, "ValidateAddress should not error")
	assert.NotNil(t, response, "ValidateAddress response should not be nil")

	assert.Equal(t, true, response.Result, "ValidateAddress response should be true")
	assert.Equal(t, "Base58check format", response.Message, "ValidateAddress response message should match")
}
