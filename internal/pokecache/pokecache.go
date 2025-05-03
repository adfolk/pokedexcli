package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	Contents map[string]cacheEntry
}

func NewCache(timeout time.Duration) *Cache {
	emptyCache := Cache{
		interval: timeout,
		Contents: map[string]cacheEntry{},
		mu:       sync.Mutex{},
	}
	go emptyCache.reapLoop()
	return &emptyCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	created := time.Now()
	newEntry := cacheEntry{createdAt: created, val: val}
	c.Contents[key] = newEntry
}

func (c *Cache) Get(key string) (res []byte, exists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Contents[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		currentTime, ok := <-ticker.C
		if !ok {
			fmt.Println("*************************")
			fmt.Printf("\n%v\n", currentTime)
			fmt.Println("*************************")
			break
		}
		c.mu.Lock()
		fmt.Println("*************************")
		fmt.Println("reapLoop is running")
		fmt.Println("*************************")
		for key, entry := range c.Contents {
			deadline := entry.createdAt.Add(c.interval)
			if currentTime.After(deadline) {
				fmt.Println("*************************")
				fmt.Printf("\ndeleting entry: %v\n", key)
				fmt.Println("*************************")
				delete(c.Contents, key)
			}
		}
		c.mu.Unlock()
	}
}
