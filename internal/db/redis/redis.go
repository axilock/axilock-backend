package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisClient(address string, password string) RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
	return RedisStore{
		client: client,
	}
}

// SetToken stores a token in Redis
func (tc *RedisStore) SetToken(ctx context.Context, key string, token any, duration time.Duration) error {
	err := tc.client.Set(ctx, key, token, duration).Err()
	if err != nil {
		return fmt.Errorf("failed to store token in Redis: %w", err)
	}
	return nil
}

// GetToken retrieves a token from Redis
func (tc *RedisStore) GetToken(ctx context.Context, key string) (any, error) {
	result, err := tc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Key does not exist
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch token from Redis: %w", err)
	}
	return result, nil
}

// DeleteToken removes a token from Redis
func (tc *RedisStore) DeleteToken(ctx context.Context, key string) error {
	err := tc.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete token from Redis: %w", err)
	}
	return nil
}
