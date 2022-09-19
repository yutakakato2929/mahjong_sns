package utility

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashStr(trg string) string {
	hashed := ""
	b := []byte(trg)
	sha256 := sha256.Sum256(b)
	hashed = hex.EncodeToString(sha256[:])
	return hashed
}
