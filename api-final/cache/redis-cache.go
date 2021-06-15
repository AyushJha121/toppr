package cache

import (
	"api-final/entity"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) TopicCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
func (cache *redisCache) Set(key string, value entity.Topic) {
	ctx := context.TODO()
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err.Error())
	}
	client.Set(ctx, key, json, cache.expires*time.Second)
}
func (cache *redisCache) Get(key string) entity.Topic {

	ctx := context.TODO()
	client := cache.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	topic := entity.Topic{}
	err = json.Unmarshal([]byte(val), &topic)
	if err != nil {
		fmt.Println(err.Error())
	}
	return topic
}
