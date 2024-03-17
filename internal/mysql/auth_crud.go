// update date: 23/12/28
// author: bocabbage
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	kerrors "kapibara-apigateway/internal/errors"
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
			logger.Warn(fmt.Sprintf("[GetPwdByAccount]NoRowsError at Query sql for: %s", sqlStr))
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
	var pwdHash string
	var err error

	result := make(map[string]string)

	sqlStr := fmt.Sprintf("SELECT `username`, `pwdHash`,`roleBitmap` FROM `users` WHERE `account` = \"%s\"", account)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&username, &pwdHash, &roleBitmap)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("[GetUsernameByAccount]NoRowsError at Query sql for: %s", sqlStr))
		}
	} else {
		result["username"] = username
		result["roleBitmap"] = roleBitmap
		result["pwdHash"] = pwdHash
	}

	return result, err
}

func RegisterNewUser(record map[string]string) (sql.Result, error) {
	var err error
	account, accountExist := record["account"]
	username, nameExist := record["username"]
	pwdHash, pwdHashExist := record["pwdHash"]

	if !accountExist || !nameExist || !pwdHashExist {
		return nil, &kerrors.KapibaraGeneralError{
			Code:    kerrors.KeyNotExistInMap,
			Message: "Key-miss in map",
		}
	}

	// [todo] Add check for getHandler
	mysqlHandler := getMySQLHandler()
	sqlStr := "INSERT INTO `users` (account, username, pwdHash) VALUES (?,?,?)"
	result, err := mysqlHandler.execute(sqlStr, account, username, pwdHash)

	return result, err
}
