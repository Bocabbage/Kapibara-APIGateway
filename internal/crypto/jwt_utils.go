package cryptoutils

import (
	mysqlsdk "kapibara-apigateway/internal/mysql"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func IsValidUser(username, password string) bool {
	pwdHash, err := mysqlsdk.GetPwdByUsername(username)

	if err != nil {
		log.Printf("IsValidUser error: %s", err)
		return false
	}

	inputPwdHashed := Sha256Hash(password)
	return pwdHash == inputPwdHashed
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			// [todo] enrich jwt-payload
			"username": username,
		},
	)

	secretKey := "mock-secret-key"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
