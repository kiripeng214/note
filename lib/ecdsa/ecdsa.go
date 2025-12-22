package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenPrivateKey() (privateKeyStr string, publicKeyStr string, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	ecPrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	privatePem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: ecPrivateKey,
	}
	privateKeyStr = string(pem.EncodeToMemory(privatePem))

	publickKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	publickPem := &pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: publickKey,
	}
	publicKeyStr = string(pem.EncodeToMemory(publickPem))

	return privateKeyStr, publicKeyStr, nil
}

func PublicKeyFromStr(str string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(str))
	if block == nil {
		return nil, errors.New("ErrAuthKeyNotPem")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	return publicKey, nil
}
