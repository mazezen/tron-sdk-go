package address

import (
	"bytes"
	"testing"

	"github.com/mazezen/tron-sdk-go/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestAddress_Scan(t *testing.T) {
	validAddress, err := Base58ToAddress("TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	b58 := validAddress.String()
	assert.Equal(t, "TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb", b58)
	t.Logf("address: %s", b58)

	tronHex := validAddress.Hex()
	assert.Equal(t, "41262f06fe4a90740371d99fc6cf448fbb097438aa", tronHex)
	t.Logf("tronHex: %s", tronHex)

	ethHex := validAddress.EthHex()
	assert.Equal(t, "0x262f06fe4a90740371d99fc6cf448fbb097438aa", ethHex)
	t.Logf("ethHex: %s", ethHex)

	base58Address := HexToBase58Address("0x262f06fe4a90740371d99fc6cf448fbb097438aa")
	assert.Equal(t, "TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb", base58Address)
	t.Logf("base58Address: %s", base58Address)

	base58Address2 := HexToBase58Address("41262f06fe4a90740371d99fc6cf448fbb097438aa")
	assert.Equal(t, "TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb", base58Address2)
	t.Logf("base58Address2: %s", base58Address2)

	// correct case
	want := validAddress
	a := &Address{}
	src := validAddress.Bytes()
	err = a.Scan(src)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !bytes.Equal(a.Bytes(), want.Bytes()) {
		t.Errorf("want %v, got %v", want, *a)
	}
	t.Logf("want %v, got %v", want, *a)
}

func TestAddress_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		address  Address
		expected bool
	}{
		{
			name: "valid address",
			address: func() Address {
				addr, _ := Base58ToAddress("TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb")
				return addr
			}(),
			expected: true,
		},
		{
			name:     "nil address",
			address:  nil,
			expected: false,
		},
		{
			name:     "empty address",
			address:  Address{},
			expected: false,
		},
		{
			name:     "wrong address",
			address:  Address{0x41, 0x00, 0x00},
			expected: false,
		},
		{
			name: "wrong prefix",
			address: func() Address {
				addr := make([]byte, AddressLength)
				addr[0] = 0x42
				return Address(addr)
			}(),
			expected: false,
		},
		{
			name: "valid mainnet address",
			address: func() Address {
				addr, _ := Base58ToAddress("TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb")
				return addr
			}(),
			expected: true,
		},
		{
			name: "base58 decode without validation - valid",
			address: func() Address {
				decode, _ := common.Decode("TDT71s4KF5fdg4799CVrTPZbhXz2StwpVb")
				//t.Logf("len decode: %d", len(decode))
				if len(decode) > 4 {
					return Address(decode[:len(decode)-4])
				}
				return Address(decode)
			}(),
			expected: true,
		},
		{
			name: "base58 decode without validation - wrong prefix",
			address: func() Address {
				data := make([]byte, 21)
				data[0] = 0x42
				for i := 1; i < 21; i++ {
					data[i] = byte(i)
				}
				encoded := common.EncodeCheck(data)
				decoded, _ := common.Decode(encoded)
				if len(decoded) > 4 {
					return Address(decoded[:len(decoded)-4])
				}
				return Address(decoded)
			}(),
			expected: false,
		},
		{
			name: "hex decode - valid TRON address",
			address: func() Address {
				hexBytes, _ := common.FromHex("41262f06fe4a90740371d99fc6cf448fbb097438aa")
				return Address(hexBytes)
			}(),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.address.IsValid(); got != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
