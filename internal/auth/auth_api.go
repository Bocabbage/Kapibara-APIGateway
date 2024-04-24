package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"kapibara-apigateway/internal/config"
	cryptoUtils "kapibara-apigateway/internal/crypto"
	mysqlsdk "kapibara-apigateway/internal/data/mysql"
	mysqlsdkMod "kapibara-apigateway/internal/data/mysql/models"
	"kapibara-apigateway/internal/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AuthLogin(c *gin.Context) {
	var err error
	account := c.PostForm("account")
	pwd := c.PostForm("password")

	// empty account/pwd
	if account == "" || pwd == "" {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "LOGIN ERROR: invalid account/password."},
		)
		return
	}

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
	bCryptTimeStart := time.Now()
	storedPwdHash := record.PwdHash
	compResult := cryptoUtils.BCryptHashCompare(storedPwdHash, pwd)
	logger.Debug(fmt.Sprintf("bcrypt: cost %.3f s.", time.Since(bCryptTimeStart).Seconds()))
	if !compResult {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "LOGIN ERROR: incorrect pwd."},
		)
		return
	}

	// generate jwt
	// [todo] enhancement: decoupling jwt-format
	genJwtTimeStart := time.Now()
	jwtToken, err := cryptoUtils.GenerateJWT(record)
	if err != nil {
		logger.Error(fmt.Sprintf("LOGIN ERROR: generate jwt-token error for [%s].", account))
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "server internal error"},
		)
		return
	}
	logger.Debug(fmt.Sprintf("gen jwt: cost %.3f s.", time.Since(genJwtTimeStart).Seconds()))
	logger.Debug(fmt.Sprintf("Login success for %s.", account))
	c.SetCookie(
		"_kapibara_access_token",
		jwtToken,
		int(config.GlobalConfig.JWTConf.Expired),
		"/",
		config.GlobalConfig.ServerConf.ServerDomain,
		true,
		true,
	)
	c.JSON(
		http.StatusOK,
		gin.H{
			"access_token": jwtToken,
			"user_name":    record.UserName,
			"uuid":         uuid.NewV4(),
		},
	)

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

	err = mysqlsdk.RegisterNewUser(&mysqlsdkMod.User{
		Account:  account,
		UserName: username,
		PwdHash:  pwdHash,
	})
	if err != nil {
		logger.Debug(fmt.Sprintf("Register new user error: [%v]", err))
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
