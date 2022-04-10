package hash

import "crypto/sha512"

func SHA512(value []byte) []byte {
	hashValue := sha512.New()
	hashValue.Write(value)
	hashResult := hashValue.Sum(nil)
	return hashResult
}

func SHA512String(value []byte) string {
	return string(SHA512(value))
}

func SHA512Verify(value []byte, hashValue string) bool {
	return SHA512String(value) == hashValue
}
