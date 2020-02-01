package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

const (
	cachePrefix = "cache."
)

type Cache interface {
	Get(key string, data interface{}) (found bool, err error)
	Set(key string, data interface{}, expires time.Duration) (err error)
}

func New(client *redis.Client) Cache {
	return &cache{client: client}
}

type cache struct {
	client *redis.Client
}

func (c *cache) Get(key string, data interface{}) (found bool, err error) {
	var cached string
	if cached, err = c.client.Get(cachePrefix + key).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			return
		}
	}
	if len(cached) == 0 {
		return
	}
	if err = json.Unmarshal([]byte(cached), data); err != nil {
		return
	}
	found = true
	return
}

func (c *cache) Set(key string, data interface{}, expires time.Duration) (err error) {
	var buf []byte
	if buf, err = json.Marshal(data); err != nil {
		return
	}
	err = c.client.Set(cachePrefix+key, string(buf), expires).Err()
	return
}
