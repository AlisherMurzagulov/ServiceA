package redis

import (
	"encoding/json"
	"gopkg.in/redis.v5"
	"serviceB/cfg"
	"time"
)

type manager struct {
	cache *redis.Client
}

type Manager interface {
	Exits(key string) bool
	Get(key string) (string, time.Duration)
	Set(key string, data interface{}, expiration time.Duration)
	Delete(key string)
}

func NewManager() Manager {

	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.AppConfigs.Cache.GetHostPort(),
		DB:       cfg.AppConfigs.Cache.DatabaseNumber,
		Password: cfg.AppConfigs.Cache.Password})

	if err := cache.Ping().Err(); err != nil {
		panic(err)
	}

	return &manager{
		cache: cache,
	}
}

func (r *manager) Exits(key string) bool {
	return r.cache.Exists(key).Val()
}

func (r *manager) Get(key string) (string, time.Duration) {
	return r.cache.Get(key).Val(), r.cache.TTL(key).Val() / time.Second
}

func (r *manager) Set(key string, data interface{}, expiration time.Duration) {
	dataBytes, _ := json.Marshal(&data)
	r.cache.Set(key, string(dataBytes), expiration*time.Second)
}

func (r *manager) Delete(key string) {
	r.cache.Del(key)
}
