package dao

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"), // no password set
		DB:           viper.GetInt("redis.db"),          // use default DB
		PoolSize:     100,
		MinIdleConns: 5,
	})
}
