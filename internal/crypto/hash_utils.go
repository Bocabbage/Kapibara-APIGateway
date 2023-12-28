package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))

	hasherInByte := hasher.Sum(nil)
	hasherString := hex.EncodeToString(hasherInByte)

	return hasherString
}
