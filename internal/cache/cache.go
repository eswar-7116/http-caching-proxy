package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.RWMutex
	entries map[string]Entry
	TTL     time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		entries: make(map[string]Entry),
		TTL:     ttl,
	}
}

func (c *Cache) Get(url string) (Entry, bool) {
	c.mu.RLock()
	entry, ok := c.entries[url]
	c.mu.RUnlock()

	if !ok {
		return Entry{}, false
	}

	if time.Now().After(entry.ExpiresAt) {
		c.mu.Lock()
		delete(c.entries, url)
		c.mu.Unlock()

		return Entry{}, false
	}

	return entry, true
}

func (c *Cache) Set(url string, entry Entry) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry.ExpiresAt = time.Now().Add(c.TTL)
	c.entries[url] = entry
}
