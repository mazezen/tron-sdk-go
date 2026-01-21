package mnemonic

import "testing"

func TestGenerate(t *testing.T) {
	t.Logf("12 mnemonic: %s", Generate12())
	t.Logf("24 mnemonic: %s", Generate24())
}
