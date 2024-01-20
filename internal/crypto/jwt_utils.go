package cryptoutils

import (
	"kapibara-apigateway/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(record map[string]string) (string, error) {
	roleBitmap, _ := strconv.ParseInt(record["roleBitmap"], 2, 64)
	currentTime := time.Now().Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": record["username"],
			"exp":      currentTime + config.GlobalConfig.JWTConf.Expired,
			"roles":    roleBitmap + config.GlobalConfig.MySQLConf.Salt,
		},
	)

	secretKey := config.GlobalConfig.JWTConf.SecretKey
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(token string) {
	// [todo] implement
}
