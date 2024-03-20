package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kapibara-apigateway/internal/config"
	cryptoUtils "kapibara-apigateway/internal/crypto"
	"kapibara-apigateway/internal/logger"
	mysqlsdk "kapibara-apigateway/internal/mysql"
	"kapibara-apigateway/internal/redis_utils"
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

	// use redis as cache
	redisClient := redis_utils.GetRedisClient()

	// get related-account
	var record map[string]string

	getCacheTimeStart := time.Now()
	recordCacheCtx, recordCacheCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
	defer recordCacheCancel()
	cacheRecord := redisClient.HGetAll(recordCacheCtx, config.AUTHCACHE_REDISKEY+":"+account)
	if cacheRecord.Err() == nil && len(cacheRecord.Val()) > 0 {
		logger.Debug(fmt.Sprintf("getrecord from redis: cost %.3f s.", time.Since(getCacheTimeStart).Seconds()))
		record = cacheRecord.Val()
	} else {
		getRecordTimeStart := time.Now()

		record, err = mysqlsdk.GetRecordByAccount(account)
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
		logger.Debug(fmt.Sprintf("getrecord from mysql: cost %.3f s.", time.Since(getRecordTimeStart).Seconds()))

		setRecordCacheCtx, setRecordCacheCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
		defer setRecordCacheCancel()
		hsetResult := redisClient.HSet(setRecordCacheCtx, config.AUTHCACHE_REDISKEY+":"+account, record)
		if err = hsetResult.Err(); err != nil {
			logger.Warn(fmt.Sprintf("[failed]set record-cache for account: [%s]: %v", account, err))
		} else {
			logger.Debug(fmt.Sprintf("[success]set record-cache for account: [%s]", account))
		}
	}

	// check the password
	// first check redis cache
	ctx, cancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
	defer cancel()
	recentCheck := redisClient.Exists(ctx, config.IS_VALID_RECENT_REDISKEY+":"+account)
	if recentCheck.Val() < 1 {
		// if not, do slow bcryptcmp
		bCryptTimeStart := time.Now()
		storedPwdHash := record["pwdHash"]
		compResult := cryptoUtils.BCryptHashCompare(storedPwdHash, pwd)
		logger.Debug(fmt.Sprintf("bcrypt: cost %.3f s.", time.Since(bCryptTimeStart).Seconds()))
		if !compResult {
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": "LOGIN ERROR: incorrect pwd."},
			)
			return
		}
		storeCtx, storeCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
		defer storeCancel()
		setResult := redisClient.Set(storeCtx, config.IS_VALID_RECENT_REDISKEY+":"+account, 1, config.REDIS_AUTH_EXPIRE)
		if e := setResult.Err(); e != nil {
			logger.Warn(fmt.Sprintf("[failed]set auth-cache for account: [%s]: %v", account, e))
		} else {
			logger.Debug(fmt.Sprintf("[success]set auth-cache for account: [%s]", account))
		}
	}

	// generate jwt
	// [todo] enhancement: decoupling jwt-format
	genJwtTimeStart := time.Now()
	record["account"] = account
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
			"user_name":    record["username"],
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
