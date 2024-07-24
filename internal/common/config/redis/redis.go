package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type redisConn struct {
	c *redis.Client
}

func (r *redisConn) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.c.Get(ctx, key)
}

func (r *redisConn) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.c.Set(ctx, key, value, expiration)
}

func (r *redisConn) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.c.Del(ctx, keys...)
}

func NewRedisConn(maxIdle int64, maxActive, DB int64, host, port, password string) (Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       int(DB),
	})

	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("error connecting to redis")
		return client, err
	}

	log.Println("connected to Redis", res)

	return &redisConn{
		client,
	}, nil
}
