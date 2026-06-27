package main

import (
	"fmt"
	"sync"
	"time"
)

// Token Bucket Rate Limiter.
// Tokens are added at a fixed rate up to a maximum capacity.
// Each request consumes one token. If no tokens are available, request is denied.
// Thread-safe via sync.Mutex.

type TokenBucket struct {
	mu         sync.Mutex
	tokens     float64
	capacity   float64
	refillRate float64   // tokens per second
	lastRefill time.Time
}

func NewTokenBucket(capacity float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:     capacity, // Start full
		capacity:   capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request is permitted (consumes a token).
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}
	return false
}

// refill adds tokens based on elapsed time since last refill.
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

// Tokens returns the current token count (for monitoring).
func (tb *TokenBucket) Tokens() float64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.refill()
	return tb.tokens
}

func main() {
	// Allow 5 requests per second, burst up to 5
	limiter := NewTokenBucket(5, 5)

	// Simulate burst of 8 requests
	fmt.Println("--- Burst of 8 requests ---")
	for i := 1; i <= 8; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request %d: allowed=%v\n", i, allowed)
	}

	// Wait for tokens to refill
	fmt.Println("\n--- Waiting 1 second for refill ---")
	time.Sleep(1 * time.Second)
	fmt.Printf("Tokens available: %.1f\n\n", limiter.Tokens())

	// Concurrent access test
	fmt.Println("--- Concurrent requests (10 goroutines) ---")
	var wg sync.WaitGroup
	allowed := 0
	denied := 0
	var countMu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if limiter.Allow() {
				countMu.Lock()
				allowed++
				countMu.Unlock()
			} else {
				countMu.Lock()
				denied++
				countMu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Allowed: %d, Denied: %d\n", allowed, denied)
}
