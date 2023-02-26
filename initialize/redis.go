package initialize

import (
	"blog/server/initialize/global"
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.GLOBAL_CONFIG.Redis.Host + ":" + strconv.Itoa(global.GLOBAL_CONFIG.Redis.Port),
        Password: global.GLOBAL_CONFIG.Redis.Password, // no password set
        DB:       global.GLOBAL_CONFIG.Redis.Db,  // use default DB
    })

	_, err := client.Ping(context.Background()).Result()
    if err != nil {
		logger := global.GLOBAL_LOG
		logger.Debug("Redis connect ping failed, err: ", zap.Any("err:", err))
		return nil
	}

	return client
}