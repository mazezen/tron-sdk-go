package common

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	decodeBytes, err := DecodeCheck("TTiWGY3myYPyjTYrBHeACfvRaGppNbLVr9")
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5", hex.EncodeToString(decodeBytes))
	t.Logf("decodeBytes: %x", decodeBytes)
	t.Logf("decodeBytes len: %d", len(decodeBytes))

	decode, err := Decode("TTiWGY3myYPyjTYrBHeACfvRaGppNbLVr9")
	t.Logf("decode len is: %d", len(decode))
	assert.NoError(t, err)
	assert.Equal(t, "41c2aa7035af217ee90e227b99f60cbf2ca45c84d5e285928a", hex.EncodeToString(decode))
	t.Logf("decode: %x", decode)

	check := EncodeCheck(decode)
	t.Logf("check: %s", check)

	base58 := Encode(decode)
	assert.NoError(t, err)
	assert.Equal(t, "TTiWGY3myYPyjTYrBHeACfvRaGppNbLVr9", base58)
	t.Logf("base58: %s", base58)
}
