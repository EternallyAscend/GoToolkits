package hash

import "crypto"
import "golang.org/x/crypto/sha3"

// https://pkg.go.dev/golang.org/x/crypto/sha3

func SHA3512(value []byte) []byte {
	hashValue := crypto.SHA3_512.New()
	hashValue.Write(value)
	hashResult := hashValue.Sum(nil)
	return hashResult
}

func SHA3512String(value []byte) string {
	return string(SHA3512(value))
}

func SHA3512Verify(value []byte, hashValue string) bool {
	return SHA3512String(value) == hashValue
}

// 非标准SHA3 Keccak

func Keccak512(value []byte) []byte {
	hashValue := sha3.NewLegacyKeccak512()
	hashValue.Write(value)
	hashResult := hashValue.Sum(nil)
	return hashResult
}

func Keccak512String(value []byte) string {
	return string(Keccak512(value))
}

func Keccak512Verify(value []byte, hashValue string) bool {
	return Keccak512String(value) == hashValue
}
