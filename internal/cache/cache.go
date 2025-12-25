package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	entries  map[string]*list.Element
	TTL      time.Duration
	lru      *list.List
	capacity int
}

func New(capacity int, ttl time.Duration) (*Cache, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("Capacity must be a positive integer")
	}

	return &Cache{
		entries:  make(map[string]*list.Element),
		TTL:      ttl,
		lru:      list.New(),
		capacity: capacity,
	}, nil
}

func (c *Cache) Get(url string) (Entry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.entries[url]
	if !ok {
		return Entry{}, false
	}
	entry := node.Value.(Entry)

	if time.Now().After(entry.ExpiresAt) {
		delete(c.entries, url)
		c.lru.Remove(node)

		return Entry{}, false
	}

	c.lru.MoveToFront(node)

	return entry, true
}

func (c *Cache) Set(url string, entry Entry) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru.Len() >= c.capacity {
		c.evict()
	}

	entry.ExpiresAt = time.Now().Add(c.TTL)
	node := c.lru.PushFront(entry)
	c.entries[url] = node
}

func (c *Cache) evict() {
	// Remove expired entries
	prevLen := c.lru.Len()
	node := c.lru.Back()
	for node != nil {
		nextNode := node.Prev()
		if time.Now().After(node.Value.(Entry).ExpiresAt) {
			delete(c.entries, node.Value.(Entry).URL)
			c.lru.Remove(node)
		}
		node = nextNode
	}

	// Remove least recently used entry
	if c.lru.Len() == prevLen {
		node = c.lru.Back()
		delete(c.entries, node.Value.(Entry).URL)
		c.lru.Remove(node)
	}
}
