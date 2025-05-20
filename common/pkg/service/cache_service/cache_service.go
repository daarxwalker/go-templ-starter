package cache_service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"time"
	
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	
	"common/pkg/config/cache_config"
	"common/pkg/env"
)

type CacheService struct {
	client          *redis.Client
	deleteBatchSize int64
}

const (
	Token = "cache_service"
)

func New(cfg *viper.Viper) *CacheService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.GetString(cache_config.Uri),
			Password: cfg.GetString(cache_config.Password),
			DB:       cfg.GetInt(cache_config.DB),
			TLSConfig: map[bool]*tls.Config{
				true: {
					MinVersion: tls.VersionTLS12,
				},
				false: nil,
			}[env.Production()],
		},
	)
	if pingErr := client.Ping(ctx).Err(); pingErr != nil {
		log.Fatalf("ping to cache failed: %v", pingErr)
	} else {
		log.Println("connected and pinged cache!")
	}
	return &CacheService{
		client:          client,
		deleteBatchSize: cfg.GetInt64(cache_config.DeleteBatchSize),
	}
}

func (s *CacheService) Client() *redis.Client {
	return s.client
}

func (s *CacheService) Exists(c context.Context, key ...string) bool {
	return s.client.Exists(c, key...).Val() == 1
}

func (s *CacheService) Get(c context.Context, key string, target any) error {
	value := s.client.Get(c, key).Val()
	if len(value) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(value), target)
}

func (s *CacheService) MustGet(c context.Context, key string, target any) {
	if err := s.Get(c, key, target); err != nil {
		panic(err)
	}
}

func (s *CacheService) Set(c context.Context, key string, value any, expiration time.Duration) error {
	valueBytes, marshalErr := json.Marshal(value)
	if marshalErr != nil {
		return marshalErr
	}
	return s.client.Set(c, key, string(valueBytes), expiration).Err()
}

func (s *CacheService) MustSet(c context.Context, key string, value any, expiration time.Duration) {
	if err := s.Set(c, key, value, expiration); err != nil {
		panic(err)
	}
}

func (s *CacheService) Destroy(c context.Context, key string) error {
	return s.Set(c, key, "", time.Millisecond)
}

func (s *CacheService) MustDestroy(c context.Context, key string) {
	if err := s.Destroy(c, key); err != nil {
		panic(err)
	}
}

func (s *CacheService) DestroyWithPattern(c context.Context, pattern string) error {
	var cursor uint64
	for {
		keys, newCursor, scanErr := s.client.Scan(c, cursor, pattern, s.deleteBatchSize).Result()
		if scanErr != nil {
			return scanErr
		}
		if len(keys) > 0 {
			_, deleteErr := s.client.Del(c, keys...).Result()
			if deleteErr != nil {
				return deleteErr
			}
		}
		cursor = newCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}

func (s *CacheService) MustDestroyWithPattern(c context.Context, pattern string) {
	if err := s.DestroyWithPattern(c, pattern); err != nil {
		panic(err)
	}
}

func (s *CacheService) Lock(c context.Context, name, ip string) error {
	key := fmt.Sprintf("mutex:%s:%s", name, ip)
	for {
		locked, setLockErr := s.client.SetNX(c, key, 1, 0).Result()
		if setLockErr != nil {
			return setLockErr
		}
		if locked {
			return nil
		}
		select {
		case <-c.Done():
			return c.Err()
		case <-time.After(100 * time.Millisecond):
			continue
		}
	}
}

func (s *CacheService) MustLock(c context.Context, name, ip string) {
	if err := s.Lock(c, name, ip); err != nil {
		panic(err)
	}
}

func (s *CacheService) Unlock(c context.Context, name, ip string) error {
	return s.client.Del(c, fmt.Sprintf("mutex:%s:%s", name, ip)).Err()
}

func (s *CacheService) MustUnlock(c context.Context, name, ip string) {
	if err := s.Unlock(c, name, ip); err != nil {
		panic(err)
	}
}
