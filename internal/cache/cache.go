package cache

import (
	"sync"
)

type Cache struct {
	mu      sync.RWMutex
	entries map[string]Entry
}

func New() *Cache {
	return &Cache{
		entries: make(map[string]Entry),
	}
}

func (c *Cache) Get(url string) (Entry, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[url]
	return entry, ok
}

func (c *Cache) Set(url string, entry Entry) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[url] = entry
}
