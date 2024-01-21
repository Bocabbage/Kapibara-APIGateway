package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Sha256Hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))

	hasherInByte := hasher.Sum(nil)
	hasherString := hex.EncodeToString(hasherInByte)

	return hasherString
}

func BCryptHash(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 12)
	return string(bytes), err
}

func BCryptHashCompare(hash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
