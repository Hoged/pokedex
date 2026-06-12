package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entry		map[string]cacheEntry
	mu     		sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache {
		entry: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)

	return cache
}
