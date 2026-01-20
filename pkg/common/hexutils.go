package common

import "encoding/hex"

type hexError struct {
	msg string
}

var EmptyString = &hexError{msg: "empty hex string"}

func (e *hexError) Error() string {
	return e.msg
}

func FromHex(s string) ([]byte, error) {
	if Hash0xPrefix(s) {
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

func Hash0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func Hex2Bytes(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
