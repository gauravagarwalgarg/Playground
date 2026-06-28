package main

import (
	"fmt"
	"sync"
	"time"
)

// Distributed Cache - LRU Cache with TTL, Eviction, and Sharding.
// Components: LRU eviction, TTL expiration, shard-based partitioning.
// Real system would use: consistent hashing, replication, gossip protocol.

// --- LRU Node (doubly-linked list) ---
type cacheEntry struct {
	key       string
	value     interface{}
	expiresAt time.Time
	prev      *cacheEntry
	next      *cacheEntry
}

// --- Cache Shard (thread-safe LRU with TTL) ---
type CacheShard struct {
	mu       sync.Mutex
	capacity int
	items    map[string]*cacheEntry
	head     *cacheEntry // most recently used
	tail     *cacheEntry // least recently used
	hits     int64
	misses   int64
}

func NewCacheShard(capacity int) *CacheShard {
	shard := &CacheShard{
		capacity: capacity,
		items:    make(map[string]*cacheEntry),
	}
	shard.head = &cacheEntry{}
	shard.tail = &cacheEntry{}
	shard.head.next = shard.tail
	shard.tail.prev = shard.head
	return shard
}

func (s *CacheShard) Get(key string) (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.items[key]
	if !exists {
		s.misses++
		return nil, false
	}

	// Check TTL
	if !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt) {
		s.removeEntry(entry)
		delete(s.items, key)
		s.misses++
		return nil, false
	}

	// Move to front (most recently used)
	s.moveToFront(entry)
	s.hits++
	return entry.value, true
}

func (s *CacheShard) Set(key string, value interface{}, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var expiresAt time.Time
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	// Update existing
	if entry, exists := s.items[key]; exists {
		entry.value = value
		entry.expiresAt = expiresAt
		s.moveToFront(entry)
		return
	}

	// Evict if at capacity
	if len(s.items) >= s.capacity {
		s.evictLRU()
	}

	// Insert new
	entry := &cacheEntry{key: key, value: value, expiresAt: expiresAt}
	s.items[key] = entry
	s.addToFront(entry)
}

func (s *CacheShard) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, exists := s.items[key]
	if !exists {
		return false
	}
	s.removeEntry(entry)
	delete(s.items, key)
	return true
}

func (s *CacheShard) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}

func (s *CacheShard) addToFront(entry *cacheEntry) {
	entry.prev = s.head
	entry.next = s.head.next
	s.head.next.prev = entry
	s.head.next = entry
}

func (s *CacheShard) removeEntry(entry *cacheEntry) {
	entry.prev.next = entry.next
	entry.next.prev = entry.prev
}

func (s *CacheShard) moveToFront(entry *cacheEntry) {
	s.removeEntry(entry)
	s.addToFront(entry)
}

func (s *CacheShard) evictLRU() {
	lru := s.tail.prev
	if lru == s.head {
		return
	}
	s.removeEntry(lru)
	delete(s.items, lru.key)
}

// --- Distributed Cache (sharded) ---
type DistributedCache struct {
	shards    []*CacheShard
	numShards int
}

func NewDistributedCache(numShards, capacityPerShard int) *DistributedCache {
	shards := make([]*CacheShard, numShards)
	for i := range shards {
		shards[i] = NewCacheShard(capacityPerShard)
	}
	return &DistributedCache{shards: shards, numShards: numShards}
}

func (dc *DistributedCache) getShard(key string) *CacheShard {
	hash := fnvHash(key)
	return dc.shards[hash%uint32(dc.numShards)]
}

func fnvHash(key string) uint32 {
	h := uint32(2166136261)
	for _, c := range key {
		h ^= uint32(c)
		h *= 16777619
	}
	return h
}

func (dc *DistributedCache) Get(key string) (interface{}, bool) {
	return dc.getShard(key).Get(key)
}

func (dc *DistributedCache) Set(key string, value interface{}, ttl time.Duration) {
	dc.getShard(key).Set(key, value, ttl)
}

func (dc *DistributedCache) Delete(key string) bool {
	return dc.getShard(key).Delete(key)
}

func (dc *DistributedCache) Stats() (totalItems int, totalHits, totalMisses int64) {
	for _, shard := range dc.shards {
		shard.mu.Lock()
		totalItems += len(shard.items)
		totalHits += shard.hits
		totalMisses += shard.misses
		shard.mu.Unlock()
	}
	return
}

func main() {
	cache := NewDistributedCache(8, 100) // 8 shards, 100 items each

	// Basic set/get
	fmt.Println("--- Basic Operations ---")
	cache.Set("user:1", map[string]string{"name": "Alice", "email": "alice@test.com"}, 5*time.Second)
	cache.Set("user:2", map[string]string{"name": "Bob"}, 5*time.Second)
	cache.Set("session:abc", "token-xyz", time.Minute)

	val, found := cache.Get("user:1")
	if !found {
		panic("FAIL: user:1 should be found")
	}
	fmt.Printf("  user:1 = %v\n", val)
	fmt.Println("PASS: basic set/get")

	// Cache miss
	_, found = cache.Get("nonexistent")
	if found {
		panic("FAIL: nonexistent should miss")
	}
	fmt.Println("PASS: cache miss")

	// TTL expiration
	fmt.Println("\n--- TTL Expiration ---")
	cache.Set("temp", "expires-fast", 10*time.Millisecond)
	time.Sleep(20 * time.Millisecond)
	_, found = cache.Get("temp")
	if found {
		panic("FAIL: expired key should not be found")
	}
	fmt.Println("PASS: TTL expiration works")

	// LRU eviction
	fmt.Println("\n--- LRU Eviction ---")
	smallCache := NewDistributedCache(1, 3) // 1 shard, capacity 3
	smallCache.Set("a", 1, 0)
	smallCache.Set("b", 2, 0)
	smallCache.Set("c", 3, 0)
	// Access "a" to make it recently used
	smallCache.Get("a")
	// Add "d" → should evict "b" (least recently used)
	smallCache.Set("d", 4, 0)

	_, foundA := smallCache.Get("a")
	_, foundB := smallCache.Get("b")
	_, foundD := smallCache.Get("d")

	if !foundA {
		panic("FAIL: 'a' should still be cached (recently used)")
	}
	if foundB {
		panic("FAIL: 'b' should be evicted (LRU)")
	}
	if !foundD {
		panic("FAIL: 'd' should be in cache")
	}
	fmt.Println("PASS: LRU eviction correct")

	// Update existing key
	fmt.Println("\n--- Update ---")
	cache.Set("user:1", map[string]string{"name": "Alice Updated"}, 5*time.Second)
	val, _ = cache.Get("user:1")
	m := val.(map[string]string)
	if m["name"] != "Alice Updated" {
		panic("FAIL: update not reflected")
	}
	fmt.Println("PASS: update works")

	// Delete
	fmt.Println("\n--- Delete ---")
	cache.Delete("user:2")
	_, found = cache.Get("user:2")
	if found {
		panic("FAIL: deleted key should not be found")
	}
	fmt.Println("PASS: delete works")

	// Concurrent access
	fmt.Println("\n--- Concurrent Access ---")
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key:%d", id%100)
			cache.Set(key, id, time.Second)
			cache.Get(key)
		}(i)
	}
	wg.Wait()

	items, hits, misses := cache.Stats()
	fmt.Printf("  Items: %d, Hits: %d, Misses: %d\n", items, hits, misses)
	hitRate := float64(hits) / float64(hits+misses) * 100
	fmt.Printf("  Hit rate: %.1f%%\n", hitRate)
	fmt.Println("PASS: concurrent access safe")

	fmt.Println("\nPASS: distributed cache complete")
}
