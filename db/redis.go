package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

//var logger = logging.MustGetLogger("cache")

var _client *redis.Client

func InitRedis(uri string, db int, password string) error {
	if uri == "" {
		return fmt.Errorf("redis uri is empty")
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
	_client.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetClient() *redis.Client {
	return _client
}
