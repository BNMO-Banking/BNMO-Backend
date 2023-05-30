package redis

import (
	"BNMO/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) RedisCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisCache) getRatesClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) SetCache(key string, value interface{}) {
	client := cache.getRatesClient()
	setter := client.Set(ctx, key, value, cache.expires*time.Second)
	if setter.Err() != nil {
		fmt.Println("Failed to set value in redis cache", setter.Err().Error())
		return
	}
	fmt.Println("Successfully set rates in redis cache")
}

// Known bugs: connecting to redis cache may fail and display an error
// i/o timeout.
// Possible fix: Restart device
func (cache *redisCache) GetCache(key string, requested string) interface{} {
	client := cache.getRatesClient()
	output, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return -1
	}

	if (key == "symbols") {
		var symbols models.SymbolsCache
		
		json.Unmarshal([]byte(output), &symbols)

		return symbols.Symbols
	} else if (key == "rates") {
		var rates models.RatesCache

		json.Unmarshal([]byte(output), &rates)

		return rates.Rates[requested]
	}

	return nil
}