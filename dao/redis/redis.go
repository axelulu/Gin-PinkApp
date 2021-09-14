package redis

import (
	"fmt"
	"web_app/settings"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	return err
}

func Close() {
	_ = client.Close()
}
