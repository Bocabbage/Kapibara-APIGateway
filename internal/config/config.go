package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// type definition
type MySQLConfig struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	DBName string
}

type LogConfig struct {
	Level       string
	ServiceName string
	// Type      string
	// LocalPath string
	// LocalName string
}

type Config struct {
	LogConf   LogConfig
	MySQLConf MySQLConfig
}

// global config objects
var GlobalConfig = new(Config)

func loadGlobalConfig() {
	// [optim] use yaml file?
	GlobalConfig.MySQLConf = MySQLConfig{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_ADDR"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	GlobalConfig.LogConf = LogConfig{
		Level: os.Getenv("LOG_LEVEL"),
		// LocalPath: os.Getenv("LOG_LOCALPATH"),
		// LocalName: os.Getenv("LOG_LOCALNAME"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	}
}

func init() {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Fatal("Load dotenv file failed.")
	}
	loadGlobalConfig()
}
