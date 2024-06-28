// redis/redis.go

package redis

import (
	"context"
	"data-manage/utils/configs"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisC struct {
	Client *redis.Client
}

var redisClusterClient *RedisC

func InitRedisClusterClient(ctx context.Context) error {
	// Load Redis configuration
	redisConfig, err := configs.LoadConfig("configs/redis.yml")
	if err != nil {
		return fmt.Errorf("failed to load Redis configuration: %v", err)
	}

	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisConfig.Redis.Host, redisConfig.Redis.Port),
	})
	// Ping the Redis server to check connectivity
	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Set the global Redis cluster client
	redisClusterClient = &RedisC{Client: client}
	return nil
}

func CloseRedis() {
	if redisClusterClient != nil && redisClusterClient.Client != nil {
		if err := redisClusterClient.Client.Close(); err != nil {
			log.Printf("Failed to close Redis client: %v", err)
		}
	}
}

func GetRedisClient() *RedisC {
	return redisClusterClient
}
