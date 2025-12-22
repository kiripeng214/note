package ecdsa_test

import (
	"lib/ecdsa"
	"testing"
)

func TestGenEcdsaPrivateKey(t *testing.T) {
	privte, public, err := ecdsa.GenPrivateKey()
	t.Log(privte, public, err)
}

func TestGenEcdsaPublicKey(t *testing.T) {
	public := `-----BEGIN EC PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMYW5N/LS9RRQJVpZcuVMpwnSvOT7
CQNs8MNue3eMkqh1I+EDvlvA4yaGBp0H2ZEMXEU4gy8Npa+BNMhmnHUJHA==
-----END EC PUBLIC KEY-----`
	key, err := ecdsa.PublicKeyFromStr(public)
	t.Log(key, err)
}
