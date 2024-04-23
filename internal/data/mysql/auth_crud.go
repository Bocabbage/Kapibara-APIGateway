package mysql

import (
	"fmt"
	"kapibara-apigateway/internal/logger"
)

func GetPwdByAccount(account string) (string, error) {
	db := getMySQLClient()
	var user Users
	if err := db.Where(&Users{Account: account}).First(&user).Error; err != nil {
		logger.Warn(fmt.Sprintf("[GetPwdByAccount]Error: %v", err))
		return "", err
	}
	return user.PwdHash, nil
}

func GetPwdStatusByAccount(account string) (string, int8, error) {
	db := getMySQLClient()
	var user Users
	if err := db.Where(&Users{Account: account}).First(&user).Error; err != nil {
		logger.Warn(fmt.Sprintf("[GetPwdStatusByAccount]Error: %v", err))
		return "", 0, err
	}

	return user.PwdHash, int8(user.Status), nil
}

func GetUsernameByAccount(account string) (string, error) {
	db := getMySQLClient()
	var user Users
	if err := db.Where(&Users{Account: account}).First(&user).Error; err != nil {
		logger.Warn(fmt.Sprintf("[GetPwdByAccount]Error: %v", err))
		return "", err
	}
	return user.UserName, nil
}

func GetRecordByAccount(account string) (*Users, error) {
	db := getMySQLClient()
	var user Users
	if err := db.Where(&Users{Account: account}).First(&user).Error; err != nil {
		logger.Warn(fmt.Sprintf("[GetRecordByAccount]Error: %v", err))
		return nil, err
	}
	return &user, nil
}

func RegisterNewUser(user *Users) error {
	db := getMySQLClient()
	result := db.Create(user)
	if err := result.Error; err != nil {
		logger.Warn(fmt.Sprintf("[RegisterNewUser]Error: %v", err))
		return err
	}
	return nil
}
