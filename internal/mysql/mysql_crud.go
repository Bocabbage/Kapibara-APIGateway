// update date: 23/12/28
// author: bocabbage
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetPwdByUsername(username string) (string, error) {
	var pwdHash string
	var err error
	sqlStr := fmt.Sprintf("SELECT `pwdHash` FROM `users` WHERE `username` = \"%s\"", username)

	mysqlHandler := getMySQLHandler()
	rowResult := mysqlHandler.queryRow(sqlStr)
	err = rowResult.Scan(&pwdHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[GetPwdByUsername]NoRowsError at Query sql for: %s", sqlStr)
		}
	}

	return pwdHash, err
}
