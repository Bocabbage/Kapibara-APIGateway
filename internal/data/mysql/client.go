package mysql

import (
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	"sync"
	"time"

	mysqld "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var singletonDb *gorm.DB

func getMySQLClient() *gorm.DB {
	once.Do(func() {
		cfg := mysqld.Config{
			User:      config.GlobalConfig.MySQLConf.User,
			Passwd:    config.GlobalConfig.MySQLConf.Passwd,
			Net:       config.GlobalConfig.MySQLConf.Net,
			Addr:      config.GlobalConfig.MySQLConf.Addr,
			DBName:    config.GlobalConfig.MySQLConf.DBName,
			ParseTime: true,
			Loc:       time.Local,
		}
		dsn := cfg.FormatDSN()
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			logger.Fatal("Wrong happened in mysqlHandler init.")
		}
		singletonDb = db
	})

	return singletonDb
}
