package redis_utils

import (
	"kapibara-apigateway/internal/config"
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once
var singletonHandler *redis.Client

func GetRedisClient() *redis.Client {
	once.Do(func() {
		// read config
		cfg := redis.Options{
			Addr:     config.GlobalConfig.RedisConf.Addr,
			Password: config.GlobalConfig.RedisConf.Password,
			DB:       config.GlobalConfig.RedisConf.DB,
		}
		// create client
		singletonHandler = redis.NewClient(&cfg)
	})
	return singletonHandler
}
