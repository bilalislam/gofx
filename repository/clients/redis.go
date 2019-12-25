package clients

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	Conn *redis.Client
}

func (rc *RedisClient) Get(id string) error {
	err := rc.Conn.Get(id).Err()
	return err
}

func (rc *RedisClient) Add(model interface{}, duration time.Duration) error {
	err := rc.Conn.Set("1", model, duration).Err()
	return err
}
