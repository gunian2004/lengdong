package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port),
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})

	// 测试 Redis 连接
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}

	log.Println("Redis 连接已建立")
}
