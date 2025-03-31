package cache

import (
	"sync"
)

type Cache struct {
	mutex    sync.RWMutex
	items    map[string]string
}

func NewCache() *Cache {
	return &Cache{
		items:    make(map[string]string),
	}
}

// Get retrieves a value for a given key
func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the item exists
	item, exists := c.items[key]
	if !exists {
		return "", false
	}

	return item, true
}

// Put inserts a new key-value pair
func (c *Cache) Put(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Insert new item
	c.items[key] = value
}
