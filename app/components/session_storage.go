package components

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

type sessionStorage struct {
	redisClient *redis.Client
	ctx         context.Context
}

type SessionStorage interface {
	PutCache(cacheName string, key string, value any, duration int) error
	GetCache(cacheName string, key string, types any) error
	RemoveCache(cacheName string, key string) error
}

func SessionStorageInit(redisClient *redis.Client) SessionStorage {
	return &sessionStorage{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

// GetCache implements SessionStorage
func (s *sessionStorage) GetCache(cacheName string, key string, types any) error {
	cacheKey := cacheName + ":" + key
	value, err := s.redisClient.Get(s.ctx, cacheKey).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(value), types)
}

// PutCache implements SessionStorage
func (s *sessionStorage) PutCache(cacheName string, key string, value any, duration int) error {
	cacheKey := cacheName + ":" + key
	v, err := json.Marshal(value)

	if err != nil {
		return err
	}

	timeDuration := time.Duration(duration*60) * time.Second
	if duration == redis.KeepTTL {
		timeDuration = redis.KeepTTL
	}

	err = s.redisClient.Set(s.ctx, cacheKey, v, timeDuration).Err()

	if err != nil {
		println("something went wrong ", err.Error())
	}

	return err
}

// RemoveCache implements SessionStorage
func (s *sessionStorage) RemoveCache(cacheName string, key string) error {
	cacheKey := cacheName + ":" + key

	err := s.redisClient.Del(s.ctx, cacheKey).Err()

	if err != nil {
		println("something went wrong ", err.Error())
	}

	return err
}
