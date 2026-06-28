package main

import (
	"fmt"
	"sync"
	"time"
)

// Rate Limiting Algorithms - Common strategies for controlling request rates.
// Covers: Token Bucket, Leaky Bucket, Fixed Window, Sliding Window Log.

// --- Token Bucket (most common) ---
// Tokens added at a fixed rate. Requests consume tokens. Allows bursts.
type TokenBucket struct {
	mu         sync.Mutex
	tokens     float64
	capacity   float64
	refillRate float64 // tokens/second
	lastRefill time.Time
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:     capacity,
		capacity:   capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

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

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

// --- Leaky Bucket ---
// Requests processed at a fixed rate. Excess requests queued up to capacity.
type LeakyBucket struct {
	mu         sync.Mutex
	queue      int
	capacity   int
	leakRate   int // requests processed per second
	lastLeak   time.Time
}

func NewLeakyBucket(capacity, leakRate int) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.leak()
	if lb.queue < lb.capacity {
		lb.queue++
		return true
	}
	return false
}

func (lb *LeakyBucket) leak() {
	now := time.Now()
	elapsed := now.Sub(lb.lastLeak).Seconds()
	leaked := int(elapsed * float64(lb.leakRate))
	if leaked > 0 {
		lb.queue -= leaked
		if lb.queue < 0 {
			lb.queue = 0
		}
		lb.lastLeak = now
	}
}

// --- Fixed Window Counter ---
// Simple counter reset every window interval.
type FixedWindow struct {
	mu          sync.Mutex
	count       int
	limit       int
	windowStart time.Time
	windowSize  time.Duration
}

func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:       limit,
		windowSize:  windowSize,
		windowStart: time.Now(),
	}
}

func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	now := time.Now()
	if now.Sub(fw.windowStart) >= fw.windowSize {
		fw.count = 0
		fw.windowStart = now
	}
	if fw.count < fw.limit {
		fw.count++
		return true
	}
	return false
}

// --- Sliding Window Log ---
// Tracks timestamps of all requests. Most accurate but memory-intensive.
type SlidingWindowLog struct {
	mu         sync.Mutex
	timestamps []time.Time
	limit      int
	windowSize time.Duration
}

func NewSlidingWindowLog(limit int, windowSize time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		limit:      limit,
		windowSize: windowSize,
	}
}

func (swl *SlidingWindowLog) Allow() bool {
	swl.mu.Lock()
	defer swl.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-swl.windowSize)

	// Remove expired entries
	valid := 0
	for _, ts := range swl.timestamps {
		if ts.After(cutoff) {
			swl.timestamps[valid] = ts
			valid++
		}
	}
	swl.timestamps = swl.timestamps[:valid]

	if len(swl.timestamps) < swl.limit {
		swl.timestamps = append(swl.timestamps, now)
		return true
	}
	return false
}

// RateLimiter interface for polymorphic usage
type RateLimiter interface {
	Allow() bool
}

func testRateLimiter(name string, rl RateLimiter, requests int) (allowed, denied int) {
	for i := 0; i < requests; i++ {
		if rl.Allow() {
			allowed++
		} else {
			denied++
		}
	}
	fmt.Printf("[%s] %d requests: %d allowed, %d denied\n", name, requests, allowed, denied)
	return
}

func main() {
	// Test Token Bucket: 10 tokens, refill 10/sec
	fmt.Println("--- Token Bucket (capacity=10, refill=10/s) ---")
	tb := NewTokenBucket(10, 10)
	allowed, denied := testRateLimiter("TokenBucket", tb, 15)
	if allowed != 10 || denied != 5 {
		panic(fmt.Sprintf("FAIL: expected 10 allowed, 5 denied, got %d/%d", allowed, denied))
	}
	fmt.Println("PASS: burst of 10, then rate limited")

	// After waiting, tokens refill
	time.Sleep(500 * time.Millisecond)
	if !tb.Allow() {
		panic("FAIL: should allow after refill")
	}
	fmt.Println("PASS: allows after refill")

	// Test Leaky Bucket: capacity=5, leak=10/sec
	fmt.Println("\n--- Leaky Bucket (capacity=5, leak=10/s) ---")
	lb := NewLeakyBucket(5, 10)
	allowed, denied = testRateLimiter("LeakyBucket", lb, 8)
	if allowed != 5 {
		panic(fmt.Sprintf("FAIL: expected 5 allowed, got %d", allowed))
	}
	fmt.Println("PASS: queued up to capacity")

	// Test Fixed Window: 5 requests per 1 second window
	fmt.Println("\n--- Fixed Window (limit=5, window=1s) ---")
	fw := NewFixedWindow(5, 1*time.Second)
	allowed, denied = testRateLimiter("FixedWindow", fw, 10)
	if allowed != 5 || denied != 5 {
		panic(fmt.Sprintf("FAIL: expected 5/5, got %d/%d", allowed, denied))
	}
	fmt.Println("PASS: fixed window enforced")

	// After window resets
	time.Sleep(1 * time.Second)
	if !fw.Allow() {
		panic("FAIL: should allow in new window")
	}
	fmt.Println("PASS: allows in new window")

	// Test Sliding Window Log: 5 requests per 1 second
	fmt.Println("\n--- Sliding Window Log (limit=5, window=1s) ---")
	swl := NewSlidingWindowLog(5, 1*time.Second)
	allowed, denied = testRateLimiter("SlidingWindowLog", swl, 8)
	if allowed != 5 || denied != 3 {
		panic(fmt.Sprintf("FAIL: expected 5/3, got %d/%d", allowed, denied))
	}
	fmt.Println("PASS: sliding window enforced")

	fmt.Println("\n--- All rate limiting algorithms working correctly ---")
}
