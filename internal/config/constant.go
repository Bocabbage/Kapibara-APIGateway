package config

import "time"

const REDIS_CLI_TIMEOUT = 1 * time.Second
const GRPC_CLI_TIMEOUT = 50 * time.Second
const REDIS_AUTH_EXPIRE = 10 * time.Minute
const AUTHCACHE_REDISKEY = "account_"
const IS_VALID_RECENT_REDISKEY = "auth_recent"
