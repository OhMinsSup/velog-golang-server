package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func CreateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
