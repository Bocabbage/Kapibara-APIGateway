package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// type definition
type MySQLConfig struct {
	User    string
	Passwd  string
	Net     string
	Addr    string
	DBName  string
	StrSalt string
	Salt    int64
}

type LogConfig struct {
	Level       string
	ServiceName string
	// Type      string
	// LocalPath string
	// LocalName string
}

type ServerConfig struct {
	ServerAddr   string
	ServerPort   string
	ServerDomain string
	CertFile     string
	KeyFile      string
}

type JWTConfig struct {
	Expired   int64
	SecretKey string
}

type CORSConfig struct {
	AllowOrigins []string
}

type MikananiConfig struct {
	GRpcServerAddr string
}

type HealthTestConfig struct {
	TestConfig string
}

type Config struct {
	LogConf      LogConfig
	MySQLConf    MySQLConfig
	ServerConf   ServerConfig
	JWTConf      JWTConfig
	CORSConf     CORSConfig
	MikananiConf MikananiConfig
	HealthConf   HealthTestConfig
}

// global config objects
var GlobalConfig = new(Config)

func loadGlobalConfig() {

	mysqlSalt, _ := strconv.ParseInt(os.Getenv("MYSQL_SALT"), 16, 64)
	jwtExpired, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRED"), 10, 64)

	GlobalConfig.MySQLConf = MySQLConfig{
		User:    os.Getenv("MYSQL_USER"),
		Passwd:  os.Getenv("MYSQL_PASSWORD"),
		Net:     "tcp",
		Addr:    os.Getenv("MYSQL_ADDR"),
		DBName:  os.Getenv("MYSQL_DATABASE"),
		StrSalt: os.Getenv("MYSQL_STRSALT"),
		Salt:    mysqlSalt,
	}
	GlobalConfig.LogConf = LogConfig{
		Level: os.Getenv("LOG_LEVEL"),
		// LocalPath: os.Getenv("LOG_LOCALPATH"),
		// LocalName: os.Getenv("LOG_LOCALNAME"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	}
	GlobalConfig.ServerConf = ServerConfig{
		ServerAddr:   os.Getenv("SERVER_ADDR"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		ServerDomain: os.Getenv("SERVER_DOMAIN"),
		CertFile:     os.Getenv("CERT_FILE"),
		KeyFile:      os.Getenv("KEY_FILE"),
	}
	GlobalConfig.JWTConf = JWTConfig{
		Expired:   jwtExpired,
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
	GlobalConfig.HealthConf = HealthTestConfig{
		TestConfig: os.Getenv("TEST_CONFIG_FLAG"),
	}
	GlobalConfig.CORSConf = CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	}
	GlobalConfig.MikananiConf = MikananiConfig{
		GRpcServerAddr: os.Getenv("MIKANANI_GRPC_ADDR"),
	}
}

func loadDotEnvOnce() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Not able to get the code-dir.")
		return
	}
	dir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(dir, "../../config/.env"))
	if err != nil {
		fmt.Println("Load dotenv file failed.")
	}
}

func init() {
	loadDotEnvOnce()
	loadGlobalConfig()
}
