package main

import (
	"fmt"
	"time"
)

// Proxy Pattern: Provides a surrogate or placeholder for another object
// to control access, add caching, logging, or lazy initialization.

// Subject interface
type Database interface {
	Query(sql string) string
	Execute(sql string) bool
}

// RealSubject - actual database
type RealDatabase struct {
	name string
}

func NewRealDatabase(name string) *RealDatabase {
	fmt.Printf("[RealDB] Connecting to %s...\n", name)
	return &RealDatabase{name: name}
}

func (db *RealDatabase) Query(sql string) string {
	return fmt.Sprintf("Result from %s: [rows for '%s']", db.name, sql)
}

func (db *RealDatabase) Execute(sql string) bool {
	fmt.Printf("[RealDB] Executing: %s\n", sql)
	return true
}

// CachingProxy - adds caching layer on top
type CachingProxy struct {
	db    *RealDatabase
	cache map[string]cachedResult
	ttl   time.Duration
}

type cachedResult struct {
	value     string
	timestamp time.Time
}

func NewCachingProxy(dbName string, ttl time.Duration) *CachingProxy {
	return &CachingProxy{
		db:    NewRealDatabase(dbName),
		cache: make(map[string]cachedResult),
		ttl:   ttl,
	}
}

func (p *CachingProxy) Query(sql string) string {
	if cached, ok := p.cache[sql]; ok {
		if time.Since(cached.timestamp) < p.ttl {
			fmt.Println("[Proxy] Cache HIT")
			return cached.value
		}
		fmt.Println("[Proxy] Cache EXPIRED")
		delete(p.cache, sql)
	}

	fmt.Println("[Proxy] Cache MISS - forwarding to real DB")
	result := p.db.Query(sql)
	p.cache[sql] = cachedResult{value: result, timestamp: time.Now()}
	return result
}

func (p *CachingProxy) Execute(sql string) bool {
	// Writes invalidate the cache
	p.cache = make(map[string]cachedResult)
	fmt.Println("[Proxy] Cache invalidated due to write")
	return p.db.Execute(sql)
}

func main() {
	var db Database = NewCachingProxy("production", 5*time.Second)

	// First query - cache miss
	r1 := db.Query("SELECT * FROM users")
	fmt.Println(r1)

	// Second query - cache hit
	r2 := db.Query("SELECT * FROM users")
	fmt.Println(r2)

	if r1 == r2 {
		fmt.Println("PASS: cached result matches")
	} else {
		panic("FAIL: cached result mismatch")
	}

	// Write invalidates cache
	db.Execute("INSERT INTO users VALUES (1, 'alice')")

	// Next query - cache miss again
	r3 := db.Query("SELECT * FROM users")
	fmt.Println(r3)
	fmt.Println("PASS: proxy pattern complete")
}
