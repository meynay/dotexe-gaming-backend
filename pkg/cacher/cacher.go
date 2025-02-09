package cacher

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cacher struct {
	redisClient *redis.Client
}

func NewCacher(addr, pass string) *Cacher {
	return &Cacher{redisClient: redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})}
}

func (c *Cacher) CacheSignInCode(number, code string) error {
	ctx := context.Background()
	err := c.redisClient.Set(ctx, number, code, time.Duration(600*time.Second)).Err()
	if err != nil {
		return fmt.Errorf("failed to cache code in Redis: %v", err)
	}

	return nil
}

func (c *Cacher) CheckCode(number, code string) (bool, error) {
	ctx := context.Background()
	cachedCode, err := c.redisClient.Get(ctx, number).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("no code found for user %s", number)
	} else if err != nil {
		return false, fmt.Errorf("failed to retrieve code from Redis: %v", err)
	}
	if cachedCode == code {
		c.redisClient.Del(ctx, number)
		return true, nil
	}
	return false, nil
}
