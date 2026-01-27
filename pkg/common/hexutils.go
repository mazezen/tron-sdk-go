package common

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type hexError struct {
	msg string
}

var EmptyString = &hexError{msg: "empty hex string"}

func (e *hexError) Error() string {
	return e.msg
}

func FromHex(s string) ([]byte, error) {
	if Has0xPrefix(s) {
		s = s[2:]
	}
	// Hexadecimal strings must be of even length to be decoded correctly into bytes
	// as every two hexadecimal characters correspond to one byte (8 bits)
	// Adding a leading 0 does not change the value itself (0x0a=0xa=10)
	// while adding a trailing 0 will multiply by 16 (becoming 0xa0=160)
	// which will significantly alter the meaning.
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

func Has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}
func Has41Prefix(s string) bool {
	return len(s) >= 2 && s[0] == '4' && s[1] == '1'
}

func HasTPrefix(s string) bool {
	return len(s) == 34 && s[0] == 'T'
}

func Hex2Bytes(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// HexStringToBytes hex string as bytes
func HexStringToBytes(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, EmptyString
	}
	return hex.DecodeString(strings.Replace(input, "0x", "", -1))
}

func HexToByte(s string) ([]byte, error) { return hex.DecodeString(s) }

func EthHexToByte(s string) ([]byte, error) {
	return hex.DecodeString(strings.Replace(s, "0x", "41", -1))
}

func BytesToHexString(bytes []byte) string {
	encode := make([]byte, len(bytes)*2)
	hex.Encode(encode, bytes)
	return string(encode)
}

func BytesToEthHexString(bytes []byte) string {
	encode := make([]byte, len(bytes)*2)
	hex.Encode(encode, bytes)
	return "0x" + string(encode[2:])
}

func HexToUint64(hex string) uint64 {
	hex = strings.Replace(hex, "0x", "", -1)
	n := new(big.Int)
	n, _ = n.SetString(hex, 16)
	return n.Uint64()
}

func Int64ToHex(b int64) string {
	formatInt := strconv.FormatInt(b, 16)
	return "0x" + formatInt
}

func HexToBigInt(hex string) (*big.Int, error) {
	if string(hex) == "" {
		hex = "0x0"
	}
	value, ok := new(big.Int).SetString(hex, 0)
	if !ok {
		return nil, fmt.Errorf("failed to parse quantity %s", hex)
	}
	return value, nil
}
