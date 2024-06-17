package cache

// Modified by https://github.com/die-net/lrucache

import (
	"container/list"
	"sync"
	"time"
)

// Option is part of Functional Options Pattern
type Option func(*LruCache)

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback = func(key any, value any)

// WithUpdateAgeOnGet update expires when Get element
func WithUpdateAgeOnGet() Option {
	return func(l *LruCache) {
		l.updateAgeOnGet = true
	}
}

// WithAge defined element max age (second)
func WithAge(maxAge int64) Option {
	return func(l *LruCache) {
		l.maxAge = maxAge
	}
}

// LruCache is a thread-safe, in-memory lru-cache that evicts the
// least recently used entries from memory when (if set) the entries are
// older than maxAge (in seconds).  Use the New constructor to create one.
type LruCache struct {
	maxAge         int64
	maxSize        int
	mu             sync.Mutex
	cache          map[any]*list.Element
	lru            *list.List // Front is least-recent
	updateAgeOnGet bool
	staleReturn    bool
	onEvict        EvictCallback
}

// New creates an LruCache
func New(options ...Option) *LruCache {
	lc := &LruCache{
		lru:   list.New(),
		cache: make(map[any]*list.Element),
	}

	for _, option := range options {
		option(lc)
	}

	return lc
}

// Get returns the any representation of a cached response and a bool
// set to true if the key was found.
func (c *LruCache) Get(key any) (any, bool) {
	entry := c.get(key)
	if entry == nil {
		return nil, false
	}
	value := entry.value

	return value, true
}

// Set stores the any representation of a response for a given key.
func (c *LruCache) Set(key any, value any) {
	expires := int64(0)
	if c.maxAge > 0 {
		expires = time.Now().Unix() + c.maxAge
	}
	c.SetWithExpire(key, value, time.Unix(expires, 0))
}

// SetWithExpire stores the any representation of a response for a given key and given expires.
// The expires time will round to second.
func (c *LruCache) SetWithExpire(key any, value any, expires time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if le, ok := c.cache[key]; ok {
		c.lru.MoveToBack(le)
		e := le.Value.(*entry)
		e.value = value
		e.expires = expires.Unix()
	} else {
		e := &entry{key: key, value: value, expires: expires.Unix()}
		c.cache[key] = c.lru.PushBack(e)

		if c.maxSize > 0 {
			if len := c.lru.Len(); len > c.maxSize {
				c.deleteElement(c.lru.Front())
			}
		}
	}

	c.maybeDeleteOldest()
}

func (c *LruCache) get(key any) *entry {
	c.mu.Lock()
	defer c.mu.Unlock()

	le, ok := c.cache[key]
	if !ok {
		return nil
	}

	if !c.staleReturn && c.maxAge > 0 && le.Value.(*entry).expires <= time.Now().Unix() {
		c.deleteElement(le)
		c.maybeDeleteOldest()

		return nil
	}

	c.lru.MoveToBack(le)
	entry := le.Value.(*entry)
	if c.maxAge > 0 && c.updateAgeOnGet {
		entry.expires = time.Now().Unix() + c.maxAge
	}
	return entry
}

// Delete removes the value associated with a key.
func (c *LruCache) Delete(key any) {
	c.mu.Lock()

	if le, ok := c.cache[key]; ok {
		c.deleteElement(le)
	}

	c.mu.Unlock()
}

func (c *LruCache) maybeDeleteOldest() {
	if !c.staleReturn && c.maxAge > 0 {
		now := time.Now().Unix()
		for le := c.lru.Front(); le != nil && le.Value.(*entry).expires <= now; le = c.lru.Front() {
			c.deleteElement(le)
		}
	}
}

func (c *LruCache) deleteElement(le *list.Element) {
	c.lru.Remove(le)
	e := le.Value.(*entry)
	delete(c.cache, e.key)
	if c.onEvict != nil {
		c.onEvict(e.key, e.value)
	}
}

type entry struct {
	key     any
	value   any
	expires int64
}
