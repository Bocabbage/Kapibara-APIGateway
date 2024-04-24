package mysql

import (
	"fmt"
	"kapibara-apigateway/internal/data/mysql/models"
	"kapibara-apigateway/internal/data/redis"
	"kapibara-apigateway/internal/logger"
)

func GetRecordByAccount(account string) (*models.User, error) {
	// Use redis cache
	userCache, _ := redis.GetRecordByAccountCache(account)
	if userCache == nil {
		return userCache, nil
	}

	db := getMySQLClient()
	var user models.User
	if err := db.Where(&models.User{Account: account}).First(&user).Error; err != nil {
		logger.Warn(fmt.Sprintf("[GetRecordByAccount]Error: %v", err))
		return nil, err
	}

	// Update redis cache
	redis.SetRecordByAccountCache(&user)

	return &user, nil
}

func RegisterNewUser(user *models.User) error {
	db := getMySQLClient()
	result := db.Create(user)
	if err := result.Error; err != nil {
		logger.Warn(fmt.Sprintf("[RegisterNewUser]Error: %v", err))
		return err
	}

	// Remove redis cache
	redis.RemoveRecordByAccountCache(user.Account)

	return nil
}
