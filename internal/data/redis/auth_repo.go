package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"kapibara-apigateway/internal/config"
	cryptoutils "kapibara-apigateway/internal/crypto"
	"kapibara-apigateway/internal/data/mysql/models"
	"kapibara-apigateway/internal/logger"
)

func GetRecordByAccountCache(account string) (*models.User, error) {
	rkey := userKeyGenFromAccount(account)
	redisCli := getRedisClient()
	ctx, recordCacheCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
	defer recordCacheCancel()

	cacheRecord := redisCli.Get(ctx, rkey)
	if err := cacheRecord.Err(); err != nil {
		logger.Error(fmt.Sprintf("[GetRecordByAccountCache][account:%s] Error: %v", account, err))
		return nil, err
	} else if len(cacheRecord.Val()) < 1 {
		logger.Info(fmt.Sprintf("[GetRecordByAccountCache][account:%s] cache missed", account))
		return nil, nil
	}

	var user models.User
	if err := json.Unmarshal([]byte(cacheRecord.Val()), &user); err != nil {
		logger.Error(fmt.Sprintf("[GetRecordByAccountCache][account:%s] Json Unmashal error: %v", account, err))
		return nil, err
	}

	return &user, nil
}

func SetRecordByAccountCache(user *models.User) error {
	rkey := userKeyGenFromAccount(user.Account)
	redisCli := getRedisClient()
	ctx, recordCacheCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
	defer recordCacheCancel()

	record, err := json.Marshal(user)
	if err != nil {
		logger.Error(fmt.Sprintf("[SetRecordByAccountCache][account:%s]Json Marshal failed: %v", user.Account, err))
	}

	res := redisCli.Set(ctx, rkey, record, config.REDIS_AUTH_EXPIRE)
	if err := res.Err(); err != nil {
		logger.Error(fmt.Sprintf("[SetRecordByAccountCache][account:%s]Error: %v", user.Account, err))
	}
	return nil
}

func RemoveRecordByAccountCache(account string) error {
	rkey := userKeyGenFromAccount(account)
	redisCli := getRedisClient()
	ctx, recordCacheCancel := context.WithTimeout(context.Background(), config.REDIS_CLI_TIMEOUT)
	defer recordCacheCancel()

	res := redisCli.Del(ctx, rkey)

	if err := res.Err(); err != nil {
		// TODO: add retry?
		logger.Error(fmt.Sprintf("[RemoveRecordByAccountCache][account:%s]Error: %v", account, err))
	}

	return nil
}

func userKeyGenFromAccount(account string) string {
	return cryptoutils.Sha256Hash(fmt.Sprintf("account[%s]:user", account))
}
