package cryptoutils

import (
	"fmt"
	"kapibara-apigateway/internal/config"
	kerrors "kapibara-apigateway/internal/errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Username string
	Exp      int64
	Roles    int64
}

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

func ParseJWT(tokenString string) (*JwtToken, error) {
	var resultJwtToken = JwtToken{
		Username: "",
		Exp:      0,
		Roles:    0,
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := config.GlobalConfig.JWTConf.SecretKey
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if resultJwtToken.Username, ok = claims["username"].(string); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: "jwttoken: username parse failed",
			}
		}

		if resultJwtToken.Exp, ok = claims["exp"].(int64); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: "jwttoken: exp parse failed",
			}
		}

		if resultJwtToken.Roles, ok = claims["roles"].(int64); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: "jwttoken: roles parse failed",
			}
		}
	} else {
		return nil, &kerrors.KapibaraGeneralError{
			Code:    kerrors.JwtTokenParseError,
			Message: "jwttoken: invalid token format",
		}
	}

	return &resultJwtToken, nil
}
