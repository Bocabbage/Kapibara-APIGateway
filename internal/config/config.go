package config

import (
	"os"
	"strconv"

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
	ServerAddr string
	ServerPort string
	IsTLS      bool
	CertFile   string
	KeyFile    string
}

type JWTConfig struct {
	Expired   int64
	SecretKey string
}

type Config struct {
	LogConf    LogConfig
	MySQLConf  MySQLConfig
	ServerConf ServerConfig
	JWTConf    JWTConfig
}

// global config objects
var GlobalConfig = new(Config)

func loadGlobalConfig() {

	isTLS, _ := strconv.ParseBool(os.Getenv("IS_TLS"))
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
		ServerAddr: os.Getenv("SERVER_ADDR"),
		ServerPort: os.Getenv("SERVER_PORT"),
		IsTLS:      isTLS,
		CertFile:   os.Getenv("CERT_FILE"),
		KeyFile:    os.Getenv("KEY_FILE"),
	}
	GlobalConfig.JWTConf = JWTConfig{
		Expired:   jwtExpired,
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

func init() {
	godotenv.Load("../../config/.env")
	loadGlobalConfig()
}
