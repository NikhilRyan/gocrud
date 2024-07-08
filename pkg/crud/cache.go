package crud

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

// InitCache initializes the in-memory cache
func InitCache(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}
