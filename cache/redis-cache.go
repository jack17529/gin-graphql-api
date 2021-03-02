package cache

import (
	"context"
	"encoding/json"
	"graphql-srv/graph/model"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {

	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (rc *redisCache) getClient() *redis.Client {
	log.Println("trying to get a client.")
	return redis.NewClient(&redis.Options{
		Addr:     "backend-redis:6379",
		Password: "",
		DB:       rc.db,
	})
}

var ctx = context.Background()

func (rc *redisCache) Set(key string, value *model.Video) error {
	client := rc.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	log.Println("trying to set cache")
	err = client.Set(ctx, key, json, rc.expires*time.Second).Err()
	if err != nil {
		return err
	}

	log.Println("Done setting cache.")
	return nil
}

func (rc *redisCache) Get(key string) (*model.Video, error) {

	log.Println("Trying to get the redis cache video")

	client := rc.getClient()
	log.Println("Connected to client.")

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	v := model.Video{}
	err = json.Unmarshal([]byte(val), &v)
	if err != nil {
		return nil, err
	}

	log.Println("Got cache")
	return &v, nil
}
