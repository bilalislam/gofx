package clients

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	Conn *redis.Client
}

func (rc *RedisClient) Get(model interface{}) error {
	err := rc.Conn.Get("test").Err()
	return err
}

func (rc *RedisClient) Add(model interface{}, duration time.Duration) error {
	err := rc.Conn.Set("test", model, duration).Err()
	return err
}
