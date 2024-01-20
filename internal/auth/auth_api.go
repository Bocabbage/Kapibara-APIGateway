package auth

import (
	"database/sql"
	"errors"
	"fmt"
	cryptoUtils "kapibara-apigateway/internal/crypto"
	"kapibara-apigateway/internal/logger"
	mysqlsdk "kapibara-apigateway/internal/mysql"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AuthLogin(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("password")

	// empty account/pwd
	if account == "" || pwd == "" {
		c.JSON(
			422,
			gin.H{"error": "LOGIN ERROR: invalid account/password."},
		)
		return
	}

	// get related-account
	record, err := mysqlsdk.GetRecordByAccount(account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(
				422,
				gin.H{"error": "LOGIN ERROR: invalid account: not exists."},
			)
		} else {
			logger.Error(fmt.Sprintf("LOGIN ERROR: user-database search error for [%s].", account))
			c.JSON(
				500,
				gin.H{"error": "server internal error"},
			)
		}
		return
	}

	// check the password
	storedPwdHash := record["pwdHash"]
	compResult := cryptoUtils.BCryptHashCompare(storedPwdHash, pwd)
	if !compResult {
		c.JSON(
			401,
			gin.H{"error": "LOGIN ERROR: incorrect pwd."},
		)
		return
	}

	// generate jwt
	// [todo] enhancement: decoupling jwt-format
	jwtToken, err := cryptoUtils.GenerateJWT(record)
	if err != nil {
		logger.Error(fmt.Sprintf("LOGIN ERROR: generate jwt-token error for [%s].", account))
		c.JSON(
			500,
			gin.H{"error": "server internal error"},
		)
	}

	c.JSON(
		200,
		gin.H{
			"access_token": jwtToken,
			"token_type":   "Bearer",
			"jit":          uuid.NewV4(),
		},
	)

}

func AuthRegister(c *gin.Context) {
	account := c.PostForm("account")
	username := c.PostForm("username")
	pwd := c.PostForm("password")

	// empty account/pwd
	if account == "" || pwd == "" {
		c.JSON(
			422,
			gin.H{"error": "REGISTER ERROR: invalid account/password."},
		)
		return
	}

	pwdHash, err := cryptoUtils.BCryptHash(pwd)
	if err != nil {
		logger.Error(fmt.Sprintf("REGISTER ERROR: bcrypthash error for [%s].", account))
		c.JSON(
			500,
			gin.H{"error": "server internal error"},
		)
		return
	}
	record := make(map[string]string)
	record["account"] = account
	record["username"] = username
	record["pwdHash"] = pwdHash

	err = mysqlsdk.RegisterNewUser(record)
	if err != nil {
		c.JSON(
			422,
			gin.H{"error": "REGISTER ERROR: user exists."},
		)
		return
	}

	c.JSON(
		200,
		gin.H{"msg": "register success."},
	)
}
