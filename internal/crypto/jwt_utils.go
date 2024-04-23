package cryptoutils

import (
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/data/mysql"
	kerrors "kapibara-apigateway/internal/errors"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Username string
	Account  string
	Exp      int64
	Roles    int64
}

func GenerateJWT(record *mysql.Users) (string, error) {
	expTime := time.Now().Add(time.Duration(config.GlobalConfig.JWTConf.Expired) * time.Second).Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": record.UserName,
			"account":  record.Account,
			"exp":      expTime,
			"roles":    record.RoleBitmap + config.GlobalConfig.MySQLConf.Salt,
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
		Account:  "",
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

		if resultJwtToken.Account, ok = claims["account"].(string); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: "jwttoken: account parse failed",
			}
		}

		if tmpExp, ok := claims["exp"].(float64); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: fmt.Sprintf("jwttoken: exp parse failed, claims: %v, %v", claims, reflect.TypeOf(claims["exp"])),
			}
		} else {
			resultJwtToken.Exp = int64(tmpExp)
		}

		if tmpRoles, ok := claims["roles"].(float64); !ok {
			return nil, &kerrors.KapibaraGeneralError{
				Code:    kerrors.JwtTokenParseError,
				Message: fmt.Sprintf("jwttoken: roles parse failed: %v, %v", claims, reflect.TypeOf(claims["roles"])),
			}
		} else {
			resultJwtToken.Roles = int64(tmpRoles)
		}

	} else {
		return nil, &kerrors.KapibaraGeneralError{
			Code:    kerrors.JwtTokenParseError,
			Message: "jwttoken: invalid token format",
		}
	}

	return &resultJwtToken, nil
}
