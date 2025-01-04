package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func CreateStringForHash(args ...string) string {
	var sb strings.Builder
	for _, str := range args {
		sb.WriteString(str)
	}
	return sb.String()
}

func CreateHashHMAC6(message string, key string) string {
	keyBytes := []byte(key)

	h := hmac.New(sha256.New, keyBytes)

	h.Write([]byte(message))

	hmacBytes := h.Sum(nil)

	return hex.EncodeToString(hmacBytes)
}
