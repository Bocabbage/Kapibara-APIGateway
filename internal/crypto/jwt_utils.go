package cryptoutils

import (
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	"strconv"
	"time"

	mysqlsdk "kapibara-apigateway/internal/mysql"

	"github.com/golang-jwt/jwt/v5"
)

func IsValidUser(account, password string) bool {
	pwdHash, status, err := mysqlsdk.GetPwdStatusByAccount(account)

	if err != nil {
		logger.Error(fmt.Sprintf("IsValidUser error: %s", err))
		return false
	}

	// [todo] enhance with enum + status-related return value
	if status != 2 {
		return false
	}

	inputPwdHashed := Sha256Hash(password + config.GlobalConfig.MySQLConf.StrSalt)
	return pwdHash == inputPwdHashed
}

func GenerateJWT(account string) (string, error) {
	record, err := mysqlsdk.GetRecordByAccount(account)
	if err != nil {
		logger.Error("Generate JWT failed")
		return "", err
	}

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
