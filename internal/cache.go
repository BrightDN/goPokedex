package internal

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries 	map[string]CacheEntry
	mutex 			*sync.Mutex
}

type CacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntries:	make(map[string]CacheEntry),
		mutex:   		&sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.cacheEntries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cacheEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheEntries, k)
		}
	}
}