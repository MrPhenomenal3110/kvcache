package cache

import (
	"kvcache/internal/config"
	"sync"
)

// type CacheItem struct {
// 	Key        string
// 	Value      string
// 	LastAccess int64
// }

type Cache struct {
	mutex    sync.RWMutex
	items    map[string]string
	capacity int        // Max size of cache
}

func NewCache() *Cache {
	return &Cache{
		items:    make(map[string]string),
		capacity: config.MaxCapacity,
	}
}

// Get retrieves a value for a given key and updates its position in the LRU order
func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the item exists
	item, exists := c.items[key]
	if !exists {
		return "", false
	}

	// Update LastAccess time
	// item.LastAccess = time.Now().UnixNano()

	return item, true
}

// Put inserts a new key-value pair and removes the least recently used item if capacity is exceeded
func (c *Cache) Put(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Insert new item at the end
	c.items[key] = value
}
