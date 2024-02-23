package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"kapibara-apigateway/internal/config"
	cryptoUtils "kapibara-apigateway/internal/crypto"
	"kapibara-apigateway/internal/logger"
	mysqlsdk "kapibara-apigateway/internal/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AuthLogin(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("password")
	cookieNeed := true

	// empty account/pwd
	if account == "" || pwd == "" {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "LOGIN ERROR: invalid account/password."},
		)
		return
	}

	// get related-account
	record, err := mysqlsdk.GetRecordByAccount(account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "LOGIN ERROR: invalid account: not exists."},
			)
		} else {
			logger.Error(fmt.Sprintf("LOGIN ERROR: user-database search error for [%s].", account))
			c.JSON(
				http.StatusInternalServerError,
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
			http.StatusUnprocessableEntity,
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
			http.StatusInternalServerError,
			gin.H{"error": "server internal error"},
		)
		return
	}

	logger.Debug(fmt.Sprintf("Login success for %s.", account))
	if cookieNeed {
		// [todo] set secure=true when https enabled
		c.SetCookie(
			"_kapibara_access_token",
			jwtToken,
			int(config.GlobalConfig.JWTConf.Expired),
			"/",
			config.GlobalConfig.ServerConf.ServerDomain,
			false,
			true,
		)
		// [todo] move user_info into response-body
		c.SetCookie(
			"_kapibara_user_info",
			record["username"],
			int(config.GlobalConfig.JWTConf.Expired),
			"/",
			config.GlobalConfig.ServerConf.ServerDomain,
			false,
			true,
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"access_token": jwtToken,
				"user_info":    record["username"],
				"token_type":   "Bearer",
				"jit":          uuid.NewV4(),
			},
		)
	}

}

func AuthRegister(c *gin.Context) {
	account := c.PostForm("account")
	username := c.PostForm("username")
	pwd := c.PostForm("password")

	logger.Debug(fmt.Sprintf("REGISTER Entered [%s].", account))

	// empty account/pwd
	if account == "" || pwd == "" {
		logger.Debug(fmt.Sprintf("REGISTER ERROR: bcrypthash error for [%s].", account))
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "REGISTER ERROR: invalid account/password."},
		)
		return
	}

	pwdHash, err := cryptoUtils.BCryptHash(pwd)
	if err != nil {
		logger.Error(fmt.Sprintf("REGISTER ERROR: bcrypthash error for [%s].", account))
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "server internal error"},
		)
		return
	}
	record := make(map[string]string)
	record["account"] = account
	record["username"] = username
	record["pwdHash"] = pwdHash

	_, err = mysqlsdk.RegisterNewUser(record)
	if err != nil {
		logger.Debug(fmt.Sprintf("Register new user error: [%v]", err))
		logger.Debug(fmt.Sprintf("REGISTER ERROR: record - [%v].", record))
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "REGISTER ERROR: user exists."},
		)
		return
	}

	logger.Debug(fmt.Sprintf("Register success for %s.", account))
	c.JSON(
		http.StatusOK,
		gin.H{"msg": "register success."},
	)
}
