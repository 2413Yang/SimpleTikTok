package cache

import (
	"context"
	"fmt"
	"sync"
	"time"
	"usersvr/config"
	"usersvr/log"

	redis "github.com/redis/go-redis/v9"
)

var (
	redisConn   *redis.Client
	redisOnce   sync.Once
	ValueExpire = time.Hour * 24 * 7
)

// initRedis 连接Redis
func initRedis() {
	redisConfig := config.GetGlobalConfig().RedisConfig
	log.Infof("redisConfig ======= %+v", redisConfig)
	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}
	res, err := redisConn.Set(context.Background(), "abc", 100, 1*time.Second).Result()
	log.Infof("res=======%v,err======%v", res, err)
	_, err = redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to ping redis,err:" + err.Error())
	}
}

func CloseRedis() {
	if redisConn != nil {
		redisConn.Close()
	}
}

// GetRedisCli 获取redis连接
func GetRedisCli() *redis.Client {
	redisOnce.Do(initRedis)

	return redisConn
}
