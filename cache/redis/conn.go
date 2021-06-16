package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	addr = "127.0.0.1:6379"
)

var (
	pool *redis.Pool
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			ticker := time.NewTicker(1 * time.Minute)
			for {
				select {
				case <-ticker.C:
					_, err := c.Do("PING")
					return err
				}
			}
		},
	}
}

func NewConn() redis.Conn {
	if pool == nil {
		pool = newPool()
	}
	return pool.Get()
}
