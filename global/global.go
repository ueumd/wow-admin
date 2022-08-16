package global

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"wow-admin/config"
)

const (
	SERVER_NAME string = "wow-admin"
)

var (
	CONFIG 			config.Config
	DB 				*sqlx.DB
	VIPER 			*viper.Viper
	RedisClient 	*redis.Client
)