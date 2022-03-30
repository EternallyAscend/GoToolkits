package RSA

import (
	"fmt"
	"testing"
)

const RSATest = "EternallyAscend"

func TestRSA(t *testing.T) {
	// Key Gen.
	privateKey, publicKey, err := GenerateRandomKeyOfRSA()
	if nil != err {
		t.Fail()
	}
	fmt.Println(privateKey)
	fmt.Println(publicKey)

	// Encrypt.
	encrypted, err := EncryptByPublicRSA(publicKey, []byte(RSATest))
	if nil != err {
		t.Fail()
	}
	fmt.Println(string(encrypted))

	// Decrypt.
	decrypted, err := DecryptByPrivateRSA(privateKey, encrypted)
	if nil != err {
		t.Fail()
	}
	fmt.Println(string(decrypted))
	if string(decrypted) != RSATest {
		t.Fail()
	}

	// 256 Sign
	sign, err := GenerateSha256SignWithPrivateKeyRSA(privateKey, []byte(RSATest))
	if nil != err {
		t.Fail()
	}
	fmt.Println(string(sign))
	err = VerifySignSha256WithPublicKeyRSA(publicKey, []byte(RSATest), sign)
	if nil != err {
		fmt.Println("Sign 256.")
		t.Fail()
	}
}
