package ECC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

const (
	EccKeySize224 = 224
	EccKeySize256 = 256
	EccKeySize384 = 384
	EccKeySize521 = 521
)

// EccKeyFileGenerate Deprecated
// 被遗弃的函数
func EccKeyFileGenerate(keySize int, outputPath string) error {
	var privateKey *ecdsa.PrivateKey
	err := errors.New("Cryptography/ECC.go/EccKeyFileGenerate: Wrong key size input. ")
	switch keySize {
	case EccKeySize224:
		privateKey, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
		break
	case EccKeySize256:
		privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		break
	case EccKeySize384:
		privateKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		break
	case EccKeySize521:
		privateKey, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		break
	default:
		return err
	}
	if nil != err {
		return err
	}
	if nil == privateKey {
		return errors.New("Cryptography/ECC.go/EccKeyFileGenerate: Private key is nil. ")
	}

	blockText, err := x509.MarshalECPrivateKey(privateKey)
	if nil != err {
		return err
	}
	block := &pem.Block{
		Type:    "ecdsa private key",
		Headers: nil,
		Bytes:   blockText,
	}
	privateKeyFile, err := os.Create(outputPath + "-EccPri.pem")
	defer func(privateKeyFile *os.File) {
		_ = privateKeyFile.Close()
	}(privateKeyFile)
	if nil != err {
		return err
	}
	err = pem.Encode(privateKeyFile, block)

	publicKey := privateKey.Public()
	blockText, err = x509.MarshalPKIXPublicKey(&publicKey)
	block = &pem.Block{
		Type:    "ecdsa public key",
		Headers: nil,
		Bytes:   blockText,
	}
	publicKeyFile, err := os.Create(outputPath + "-EccPub.pem")
	defer func(publicKeyFile *os.File) {
		_ = publicKeyFile.Close()
	}(publicKeyFile)
	if nil != err {
		return err
	}
	return pem.Encode(publicKeyFile, block)
}
