package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/corrots/cloud-storage/db"
	"github.com/go-redis/redis/v8"
)

var keyIsNullError = errors.New("key is null")

type CacheService struct {
	cli *redis.Client
}

func NewCacheService() *CacheService {
	return &CacheService{cli: db.GetClient()}
}

func (svc *CacheService) GetI(key string, i interface{}) error {
	if key == "" {
		return keyIsNullError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stringCmd := svc.cli.Get(ctx, key)
	if stringCmd.Err() != nil {
		if stringCmd.Val() == "" {
			return nil
		}
		return stringCmd.Err()
	}

	bytes, _ := stringCmd.Bytes()
	if err := json.Unmarshal(bytes, i); err != nil {
		return err
	}

	return nil
}

func (svc *CacheService) Set(key string, val interface{}) error {
	return svc.SetWithTTL(key, val, 10*time.Minute)
}

func (svc *CacheService) SetWithTTL(key string, val interface{}, ttl time.Duration) error {
	if key == "" {
		return keyIsNullError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	bytes, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("json marshal err: %v\n", err)
	}

	statusCmd := svc.cli.Set(ctx, key, bytes, ttl)
	return statusCmd.Err()
}
