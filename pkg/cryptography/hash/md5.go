package hash

import "crypto/md5"

// https://www.jianshu.com/p/5b6f7110eb52

func MD5(value []byte) []byte {
	hashValue := md5.New()
	hashValue.Write(value)
	hashResult := hashValue.Sum(nil)
	return hashResult
}

func MD5String(value []byte) string {
	return string(MD5(value))
}

func MD5Verify(value []byte, hashValue string) bool {
	return MD5String(value) == hashValue
}
