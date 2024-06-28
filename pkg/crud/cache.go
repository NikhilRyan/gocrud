package crud

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

func InitCache(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}
