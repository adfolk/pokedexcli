package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	val, ok := cache.entries[key]
	return val.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	dethKlok := time.NewTicker(interval)
	for range dethKlok.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
