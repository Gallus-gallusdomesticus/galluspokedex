package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interv time.Duration) *Cache {
	var cac Cache

	m := make(map[string]cacheEntry)
	cac.cache = m
	cac.interval = interv

	go cac.reapLoop()

	return &cac
}

func (c *Cache) Add(key string, val []byte) { //Add an entry on c.cache
	c.mu.Lock()
	var ce cacheEntry
	ce.val = val
	ce.createdAt = time.Now()
	c.cache[key] = ce
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) { //get the value in c.cache
	c.mu.RLock()
	ce, ok := c.cache[key]
	c.mu.RUnlock()
	if ok {
		return ce.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop() { //delete entry that is bigger than the interval to avoid clogging
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for i, ca := range c.cache {
			if time.Since(ca.createdAt) > c.interval {
				delete(c.cache, i)
			}
		}
		c.mu.Unlock()
	}

}
