package impl

import (
	"be/models/domains"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisServiceImpl struct {
	rdb    *redis.Client
	config domains.RedisConfig
}

func NewRedisServiceImpl(rdb *redis.Client, config domains.RedisConfig) *RedisServiceImpl {
	return &RedisServiceImpl{rdb: rdb, config: config}
}

func (serv *RedisServiceImpl) SetData(key string, value interface{}, durationMinute int) error {
	ctx := context.Background()

	expiration := time.Duration(durationMinute) * time.Minute

	err := serv.rdb.Set(ctx, serv.config.Prefix+key, value, expiration).Err()
	return err
}

func (serv *RedisServiceImpl) GetData(key string) (interface{}, error) {
	ctx := context.Background()

	val, err := serv.rdb.Get(ctx, serv.config.Prefix+key).Result()
	if errors.Is(redis.Nil, err) {
		return "", fmt.Errorf("key '%s' does not exist", key)
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (serv *RedisServiceImpl) DeleteData(key string) error {
	ctx := context.Background()

	err := serv.rdb.Del(ctx, serv.config.Prefix+key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (serv *RedisServiceImpl) DeleteAllWithPrefix(prefix string) error {
	ctx := context.Background()

	var cursor uint64
	var keys []string
	var err error

	for {
		keys, cursor, err = serv.rdb.Scan(ctx, cursor, "*"+prefix+"*", 100).Result()
		if err != nil {
			return fmt.Errorf("failed to scan keys with prefix '%s': %w", prefix, err)
		}

		if len(keys) > 0 {
			err := serv.rdb.Del(ctx, keys...).Err()
			if err != nil {
				log.Printf("Warning: failed to delete some keys with prefix '%s': %v", prefix, err)
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (serv *RedisServiceImpl) ClearAllData() error {
	ctx := context.Background()
	return serv.rdb.FlushDB(ctx).Err()
}
