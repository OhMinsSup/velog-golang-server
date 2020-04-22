package helpers

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func EscapeForUrl(text string) string {
	var r1 = regexp.MustCompile(`/[^0-9a-zA-Zㄱ-힣.\u3000-\u303f\u3040-\u309f\u30a0-\u30ff\uff00-\uff9f\u4e00-\u9faf\u3400-\u4dbf -]/g`)
	var r2 = regexp.MustCompile(`/ /g`)
	var r3 = regexp.MustCompile(`/--+/g`)
	var r4 = regexp.MustCompile(`/\.+$/`)
	return r4.ReplaceAllString(r3.ReplaceAllString(r2.ReplaceAllString(strings.TrimSpace(r1.ReplaceAllString(text, "")), "-"), "-"), "")
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateStringName(length int) string {
	return StringWithCharset(length, charset)
}
