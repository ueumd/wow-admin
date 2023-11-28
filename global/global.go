package global

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

const (
	SERVER_NAME string = "wow-admin"
)

var (
	DB          *sqlx.DB
	RedisClient *redis.Client
)
