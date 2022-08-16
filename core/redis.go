package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"wow-admin/global"
)
const UserWxSessionKey  = "user:wx:sessions"

func RedisInit() error {
	_cfg := global.CONFIG.Redis
	var err error
	global.RedisClient, err = InitRedis(_cfg.Address, _cfg.Password, _cfg.Db, _cfg.PoolSize)
	return err
}

func InitRedis(addr string, password string, db , poolSize int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:               addr,
		OnConnect: func(conn *redis.Conn) error {
			return conn.Ping().Err()
		},
		Password:           password,
		DB:                 db,
		PoolSize:           poolSize,
		MaxRetries:         3,
	})

	err := client.Ping().Err()
	return client, err
}


func SetUserWXSessionKey(userId int, value string) error {
	return global.RedisClient.HSet(UserWxSessionKey, fmt.Sprintf("%d", userId), value).Err()
}

func GetUserWxSession(userId int) (string,error) {
	return global.RedisClient.HGet(UserWxSessionKey, fmt.Sprintf("%d", userId)).Result()
}
