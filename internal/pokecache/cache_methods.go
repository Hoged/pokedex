package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
    defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
    defer c.mu.Unlock()
	val, exists := c.entry[key]
	return val.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entry {
		if time.Since(entry.createdAt) > interval {
			delete(c.entry, key)
			
			}
		}
		c.mu.Unlock()
	}
	
}