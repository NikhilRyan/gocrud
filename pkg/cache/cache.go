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

func GetMemCache() *cache.Cache {
	return memCache
}

// ReadFromCache handles read requests with Redis and in-memory caching
func ReadFromCache(keyPattern string, dataStruct interface{}, repoFunc func() (interface{}, error)) (interface{}, error) {
	key, err := generateKey(keyPattern, dataStruct)
	if err != nil {
		return nil, err
	}

	// Try to get from in-memory cache
	if cachedData, found := memCache.Get(key); found {
		return cachedData, nil
	}

	// Try to get from Redis
	result, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Cache miss, call the repository function
		data, err := repoFunc()
		if err != nil {
			return nil, err
		}

		// Convert data to string for storing in Redis
		dataStr, err := convertToString(data)
		if err != nil {
			return nil, err
		}

		// Store the result in Redis and in-memory cache
		if err := rdb.Set(ctx, key, dataStr, 0).Err(); err != nil {
			return nil, err
		}
		memCache.Set(key, data, cache.DefaultExpiration)
		return data, nil
	} else if err != nil {
		return nil, err
	}

	// Cache hit from Redis, convert string back to original type and store in in-memory cache
	data, err := convertFromString(result, dataStruct)
	if err != nil {
		return nil, err
	}
	memCache.Set(key, data, cache.DefaultExpiration)
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

func convertToString(data interface{}) (string, error) {
	// Convert data to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func convertFromString(dataStr string, dataStruct interface{}) (interface{}, error) {
	// Unmarshal the JSON string back to the original type
	result := reflect.New(reflect.TypeOf(dataStruct).Elem()).Interface()
	if err := json.Unmarshal([]byte(dataStr), result); err != nil {
		return nil, err
	}
	return result, nil
}
