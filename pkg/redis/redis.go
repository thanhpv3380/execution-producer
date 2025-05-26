package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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

func TTL(key string) (time.Duration, error) {
	return Client.TTL(Ctx, key).Result()
}

func PushToQueue(queueName string, data interface{}) error {
	return Client.LPush(Ctx, queueName, data).Err()
}

func Set(key string, data interface{}, expiration time.Duration) error {
	return Client.Set(Ctx, key, data, expiration).Err()
}

func Get(key string) (string, error) {
	data, err := Client.Get(Ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return "NOT_FOUND", err
		}

		return "", err
	}

	return data, nil
}

func HSet(hkey string, key string, data interface{}) error {
	return Client.HSet(Ctx, hkey, key, data).Err()
}

func HGet(hkey string, key string) (string, error) {
	data, err := Client.HGet(Ctx, hkey, key).Result()

	if err != nil {
		if err == redis.Nil {
			return "NOT_FOUND", err
		}

		return "", err
	}

	return data, nil
}

func BLPop(queue string) (string, error) {
	result, err := Client.BLPop(Ctx, 0*time.Second, queue).Result()

	if err != nil {
		return "", err
	}

	if len(result) < 2 {
		return "", fmt.Errorf("unexpected BLPOP result: %v", result)
	}

	return result[1], nil
}
