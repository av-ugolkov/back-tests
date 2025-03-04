package main

import (
	"fmt"
	"runtime"
	"sync"
	"weak"
)

type Cache struct {
	sync.Mutex
	items map[string]weak.Pointer[string] // Weak references to cached strings
}

func NewCache() *Cache {
	return &Cache{items: make(map[string]weak.Pointer[string])}
}

func (c *Cache) Get(key string) *string {
	c.Lock()
	defer c.Unlock()

	if weakPtr, ok := c.items[key]; ok {
		if val := weakPtr.Value(); val != nil {
			return val
		}
		delete(c.items, key)
	}
	return nil
}

func (c *Cache) Set(key string, value string) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = weak.Make(&value)
}

func main() {
	cache := NewCache()
	cache.Set("user:123", "John Doe")

	fmt.Println("Cached:", *cache.Get("user:123"))

	runtime.GC()

	if cache.Get("user:123") == nil {
		fmt.Println("Cache entry removed by GC")
	} else {
		fmt.Println("Still in cache:", *cache.Get("user:123"))
	}
}
