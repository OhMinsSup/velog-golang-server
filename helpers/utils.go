package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// CreateHash 해시를 생성하는 함수
func CreateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

// GetEnvWithKey process.env 에 key값과 일치하는 값을 가져온다
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
