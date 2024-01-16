// update date: 23/12/28
// author: bocabbage
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"kapibara-apigateway/internal/logger"
	"strconv"
)

func GetPwdByAccount(account string) (string, error) {
	var pwdHash string
	var err error
	sqlStr := fmt.Sprintf("SELECT `pwdHash` FROM `users` WHERE `account` = \"%s\"", account)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&pwdHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("[GetPwdByAccount]NoRowsError at Query sql for: %s", sqlStr))
		}
	}

	return pwdHash, err
}

func GetPwdStatusByAccount(account string) (string, int8, error) {
	var pwdHash string
	var status string
	var err error
	sqlStr := fmt.Sprintf("SELECT `pwdHash`, `status` FROM `users` WHERE `account` = \"%s\"", account)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&pwdHash, &status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("[GetPwdByAccount]NoRowsError at Query sql for: %s", sqlStr))
		}
	}

	statusInt, _ := strconv.ParseInt(status, 10, 64)

	return pwdHash, int8(statusInt), err
}

func GetUsernameByAccount(account string) (string, error) {
	var username string
	var err error
	sqlStr := fmt.Sprintf("SELECT `username` FROM `users` WHERE `account` = \"%s\"", account)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("[GetUsernameByAccount]NoRowsError at Query sql for: %s", sqlStr))
		}
	}

	return username, err
}

func GetRecordByAccount(account string) (map[string]string, error) {
	var username string
	var roleBitmap string
	var err error

	result := make(map[string]string)

	sqlStr := fmt.Sprintf("SELECT `username`, `roleBitmap` FROM `users` WHERE `account` = \"%s\"", account)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&username, &roleBitmap)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("[GetUsernameByAccount]NoRowsError at Query sql for: %s", sqlStr))
		}
	} else {
		result["username"] = username
		result["roleBitmap"] = roleBitmap
	}

	return result, err
}
