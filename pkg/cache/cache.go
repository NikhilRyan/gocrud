package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

var (
	rdb      *redis.Client
	memCache *cache.Cache
	ctx      = context.Background()
)

// InitRedis initializes the Redis client
func InitRedis(addr, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// InitMemCache initializes the in-memory cache
func InitMemCache(defaultExpiration, cleanupInterval time.Duration) {
	memCache = cache.New(defaultExpiration, cleanupInterval)
}

// Get retrieves an item from the cache
func Get(key string) (interface{}, bool) {
	// Try to get from in-memory cache
	if cachedData, found := memCache.Get(key); found {
		return cachedData, true
	}

	// Try to get from Redis
	result, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, false
	} else if err != nil {
		return nil, false
	}

	// Convert string back to interface
	var data interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return nil, false
	}

	// Store in in-memory cache
	memCache.Set(key, data, cache.DefaultExpiration)
	return data, true
}

// Set stores an item in both the in-memory cache and Redis
func Set(key string, data interface{}) error {
	// Store in in-memory cache
	memCache.Set(key, data, cache.DefaultExpiration)

	// Convert interface to string for storing in Redis
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Store in Redis
	return rdb.Set(ctx, key, jsonData, 0).Err()
}

// ReadFromCache handles read requests with Redis and in-memory caching
func ReadFromCache(keyPattern string, dataStruct interface{}, repoFunc func() (interface{}, error)) (interface{}, error) {
	key, err := generateKey(keyPattern, dataStruct)
	if err != nil {
		return nil, err
	}

	// Try to get from cache
	cachedData, found := Get(key)
	if found {
		return cachedData, nil
	}

	// Cache miss, call the repository function
	data, repoErr := repoFunc()
	if repoErr != nil {
		return nil, repoErr
	}

	// Store the result in cache
	if err := Set(key, data); err != nil {
		return nil, err
	}

	return data, nil
}

func generateKey(keyPattern string, dataStruct interface{}) (string, error) {
	v := reflect.ValueOf(dataStruct)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return "", fmt.Errorf("dataStruct must be a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()

	key := keyPattern
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("redis")
		if tag == "" {
			continue
		}

		fieldValue := v.Field(i).Interface()
		key = strings.Replace(key, "%"+tag+"%", fmt.Sprintf("%v", fieldValue), -1)
	}

	if strings.Contains(key, "%") {
		return "", fmt.Errorf("not all placeholders in key pattern were filled")
	}

	return key, nil
}
