package RSA

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateRandomKeyOfRSA() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if nil != err {
		return nil, nil, err
	}
	return privateKey, &(privateKey.PublicKey), nil
}

func EncryptByPublicRSA(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
}

func DecryptByPrivateRSA(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return privateKey.Decrypt(nil, data, &rsa.OAEPOptions{
		Hash:  crypto.SHA256,
		Label: nil,
	})
}

func GenerateSha256SignWithPrivateKeyRSA(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashValue := sha256.New()
	_, err := hashValue.Write(data)
	if nil != err {
		return nil, err
	}
	hashSum := hashValue.Sum(nil)
	return rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashSum, nil)
}

func VerifySignSha256WithPublicKeyRSA(publicKey *rsa.PublicKey, data []byte, sign []byte) error {
	hashValue := sha256.New()
	_, err := hashValue.Write(data)
	if nil != err {
		return err
	}
	hashSum := hashValue.Sum(nil)
	return rsa.VerifyPSS(publicKey, crypto.SHA256, hashSum, sign, nil)
}

// TODO Finish EncryptByPrivateRSA and DecryptByPublicKeyRSA with reference behind. Actually, modify the rsa.OAEP or rsa.PKCS1v15.
// https://github.com/wenzhenxi/gorsa/blob/528c7050d7030d782cfbcfa37103c68e80827fec/rsa_ext.go#L207

func EncryptByPrivateKeyRSA(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return nil, nil
}

// https://github.com/wenzhenxi/gorsa/blob/528c7050d7030d782cfbcfa37103c68e80827fec/rsa_ext.go#L176

func DecryptByPublicKeyRSA(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return nil, nil
}
