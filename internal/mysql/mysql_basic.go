// update date: 23/12/28
// author: bocabbage
// [todo]
// 1. connection pool? short connect?
// 2. retry & reconnect
package mysql

import (
	"database/sql"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var once sync.Once
var singletonHandler *mysqlHandler

type mysqlHandler struct {
	cfg mysql.Config
	db  *sql.DB
}

func (mh *mysqlHandler) queryRow(sqlStr string) *sql.Row {
	row := mh.db.QueryRow(sqlStr)
	return row
}

func (mh *mysqlHandler) query(sqlStr string) (*sql.Rows, error) {
	rows, err := mh.db.Query(sqlStr)
	return rows, err
}

func (mh *mysqlHandler) execute(sqlStr string) (int64, error) {
	// todo: implement
	var successInsert int64
	var err error

	return successInsert, err
}

func getMySQLHandler() *mysqlHandler {
	once.Do(func() {
		cfg := mysql.Config{
			User:   config.GlobalConfig.MySQLConf.User,   // os.Getenv("MYSQL_USER"),
			Passwd: config.GlobalConfig.MySQLConf.Passwd, // os.Getenv("MYSQL_PASSWORD"),
			Net:    config.GlobalConfig.MySQLConf.Net,    // "tcp",
			Addr:   config.GlobalConfig.MySQLConf.Addr,   // os.Getenv("MYSQL_ADDR"),
			DBName: config.GlobalConfig.MySQLConf.DBName, // os.Getenv("MYSQL_DATABASE"),
		}
		var err error
		var db *sql.DB
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			logger.Fatal("Wrong happened in mysqlHandler init.")
		}
		singletonHandler = &mysqlHandler{cfg: cfg, db: db}
	})
	return singletonHandler
}
