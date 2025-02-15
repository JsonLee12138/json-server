package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256String(str string) string {
	b32 := sha256.Sum256([]byte(str))
	return hex.EncodeToString(b32[:])
}
