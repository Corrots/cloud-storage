package db

import (
	"context"
	"time"

	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/go-redis/redis/v8"
)

var logger = logging.MustGetLogger("cache")

var _client *redis.Client

func InitRedis(uri string, db int, password string) error {
	if uri == "" {
		logger.Fatal("redis uri is empty")
	}

	_client = redis.NewClient(&redis.Options{
		Network:      "",
		Addr:         uri,
		OnConnect:    nil,
		Username:     "",
		Password:     password,
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     16,
		MinIdleConns: 4,
	})

	_, err := _client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetClient() *redis.Client {
	return _client
}
