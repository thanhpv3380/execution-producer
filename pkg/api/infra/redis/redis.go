package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func NewClient(addr, password string) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	_, err := Client.Ping(Ctx).Result()
	return err
}

func PushToQueue(queueName string, data interface{}) error {
	return Client.LPush(Ctx, queueName, data).Err()
}

func HSet(hkey string, key string, data interface{}) error {
	return Client.HSet(Ctx, hkey, key, data).Err()
}
